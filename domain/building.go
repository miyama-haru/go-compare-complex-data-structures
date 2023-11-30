package domain

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
