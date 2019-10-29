package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_dev/chartRoom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //缓存
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("等待客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn连接，就不会阻塞
	n, err := this.Conn.Read(this.Buf[:4]) //conn.Read将conn中的内容读取到buf[:4]这个切片中
	if n != 4 || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	var pkgLen uint32
	//binary.BigEndia.Uint32 将byte切片转换成uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen],&mes) err=", err)
		return
	}
	return

}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var bytes [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(len) err=", err)
		return
	}
	//fmt.Println("客户端，发送消息的长度ok...")
	//发送信息本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
	}

	return
}
