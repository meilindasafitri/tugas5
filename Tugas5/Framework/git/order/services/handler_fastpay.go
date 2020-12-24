package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	cm "Tugas5/Framework/git/order/common"
	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) FastPayHandler(ctx context.Context, req cm.FastPayRequest) (res cm.FastPayResponse) {
	defer panicRecovery()
	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database
	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)
	db, err = sql.Open("mysql", mySQL)
	if err != nil {
		panic(err.Error())
	}

	var merchId = req.Merchant_ID
	var sign = req.Signature
	var fastPayResponse  = res
	var paymentChanel cm.PaymentChannel

	sql := `SELECT DISTINCT
			IFNULL(merchant_id,'') merchant_id,
			IFNULL(merchant,'') merchant
		FROM trans WHERE merchant_id = ? and signature = ?`
	result, err := db.Query(sql, merchId, sign)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&fastPayResponse.Merchant_ID, &fastPayResponse.Merchant)
		if err != nil {
			panic(err.Error())
		}
	
		sqlDetial := `SELECT
				IFNULL(code,'') code,
				IFNULL(name,'') name
			FROM trans WHERE merchant_id = ? and signature = ?`
		resultDetail, errDet := db.Query(sqlDetial, merchId, sign)
		defer resultDetail.Close()
		if errDet != nil {
			panic(err.Error())
		}
		for resultDetail.Next() {	
			err := resultDetail.Scan(&paymentChanel.PgCode, &paymentChanel.PgName)
			if err != nil {
				panic(err.Error())
			}
			fastPayResponse.PaymentChan = append(fastPayResponse.PaymentChan, paymentChanel)	
		}
	}

	if fastPayResponse.Merchant_ID != "" {
		res = fastPayResponse
		res.Response = "List Payment Channel"
		res.ResponseCode = strconv.Itoa(http.StatusOK)
		res.ResponseDesc = "Sukses ambil data"
	} else {
		res.Response = "Gagal"
		res.ResponseCode = strconv.Itoa(http.StatusNotFound)
		res.ResponseDesc = "Gagal ambil data"
	}
	return
}
