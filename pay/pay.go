package pay

import (
	. "airmart_pay/pay/base"
	_ "airmart_pay/pay/chinaums"
)

func GetAll() []PayInstance {
	pays := make([]PayInstance, 0)

	Lock.Lock()
	for index := range AllPay {
		pays = append(pays, AllPay[index])
	}
	Lock.Unlock()

	return pays
}

func GetCoinInstance(payID int64) (PayInstance, bool) {
	Lock.Lock()
	pay, ok := AllPay[payID]
	Lock.Unlock()

	return pay, ok
}
