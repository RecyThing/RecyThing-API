package database

import (
	"fmt"
	"recything/app/config"

	achievement "recything/features/achievement/model"
	admin "recything/features/admin/model"
	article "recything/features/article/model"
	daily "recything/features/daily_point/model"
	droppoint "recything/features/drop-point/model"
	faq "recything/features/faq/model"
	mission "recything/features/mission/model"
	recybot "recything/features/recybot/model"
	report "recything/features/report/model"
	trashCategory "recything/features/trash_category/model"
	user "recything/features/user/model"
	voucher "recything/features/voucher/model"
	trashExchange "recything/features/trash_exchange/model"
	dropPoint "recything/features/drop-point/model"

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
	// InitMigrationMysql(db)
	return db
}

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&achievement.Achievement{})
	db.AutoMigrate(&user.Users{}, &user.UserDailyPoints{})
	db.AutoMigrate(&admin.Admin{})
	db.AutoMigrate(&report.Report{}, &report.Image{})
	db.AutoMigrate(&recybot.Recybot{})
	db.AutoMigrate(&faq.Faq{})
	db.AutoMigrate(&trashCategory.TrashCategory{})
	db.AutoMigrate(&voucher.Voucher{})
	db.AutoMigrate(&article.Article{})
	db.AutoMigrate(&droppoint.DropPoints{}, &droppoint.Schedules{})
	db.AutoMigrate(&daily.DailyPoint{})
	db.AutoMigrate(&dropPoint.DropPoints{}, &dropPoint.Schedules{})
	db.AutoMigrate(&trashExchange.TrashExchange{}, trashExchange.TrashExchangeDetail{})
	db.AutoMigrate(&mission.Mission{}, &mission.MissionStage{}, &mission.ClaimedMission{})
}
