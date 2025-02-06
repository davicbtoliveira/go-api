package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepoitory struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepoitory {
	return ProductRepoitory{
		connection: connection,
	}
}

func (pr *ProductRepoitory) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	} 

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
		return []model.Product{}, err
		}

		productList =  append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepoitory) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
	 "(product_name, price) " + 
	 " VALUES ($1, $2) RETURNING id")
	 if err != nil {
		 fmt.Println(err)
		 return 0, err
	 }

	 err = query.QueryRow(product.Name, product.Price).Scan(&id)
	 if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
} 

func (pr *ProductRepoitory) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * from product where id = $1")
	if (err != nil) {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if (err == sql.ErrNoRows) {
			return nil, nil
		}
	}

	query.Close()
	return &produto, nil
}