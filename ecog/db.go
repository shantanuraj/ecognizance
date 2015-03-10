package main

import (
	_ "fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

var (
	EcogDb = SetupDB()
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=eco sslmode=disable")
	PanicIf(err)
	return db
}

func SaveToDb(edition *Edition) {
	saveEdition(edition)
	saveEditorial(edition.Title, edition.Editorial)
	saveFeatured(edition.Title, edition.Featured)
	for i, article := range edition.Articles {
		saveArticle(edition.Title, i, article)
	}
	for i, ecohustle := range edition.Ecohustle {
		saveEcohustle(edition.Title, i, ecohustle)
	}
}

func saveEdition(edition *Edition) {
	rows, err := EcogDb.Query(`insert into
							editions (title, cover, articles, ecohustles)
							values ($1, $2, $3, $4)`,
		edition.Title, edition.Cover, len(edition.Articles), len(edition.Ecohustle))
	PanicIf(err)
	defer rows.Close()
}

func saveEditorial(edition string, e Article) {
	rows, err := EcogDb.Query(`insert into
							editorial (edition, title, author, email, post, gravatar)
							values ($1, $2, $3, $4, $5, $6)`,
		edition, e.Title, e.Author, e.Email, e.Post, e.Gravatar)
	PanicIf(err)
	defer rows.Close()
}

func saveFeatured(edition string, f Article) {
	rows, err := EcogDb.Query(`insert into
							featured (edition, title, author, email, post, gravatar)
							values ($1, $2, $3, $4, $5, $6)`,
		edition, f.Title, f.Author, f.Email, f.Post, f.Gravatar)
	PanicIf(err)
	defer rows.Close()
}

func saveArticle(edition string, count int, a Article) {
	rows, err := EcogDb.Query(`insert into
							articles (edition, count, title, author, email, post, gravatar)
							values ($1, $2, $3, $4, $5, $6, $7)`,
		edition, count, a.Title, a.Author, a.Email, a.Post, a.Gravatar)
	PanicIf(err)
	defer rows.Close()
}

func saveEcohustle(edition string, count int, e Ecohustle) {
	rows, err := EcogDb.Query(`insert into
							ecohustles (edition, count, image, caption)
							values ($1, $2, $3, $4)`,
		edition, count, e.Image, e.Caption)
	PanicIf(err)
	defer rows.Close()
}

func GetEditionByTitle(title string) *Edition {
	e := new(Edition)
	articles, ecohustles := new(int), new(int)

	rows, err := EcogDb.Query(`select title, cover, articles, ecohustles
							from editions where title = $1`, title)
	PanicIf(err)
	defer rows.Close()
	if rows.Next() {
		PanicIf(rows.Err())
		err := rows.Scan(&e.Title, &e.Cover, articles, ecohustles)
		PanicIf(err)
	}

	getEditorial(e)
	getFeatured(e)
	getArticles(e, *articles)
	getEcohustle(e, *ecohustles)

	return e
}

func getEditorial(e *Edition) {
	rows, err := EcogDb.Query(`select title, author, email, post, gravatar
							from editorial where edition = $1`, e.Title)
	PanicIf(err)
	defer rows.Close()

	e.Editorial = Article{}

	if rows.Next() {
		PanicIf(rows.Err())
		err := rows.Scan(&e.Editorial.Title, &e.Editorial.Author, &e.Editorial.Email, &e.Editorial.Post,
			&e.Editorial.Gravatar)
		PanicIf(err)
		e.Editorial.PostHTML = GetHTML(e.Editorial.Post)
	}

}

func getFeatured(e *Edition) {
	rows, err := EcogDb.Query(`select title, author, email, post, gravatar
							from featured where edition = $1`, e.Title)
	PanicIf(err)
	defer rows.Close()

	e.Featured = Article{}

	if rows.Next() {
		PanicIf(rows.Err())
		err := rows.Scan(&e.Featured.Title, &e.Featured.Author, &e.Featured.Email, &e.Featured.Post,
			&e.Featured.Gravatar)
		PanicIf(err)
		e.Featured.PostHTML = GetHTML(e.Featured.Post)
	}

}

func getArticles(e *Edition, count int) {
	rows, err := EcogDb.Query(`select title, author, email, post, gravatar
							from articles where edition = $1`, e.Title)

	PanicIf(err)
	defer rows.Close()

	e.Articles = make([]Article, count)

	i := 0
	for rows.Next() {
		PanicIf(rows.Err())
		err := rows.Scan(&e.Articles[i].Title, &e.Articles[i].Author, &e.Articles[i].Email, &e.Articles[i].Post,
			&e.Articles[i].Gravatar)
		PanicIf(err)
		e.Articles[i].PostHTML = GetHTML(e.Articles[i].Post)
		i += 1
	}
}

func getEcohustle(e *Edition, count int) {
	rows, err := EcogDb.Query(`select image, caption
							from ecohustles where edition = $1`, e.Title)

	PanicIf(err)
	defer rows.Close()

	e.Ecohustle = make([]Ecohustle, count)

	i := 0
	for rows.Next() {
		PanicIf(rows.Err())
		err := rows.Scan(&e.Ecohustle[i].Image, &e.Ecohustle[i].Caption)
		i += 1
		PanicIf(err)
	}
}

func GetAllMagazines() []Edition {
	rows, err := EcogDb.Query(`select title, cover
							from editions`)

	PanicIf(err)
	defer rows.Close()

	editions := []Edition{}
	for rows.Next() {
		PanicIf(rows.Err())
		edition := Edition{}
		err := rows.Scan(&edition.Title, &edition.Cover)
		PanicIf(err)
		editions = append(editions, edition)
	}

	return editions
}
