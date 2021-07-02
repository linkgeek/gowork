package main

import (
	"fmt"
	"oop/model"
)

func main() {
	// user := model.userInfo{
	// 	Name:   "giant",
	// 	Age:    19,
	// 	Height: 173,
	// }

	//user := model.NewUserInfo("giant", 17, 178)
	//fmt.Print(user.Name, user.Age)

	// goods := &model.Product{}
	// goods.SetName("钻石")
	// goods.SetPrice(11)

	// fmt.Print(goods.GetName(), goods.GetPrice())

	wxpay := &model.WxPay{
		PaymentArgs: model.PaymentArgs{
			AppID: "wxappid",
			MchID: "wxappid",
			Token: "wxappid",
		},
		WxOpenID: "wxopenid",
	}

	fmt.Print(wxpay.AppID, "\n")

	alipay := &model.AliPay{
		PaymentArgs: model.PaymentArgs{
			AppID: "aliappid",
			MchID: "aliappid",
			Token: "aliappid",
		},
		AliOpenID: "aliopenid",
	}

	//fmt.Print(alipay.AppID, "\n")
	//fmt.Print(alipay.PaymentArgs.AppID, "\n")

	superpay := model.PaymentArgs{
		AppID: "superappid",
		MchID: "superappid",
		Token: "superappid",
	}

	superpay.Info()
	alipay.Info()
}
