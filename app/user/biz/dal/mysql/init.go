package mysql

import (
	"fmt"
	"os"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	//测试使用
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//运行使用
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.User{})
		DB.AutoMigrate( // nolint: errcheck
			&model.User{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `user` (`id`,`created_at`,`updated_at`,`email`,`password_hashed`) VALUES (1,'2023-12-26 09:46:19.852','2023-12-26 09:46:19.852','123@admin.com','$2a$10$jTvUFh7Z8Kw0hLV8WrAws.PRQTeuH4gopJ7ZMoiFvwhhz5Vw.bj7C')")
		}
	}
}
