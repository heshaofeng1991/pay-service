package pay

import (
	"pay-service/pay/base"
	_ "pay-service/pay/chinaums"
	_ "pay-service/pay/other_pay"
)

func GetAll() []base.PayInstance {
	pays := make([]base.PayInstance, 0)

	base.Lock.Lock()
	for index := range base.AllPay {
		pays = append(pays, base.AllPay[index])
	}
	base.Lock.Unlock()

	return pays
}

func GetCoinInstance(payID int64) (base.PayInstance, bool) {
	base.Lock.Lock()
	pay, ok := base.AllPay[payID]
	base.Lock.Unlock()

	return pay, ok
}
