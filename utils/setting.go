package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查路径", err)
	}
	LoadServer(file)
	LoadDataBase(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	// Section 在ini文件内寻找 [name] 分区 , 在分区中寻找键 name , 若没有键值则默认为 defaultVal
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("28ufj23e")
}

func LoadDataBase(file *ini.File) {
	// Section 在ini文件内寻找 [name] 分区 , 在分区中寻找键 name , 若没有键值则默认为 defaultVal
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("server").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("xoS8YeMaACMgXABL15YBCFERJdeKRRVR_g4adxsy")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("yi1FH5W4LCRCfc1PuOvN2ioMbJau2Ji4O7cepkgS")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("spxzxblog")
	QiniuSever = file.Section("qiniu").Key("QiniuSever").MustString("http://rbjlfzm9o.hn-bkt.clouddn.com")

}
