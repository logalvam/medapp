package main

import (
	"fmt"
	"go/login"
	"go/stock"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Api Start (+)")
	http.HandleFunc("/insertMedicineBrand", stock.InsertBrand)
	http.HandleFunc("/insertMedicineName", stock.AddMedicine)
	http.HandleFunc("/loginValidation", login.LoginValidation)
	http.HandleFunc("/loginhistoryInsert", login.LoginHistory)
	http.ListenAndServe(":8260", nil)
}
