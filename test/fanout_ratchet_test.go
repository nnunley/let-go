package test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// lgBin is the lg binary built once for all fanout-ratchet integration tests.
var lgBin string

// repoRoot is the absolute path to the repository root (tests run from test/).
var repoRoot string

func TestMain(m *testing.M) {
	tmp, err := os.MkdirTemp("", "lg-fanout-*")
	if err != nil {
		panic(err)
	}
	lgBin = filepath.Join(tmp, "lg-bin")
	build := exec.Command("go", "build", "-o", lgBin, "../")
	build.Dir = "."
	if out, err := build.CombinedOutput(); err != nil {
		panic("build lg: " + err.Error() + "\n" + string(out))
	}
	repoRoot, _ = filepath.Abs("..")
	code := m.Run()
	os.RemoveAll(tmp)
	os.Exit(code)
}

// runRatchet runs scripts/fanout-ratchet.lg with args from repo root; returns
// combined output and exit code.
func runRatchet(t *testing.T, args ...string) (string, int) {
	t.Helper()
	full := append([]string{"scripts/fanout-ratchet.lg"}, args...)
	cmd := exec.Command(lgBin, full...)
	cmd.Dir = repoRoot
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	code := 0
	if ee, ok := err.(*exec.ExitError); ok {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("run ratchet: %v\noutput:\n%s", err, out.String())
	}
	return out.String(), code
}

// writeModule writes <dir>/<module>/<module>.go with exactly `size` bytes
// (size-1 'a's + trailing newline). file-size in the script == chars == bytes.
func writeModule(t *testing.T, treeDir, module string, size int) {
	t.Helper()
	d := filepath.Join(treeDir, module)
	if err := os.MkdirAll(d, 0755); err != nil {
		t.Fatal(err)
	}
	body := strings.Repeat("a", size-1) + "\n"
	if err := os.WriteFile(filepath.Join(d, module+".go"), []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func writeBaseline(t *testing.T, path, edn string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(edn), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestFanoutShow(t *testing.T) {
	tree := t.TempDir()
	writeModule(t, tree, "a", 1000)
	out, code := runRatchet(t, "show", "--no-regen", "--no-wireup", "--tree-dir", tree)
	if code != 0 {
		t.Fatalf("show exit=%d, want 0\n%s", code, out)
	}
	if !strings.Contains(out, ":total-bytes 1000") {
		t.Errorf("expected :total-bytes 1000 in output:\n%s", out)
	}
	if !strings.Contains(out, "\"a\"") {
		t.Errorf("expected module \"a\" in output:\n%s", out)
	}
}
