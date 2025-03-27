package Library

import "fmt"

type Library struct {
	BooksPresent []*Book
	LibraryUsers []*User
}


func (lib *Library) AddBookToLib(b *Book) string{
	lib.BooksPresent = append(lib.BooksPresent, b)
	return "Book added in the library"
}

func (lib *Library) ShowBooksUsers() {
	fmt.Println("Books present in the library :-> ")
	for _, book := range lib.BooksPresent {
		fmt.Printf("Title %s \nAuthor %s\n ID %s, Borrowed %t \n",book.Title, book.Author, book.Id, book.IsBorrowed)
	}
	 fmt.Println("Users present in the library :-> ")
	for _, user := range lib.LibraryUsers {
		fmt.Printf("Name %s \nBorrowed Books %d\n", user.UserName, len(user.BorrowedBooks))
	}

}

func (lib *Library) HandleBookTransactions(id string, query string)string {
	for _, elem := range lib.BooksPresent {
		if elem.Id == id {
			if elem.IsBorrowed && query == "borrow"{
				return "Book is already borrowed"
			}else if !elem.IsBorrowed && query == "return"{
				return "Book is not currently borrowed"
			}
			elem.BookTransaction(query)
			return fmt.Sprintf("Book %s is now %s ", elem.Title, query)
		}
	}
	return "Invalid Book ID"
}
