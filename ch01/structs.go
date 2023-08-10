package main

import "fmt"

type Test struct {
	field int
}

func (this Test) WrongModifier(val int) {
	this.field = val
}

func (this *Test) Modifier(val int) {
	this.field = val
}

func Detect(i interface{}) (ty string) {
	switch v := i.(type) {

	case int, int8, int16, int32, int64:
		ty = fmt.Sprint(v, " is an ", "Integer")
	case string:
		ty = fmt.Sprint(v, " is a ", "String")
	case float64, float32:

		ty = fmt.Sprint(v, " is a ", "Float")
	//etc
	default:
		ty = "Whatever"
	}

	return
}
func main() {
	x := Test{field: 3}
	x.WrongModifier(10)
	fmt.Println(x)

	x.Modifier(22)
	fmt.Println(Detect(x.field))
}
