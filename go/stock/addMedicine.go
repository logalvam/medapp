package stock

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type MedicineStruct struct {
	Medicinename string  `json:"medName"`
	Quantity     int     `json:"medQuantity"`
	Unitprice    float64 `json:"unitprice"`
	CreatedBy    string  `json:"createdBy"`
	CreatedDate  string  `json:"createdDate"`
	UpdatedBy    string  `json:"updatedBy"`
	UpdatedDate  string  `json:"updatedDate"`
}

type MedicineResp struct {
	MedicineArr []MedicineStruct `json:"medicineArr"`
	Status      string           `json:"status"`
	ErrMsg      string           `json:"errMeg"`
}

func AddMedicine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insert Medicine Name")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "PUT" {
		var MedicineRec MedicineStruct
		var resp MedicineResp
		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			err := json.Unmarshal(body, &MedicineRec)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
			}
			MedicineName, err := InsertMedicine(MedicineRec)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Println(resp)
				fmt.Fprintln(w, resp)
			} else {
				data, err := json.Marshal(MedicineName)
				if err != nil {
					resp.Status = "E"
					resp.ErrMsg = err.Error()
					fmt.Println(resp)
					fmt.Fprintln(w, resp)
				} else {
					fmt.Println("Stock is inserted")
					fmt.Fprintln(w, "Success")
					fmt.Fprintln(w, string(data))

				}

			}

		}
	}
}

func InsertMedicine(MedicineRec MedicineStruct) (MedicineResp, error) {
	var resp MedicineResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()
		medicineInsert := `insert into loganathan.medapp_stock(medicine_name,quantity,unit_price,created_by,created_date,updated_by,updated_date)values(?,?,?,?,?,?,?)`
		_, err := db.Exec(medicineInsert, &MedicineRec.Medicinename, &MedicineRec.Quantity, &MedicineRec.Unitprice, &MedicineRec.CreatedBy, &MedicineRec.CreatedDate, &MedicineRec.UpdatedBy, &MedicineRec.UpdatedDate)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			resp.MedicineArr = append(resp.MedicineArr, MedicineRec)
			resp.ErrMsg = "Success"
			resp.Status = "S"
		}

	}
	return resp, err
}
