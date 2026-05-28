//go:build bootstrap

// lgbgen compiles core.lg and all embedded namespaces into a pre-compiled .lgb bundle,
// or with --target=go, generates Go source files from the Go-lowering pipeline.
//
// Usage:
//
//	go run -tags bootstrap ./cmd/lgbgen [output-path]              # .lgb bundle (default)
//	go run -tags bootstrap ./cmd/lgbgen --target=go <out-dir>       # Go source bootstrap
//
// Default .lgb output: pkg/rt/core_compiled.lgb (when run from repo root)
// Default Go output dir: pkg/rt/core_go_lowered
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nooga/let-go/pkg/bytecode"
	"github.com/nooga/let-go/pkg/compiler"
	"github.com/nooga/let-go/pkg/rt"
	"github.com/nooga/let-go/pkg/vm"
)

// pipelineSourceFingerprint hashes every embedded .lg file that participates
// in the Go-lowering pipeline. Any change here invalidates the cache for
// every generated file (a pipeline change affects all lowerings).
func pipelineSourceFingerprint() string {
	h := sha256.New()
	srcs := []*string{
		&rt.CoreSrc, // build/optimize/lower live here via macroexpansion
		&rt.IRZipperSrc, &rt.IRPassesSrc, &rt.IRDominanceSrc,
		&rt.IRLowerSrc, &rt.IRLowerGoSrc,
		&rt.IRPassDCESrc, &rt.IRPassConstFoldSrc, &rt.IRPassMutabilitySrc,
		&rt.IRPassCSESrc, &rt.IRPassTypeInferSrc, &rt.IRPassLICMSrc,
		&rt.IRPassInferArgTypesSrc, &rt.IRBuildSrc, &rt.IRValidateSrc,
		&rt.IRPassPipelineSrc, &rt.IRDumpSrc, &rt.IRPassTraceSrc,
		&rt.GraphSrc, &rt.IRDataSrc,
	}
	for _, s := range srcs {
		if s != nil && *s != "" {
			h.Write([]byte(*s))
			h.Write([]byte{0}) // separator so concatenation can't collide
		}
	}
	return hex.EncodeToString(h.Sum(nil))
}

// sha256Hex returns the hex-encoded SHA-256 of s.
func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// headerHashes captures the cache fingerprint a generated file declares
// in its leading comments. A successful parse + match against the current
// inputs lets the generator skip lowering entirely.
type headerHashes struct {
	nsHash       string
	pipelineHash string
}

// readHeaderHashes parses the cache header from an existing generated file.
// Returns nil when the file doesn't exist, can't be read, or doesn't carry
// both expected gogen.* lines within the first handful of lines.
func readHeaderHashes(path string) *headerHashes {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var h headerHashes
	for i := 0; i < 16 && scanner.Scan(); i++ {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "// gogen.nsHash:"):
			h.nsHash = strings.TrimSpace(strings.TrimPrefix(line, "// gogen.nsHash:"))
		case strings.HasPrefix(line, "// gogen.pipelineHash:"):
			h.pipelineHash = strings.TrimSpace(strings.TrimPrefix(line, "// gogen.pipelineHash:"))
		case strings.HasPrefix(line, "package "):
			// Header section ended; stop scanning even if we haven't found
			// both hashes (older files predate the header — treat as miss).
			i = 16
		}
	}
	if h.nsHash == "" || h.pipelineHash == "" {
		return nil
	}
	return &h
}

// embeddedNS lists all embedded namespaces in compilation order.
// Dependencies must come before dependents (test depends on walk).
var embeddedNS = []struct {
	name string
	src  *string
}{
	{"core", &rt.CoreSrc},
	{"walk", &rt.WalkSrc},
	{"string", &rt.StringSrc},
	{"set", &rt.SetSrc},
	{"pprint", &rt.PprintSrc},
	{"edn", &rt.EdnSrc},
	{"io", &rt.IoSrc},
	{"async", &rt.AsyncSrc},
	{"test", &rt.TestSrc}, // depends on walk — must come after
	// ir.data is loaded from source on demand (like `zip` and `data`)
	// because precompiled ns chunks only replay nil stubs for defn,
	// not function bodies — the intern block at the bottom of data.lg
	// must run with live function values, which only happens via the
	// source-load path in the resolver's loadEmbedded.
	{"ir.zipper", &rt.IRZipperSrc},
	{"ir.passes", &rt.IRPassesSrc},
	{"ir.dominance", &rt.IRDominanceSrc},
	{"ir.lower", &rt.IRLowerSrc},
	{"ir.lower-go", &rt.IRLowerGoSrc},
	{"ir.passes.dce", &rt.IRPassDCESrc},
	{"ir.passes.constfold", &rt.IRPassConstFoldSrc},
	{"ir.passes.mutability", &rt.IRPassMutabilitySrc},
	{"ir.passes.cse", &rt.IRPassCSESrc},
	{"ir.passes.typeinfer", &rt.IRPassTypeInferSrc},
	{"ir.passes.licm", &rt.IRPassLICMSrc},
	{"ir.passes.infer-arg-types", &rt.IRPassInferArgTypesSrc},
	{"graph", &rt.GraphSrc},
	{"ir.build", &rt.IRBuildSrc},
	{"ir.validate", &rt.IRValidateSrc},
	{"ir.passes.pipeline", &rt.IRPassPipelineSrc},
	{"ir.dump", &rt.IRDumpSrc},
	{"ir.passes.trace", &rt.IRPassTraceSrc},
	// zip and data are loaded from source on demand (precompiled ns chunks
	// only replay nil stubs for defn, not the actual function bodies)
}

func main() {
	// Parse mode: --target=go <out-dir> or default bytecode mode
	targetGo := false
	outPath := "pkg/rt/core_compiled.lgb"
	goOutDir := "pkg/rt/core_go_lowered"

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--target=go":
			targetGo = true
			if i+1 < len(args) {
				goOutDir = args[i+1]
				i++
			}
		default:
			outPath = args[i]
		}
	}

	// rt.init() has already run — native builtins are registered in CoreNS.
	consts := vm.NewConsts()

	// Phase 0: bootstrap ir.data (same for both modes).
	{
		coreNS := rt.NS(rt.NameCoreNS)
		c := compiler.NewCompiler(consts, coreNS)
		c.SetSource("<embedded:ir.data:lgbgen-bootstrap>")
		if _, _, err := c.CompileMultiple(strings.NewReader(rt.IRDataSrc)); err != nil {
			fmt.Fprintf(os.Stderr, "ir.data bootstrap compilation failed: %v\n", err)
			os.Exit(1)
		}
	}

	// Phase 1: compile all namespaces from source (bytecode, sets up VM state).
	nsChunks := make(map[string]*vm.CodeChunk)
	nsOrder := make([]string, 0, len(embeddedNS))

	for _, ns := range embeddedNS {
		src := *ns.src
		if src == "" {
			continue
		}
		coreNS := rt.NS(rt.NameCoreNS)
		c := compiler.NewCompiler(consts, coreNS)
		c.SetSource("<embedded:" + ns.name + ">")

		chunk, _, err := c.CompileMultiple(strings.NewReader(src))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s compilation failed: %v\n", ns.name, err)
			os.Exit(1)
		}
		nsChunks[ns.name] = chunk
		nsOrder = append(nsOrder, ns.name)
		// Meta-circular self-hosting (LGGOIR_SELFHOST=1) clobbers the
		// just-replayed Lisp vars with their NativeFn counterparts so
		// subsequent --target=go pipeline invocations run native. Disabled
		// by default while the underlying gogen-emitted dispatch path
		// still has pre-existing bugs (nil-pointer chains via reflective
		// BoxNativeFn closures inside toposort-by → reduce → native fn).
		// See task tracking for the parity hunt.
		if os.Getenv("LGGOIR_SELFHOST") != "" {
			if targetNS := rt.LookupNS(ns.name); targetNS != nil {
				rt.ApplyGoOverrides(targetNS)
			}
		}
		if targetGo {
			fmt.Printf("  compiled %-20s (%d bytecode + lowering to Go)\n", ns.name, len(chunk.Code())*4)
		} else {
			fmt.Printf("  compiled %-10s (%d bytes bytecode)\n", ns.name, len(chunk.Code())*4)
		}
	}
	if targetGo && os.Getenv("LGGOIR_SELFHOST") != "" {
		if coreNS := rt.LookupNS(rt.NameCoreNS); coreNS != nil {
			rt.ApplyGoOverrides(coreNS)
		}
	}

	if targetGo {
		runGoTarget(goOutDir)
		return
	}

	// Bytecode mode: write .lgb bundle.
	f, err := os.Create(outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %s: %v\n", outPath, err)
		os.Exit(1)
	}
	defer f.Close()

	if err := bytecode.EncodeBundleOrdered(f, consts, nsChunks, nsOrder); err != nil {
		fmt.Fprintf(os.Stderr, "encode failed: %v\n", err)
		os.Exit(1)
	}

	fi, _ := f.Stat()
	fmt.Printf("wrote %s (%d bytes, %d consts, %d namespaces)\n",
		outPath, fi.Size(), len(consts.Values()), len(nsChunks))
}

// nsToGoPkgName returns the Go package name for a namespace — the LAST
// dotted segment, with hyphens replaced. ir.dominance → "dominance",
// ir.passes.dce → "dce". This mirrors Go's convention of one short
// package name per directory.
func nsToGoPkgName(nsName string) string {
	last := nsName
	if i := strings.LastIndex(nsName, "."); i >= 0 {
		last = nsName[i+1:]
	}
	return strings.ReplaceAll(last, "-", "_")
}

// nsToGoPkgPath returns the directory path components for a namespace.
// ir.dominance → ["ir", "dominance"], ir.passes.dce → ["ir", "passes", "dce"].
// Each component has hyphens replaced with underscores.
func nsToGoPkgPath(nsName string) []string {
	parts := strings.Split(nsName, ".")
	out := make([]string, len(parts))
	for i, p := range parts {
		out[i] = strings.ReplaceAll(p, "-", "_")
	}
	return out
}

// runGoTarget re-pipelines each namespace's defn forms through the Go
// lowering pipeline and writes .go source files to outDir.
func runGoTarget(outDir string) {
	pipelineVar := rt.LookupVar("ir.passes.pipeline", "lower-ns-to-go")
	if pipelineVar == nil {
		fmt.Fprintf(os.Stderr, "fatal: ir.passes.pipeline/lower-ns-to-go not found\n")
		os.Exit(1)
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "create output dir %s: %v\n", outDir, err)
		os.Exit(1)
	}

	// Each generated file carries its own input fingerprint in leading
	// comments — no external manifest file. On regen, parse the existing
	// file's header; if both hashes match, skip the lowering pipeline
	// entirely. Self-describing files survive moves, renames, and partial
	// deletions without a separate index getting out of sync.
	pipelineFP := pipelineSourceFingerprint()

	for _, ns := range embeddedNS {
		src := *ns.src
		if src == "" {
			continue
		}

		// Read forms from source and pick out defn forms only.
		// defmacro forms are skipped for now — their bodies are macro
		// template construction code that doesn't lower cleanly.
		r := compiler.NewLispReader(strings.NewReader(src), "<embedded:"+ns.name+">")
		var defnForms []vm.Value
		for {
			form, err := r.Read()
			if err != nil {
				if strings.Contains(err.Error(), "EOF") {
					break
				}
				fmt.Fprintf(os.Stderr, "%s: read error: %v\n", ns.name, err)
				os.Exit(1)
			}
			if isDefnOnly(form) {
				defnForms = append(defnForms, form)
			}
		}

		if len(defnForms) == 0 {
			fmt.Printf("  skip %-20s (no defn forms)\n", ns.name)
			continue
		}

		// Each namespace maps to its own Go package — directory mirrors the
		// dotted ns path, package name is the last segment. ir.passes.dce
		// becomes outDir/ir/passes/dce/ with package dce.
		pkgName := nsToGoPkgName(ns.name)
		pathParts := nsToGoPkgPath(ns.name)
		pkgDir := filepath.Join(append([]string{outDir}, pathParts...)...)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "create pkg dir %s: %v\n", pkgDir, err)
			os.Exit(1)
		}
		filename := filepath.Join(pkgDir, pkgName+".go")

		// Cache check: skip lowering when the existing file's header records
		// the same nsHash AND pipelineHash. The Go lowering pipeline is the
		// slow part (>10s per heavy ns), so this is the single biggest
		// dev-loop speedup.
		nsHash := sha256Hex(src)
		if existing := readHeaderHashes(filename); existing != nil &&
			existing.nsHash == nsHash && existing.pipelineHash == pipelineFP {
			fmt.Printf("  cache %-40s (unchanged)\n", filename)
			continue
		}

		// Batch all forms into one namespace file.
		formsVec := vm.NewPersistentVector(defnForms)
		args := []vm.Value{
			vm.String(pkgName),
			vm.Symbol(ns.name),
			formsVec,
		}
		result, err := pipelineVar.Invoke(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: Go lowering failed: %v\n", ns.name, err)
			continue
		}
		goSrc, ok := result.(vm.String)
		if !ok {
			fmt.Fprintf(os.Stderr, "%s: expected string result, got %T\n", ns.name, result)
			continue
		}
		// Header carries: build tag, generator marker, and the input/pipeline
		// hashes — read by the cache check above on subsequent runs.
		header := "//go:build gogen_ir\n\n" +
			"// Code generated by lgbgen --target=go. DO NOT EDIT.\n" +
			"// gogen.nsHash:       " + nsHash + "\n" +
			"// gogen.pipelineHash: " + pipelineFP + "\n\n"
		if err := os.WriteFile(filename, []byte(header+string(goSrc)), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "%s: write %s: %v\n", ns.name, filename, err)
			os.Exit(1)
		}
		fmt.Printf("  wrote %s (%d bytes, %d defns)\n", filename, len(goSrc), len(defnForms))
	}
}

// isDefnLike returns true if form is a list beginning with defn or defmacro.
func isDefnLike(form vm.Value) bool {
	return isDefnOnly(form) || isDefmacroForm(form)
}

func isDefmacroForm(form vm.Value) bool {
	list, ok := form.(vm.Sequable)
	if !ok {
		return false
	}
	seq := list.Seq()
	if seq == nil {
		return false
	}
	head := seq.First()
	sym, ok := head.(vm.Symbol)
	if !ok {
		return false
	}
	return string(sym) == "defmacro"
}

// isDefnOnly returns true if form is a list beginning with defn (not defmacro).
func isDefnOnly(form vm.Value) bool {
	list, ok := form.(vm.Sequable)
	if !ok {
		return false
	}
	seq := list.Seq()
	if seq == nil {
		return false
	}
	head := seq.First()
	sym, ok := head.(vm.Symbol)
	if !ok {
		return false
	}
	return string(sym) == "defn" || string(sym) == "defn-"
}
