package login

import (
	"encoding/json"
	"fmt"
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

type UserStruct struct {
	UserId string `json:"userid"`
}

func LoginHistory(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login History Insert")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var HistoryRec HistoryStruct
		var UserRec UserStruct
		var resp LoginResp
		resp.Status = "S"
		User := r.Header.Get("USER")
		fmt.Println("****")
		fmt.Println(User)
		UserRec.UserId = User
		HistoryRec.UserId = User

		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	resp.Status = "E"
		// 	resp.ErrMsg = err.Error()
		// 	fmt.Fprintln(w, resp)
		// 	fmt.Println(resp)
		// } else {
		// err := json.Unmarshal(body, &HistoryRec)
		// if err != nil {
		// 	resp.Status = "E"
		// 	resp.ErrMsg = err.Error()
		// 	fmt.Fprintln(w, resp)
		// 	fmt.Println(resp)
		// } else {
		Result, err := insertLoginHistory(User)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			} else {
				fmt.Fprintln(w, string(data))

			}
		}
		// }

	}
	// }
}

func insertLoginHistory(User string) (HistoryResp, error) {
	var resp HistoryResp
	// var UserRec UserStruct

	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()
		lLoginHis := `insert into loganathan.medapp_login_history(login_date,login_time,logout_time,created_by,created_date,updated_by,updated_date,user_id)values(curdate(),Now(),null,?,curdate(),?,curdate(),?)`
		_, err := db.Query(lLoginHis, User, User, User)
		if err != nil {
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			resp.Status = "S"
			fmt.Println(resp)
		}
	}
	return resp, err
}
