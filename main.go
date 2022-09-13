package main

import (
	"fileMethods"
	"fmt"
	"matrix"
)

func main() {
	m := matrix.NewMatrix("matrix1")
	m.ParseMatrixSize(fileMethods.ReadFile("matrix1Size"))
	m.ParseMatrix(fileMethods.ReadFile("matrix1"))
	m.MultiplyMatrix()
	fmt.Println(m)
}
