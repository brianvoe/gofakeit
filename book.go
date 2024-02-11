package gofakeit

func BookTitle() string { return bookTitle(GlobalFaker) }

func (f *Faker) BookTitle() string { return bookTitle(f) }

func bookTitle(f *Faker) string { return getRandValue(f, []string{"book", "title"}) }

func BookAuthor() string { return bookAuthor(GlobalFaker) }

func (f *Faker) BookAuthor() string { return bookAuthor(f) }

func bookAuthor(f *Faker) string { return getRandValue(f, []string{"book", "author"}) }

func BookGenre() string { return bookGenre(GlobalFaker) }

func (f *Faker) BookGenre() string { return bookGenre(f) }

func bookGenre(f *Faker) string { return getRandValue(f, []string{"book", "genre"}) }

type BookInfo struct {
	Title  string `json:"title" xml:"name"`
	Author string `json:"author" xml:"author"`
	Genre  string `json:"genre" xml:"genre"`
}

func Book() *BookInfo { return book(GlobalFaker) }

func (f *Faker) Book() *BookInfo { return book(f) }

func book(f *Faker) *BookInfo {
	return &BookInfo{
		Title:  bookTitle(f),
		Author: bookAuthor(f),
		Genre:  bookGenre(f),
	}
}

func addBookLookup() {
	AddFuncLookup("book", Info{
		Display:     "Book",
		Category:    "book",
		Description: "Written or printed work consisting of pages bound together, covering various subjects or stories",
		Example: `{
	"title": "Anna Karenina",
	"author": "Toni Morrison",
	"genre": "Thriller"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return book(f), nil
		},
	})

	AddFuncLookup("booktitle", Info{
		Display:     "Title",
		Category:    "book",
		Description: "The specific name given to a book",
		Example:     "Hamlet",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bookTitle(f), nil
		},
	})

	AddFuncLookup("bookauthor", Info{
		Display:     "Author",
		Category:    "book",
		Description: "The individual who wrote or created the content of a book",
		Example:     "Mark Twain",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bookAuthor(f), nil
		},
	})

	AddFuncLookup("bookgenre", Info{
		Display:     "Genre",
		Category:    "book",
		Description: "Category or type of book defined by its content, style, or form",
		Example:     "Adventure",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bookGenre(f), nil
		},
	})
}
