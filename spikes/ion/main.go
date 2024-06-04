package main

import (
	"fmt"

	ion "github.com/amazon-ion/ion-go/ion"
)

type T struct {
	A string
	B struct {
		RenamedC int   `ion:"C"`
		D        []int `ion:",omitempty"`
	}
}

func main() {
	t := T{}

	err := ion.Unmarshal([]byte(`{A:"Ion!",B:{C:2,D:[3,4]}}`), &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	t.A = "Shiny!"
	t.B.RenamedC = 42
	t.B.D = []int{6, 9}

	text, err := ion.MarshalText(&t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- text:\n%s\n\n", string(text))

	binary, err := ion.MarshalBinary(&t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- binary:\n%X\n\n", binary)
}
