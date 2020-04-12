package handlers

type BooksType struct {
	BookID  uint64 `db:"bookID" json:"id"`
	Title   string `db:"title" json:"title"`
	Authors string `db:"authors" json:"authors"`
}
