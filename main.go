package main

import (
	"database/sql"
	db2 "go-hexa/adpters/db"
	"go-hexa/aplication"
)

func main() {
	db, _ := sql.Open("sqlite3", "./db.sqlite")

	sqliteProductRepository := db2.NewSqliteProductRepository(db)
	productService := aplication.NewProductService(sqliteProductRepository)

	productService.Create("Iphone dos novos", 50)

}
