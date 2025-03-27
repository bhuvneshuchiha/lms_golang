package main

import (
	"fmt"

	"www.github.com/bhuvneshuchiha/lms/Library"
)

var lib Library.Library = Library.Library{
	BooksPresent: make([]*Library.Book, 0),
	LibraryUsers: make([]*Library.User, 0),
}

func main() {

	book1 := Library.Book{
		Id: "1",
		Author: "James Huan",
		Title: "That's my boy",
		IsBorrowed: false,
	}

	book2 := Library.Book{
		Id: "2",
		Author: "Hank",
		Title: "Black Cops",
		IsBorrowed: false,
	}
	book3 := Library.Book{
		Id: "3",
		Author: "Jenkins",
		Title: "Hello Brother",
		IsBorrowed: false,
	}

	user1 := Library.User{
		UserID: "1",
		UserName: "Mark",
		BorrowedBooks: []*Library.Book{},
	}
	user2 := Library.User{
		UserID: "2",
		UserName: "Jane",
		BorrowedBooks: []*Library.Book{},
	}

	lib.BooksPresent = append(lib.BooksPresent, &book1)
	lib.BooksPresent = append(lib.BooksPresent, &book2)
	lib.BooksPresent = append(lib.BooksPresent, &book3)

	bookStorage := make(map[string]*Library.Book)
	for _, elem := range lib.BooksPresent{
		bookStorage[elem.Id] = elem
	}

	lib.LibraryUsers = append(lib.LibraryUsers, &user1)
	lib.LibraryUsers = append(lib.LibraryUsers, &user2)

	fmt.Println("I am here to borrow a book")

	booksPres := lib.BooksPresent
	for _, elem := range booksPres {
		fmt.Printf("Author : %s,\nTitle : %s,\nIsBorrowed : %t\n", elem.Author, elem.Title, elem.IsBorrowed)
	}

	user2.BorrowBook("1", bookStorage)
	user2BorrowedBooks := user2.BorrowedBooks
	fmt.Println("List of books borrowed by user 2")
	for k,v := range user2BorrowedBooks{
		fmt.Printf("Book:- %d , Author : %s , Title : %s\n", k+1, v.Author, v.Title)
	}
	lib.ShowBooksUsers()

	status := lib.HandleBookTransactions("1", "return")
	fmt.Printf("Status :-> %s ", status)
}
