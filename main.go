package main

import (
    "fmt"
    "github.com/jenkins_golang_test/operator"
)

func main() {
    fmt.Println("Hello World")
    a := 1.0
    b := 2.0
    sum, _ := operator.Add(a, b)
    fmt.Printf("a + b = %f", sum)
}
