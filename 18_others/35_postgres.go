package others

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "productsdb"
)

var db *sql.DB

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
}

func InsertProduct(p Product) {
	res, err := db.Exec("INSERT INTO products (title, description, price) VALUES ($1, $2, $3)", p.Title, p.Description, p.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffecced, err := res.RowsAffected()
	fmt.Printf("Inserted %d rows\n", rowsAffecced)

}

func UpdateProduct(p Product) {
	res, err := db.Exec("UPDATE products SET title=$1, description=$2, price=$3 WHERE id=$4", p.Title, p.Description, p.Price, p.ID)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffecced, err := res.RowsAffected()
	fmt.Printf("Updated %d rows\n", rowsAffecced)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows")
			return
		} else {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("ID: %d, Title: %s, Description: %s, Price: %f\n", p.ID, p.Title, p.Description, p.Price)
		products = append(products, &p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range products {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Price: %.2f\n", p.ID, p.Title, p.Description, p.Price)
	}

}

func GetProductById(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM products WHERE id=$1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Fatal("No rows")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Println("Product title is", product)
	}
}

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	// product := Product{
	// 	Title:       "Product 2",
	// 	Description: "Description 2",
	// 	Price:       9999.99,
	// }

	// InsertProduct(product)
	// GetProducts()
	//GetProductById(2)

	updatedProduct := Product{
		ID:          3,
		Title:       "Product updated",
		Description: "Description updated",
	}
	UpdateProduct(updatedProduct)
}
