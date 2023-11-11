package entity

type RecybotRepositoryInterface interface {
	Create(recybot RecybotCore) (RecybotCore, error)
	Update(idData string) (RecybotCore, error)
	Delete(idData string) (RecybotCore, error)
}

type RecybotServiceInterface interface {
	CreateData(recybot RecybotCore) (RecybotCore, error)
	UpdateData(idData string) (RecybotCore, error)
	Delete(idData string) (RecybotCore, error)
}
