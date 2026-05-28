package rt

import (
	"testing"

	"github.com/nooga/let-go/pkg/vm"
)

func TestApplyGoOverrides_OverridesNSVars(t *testing.T) {
	nsName := "rt.test.gogen.override"
	delete(nsRegistry, nsName)
	delete(pendingGoOverrides, nsName)
	defer delete(nsRegistry, nsName)
	defer delete(pendingGoOverrides, nsName)

	ns := vm.NewNamespace(nsName)
	nsRegistry[nsName] = ns
	ns.Def("greeting", vm.String("from-bytecode"))

	marker, _ := vm.NativeFnType.Wrap(func(_ []vm.Value) (vm.Value, error) {
		return vm.String("from-go"), nil
	})
	RegisterGoOverrides(nsName, map[string]vm.Value{"greeting": marker})

	ApplyGoOverrides(ns)

	v := ns.LookupLocal(vm.Symbol("greeting"))
	if v == nil {
		t.Fatalf("greeting not bound after override")
	}
	got := v.Deref()
	if got != marker {
		t.Fatalf("expected override fn, got %v", got)
	}

	// Second apply is a no-op: pending map is drained.
	if _, present := pendingGoOverrides[nsName]; present {
		t.Fatalf("expected pendingGoOverrides[%q] to be drained", nsName)
	}
}

func TestApplyGoOverrides_NoopWhenEmpty(t *testing.T) {
	nsName := "rt.test.gogen.override.empty"
	delete(nsRegistry, nsName)
	defer delete(nsRegistry, nsName)

	ns := vm.NewNamespace(nsName)
	nsRegistry[nsName] = ns
	ApplyGoOverrides(ns)
	ApplyGoOverrides(nil)
}
