package jot

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func ExampleFile_Generate() {
	File("hello").
		AddStruct(Struct("HelloPrinter").
			AddField(Field("hello", Type(reflect.TypeOf((*string)(nil)).Elem())).SetDoc("This is a field.")).
			// We could also init HelloPrinters spec outside of this chain and use
			// Struct.Type() to get the type.
			AddMethodRecv("r", true, Function("Print").
				AddParameter("who", Type(reflect.TypeOf((*string)(nil)).Elem())).
				AddReturnType(Type(reflect.TypeOf((*string)(nil)).Elem())).
				AddCodeFmt(`return fmt.Sprintf("%s %s\n", r.hello, {0})`, "who"))).
		Generate(os.Stdout)
	// Output:
	// package hello
	//
	// import "fmt"
	//
	// type HelloPrinter struct {
	// 	// This is a field.
	//	hello string
	// }
	//
	// func (r *HelloPrinter) Print(who string) string {
	// 	return fmt.Sprintf("%s %s\n", r.hello, who)
	// }
	//
}

type GenerateSpec struct {
	Cmd []interface{}
}

func (s *GenerateSpec) Write(w io.Writer) error {
	io.WriteString(w, fmt.Sprintf("//go:generate %s", s.Cmd...))
	return nil
}

func Generate(args ...interface{}) *GenerateSpec {
	return &GenerateSpec{Cmd: args}
}

func ExampleFile_Generate_customSpec() {
	File("plugin_example").
		AddSpec(Generate("ls")).
		Generate(os.Stdout)
	// Output:
	// package plugin_example
	//
	// //go:generate ls
}
