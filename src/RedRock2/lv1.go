package main

import (
	"fmt"

)

type Author struct {
	Name string				//名字
	VIP bool 				//是否是高贵的带会员
	Icon string 			//头像
	Signature string 		//签名
	Focus int 				//关注人数
}


func main() {
	ms := new(Author)
	ms.Name = "毕导THU"
	ms.VIP = true
	ms.Icon= "不会"
	ms.Signature="数理化狂热爱好者"
	ms.Focus=2960000
	fmt.Printf("The author name is: %s\n", ms.Name)
	fmt.Printf("vip: %t\n", ms.VIP)
	fmt.Printf("The icon is: %s\n", ms.Icon)
	fmt.Printf("The signature is: %s\n", ms.Signature)
	fmt.Printf("The icon is: %d\n", ms.Focus)

}