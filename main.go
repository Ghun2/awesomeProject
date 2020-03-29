package main

import (
	"fmt"
	"github/Ghun2/awesomeProject/router"
)

func main() {
	fmt.Println("Hello ji hun")

	e := router.New()

	e.Logger.Fatal(e.Start(":7777"))
}
