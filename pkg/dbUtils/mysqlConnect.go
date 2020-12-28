package dbUtils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	fmt.Println("Connecting to mysql...")
	db, err := sql.Open("mysql", "apper:app123@tcp(192.168.222.62:3306)/business")
	defer db.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	query := "select id, storeId, productId, productName, eachPrice, quantity, totalPrice from orders where storeId=101;"
	result, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var id int
		var stId int
		var prId int
		var prdName string
		var ehPrc float32
		var qty int
		var totPrc float32
		err = result.Scan(&id, &stId, &prId, &prdName, &ehPrc, &qty, &totPrc)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Println("query result\n", id, stId, prId, prdName, ehPrc, qty, totPrc)
	}
}
