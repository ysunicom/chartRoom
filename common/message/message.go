package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}
type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回的状态确 500:用户未注册 200:用户登陆成功
	Error string `json:"error"` //返回的错误信息
}

type RegisterMes struct {
	//
}
