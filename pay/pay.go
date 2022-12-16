package pay

import (
	. "airmart_pay/pay/base"
	_ "airmart_pay/pay/whole_pay"
)

func GetAll() []PayInstance {
	coins := make([]PayInstance, 0)
	Lock.Lock()
	for index := range AllPay {
		coins = append(coins, AllPay[index])
	}
	Lock.Unlock()
	return coins
}

func GetCoinInstance(payID int64) (PayInstance, bool) {
	Lock.Lock()
	coin, ok := AllPay[payID]
	Lock.Unlock()
	return coin, ok
}
