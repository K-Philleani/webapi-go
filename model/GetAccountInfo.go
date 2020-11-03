package model

import "webapi-go/database"

type Account struct {
	Id 			int 	`json:"id"`
	UserAccount string 	`json:"user_account"`
	UserPwd     string  `json:"user_pwd"`
	UserPhone  string   `json:"user_phone"`
}

func (a *Account) GetAccountInfo() (rowList []Account, err error){
	sqlStr := "select id, user_account, user_pwd, user_phone from account where id > 0"
	rows, err := database.DB.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var account Account
		err = rows.Scan(&account.Id, &account.UserAccount, &account.UserPwd, &account.UserPhone)
		if err != nil {
			return
		}
		rowList = append(rowList, account)
	}
	return
}
