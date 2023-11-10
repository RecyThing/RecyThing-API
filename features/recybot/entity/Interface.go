package entity

type RecybotRepositoryInterface interface{
	Create(recybot RecybotCore)(RecybotCore, error)
}


type RecybotServiceInterface interface{
	Create(recybot RecybotCore)(RecybotCore, error)
}