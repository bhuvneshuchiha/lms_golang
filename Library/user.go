package Library

type User struct {
	UserID        string
	UserName      string
	BorrowedBooks []*Book
}

func CreateUser(id, name string) User {
	user := User{
		UserID:        id,
		UserName:      name,
		BorrowedBooks: []*Book{},
	}
	return user
}

type BookActions interface {
	BorrowBook(id string, bookStorage map[string]*Book) bool
	ReturnBook(id string, bookStorage map[string]*Book) bool
}

func (u *User) BorrowBook(id string, bookStorage map[string]*Book) bool {
	borrowedBook, ok := bookStorage[id]
	if !ok || borrowedBook.IsBorrowed {
		return false
	}
	borrowedBook.BookTransaction("borrow")
	u.BorrowedBooks = append(u.BorrowedBooks, borrowedBook)
	return true
}

func (u *User) ReturnBook(id string, bookStorage map[string]*Book) bool {
	book, ok := bookStorage[id]
	if !ok || !book.IsBorrowed {
		return false
	}
	book.BookTransaction("return")

	for index, elem := range u.BorrowedBooks {
		if elem.Id == id {
			u.BorrowedBooks = append(u.BorrowedBooks[:index], u.BorrowedBooks[index+1:]...)
			return true
		}
	}
	return false
}
