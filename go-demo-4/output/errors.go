package output

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintError(value any) {
	val, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", val)
	}
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	default:
		color.Red("123123123")
	}
}

func sum[T int | float64 | string, V int](a, b T) (T, V) {
	// value, ok := x.(int)
	switch d := any(a).(type) {
	case string:
		fmt.Println(d)
	}
	return a + b, 2
}

type List[T int | string] struct {
	elements []T
}

func (l *List[T]) Add(element T) {

}
