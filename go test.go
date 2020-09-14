package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Liste struct {
	A,B,C string
}
func main() {
	rand.Seed(time.Now().UnixNano())
	nombre :=  rand.Intn(100)

	fmt.Println(nombre)

}
