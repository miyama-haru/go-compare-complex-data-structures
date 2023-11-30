package cases

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "go-compare-complex-data-structures/domain"
)

func TestMakeBuilding(t *testing.T) {
    want := MakeBuilding()
    got  := makeBuildingFixture()

    if !assert.Equal(t, got.Name, want.Name) {
        t.Errorf("Invalid struct value: wanted name = %v but got %v", want.Name, got.Name)
    }

    if len(got.Floors) != len(want.Floors) {
        t.Errorf("Invalid floors count. Wanted %d but got %d", len(want.Floors), len(got.Floors))
    }

    for _, gotFloor := range got.Floors {
        for _, wantFloor := range want.Floors {
            if gotFloor.Number != wantFloor.Number {
                continue
            }

            if len(gotFloor.Apts) != len(wantFloor.Apts) {
                t.Errorf("Invalid Apts count for floor %d: wanted %d but got %d", gotFloor.Number, len(wantFloor.Apts), len(gotFloor.Apts))
            }

            for _, gotApt := range gotFloor.Apts {
                for _, wantApt := range wantFloor.Apts {
                    if gotApt.Door != wantApt.Door {
                        continue
                    }

                    if !assert.ElementsMatch(t, gotApt.People, wantApt.People) {
                        t.Errorf("Invalid Apt values: wanted %+v but got %+v", wantApt.People, gotApt.People)
                    }
                }
            }
        }
    }
}

func makeBuildingFixture() domain.Building {
    return domain.Building{
        Name: "G5",
        Floors: []domain.Floor{
            {
                Number: 1,
                Apts: []domain.Apt{
                    {
                        Door: 4,
                        People: []domain.Person{
                            {
                                Name: "Amelia",
                                Age: 42,
                            },
                            {
                                Name: "Luciano",
                                Age: 40,
                            },
                        },
                    },
                    {
                        Door: 3,
                        People: []domain.Person{
                            {
                                Name: "Lilian",
                                Age: 26,
                            },
                            {
                                Name: "Barbara",
                                Age: 25,
                            },
                        },
                    },
                    {
                        Door: 5,
                        People: []domain.Person{
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
                Number: 0,
                Apts: []domain.Apt{
                    {
                        Door: 1,
                        People: []domain.Person{
                            {
                                Name: "Pedro",
                                Age: 20,
                            },
                            {
                                Name: "Matheus",
                                Age: 7,
                            },
                            {
                                Name: "Ana",
                                Age: 23,
                            },
                        },
                    },
                    {
                        Door: 2,
                        People: []domain.Person{
                            {
                                Name: "Marcela",
                                Age: 26,
                            },
                            {
                                Name: "Miranda",
                                Age: 32,
                            },
                            {
                                Name: "Caio",
                                Age: 12,
                            },
                        },
                    },
                    {
                        Door: 0,
                        People: []domain.Person{
                            {
                                Name: "Haruka",
                                Age: 21,
                            },
                            {
                                Name: "Miyama",
                                Age: 21,
                            },
                            {
                                Name: "Danyela",
                                Age: 19,
                            },
                        },
                    },
                },
            },
            {
                Number: 2,
                Apts: []domain.Apt{
                    {
                        Door: 6,
                        People: []domain.Person{
                            {
                                Name: "Julia",
                                Age: 17,
                            },
                            {
                                Name: "Rosa",
                                Age: 53,
                            },
                        },
                    },
                    {
                        Door: 8,
                        People: []domain.Person{
                            {
                                Name: "Robe",
                                Age: 23,
                            },
                            {
                                Name: "Gwen",
                                Age: 19,
                            },
                            {
                                Name: "Miles",
                                Age: 17,
                            },
                        },
                    },
                    {
                        Door: 7,
                        People: []domain.Person{
                            {
                                Name: "Caique",
                                Age: 19,
                            },
                            {
                                Name: "Bianca",
                                Age: 18,
                            },
                            {
                                Name: "Marcia",
                                Age: 21,
                            },
                        },
                    },
                },
            },
        },
    }
}
