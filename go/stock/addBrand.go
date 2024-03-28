package stock

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func LocalDBConnect() (*sql.DB, error) {
	log.Println("Local DBConnect +")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "loganathan")
	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Println("open connection failed:", err.Error())
		return db, err
	}
	log.Println("local DBConnect -")
	return db, nil
}

type BrandStruct struct {
	Medicinename string `json:"medName"`
	Brand        string `json:"medBrand"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}

type BrandResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func InsertBrand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insert Medicine Name")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "PUT" {
		var BrandRec BrandStruct
		var resp BrandResp
		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			err := json.Unmarshal(body, &BrandRec)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
			}
			brand, err := MedicineBrand(BrandRec)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Println(resp)
				fmt.Fprintln(w, resp)
			}
			data, err := json.Marshal(brand)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Println(resp)
				fmt.Fprintln(w, resp)
			} else {
				fmt.Print("insert brand")
				fmt.Fprintln(w, string(data))
			}
		}
	}
}

func MedicineBrand(BrandRec BrandStruct) (BrandResp, error) {
	var resp BrandResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()
		brandinsert := `insert into loganathan.medapp_medicine_master(medicine_name,brand,created_by,created_date,updated_by,updated_date)values(?,?,?,?,?,?)`
		_, err := db.Exec(brandinsert, &BrandRec.Medicinename, &BrandRec.Brand, &BrandRec.CreatedBy, &BrandRec.CreatedDate, &BrandRec.UpdatedBy, &BrandRec.UpdatedDate)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			resp.ErrMsg = "Success"
			resp.Status = "S"
		}

	}
	return resp, err
}
