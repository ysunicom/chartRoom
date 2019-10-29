package main

import (
	"fmt"
	"go_dev/chartRoom/common/message"
	"go_dev/chartRoom/server/process"
	"go_dev/chartRoom/server/utils"
	"io"
	"net"
)

//先创建processer结构构

type Processer struct {
	Conn net.Conn
}

func (this *Processer) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录消息
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册消息
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processer) control() (err error) {
	//读客户端发送的信息
	for {
		//这里我们将读取数据封装成一个函数，
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
			//return
		}
		//fmt.Println("mes=", mes)
		err = this.ServerProcessMes(&mes)

		if err != nil {
			fmt.Println("serverProcessMes err=", err)
			return err
		}

	}
}
