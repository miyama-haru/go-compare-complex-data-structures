package main

import (
    "fmt"

    "go-compare-complex-data-structures/cases"
)

func main() {
    newBuilding := cases.MakeBuilding()

    fmt.Printf("newBuilding: \n%+v", newBuilding)
}
