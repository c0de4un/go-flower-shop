package main

import (
	"fmt"

	"github.com/c0de4un/go-flower-shop/internal/logging"
)

func main() {
	logging.InitializeLogger()

	fmt.Println("Hello World !")
}
