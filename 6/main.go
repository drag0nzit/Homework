package main

import (
	"fmt"
	"sort"
	"strings"
)

type Book struct {
	ID        int
	Title     string
	Author    string
	Year      int
	Available bool
	Borrowed  int
}

type Reader struct {
	ID    int
	Name  string
	Books []int
}

type Library struct {
	Books    []Book
	Readers  []Reader
	MaxBooks int
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) DeleteBook(id int) {
	for i, b := range l.Books {
		if b.ID == id {
			if !b.Available {
				fmt.Println("Книга выдана")
				return
			}

			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return
		}
	}
}

func (l *Library) FindBook(id int) *Book {
	for i := range l.Books {
		if l.Books[i].ID == id {
			return &l.Books[i]
		}
	}
	return nil
}

func (l *Library) SearchBook(text string) {
	for _, b := range l.Books {
		if strings.Contains(strings.ToLower(b.Title), strings.ToLower(text)) ||
			strings.Contains(strings.ToLower(b.Author), strings.ToLower(text)) {
			fmt.Println(b.ID, b.Title, b.Author, b.Year)
		}
	}
}

func (l *Library) AddReader(reader Reader) {
	l.Readers = append(l.Readers, reader)
}

func (l *Library) FindReader(id int) *Reader {
	for i := range l.Readers {
		if l.Readers[i].ID == id {
			return &l.Readers[i]
		}
	}
	return nil
}

func (l *Library) GiveBook(readerID, bookID int) {
	reader := l.FindReader(readerID)
	book := l.FindBook(bookID)

	if reader == nil || book == nil {
		fmt.Println("Читатель или книга не найдены")
		return
	}

	if !book.Available {
		fmt.Println("Книга уже выдана")
		return
	}

	if len(reader.Books) >= l.MaxBooks {
		fmt.Println("Достигнут лимит книг")
		return
	}

	book.Available = false
	book.Borrowed++

	reader.Books = append(reader.Books, bookID)

	fmt.Println("Книга выдана")
}

func (l *Library) ReturnBook(readerID, bookID int) {
	reader := l.FindReader(readerID)
	book := l.FindBook(bookID)

	if reader == nil || book == nil {
		return
	}

	for i, id := range reader.Books {
		if id == bookID {
			reader.Books = append(
				reader.Books[:i],
				reader.Books[i+1:]...,
			)

			book.Available = true
			fmt.Println("Книга возвращена")
			return
		}
	}
}

func (l *Library) SortByTitle() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Title < l.Books[j].Title
	})
}

func (l *Library) SortByAuthor() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Author < l.Books[j].Author
	})
}

func (l *Library) SortByYear() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Year < l.Books[j].Year
	})
}

func (l *Library) Statistics() {
	var activeReader Reader
	maxBooks := -1

	for _, r := range l.Readers {
		if len(r.Books) > maxBooks {
			maxBooks = len(r.Books)
			activeReader = r
		}
	}

	var popular Book
	maxBorrowed := -1

	for _, b := range l.Books {
		if b.Borrowed > maxBorrowed {
			maxBorrowed = b.Borrowed
			popular = b
		}
	}

	if maxBooks >= 0 {
		fmt.Println("Самый активный читатель:", activeReader.Name)
	}

	if maxBorrowed > 0 {
		fmt.Println("Самая популярная книга:", popular.Title)
	} else {
		fmt.Println("Книги ещё не выдавались")
	}
}

func main() {
	lib := Library{MaxBooks: 3}

	lib.AddBook(Book{1, "Война и мир", "Толстой", 1869, true, 0})
	lib.AddBook(Book{2, "Преступление и наказание", "Достоевский", 1866, true, 0})
	lib.AddBook(Book{3, "Мастер и Маргарита", "Булгаков", 1967, true, 0})

	lib.AddReader(Reader{1, "Алексей", nil})
	lib.AddReader(Reader{2, "Иван", nil})

	lib.GiveBook(1, 1)
	lib.GiveBook(1, 2)
	lib.GiveBook(2, 1)

	fmt.Println("\nПоиск:")
	lib.SearchBook("Булгаков")

	fmt.Println("\nСтатистика:")
	lib.Statistics()

	fmt.Println("\nВозврат книги:")
	lib.ReturnBook(1, 1)
}
