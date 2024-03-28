package main

import (
	"fmt"
	"go/bills"
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
	http.HandleFunc("/viewMedicine", stock.ViewMedicine)
	http.HandleFunc("/customerBill", bills.CustomerBill)
	http.HandleFunc("/addNewUser", login.AddNewUser)
	http.ListenAndServe(":8260", nil)
}
