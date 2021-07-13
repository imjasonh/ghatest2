package main

func main() {
        _ = New[*foo]()
}

type i interface { Foo() }
type foo struct {}
func (f *foo) Foo() {} // satisfy interface

func New[T i]() str[T] {
        return str[T]{}
}

type str[T i] struct {
        s string
}

