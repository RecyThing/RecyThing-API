package entity

type RecybotRepositoryInterface interface {
	Create(recybot RecybotCore) (RecybotCore, error)
	Update(idData string, data RecybotCore) (RecybotCore, error)
	Delete(idData string)  error
	GetAll() ([]RecybotCore, error)
	GetById(idData string) (RecybotCore, error)
}

type RecybotServiceInterface interface {
	CreateData(recybot RecybotCore) (RecybotCore, error)
	UpdateData(idData string, data RecybotCore) (RecybotCore, error)
	DeleteData(idData string)  error
	GetAllData() ([]RecybotCore, error)
	GetById(idData string) (RecybotCore, error)
	GetPrompt(question string) (string, error)
}
