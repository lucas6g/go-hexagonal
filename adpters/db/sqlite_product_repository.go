package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go-hexa/aplication"
)

type SqliteProductRepository struct {
	db *sql.DB //dependencia da conexao com o banco
}

// funcao que retorna a instancia do banco
func NewSqliteProductRepository(db *sql.DB) *SqliteProductRepository {
	return &SqliteProductRepository{db: db}
}

func (s *SqliteProductRepository) FindById(id string) (aplication.ProductInterface, error) {
	var product aplication.Product
	smtp, err := s.db.Prepare("select id,name,price,status from products where id=?;")

	if err != nil {
		return nil, err
	}
	//setando os valores
	err = smtp.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	//retorna o ponteiro para o obejeto na memoria
	return &product, nil

}

func (s *SqliteProductRepository) insert(product aplication.ProductInterface) (aplication.ProductInterface, error) {

	smtp, err := s.db.Prepare("insert into products (id,name,price,status) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = smtp.Exec( //_ ignorando o result que o exec retorna
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}

	err = smtp.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *SqliteProductRepository) update(product aplication.ProductInterface) (aplication.ProductInterface, error) {

	_, err := s.db.Exec("update  products set name= ?, price =?, status =? where id = ?; ", product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())

	if err != nil {
		return nil, err //pra retornar o obejeto em branco uso  nil e o erro
	}

	return product, nil //retornando o objeto que foi atualizado e branco para o erro
}

func (s *SqliteProductRepository) Save(product aplication.ProductInterface) (aplication.ProductInterface, error) {
	var rows int
	s.db.QueryRow("select id from products where id = ?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := s.insert(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := s.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
