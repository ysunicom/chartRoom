package main

import (
	"fmt"
	"go_dev/chartRoom/client/process"
	"os"
)

var userId int
var userPwd string

func main() {
	//用户选择
	var key int
	//是否继续显示菜单
	//var loop bool = true
	for true {
		fmt.Println("--------------------欢迎登陆多人聊天系统---------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出聊天室")
		fmt.Println("请选择(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId) //还要加个判断，如果用户输入的不是数字，要求重新输入
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败")
			}

		case 2:
			fmt.Println("注册用户")
			// loop = false
		case 3:
			fmt.Println("退出聊天室")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

}
