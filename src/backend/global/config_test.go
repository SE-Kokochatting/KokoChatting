package global

import (
	"testing"
)

// 本测试函数测试不通过是正常的，因为在测试函数运行和main.go运行时，config文件夹的相对路径不同
func TestConfig_GetConfigByPath(t *testing.T) {
	host,err := GetGlobalConfig().GetConfigByPath("mysql.host")
	if err != nil{
		t.Errorf(err.Error())
		t.Fail()
	}
	if host != "localhost"{
		t.Errorf("get mysql host config error")
	}

}