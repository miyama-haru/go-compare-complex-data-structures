package cases

import (
    "testing"

    "go-compare-complex-data-structures/domain"
)

func TestMakeBookInventory(t *testing.T) {
    
}

func makeBookInventoryFixture() domain.BookInventory {
    return domain.BookInventory{
        ReferenceDate: "2023-11-30",
        TotalAvailable: 2547,
        TotalWorth: 51449.4,
        Books: []domain.Book{
            {
                Format: domain.AUDIOBOOK,
                Genre: domain.FANTASY,
                Data: []domain.BookData{
                    {
                        Title: "The dragon of the blue sword: a King's Field story",
                        DateOfRelease: "2007-03-16",
                        WordCount: 757033,
                        Authors: []domain.Author{
                            {
                                FirstName: "Naotoshi",
                                LastName: "Zin",
                                DoB: "1974-01-01",
                            },
                            {
                                FirstName: "Satoru",
                                LastName: "Yanagi",
                                DoB: "1974-01-01",
                            },
                        },
                    },
                },
            },
            {
                Format: domain.EBOOK,
                Genre: domain.EDUCATIONAL,
                Data: []domain.BookData{
                    {
                        Title: "Dark Souls II is the best Dark Souls",
                        DateOfRelease: "2022-11-29",
                        WordCount: 68837373,
                        Authors: []domain.Author{
                            {
                                FirstName: "Naotoshi",
                                LastName: "Zin",
                                DoB: "1974-01-01",
                            },
                            {
                                FirstName: "Hidetaka",
                                LastName: "Miyazaki",
                                DoB: "1974-01-01",
                            },
                        },
                    },
                    {
                        Title: "Complex data structures in Go",
                        DateOfRelease: "2023-11-29",
                        WordCount: 15300,
                        Authors: []domain.Author{
                            {
                                FirstName: "Haruka",
                                LastName: "Miyama",
                                DoB: "2000-01-01",
                            },
                        },
                    },
                },
            },
        },
    }
}
