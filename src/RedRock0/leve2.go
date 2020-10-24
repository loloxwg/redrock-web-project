package main

import "fmt"

func main()  {
	var name string
	var pwd string
	fmt.Printf("请输入账号:")
	fmt.Scanln(&name)
	fmt.Printf("请输入密码：")
	fmt.Scanln(&pwd)
	if name=="xwg"&&pwd=="123456"{
		fmt.Print("登录成功")
	}else {
		fmt.Print("密码错误")
	}

}
