package wasmhost

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

func TestHostEvalCanRequireIRPipeline(t *testing.T) {
	consts := vm.NewConsts()
	ctx := compiler.NewCompiler(consts, rt.NS("user"))
	NewResolver(ctx)
	host := New(consts)

	var frames []Frame
	host.Handle(Request{
		ID:      "probe",
		Session: "test",
		Op:      "eval",
		Code:    "(require 'ir.passes.pipeline)\n:ok",
	}, func(frameJSON string) {
		frame, err := DecodeFrame([]byte(frameJSON))
		if err != nil {
			t.Fatalf("DecodeFrame: %v", err)
		}
		frames = append(frames, frame)
	})

	for _, frame := range frames {
		if frame.Kind == "err" {
			t.Fatalf("unexpected error frame: %s", frame.Text)
		}
	}
	if got := lastValueFrame(frames); got != ":ok" {
		t.Fatalf("last value = %q, want :ok", got)
	}
	if v := rt.LookupVar("ir.passes.pipeline", "compile-form"); v == nil {
		t.Fatal("ir.passes.pipeline/compile-form var not found after require")
	}
}

func TestHostEvalEmitsRequestedIRArtifacts(t *testing.T) {
	consts := vm.NewConsts()
	ctx := compiler.NewCompiler(consts, rt.NS("user"))
	NewResolver(ctx)
	host := New(consts)

	frames := runRequest(t, host, Request{
		ID:      "inspect",
		Session: "test",
		Op:      "eval",
		Code:    "(defn add [x y] (+ x y))",
		Inspect: InspectOptions{
			Bytecode:          true,
			IR:                true,
			OptimizedBytecode: true,
			LoweredGo:         true,
		},
	})

	var sawBytecode, sawIR, sawOptBytecode, sawLowered bool
	for _, frame := range frames {
		if frame.Kind != "artifact" {
			continue
		}
		switch frame.Artifact {
		case "bytecode":
			sawBytecode = true
		case "ir":
			sawIR = true
			if got, ok := frame.Content.(string); !ok || !strings.Contains(got, "fn add") {
				t.Fatalf("ir artifact = %#v, want dump containing fn add", frame.Content)
			}
		case "optimized_bytecode":
			sawOptBytecode = true
		case "lowered_go":
			sawLowered = true
			if got, ok := frame.Content.(string); !ok || !strings.Contains(got, "func add(") {
				t.Fatalf("lowered_go artifact = %#v, want Go decl containing func add(", frame.Content)
			}
		}
	}
	if !sawIR {
		got, err := host.inspectIR("user", "(defn add [x y] (+ x y))")
		t.Fatalf("missing ir artifact; inspectIR err=%v dump=%q", err, got)
	}
	if !sawLowered {
		got, err := host.inspectLoweredGo("user", "(defn add [x y] (+ x y))")
		t.Fatalf("missing lowered_go artifact; inspectLoweredGo err=%v decl=%q", err, got)
	}
	if !sawBytecode || !sawIR || !sawOptBytecode || !sawLowered {
		t.Fatalf("artifact coverage bytecode=%v ir=%v optimized_bytecode=%v lowered_go=%v", sawBytecode, sawIR, sawOptBytecode, sawLowered)
	}
}

func TestHostEvalSkipsIRArtifactsForNonLowerableForms(t *testing.T) {
	consts := vm.NewConsts()
	ctx := compiler.NewCompiler(consts, rt.NS("user"))
	NewResolver(ctx)
	host := New(consts)

	frames := runRequest(t, host, Request{
		ID:      "mixed",
		Session: "test",
		Op:      "eval",
		Code:    "(defn add [x y] (+ x y))\n(add 1 2)",
		Inspect: InspectOptions{
			IR:                true,
			OptimizedBytecode: true,
			LoweredGo:         true,
		},
	})

	for _, frame := range frames {
		if frame.Kind == "err" && strings.Contains(frame.Text, "nth index out of bounds") {
			t.Fatalf("non-lowerable form should not trigger IR compile-form crash: %s", frame.Text)
		}
	}
	var irArtifactsOnSecond int
	for _, frame := range frames {
		if frame.Kind == "artifact" && frame.FormIndex == 2 {
			irArtifactsOnSecond++
		}
	}
	if irArtifactsOnSecond != 0 {
		t.Fatalf("second non-lowerable form emitted %d artifacts, want 0", irArtifactsOnSecond)
	}
}

func TestHostEvalEmitsArtifactErrorsForFailedExtraction(t *testing.T) {
	consts := vm.NewConsts()
	ctx := compiler.NewCompiler(consts, rt.NS("user"))
	NewResolver(ctx)
	host := New(consts)

	frames := runRequest(t, host, Request{
		ID:      "artifact-error",
		Session: "test",
		Op:      "eval",
		Code:    "(+ 1 2)",
		Inspect: InspectOptions{
			OptimizedBytecode: true,
		},
	})

	var sawArtifactErr bool
	for _, frame := range frames {
		if frame.Kind == "err" && frame.Artifact == "optimized_bytecode" {
			sawArtifactErr = true
			if !strings.Contains(frame.Text, "unsupported") {
				t.Fatalf("artifact error text = %q, want unsupported extraction failure", frame.Text)
			}
		}
	}
	if !sawArtifactErr {
		t.Fatal("missing optimized_bytecode extraction error frame")
	}
}

func runRequest(t *testing.T, host *Host, req Request) []Frame {
	t.Helper()
	var frames []Frame
	host.Handle(req, func(frameJSON string) {
		frame, err := DecodeFrame([]byte(frameJSON))
		if err != nil {
			t.Fatalf("DecodeFrame: %v", err)
		}
		frames = append(frames, frame)
	})
	return frames
}

func DecodeFrame(bs []byte) (Frame, error) {
	var frame Frame
	err := json.Unmarshal(bs, &frame)
	return frame, err
}

func lastValueFrame(frames []Frame) string {
	last := ""
	for _, frame := range frames {
		if frame.Kind == "value" {
			last = strings.TrimSpace(frame.Value)
		}
	}
	return last
}
