package domain

// Book formats here considered
const (
    EBOOK     = "digital"
    PHYSICAL  = "paper"
    AUDIOBOOK = "audio"
)

// Genres of books here considered
const (
    EDUCATIONAL = "edu"
    SCI_FI      = "scifi"
    FANTASY     = "fant"
)

type BookInventory struct {
    ReferenceDate  string
    TotalAvailable int
    TotalWorth     float32
    Books          []Book
}

type Book struct {
    Format string
    Genre  string
    Data   []BookData
}

type BookData struct {
    Title         string
    DateOfRelease string
    WordCount     int
    Authors       []Author
}

type Author struct {
    FirstName string
    LastName  string
    DoB       string
}
