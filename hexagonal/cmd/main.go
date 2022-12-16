package main

import (
	"fmt"

	"github.com/rashed091/architecture-design/hexagonal/internal/adapters/core/arithmetic"
)

func main(){
	arithAdapter := arithmetic.NewAdapter();
	result, err := arithAdapter.Addition(2, 3)
	if err != nil {
		fmt.Println("Found an error.")
	}
	fmt.Println(result)
}
