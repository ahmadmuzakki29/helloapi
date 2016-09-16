package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)
var db *sql.DB
func init(){
	dbcon, err := sql.Open("postgres", "postgres://tokopedia:456tokopedia@192.168.100.126/tokopedia-dev-db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = dbcon
}

func GetProduct() *sql.Rows{
	rows, err := db.Query("SELECT product_id,product_name FROM ws_product limit 100")
	if err != nil{
		log.Fatal(err)
	}
	return rows
}

func GetProductDetail() *sql.Rows{
	
}