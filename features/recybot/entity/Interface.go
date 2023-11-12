package entity

type RecybotRepositoryInterface interface {
	Create(recybot RecybotCore) (RecybotCore, error)
	Update(idData string, recybot RecybotCore) (RecybotCore, error)
	Delete(idData string)  error
	SelectAll() ([]RecybotCore, error)
	SelectById(idData string) (RecybotCore, error)
}

type RecybotServiceInterface interface {
	CreateData(recybot RecybotCore) (RecybotCore, error)
	UpdateData(idData string, recybot RecybotCore) (RecybotCore, error)
	DeleteData(idData string)  error
	SelectAllData() ([]RecybotCore, error)
	SelectById(idData string) (RecybotCore, error)
}
