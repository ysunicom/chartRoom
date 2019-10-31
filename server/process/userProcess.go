package process

import (
	"encoding/json"
	"fmt"
	"go_dev/chartRoom/common/message"
	"go_dev/chartRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data,并直接反序列成loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) err=", err)
		return
	}
	//声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//声明一个loginResMes
	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请注册"
	}
	//序列化loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err=", err)
		return
	}
	//4.将data赋值给resMes.Data
	resMes.Data = string(data)
	//5.对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err=", err)
		return
	}
	//6.发送data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn, data) err=", err)
		return
	}
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message)(err error){
	return err
}
