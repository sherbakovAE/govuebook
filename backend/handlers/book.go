package handlers

import (
	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
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
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	itemsPerPage, err := strconv.Atoi(c.QueryParam("itemsPerPage"))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	var rBooks []BooksType
	rows, err := db.Queryx("SELECT  bookID, title, authors FROM  books ORDER BY bookID LIMIT ? OFFSET ? ", itemsPerPage, itemsPerPage*(page-1))
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

	var count int
	err = db.Get(&count, "SELECT count(*) FROM books")
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	c.Response().Header().Set("X-Total-Count", strconv.Itoa(count))
	return c.JSON(http.StatusOK, rBooks)
}
