package document

import (
	"github.com/kataras/golog"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var globalDB *gorm.DB = nil
var globalESClient *elastic.Client

const (
	productMysqlUrl = "ad_service_rw:53f7MrwHDBKgG6l@tcp(10.18.34.229:13306)/ad_service?charset=utf8mb4&parseTime=True&loc=Local"
	testMysqlUrl    = "ad_service_rw:m03663LR04K1Ax5@tcp(10.18.54.157:3306)/ad_service?charset=utf8mb4&parseTime=True&loc=Local"
	productESUrl    = "http://10.19.37.12:9200"
	testESUrl       = "http://10.19.15.24:9200"
)

var inited = false
var mysqlUrl = ""
var esUrl = ""

func InitEnv(env int) {
	switch env {
	case 1:
		mysqlUrl = testMysqlUrl
		esUrl = testESUrl
		inited = true
	case 2:
		mysqlUrl = productMysqlUrl
		esUrl = productESUrl
		inited = true
	}

}

func initDB() {
	dsn := mysqlUrl
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	globalDB = db
}

func GetDB() *gorm.DB {
	if !inited {
		panic("环境未被初始化")
	}
	if globalDB == nil {
		initDB()
		//panic("database init error,please check datastore\\service\\base_service.go")
	}
	return globalDB
}

func GetESClient() *elastic.Client {
	if !inited {
		panic("环境未被初始化")
	}
	if globalESClient == nil {
		client, err := elastic.NewClient(elastic.SetURL(esUrl))
		if err != nil {
			golog.Errorf("connect es error:%v", err)
			panic(err)
		}
		globalESClient = client
		//panic("database init error,please check datastore\\service\\base_service.go")
	}
	return globalESClient
}
