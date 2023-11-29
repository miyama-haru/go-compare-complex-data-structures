package main

import (
    "fmt"
)

type Building struct {
    Name string
    Floors []Floor
}

type Floor struct {
    Number int
    Apts []Apt
}

type Apt struct {
    Door int
    People []Person
}

type Person struct {
    Name string
    Age int
}

func makeBuilding() Building {
    return Building{
        Name: "G5",
        Floors: []Floor{
            {
                Number: 0,
                Apts: []Apt{
                    {
                        Door: 0,
                        People: []Person{
                            {
                                Name: "Miyama",
                                Age: 21,
                            },
                            {
                                Name: "Danyela",
                                Age: 19,
                            },
                            {
                                Name: "Haruka",
                                Age: 21,
                            },
                        },
                    },
                    {
                        Door: 1,
                        People: []Person{
                            {
                                Name: "Pedro",
                                Age: 20,
                            },
                            {
                                Name: "Ana",
                                Age: 23,
                            },
                            {
                                Name: "Matheus",
                                Age: 7,
                            },
                        },
                    },
                    {
                        Door: 2,
                        People: []Person{
                            {
                                Name: "Miranda",
                                Age: 32,
                            },
                            {
                                Name: "Marcela",
                                Age: 26,
                            },
                            {
                                Name: "Caio",
                                Age: 12,
                            },
                        },
                    },
                },
            },
            {
                Number: 1,
                Apts: []Apt{
                    {
                        Door: 3,
                        People: []Person{
                            {
                                Name: "Barbara",
                                Age: 25,
                            },
                            {
                                Name: "Lilian",
                                Age: 26,
                            },
                        },
                    },
                    {
                        Door: 4,
                        People: []Person{
                            {
                                Name: "Luciano",
                                Age: 40,
                            },
                            {
                                Name: "Amelia",
                                Age: 42,
                            },
                        },
                    },
                    {
                        Door: 5,
                        People: []Person{
                            {
                                Name: "Vincent",
                                Age: 22,
                            },
                            {
                                Name: "Julian",
                                Age: 23,
                            },
                        },
                    },
                },
            },
            {
                Number: 2,
                Apts: []Apt{
                    {
                        Door: 6,
                        People: []Person{
                            {
                                Name: "Rosa",
                                Age: 53,
                            },
                            {
                                Name: "Julia",
                                Age: 17,
                            },
                        },
                    },
                    {
                        Door: 7,
                        People: []Person{
                            {
                                Name: "Bianca",
                                Age: 18,
                            },
                            {
                                Name: "Caique",
                                Age: 19,
                            },
                            {
                                Name: "Marcia",
                                Age: 21,
                            },
                        },
                    },
                    {
                        Door: 8,
                        People: []Person{
                            {
                                Name: "Gwen",
                                Age: 19,
                            },
                            {
                                Name: "Miles",
                                Age: 17,
                            },
                            {
                                Name: "Robe",
                                Age: 23,
                            },
                        },
                    },
                },
            },
        },
    }
}

func main() {
    newBuilding := makeBuilding()

    fmt.Printf("newBuilding: \n%+v", newBuilding)
}
