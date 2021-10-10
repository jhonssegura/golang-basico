package main

import (
	"fmt"
	"go-basico-platzi/src/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)
}
