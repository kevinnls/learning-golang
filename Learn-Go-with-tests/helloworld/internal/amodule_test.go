package amodule
import "testing"

func TestModularHello(t *testing.T) {
	got := ModularHello()
	want := "Hello from amodule"
	if got != want {
		t.Errorf("oops")
	}
}
