package rt

import (
	"strings"
	"testing"

	"github.com/nooga/let-go/pkg/vm"
)

// TestOpcodeDisassemblyCoverage asserts the disassembler covers exactly the set
// of opcodes the VM defines: every opcode 0..OP_DIV has a real mnemonic in the
// name table and a defined instruction stride, and there are no stale
// name-table entries past the last opcode. This is the guardrail that makes
// opcode enum changes (e.g. the v3 trace-opcode removal + renumber) verifiable:
// if the enum, the OpcodeToString name slice, and opcodeStride ever drift out
// of alignment, this fails.
func TestOpcodeDisassemblyCoverage(t *testing.T) {
	for op := int32(0); op <= vm.OP_DIV; op++ {
		name := vm.OpcodeToString(op)
		if strings.Contains(name, "???") {
			t.Errorf("opcode %d has no name-table entry (OpcodeToString returned %q)", op, name)
		}
		if s := opcodeStride(op); s < 1 {
			t.Errorf("opcode %d (%s) has invalid stride %d", op, strings.TrimSpace(name), s)
		}
	}

	// The name table must not carry entries past the last real opcode — a stale
	// trailing entry would mean the enum and the table have drifted.
	if beyond := vm.OpcodeToString(vm.OP_DIV + 1); !strings.Contains(beyond, "???") {
		t.Errorf("expected opcode %d (past OP_DIV) to be unknown, got %q — stale name-table entry?",
			vm.OP_DIV+1, beyond)
	}
}
