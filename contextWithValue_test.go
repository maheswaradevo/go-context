package golangcontext

import (
	"context"
	"fmt"
	"testing"
)

//TestContextWithValue -> Inserting value to a context from parent to child
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	//Get value from context
	fmt.Println(contextF.Value("f"))	//Value F
	fmt.Println(contextF.Value("c"))	//Value C
	fmt.Println(contextF.Value("b"))	//Value <nil> -> There's no "b" value in parent F
	fmt.Println(contextA.Value("b"))	//Value <nil> -> Cannot get value from parent to child
}