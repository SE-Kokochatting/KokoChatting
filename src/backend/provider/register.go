package provider

type RegisterProvider struct {
	mysqlProvider
}

// 操作数据库表函数


func NewRegisterProvider() *RegisterProvider {
	return &RegisterProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}