package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	fmt.Println("hello world !")
	uuidWithHypen := uuid.New()
	fmt.Println(uuidWithHypen)

}
