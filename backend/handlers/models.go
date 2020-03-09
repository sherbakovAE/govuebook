package handlers

type BooksType struct {
	BookID  uint64 `db:"bookID"`
	Title   string `db:"title"`
	Authors string `db:"authors"`
}
