package test

import (
	"fmt"

	"github.com/google/uuid"
)

func Hello() {
	genCode, err := uuid.NewRandom()
	if err == nil {
		fmt.Println("Hello from Foo", genCode.String())
	}
}
