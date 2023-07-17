package monkey

import (
	"testing"
)

func fnFoo() string {
	return "foo"
}

var varFoo = "foo"

type foo struct{}

func (f *foo) PublicFoo(x string) string {
	return "foo " + x
}

func (f *foo) privateFoo(x string) string {
	return "foo " + x
}

func TestFunc(t *testing.T) {
	p := Func(nil, fnFoo, func() string { return "bar" })
	defer p.Reset()

	if fnFoo() != "bar" {
		t.Error("patch failed")
	}
}

func TestVar(t *testing.T) {
	p := Var(nil, &varFoo, "bar")
	defer p.Reset()

	if varFoo != "bar" {
		t.Error("patch failed")
	}
}

func TestMethod(t *testing.T) {
	t.Run("public method", func(t *testing.T) {
		f := new(foo)
		p := Method(nil, f, f.PublicFoo, func(x string) string { return "bar " + x })
		defer p.Reset()

		if f.PublicFoo("x") != "bar x" {
			t.Error("patch failed")
		}
	})

	t.Run("private method", func(t *testing.T) {
		f := new(foo)
		p := Method(nil, f, f.privateFoo, func(x string) string { return "bar " + x })
		defer p.Reset()

		if f.privateFoo("x") != "bar x" {
			t.Error("patch failed")
		}
	})
}
