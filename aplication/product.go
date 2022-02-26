package aplication

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/satori/go.uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float32
}

// ProductServiceInterface  application service
type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float32) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	FindById(id string) (ProductInterface, error)
}
type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductRepository interface {
	ProductReader
	ProductWriter
}

const DISABLED = "disabled"
const ENABLED = "enabled"

type Product struct {
	//tags sao anotations `` = anotations
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float32 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

//convesao no go criar a instancia do obejeto dentro da classe

func NewProduct() *Product {
	product := Product{
		Id:     uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product //retornado o ponteiro do obejeto
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price most be big then zero")
}
func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price most be equal to zero")
}
func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED

	}

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("status most be enebled or desable")
	}

	if p.Price < 0 {
		return false, errors.New("the price most be 0 or big than zero")
	}

	_, err := govalidator.ValidateStruct(p) //retorna os campos que estao dando erro
	if err != nil {
		return false, err
	}
	return true, nil //returnar nil quando nao tem erro
}

func (p *Product) GetId() string {
	return p.Id

}
func (p *Product) GetStatus() string {
	return p.Status
}
func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) GetName() string {
	return p.Name
}
