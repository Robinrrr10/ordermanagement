package dbUtils

import (
	"database/sql"
	"fmt"

	"../utils"
	_ "github.com/go-sql-driver/mysql"
)

func getDBUrl() string {
	fmt.Println("Connecting to mysql...")
	dbUrl := utils.DBuser + ":" + utils.DBpass + "@tcp(" + utils.DBhost + ":" + utils.DBport + ")/" + utils.DBname
	fmt.Println("DB URL is:" + dbUrl)
	return dbUrl
}

func Connect() {
	fmt.Println("Connecting to mysql...")
	dbUrl := utils.DBuser + ":" + utils.DBpass + "@tcp(" + utils.DBhost + ":" + utils.DBport + ")/" + utils.DBname
	//fmt.Println("DB URL is:" + dbUrl)
	db, err := sql.Open("mysql", dbUrl)
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

func Fetch(query string) *sql.Rows {
	fmt.Println("Query is:", query)
	db, err := sql.Open("mysql", getDBUrl())
	defer db.Close()
	if err != nil {
		fmt.Println("Error:", err)
	}
	result, err2 := db.Query(query)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	return result
}

func Update(query string) sql.Result {
	fmt.Println("Query is:", query)
	db, err := sql.Open("mysql", getDBUrl())
	defer db.Close()
	if err != nil {
		fmt.Println("Error:", err)
	}
	result, err2 := db.Exec(query)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	return result
}
