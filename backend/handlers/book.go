package handlers

import (
	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var db *sqlx.DB

func init() {
	db = initDB("books.db")
}

func initDB(filepath string) *sqlx.DB {
	//откроем файл или создадим его
	db, err := sqlx.Open("sqlite3", filepath)
	// проверяем ошибки и выходим при их наличии
	if err != nil {
		panic(err)
	}
	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}

	return db
}

func Books(c echo.Context) error {
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	var rBooks []BooksType
	rows, err := db.Queryx("SELECT bookID, title, authors FROM  books")
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	for rows.Next() {
		var b BooksType
		err = rows.StructScan(&b)
		if err != nil {
			return err
		}
		rBooks = append(rBooks, b)
	}

	return c.JSON(http.StatusOK, rBooks)

}
