package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func soma(a, b int) (int, bool) {
	if a > b {
		return a + b, true
	}
	return a + b, false
}

type ProductOld struct {
	Name	string
	Description string
	PurchasePrice float64
}

func (p ProductOld) getSalesPrice() float64 {
	p.Name = "Wesley"
	fmt.Println("nome de p:", p.Name);
	return p.PurchasePrice * 2
}

func (p *ProductOld) getSalesPrice2() float64 { // * ponteiro
	p.Name = "Wesley"
	return p.PurchasePrice * 2
}

func mainOld() {
	resultado, check := soma(1, 2)
	fmt.Println(resultado, check)

	product1 := ProductOld {
		Name: "laptop",
		Description: "notebook vaio",
		PurchasePrice: 1299.99,
	}
	fmt.Println(product1.getSalesPrice())
	fmt.Println(product1.Name)
	fmt.Println(product1.getSalesPrice2())
	fmt.Println(product1.Name)

	http.HandleFunc("/", hello)
	http.HandleFunc("/products", createProductOld)
	http.ListenAndServe(":9191", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func createProductOld(w http.ResponseWriter, r *http.Request) {
	product := ProductOld{Name: "iPhone", Description: "smartphone apple", PurchasePrice: 1599.99}
	persistProductOld(product)
	w.Write([]byte("Created product"))
}

func persistProductOld(product ProductOld) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("Insert into products(name, description, purchase_price) values($1, $2, $3)")
	if err != nil {
		panic(err)
	}
	stmt.Exec(product.Name, product.Description, product.PurchasePrice)
}

// novo main com echo //

type Product struct {
	Name	string `json:"nome"`
	Price float64 `json:"preco"`
}

var products []Product

func generateProducts() {
	p1 := Product{Name: "Product 1", Price: 1.99}
	p2 := Product{Name: "Product 2", Price: 2.99}
	products = append(products, p1, p2);
}

func main() {
	generateProducts()
	e := echo.New()
	e.GET("/products", listProducts)
	e.POST("/products", createProduct)
	e.Logger.Fatal(e.Start(":9191"))
}

func listProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func createProduct(c echo.Context) error {
	product := Product{}
	c.Bind(&product)
	err := persistProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, product)
}

func persistProduct(product Product) error {
	db, err := sql.Open("sqlite3", "test2.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("Insert into products(name, price) values($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}
