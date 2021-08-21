package main

import (
	"fmt"
	"github.com/RamVellanki/habittracker/router"
)

func main() {
	router.RegisterRouter()
	fmt.Println("Hello Golang!")
}
