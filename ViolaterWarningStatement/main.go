package main

import "fmt"

func main() {
	appName := "卓組くん"
	contactMail := "kyota.HelloWorld@gmail.com"
	contactNum := "080-2028-3797"

	fmt.Printf(
		"不正を検知しました。アプリケーションの譲渡、販売、複製は禁止されています。\n"+
			"利用権限が剥奪されたため、現在「%s」の利用はできません。\n"+
			"\n"+
			"心当たりのない場合はお手数ですが、下記連絡先までご連絡をお願いいたします。\n"+
			"Tel:  %s\nMail: %s\n",
		appName, contactMail, contactNum)
}
