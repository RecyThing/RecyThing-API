package repository

import (
	"errors"
	"recything/features/daily_point/entity"
	"recything/features/daily_point/model"
	muser "recything/features/user/model"
	"time"

	"gorm.io/gorm"
)

type dailyPointRepository struct {
	db *gorm.DB
}

func NewDailyPointRepository(db *gorm.DB) entity.DailyPointRepositoryInterface {
	return &dailyPointRepository{
		db: db,
	}
}

// DailyClaim implements entity.DailyPointRepositoryInterface.
func (dailyPoint *dailyPointRepository) DailyClaim(userId string) (error) {
	userDaily := new(muser.UserDailyPoints)
	userProf := new (muser.Users)
	pointData := new(model.DailyPoint)
	tx := dailyPoint.db.Begin()

	userDaily.UserID = userId
	
	//melakukan pengecekan untuk menghitung hari claim
	var countData int64
	err := tx.Model(&muser.UserDailyPoints{}).Where("users_id = ?", userId).Count(&countData)
	if err != nil {
		tx.Rollback()
		return err.Error
	}
	dpId := countData + 1

	//melakukakn pengecekan agar user tidak claim 2 kali sehari 
	today := time.Now().Truncate(24 * time.Hour)
	var existingCount int64
	err = tx.Model(&muser.UserDailyPoints{}).Where("users_id = ? AND created_at = ?", userId, today).Count(&existingCount)
	if err != nil {
		tx.Rollback()
		return err.Error
	}

	//apabila lebih dari 1 maka user telah melakukan claim
	if existingCount > 0 {
		tx.Rollback()
		return errors.New("user telah melakukan claim hari ini")
	}

	//pengecekan apabila telah melakukan daily 7 kali
	if dpId == 8{
		err := tx.Where("users_id = ?", userId).Delete(&muser.UserDailyPoints{})
		if err != nil {
			tx.Rollback()
			return err.Error
		}
	}

	//melakukan pengecekan id pada dailypoint untuk mengambil point
	err = tx.Where("id = ?", dpId).First(&pointData)
	if err != nil {
		tx.Rollback()
		return err.Error
	}

	userDaily.DailyPointID = int(dpId)
	userDaily.CreatedAt = time.Now().Truncate(24 * time.Hour)
	userProf.Point = pointData.Point

	return nil
}

// PostWeekly implements entity.DailyPointRepositoryInterface.
func (daily *dailyPointRepository) PostWeekly() error {
	dailyPointData := []entity.DailyPointCore{}

	hari1 := entity.DailyPointCore{Point: 100, Description: "hari 1"}
	dailyPointData = append(dailyPointData, hari1)

	hari2 := entity.DailyPointCore{Point: 150, Description: "hari 2"}
	dailyPointData = append(dailyPointData, hari2)

	hari3 := entity.DailyPointCore{Point: 200, Description: "hari 3"}
	dailyPointData = append(dailyPointData, hari3)

	hari4 := entity.DailyPointCore{Point: 250, Description: "hari 4"}
	dailyPointData = append(dailyPointData, hari4)

	hari5 := entity.DailyPointCore{Point: 300, Description: "hari 5"}
	dailyPointData = append(dailyPointData, hari5)

	hari6 := entity.DailyPointCore{Point: 350, Description: "hari 6"}
	dailyPointData = append(dailyPointData, hari6)

	hari7 := entity.DailyPointCore{Point: 400, Description: "hari 7"}
	dailyPointData = append(dailyPointData, hari7)

	post := entity.ListDailyPointCoreToDailyPointModel(dailyPointData)

	tx := daily.db.Create(post)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}