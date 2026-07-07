package bytecode

import (
	"bytes"
	"testing"

	"github.com/nooga/let-go/pkg/vm"
)

// TestRemapLegacyChunksV2ToV3 exercises the frozen v1/v2 → v3 opcode remap over
// a hand-built v2 code array. Uses literal v2 opcode NUMBERS (not vm.OP_*, which
// now describe v3) so it pins the migration independently of the live enum.
func TestRemapLegacyChunksV2ToV3(t *testing.T) {
	const hi = 5 << 16 // a packed high-bits payload (e.g. stack-pointer hint)

	// v2 numbering: TRACE_ENABLE=18, TRACE_DISABLE=19, MAKE_MULTI_ARITY=20,
	// TAIL_CALL=21, TRY_PUSH=22, ADD=25, RECUR=16, LOAD_CONST=1, NOOP=0, RETURN=4.
	v2 := []int32{
		0,     // NOOP                      -> 0
		1, 25, // LOAD_CONST, arg=25        -> 1, 25 (arg must NOT remap)
		18,          // TRACE_ENABLE              -> 0 (NOOP)
		19,          // TRACE_DISABLE             -> 0 (NOOP)
		20 | hi, 99, // MAKE_MULTI_ARITY|hi, arg -> 18|hi, 99 (high bits kept)
		22, 30, 40, // TRY_PUSH, a=30, b=40      -> 20, 30, 40 (args kept)
		16, 7, 21, 8, // RECUR, a,b,c            -> 16, 7, 21, 8 (stride-4 args kept)
		25, // ADD                       -> 23
		4,  // RETURN                    -> 4
	}
	want := []int32{
		0,
		1, 25,
		0,
		0,
		18 | hi, 99,
		20, 30, 40,
		16, 7, 21, 8,
		23,
		4,
	}

	chunk := vm.NewCodeChunk(vm.NewConsts())
	chunk.Append(v2...)
	remapLegacyChunks([]*vm.CodeChunk{chunk})

	got := chunk.Code()
	if len(got) != len(want) {
		t.Fatalf("code length changed: got %d, want %d (remap must preserve width)", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("code[%d]: got %d, want %d", i, got[i], want[i])
		}
	}
}

// TestV3RoundTripIdentity confirms a v3-encoded module decodes to identical
// bytecode with no remap (the encoder writes the current FormatVersion, the
// decoder takes the v3 path).
func TestV3RoundTripIdentity(t *testing.T) {
	if FormatVersion != 3 {
		t.Fatalf("this test assumes FormatVersion 3, got %d", FormatVersion)
	}
	consts := vm.NewConsts()
	chunk := vm.NewCodeChunk(consts)
	// A mix spanning the renumber boundary in v3 terms.
	want := []int32{
		vm.OP_LOAD_CONST, 0,
		vm.OP_MAKE_MULTI_ARITY, 1,
		vm.OP_TAIL_CALL, 2,
		vm.OP_ADD,
		vm.OP_RETURN,
	}
	chunk.Append(want...)
	chunk.SetMaxStack(4)

	fn := vm.MakeFunc(1, false, chunk)
	fn.SetName("rt-fn")
	b := NewModuleBuilder()
	b.AddChunk(chunk)
	b.AddConst(fn)

	var buf bytes.Buffer
	if err := Encode(&buf, b.Build()); err != nil {
		t.Fatalf("Encode: %v", err)
	}
	decoded, err := Decode(&buf)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	gotFn, ok := decoded.Consts[0].(*vm.Func)
	if !ok {
		t.Fatalf("expected *Func, got %T", decoded.Consts[0])
	}
	got := gotFn.Chunk().Code()
	if len(got) != len(want) {
		t.Fatalf("code length: got %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("code[%d]: got %d, want %d", i, got[i], want[i])
		}
	}
}
