package provider


type ExampleProvider struct{
	mysqlProvider
}

func (prd *ExampleProvider) ExampleCRUD(args ...interface{}) error {
	prd.mysqlDb.Table("example").Where("dummy = ?",1)
	if prd.mysqlDb.Error != nil{
		// log
	}
	return prd.mysqlDb.Error
}

func NewExampleProvider()*ExampleProvider{
	return &ExampleProvider{
		mysqlProvider{},
	}
}