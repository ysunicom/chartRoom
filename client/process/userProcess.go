package process

import (
	"encoding/json"
	"fmt"
	"go_dev/chartRoom/client/utils"
	"go_dev/chartRoom/common/message"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//1.连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mess message.Message
	mess.Type = message.LoginMesType
	//3.创建一个LoginMes 结构构
	loginMes := message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data 赋给mess.Data字段
	mess.Data = string(data)
	//6.将mes进行序列化
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7.这时候data就是我们要发送的消息
	//7.1先把data的长度发送给server
	//先获取data的长度->转成一个表示长度的byte切片
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	//读取服务器发回来的响应消息
	mess, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mess.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登陆成功")
		//1.显示我们的登录成功的菜单
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}

	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
