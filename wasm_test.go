package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestWasmGoBuildArgsDefault(t *testing.T) {
	got := wasmGoBuildArgs("/tmp/app.wasm", "")
	want := []string{"build", "-o", "/tmp/app.wasm", "."}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("args = %#v, want %#v", got, want)
	}
}

func TestWasmGoBuildArgsWithTags(t *testing.T) {
	got := wasmGoBuildArgs("/tmp/app.wasm", "gogen_ir")
	want := []string{"build", "-tags", "gogen_ir", "-o", "/tmp/app.wasm", "."}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("args = %#v, want %#v", got, want)
	}
}

func TestWasmGoBuildArgsWithMultiWordTags(t *testing.T) {
	got := wasmGoBuildArgs("/tmp/app.wasm", "gogen_ir,foo bar")
	want := []string{"build", "-tags", "gogen_ir,foo bar", "-o", "/tmp/app.wasm", "."}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("args = %#v, want %#v", got, want)
	}
}

func TestHasBuildTag(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		tag  string
		want bool
	}{
		{"", "gogen_ir", false},
		{"gogen_ir", "gogen_ir", true},
		{"foo gogen_ir", "gogen_ir", true},
		{"foo,gogen_ir,bar", "gogen_ir", true},
		{"foo\tbar\nbaz", "gogen_ir", false},
		{"gogen", "gogen_ir", false},
	} {
		if got := hasBuildTag(tc.raw, tc.tag); got != tc.want {
			t.Fatalf("hasBuildTag(%q, %q) = %v, want %v", tc.raw, tc.tag, got, tc.want)
		}
	}
}

func TestListGogenIRPackages(t *testing.T) {
	root := t.TempDir()
	mustMkdirAll(t, filepath.Join(root, "pkg", "rt", "core_go_lowered", "ir", "passes", "pipeline"))
	mustMkdirAll(t, filepath.Join(root, "pkg", "rt", "core_go_lowered", "core"))
	mustWriteFile(t, filepath.Join(root, "pkg", "rt", "core_go_lowered", "ir", "passes", "pipeline", "pipeline.go"), "package pipeline\n")
	mustWriteFile(t, filepath.Join(root, "pkg", "rt", "core_go_lowered", "core", "core.go"), "package core\n")
	mustMkdirAll(t, filepath.Join(root, "pkg", "rt", "core_go_lowered", "empty"))

	got, err := listGogenIRPackages(root)
	if err != nil {
		t.Fatalf("listGogenIRPackages: %v", err)
	}
	want := []string{"core", "ir/passes/pipeline"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("packages = %#v, want %#v", got, want)
	}
}

func TestWriteWasmGogenIRWireup(t *testing.T) {
	tmp := t.TempDir()
	src := t.TempDir()
	mustMkdirAll(t, filepath.Join(src, "pkg", "rt", "core_go_lowered", "ir", "passes", "pipeline"))
	mustWriteFile(t, filepath.Join(src, "pkg", "rt", "core_go_lowered", "ir", "passes", "pipeline", "pipeline.go"), "package pipeline\n")

	if err := writeWasmGogenIRWireup(tmp, src); err != nil {
		t.Fatalf("writeWasmGogenIRWireup: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(tmp, "main_gogen_ir.go"))
	if err != nil {
		t.Fatalf("read wireup: %v", err)
	}
	got := string(data)
	for _, want := range []string{
		"//go:build gogen_ir",
		"package main",
		"_ \"github.com/nooga/let-go/pkg/rt/core_go_lowered/ir/passes/pipeline\"",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("wireup missing %q", want)
		}
	}
}

func mustMkdirAll(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}

func mustWriteFile(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
