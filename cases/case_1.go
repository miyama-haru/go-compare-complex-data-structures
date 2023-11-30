package cases

import "go-compare-complex-data-structures/domain"

func MakeBuilding() domain.Building {
    return domain.Building{
        Name: "G5",
        Floors: []domain.Floor{
            {
                Number: 0,
                Apts: []domain.Apt{
                    {
                        Door: 0,
                        People: []domain.Person{
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
                        People: []domain.Person{
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
                        People: []domain.Person{
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
                Apts: []domain.Apt{
                    {
                        Door: 3,
                        People: []domain.Person{
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
                        People: []domain.Person{
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
                Number: 2,
                Apts: []domain.Apt{
                    {
                        Door: 6,
                        People: []domain.Person{
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
                        People: []domain.Person{
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
                        People: []domain.Person{
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
