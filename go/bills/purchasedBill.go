package bills

// this this code is used to append the purchased items to the font there displayed in preview

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PurchasedStruct struct {
	Billno       int     `json:"billno"`
	BillDate     string  `json:"billdate"`
	MedicineName string  `json:"medName"`
	Quantity     int     `json:"medQuantity"`
	Amount       float64 `json:"amount"`
}

type PurchasedItem struct {
	Item string `json:"item"`
}

type PurchasedResp struct {
	PurchasedArr []PurchasedStruct `json:"purchasedArr"`
	Status       string            `json:"status"`
	ErrMsg       string            `json:"errMeg"`
}

func CustomerBill(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Customer Bill")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		fmt.Fprintln(w, "Customer Bill")
		var resp PurchasedResp
		var item PurchasedItem

		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp.ErrMsg = err.Error()
			resp.Status = "S"
		} else {
			err := json.Unmarshal(body, &item)
			if err != nil {
				resp.ErrMsg = err.Error()
				resp.Status = "E"
			} else {
				CustomerBill, err := PurchasedItems(item)
				if err != nil {
					resp.ErrMsg = err.Error()
					resp.Status = "E"
				} else {
					data, err := json.Marshal(CustomerBill)
					if err != nil {
						resp.ErrMsg = err.Error()
						resp.Status = "E"
					} else {
						fmt.Println("purchased Items")
						fmt.Fprintln(w, string(data))
					}
				}
			}
		}

	}
}

func PurchasedItems(item PurchasedItem) (PurchasedResp, error) {
	var resp PurchasedResp
	var PurchasedRec PurchasedStruct
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		resp.Status = "S"
		resp.ErrMsg = err.Error()
	} else {
		defer db.Close()
		CustomerBill := `select mbm.bill_no ,mbm.bill_date ,mbd.medicine_name,mbd.quantity  , mbm.bill_amount  from  loganathan.medapp_bill_master mbm  join loganathan.medapp_bill_details mbd 
		on mbm.bill_no = mbd.bill_no		
		where  mbd.medicine_name  = ?;`
		rows, err := db.Query(CustomerBill, &item.Item)
		if err != nil {
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			for rows.Next() {
				err = rows.Scan(&PurchasedRec.Billno, &PurchasedRec.BillDate, &PurchasedRec.MedicineName, &PurchasedRec.Quantity, &PurchasedRec.Amount)
				if err != nil {
					resp.ErrMsg = err.Error()
					resp.Status = "E"
				} else {
					resp.PurchasedArr = append(resp.PurchasedArr, PurchasedRec)
				}
			}
		}
	}
	return resp, err

}
