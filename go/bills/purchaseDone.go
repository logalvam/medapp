package bills

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type BillDetailsStruct struct {
// 	BillNo       int     `json:"billno"`
// 	MedicineName string  `json:"medName"`
// 	Quantity     int     `json:"quantity"`
// 	UnitPrice    float64 `json:"unitprice"`
// 	Amount       float64 `json:"amount"`
// 	CreatedBy    string  `json:"createdBy"`
// 	CreatedDate  string  `json:"createdDate"`
// 	UpdatedBy    string  `json:"updatedBy"`
// 	UpdatedDate  string  `json:"updatedDate"`
// }

// type BillDetailsResp struct {
// 	Status string `json:"status"`
// 	ErrMsg string `json:"errMeg"`
// }

// type BillMasterStruct struct {
// 	BillNo      int     `json:"billno"`
// 	BillDate    string  `json:"billdate"`
// 	Amount      float64 `json:"amount"`
// 	BillGst     float64 `json:"billGst"`
// 	NetPrice    float64 `json:"netPrice"`
// 	UserId      string  `json:"userId"`
// 	CreatedBy   string  `json:"createdBy"`
// 	CreatedDate string  `json:"createdDate"`
// 	UpdatedBy   string  `json:"updatedBy"`
// 	UpdatedDate string  `json:"updatedDate"`
// }

// type BillMasterResp struct {
// 	Status string `json:"status"`
// 	ErrMsg string `json:"errMeg"`
// }

// type BillReq struct {
// 	MedicineName string `json:"medName"`
// 	Quantity     int    `json:"quantity"`
// }

// func BillMasterDetails(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Purchased Done")
// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
// 	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

// 	var BillMasRec BillMasterStruct
// 	var Mresp BillMasterResp
// 	var Dresp BillDetailsResp
// 	var BillRec BillReq
// 	Dresp.Status = "S"
// 	Mresp.Status = "S"
// 	if err != nil {
// 		Dresp.ErrMsg = err.Error()
// 		Dresp.Status = "E"
// 		Mresp.ErrMsg = err.Error()
// 		Mresp.Status = "E"
// 		fmt.Fprintln(w, Dresp)
// 		fmt.Println(Dresp)
// 		fmt.Fprintln(w, Mresp)
// 		fmt.Println(Mresp)
// 	} else {
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			Dresp.ErrMsg = err.Error()
// 			Dresp.Status = "E"
// 			Mresp.ErrMsg = err.Error()
// 			Mresp.Status = "E"
// 			fmt.Fprintln(w, Dresp)
// 			fmt.Println(Dresp)
// 			fmt.Fprintln(w, Mresp)
// 			fmt.Println(Mresp)
// 		} else {
// 			err := json.Unmarshal(body, &BillRec)
// 			if err != nil {
// 				Dresp.ErrMsg = err.Error()
// 				Dresp.Status = "E"
// 				Mresp.ErrMsg = err.Error()
// 				Mresp.Status = "E"
// 				fmt.Fprintln(w, Dresp)
// 				fmt.Println(Dresp)
// 				fmt.Fprintln(w, Mresp)
// 				fmt.Println(Mresp)
// 			} else {
// 				MasterResult ,err:= BillMaster(BillRec)
// 				DetailsResult ,err:= BillDetails(BillRec)
// 			}
// 		}
// 	}

// }

// func BillMaster(BillRec BillReq) (BillMasterStruct, error) {

// 	var BillMasRec BillMasterStruct
// 	var Mresp BillMasterResp
// 	Mresp.Status = "S"

// 	db, err := LocalDBConnect()
// 	if err != nil {
// 		Mresp.ErrMsg = err.Error()
// 		Mresp.Status = "E"
// 	} else {
// 		defer db.Close()

// 		lBillDetails := ``
// 	}

// 	return BillMasRec, err
// }
