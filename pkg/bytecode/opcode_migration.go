package bytecode

import "github.com/nooga/let-go/pkg/vm"

// Frozen opcode migration: v1/v2 → v3.
//
// v1/v2 bundles used an opcode numbering that included OP_TRACE_ENABLE (18) and
// OP_TRACE_DISABLE (19). v3 retired both, shifting every opcode >= 20 down by 2.
// The decoder reads a chunk's code as a flat []int32 where opcode words are
// interleaved with argument words, so the remap must walk instruction by
// instruction using the FROZEN v2 stride table below — never rewriting an
// argument word.
//
// This table is intentionally decoupled from the live vm.OP_* constants (which
// now describe v3). It is frozen: do not update it when opcodes change again;
// add a v3→v4 table instead.
const (
	v2opTraceEnable  int32 = 18
	v2opTraceDisable int32 = 19
	v2opFirstShifted int32 = 20 // OP_MAKE_MULTI_ARITY in v2 numbering
)

// v2Stride returns the width, in int32 words, of a v2 opcode's instruction.
// Mirrors pkg/rt/disasm.go opcodeStride as of FormatVersion 2.
func v2Stride(op int32) int {
	switch op & 0xff {
	case 16: // OP_RECUR (offset, argc)
		return 4
	case 22: // OP_TRY_PUSH (catchOffset, finallyOffset)
		return 3
	case 1, 2, 3, 5, 6, 7, 9, 10, 12, 14, 17, 20, 21:
		// LOAD_CONST, LOAD_ARG, INVOKE, BRANCH_TRUE, BRANCH_FALSE, JUMP,
		// POP_N, DUP_NTH, LOAD_VAR, LOAD_CLOSEDOVER, RECUR_FN,
		// MAKE_MULTI_ARITY, TAIL_CALL — one inline argument word.
		return 2
	default:
		return 1
	}
}

// remapV2Opcode maps a v2 opcode byte to its v3 value. TRACE_ENABLE/DISABLE
// become NOOP (also stride 1, so instruction width is preserved).
func remapV2Opcode(op int32) int32 {
	switch {
	case op == v2opTraceEnable || op == v2opTraceDisable:
		return 0 // OP_NOOP
	case op >= v2opFirstShifted:
		return op - 2
	default:
		return op
	}
}

// remapLegacyChunks rewrites the opcode word of every instruction in each chunk
// from v1/v2 numbering to v3, in place. It walks by v2 stride so argument words
// are left untouched, and only the low byte (the opcode) of an opcode word is
// changed — packed high bits (e.g. the stack-pointer hint) are preserved. Every
// remap preserves instruction width (TRACE→NOOP is stride 1), so IPs, jump
// targets, and source-map offsets remain valid.
func remapLegacyChunks(chunks []*vm.CodeChunk) {
	for _, ch := range chunks {
		code := ch.Code()
		for i := 0; i < len(code); {
			op := code[i] & 0xff
			stride := v2Stride(op)
			if stride < 1 {
				stride = 1
			}
			code[i] = (code[i] &^ 0xff) | remapV2Opcode(op)
			i += stride
		}
	}
}
