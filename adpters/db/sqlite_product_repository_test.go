package db_test

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/require"
	"go-hexa/adpters/db"
	"go-hexa/aplication"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `create table products ("id"string ,"name" string,"price" float ,"status" string);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	query := `insert into products (id ,name,price,status) values("uuid","product 1",30,"disabled");`

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()

}

func TestSqliteProductRepository_FindById(t *testing.T) {
	setUp()
	// comando que vai esperar tudo na funcao rodar para ser executado no final
	defer Db.Close()

	sqliteProductRepository := db.NewSqliteProductRepository(Db)

	product, err := sqliteProductRepository.FindById("uuid")

	require.Nil(t, err)
	fmt.Println(product)

	require.Equal(t, "product 1", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())

}

func TestSqliteProductRepository_Save(t *testing.T) {
	setUp()

	// comando que vai esperar tudo na funcao rodar para ser executado no final
	defer Db.Close()
	sqliteProductRepository := db.NewSqliteProductRepository(Db)

	product := aplication.NewProduct()
	product.Name = "product 1"
	product.Price = 30

	result, err := sqliteProductRepository.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, product.Status, result.GetStatus())

	product.Name = "product 1"
	product.Status = "enabled"
	product.Price = 35

	result, err = sqliteProductRepository.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, product.Status, result.GetStatus())

}
