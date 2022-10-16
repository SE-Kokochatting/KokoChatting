package provider

type RegisterProvider struct {
	mysqlProvider
}

func NewRegisterProvider() *RegisterProvider {
	return &RegisterProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}