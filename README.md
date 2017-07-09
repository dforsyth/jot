# jot

jot is a simple and incomplete API for programmatically generating Go code.

```go
jot.File("hello").
    AddStruct(jot.Struct("HelloPrinter").
        AddField(Field("hello", Type(reflect.TypeOf((*string)(nil)).Elem())).SetDoc("This is a field.")).
        // We could also init HelloPrinters spec outside of this chain and use
        // Struct.Type() to get the type.
        AddMethodRecv("r", true, jot.Function("Print").
            AddParameter("who", jot.Type(reflect.TypeOf((*string)(nil)).Elem())).
            AddReturnType(jot.Type(reflect.TypeOf((*string)(nil)).Elem())).
            AddCodeFmt(`return fmt.Sprintf("%s %s\n", r.hello, {0})`, "who"))).
    Generate(os.Stdout)
```
which generates:
```go
package hello

import "fmt"

type HelloPrinter struct{
	// This is a field.
	hello string
}

func (r *HelloPrinter) Print(who string) string {
	return fmt.Sprintf("%s %s\n", r.hello, who)
}
```

jot also allows the user to plug in their own generators that satisfy the Spec
interface:
```go
type GenerateSpec struct {
    Cmd []interface{}
}

func (s *GenerateSpec) Write(w io.Writer) error {
    io.WriteString(w, fmt.Sprintf("//go:generate %s", s.Cmd...))
}

func Generate(args ...interface{}) Spec {
    return &GenerateSpec{Cmd: args}
}

jot.File("plugin_example").
	AddSpec(Generate("ls")).
	Generate(os.Stdout)
```
which generates:
```go
package plugin_example

//go:generate ls
```
