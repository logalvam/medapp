package login

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HistoryStruct struct {
	UserId      string `json:"userid"`
	LoginDate   string `json:"logindate"`
	LoginTime   string `json:"logintime"`
	LogoutTime  string `json:"logouttime"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}
type HistoryResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func LoginHistory(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Login Validation")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var HistoryRec HistoryStruct
		var resp LoginResp
		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			err := json.Unmarshal(body, &HistoryRec)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			}
			data := insertLoginHistory(HistoryRec)
			if err != nil {
				fmt.Fprintln(w, "error in get data marchal"+err.Error())
			} else {
				fmt.Fprintln(w, data)
			}
		}
	}
}

func insertLoginHistory(HistoryRec HistoryStruct) HistoryResp {
	var resp HistoryResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()
		lLoginHis := `insert into loganathan.medapp_login_history(login_date,login_time,logout_time,created_by,created_date,updated_by,updated_date,user_id)values(curdate(),Now(),null,?,curdate(),?,curdate(),?)`
		_, err := db.Query(lLoginHis, &HistoryRec.CreatedBy, &HistoryRec.UpdatedBy, &HistoryRec.UserId)
		if err != nil {
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			resp.Status = "S"
			fmt.Println(resp)
		}
	}
	return resp
}
