package database

import (
	"fmt"
	"recything/app/config"

	admin "recything/features/admin/model"
	recybot "recything/features/recybot/model"
	report "recything/features/report/model"
	user "recything/features/user/model"
	article "recything/features/article/model"
	trashCategory "recything/features/trash_category/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitMigrationMysql(db)
	return db
}

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&admin.Admin{})
	db.AutoMigrate(&report.Report{}, &report.Image{})
	db.AutoMigrate(&recybot.Recybot{})
	db.AutoMigrate(&article.Article{})
	db.AutoMigrate(&trashCategory.TrashCategory{})

}
