package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"

	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
)

func initDatabase() error {
	fmt.Println(os.Getenv("DB_USER"))

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_DATABASE"),
	))

	if err != nil {
		return err
	}

	defer db.Close()

	for i := 1; i < 10; i++ {
		q := "INSERT INTO `product` (name, description, price, rate, image) VALUES (?, ?, ?, ?, ?)"
		insert, err := db.Prepare(q)

		if err != nil {
			return err
		}

		product := product{
			Name:        faker.Word(),
			Description: faker.Paragraph(),
			Price:       10 + rand.Float32()*(100-10),
			Rate:        0 + rand.Float32()*(5-0),
			Image:       fmt.Sprintf("http://lorempixel.com/200/200?%s", faker.UUIDDigit()),
		}

		_, err = insert.Exec(product.Name, product.Description, product.Price, product.Rate, product.Image)

		if err != nil {
			return err
		}

	}

	return nil
}

func getAllData() ([]product, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_DATABASE"),
	))

	if err != nil {
		return nil, err
	}

	var products []product

	rows, err := db.Query("SELECT * FROM `product`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		var product product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Rate, &product.Image)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func getDataById(id int) ([]product, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_DATABASE"),
	))

	if err != nil {
		return nil, err
	}

	var products []product

	q := "SELECT * FROM `product` WHERE `id` = ? "
	rows, err := db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		var product product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Rate, &product.Image)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// Delete all people above `age`
func deleteDataById(id int) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_DATABASE"),
	))

	if err != nil {
		return err
	}

	q := "DELETE FROM `product` WHERE `id` = ?"
	drop, err := db.Prepare(q)

	if err != nil {
		return err
	}

	defer drop.Close()
	defer db.Close()

	_, err = drop.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// Update a persons `age` based on `name`
func updateProductById(id int, product product) ([]product, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_DATABASE"),
	))

	if err != nil {
		return nil, err
	}

	q := "UPDATE `product` SET `name` = ?, `description` = ?, `price` = ? WHERE `id` = ?"

	update, err := db.Prepare(q)

	if err != nil {
		return nil, err
	}

	_, err = update.Exec(product.Name, product.Description, product.Price, id)

	if err != nil {
		return nil, err
	}

	defer update.Close()
	defer db.Close()

	products, _ := getDataById(id)

	return products, nil

}
