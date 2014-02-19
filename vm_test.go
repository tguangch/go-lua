package lua

import "testing"

func TestVm(t *testing.T) {
	l := NewState()
	BaseOpen(l)
	LoadFile(l, "fixtures/fib.lua", "t")
	Call(l, 0, 0)
}

func TestConcat(t *testing.T) {
	l := NewState()
	BaseOpen(l)
	LoadString(l, "print('hello'..'world')")
	Call(l, 0, 0)
}

func TestProtectedCall(t *testing.T) {
	l := NewState()
	OpenLibraries(l)
	LoadString(l, "assert(not pcall(bit32.band, {}))")
	Call(l, 0, 0)
}
