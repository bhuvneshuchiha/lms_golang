package Library

type Book struct {
	Id         string
	Title      string
	Author     string
	IsBorrowed bool
}

func CreateBook(id, title, author string) Book {
	bookCreated := Book{
		Id:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
	return bookCreated
}

func CheckBook(id string, bookStorage map[string]Book) bool {
	_, ok := bookStorage[id]
	if ok {
		return true
	}
	return false
}

func (b *Book) BookTransaction(transactionQuery string) string {
	if b.IsBorrowed == false && transactionQuery == "borrow" {
		b.IsBorrowed = true
		return "borrowed"
	}
	if b.IsBorrowed && transactionQuery == "return" {
		b.IsBorrowed = false
		return "returned"
	}
	return "Invalid transactionQuery"
}
