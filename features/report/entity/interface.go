package entity

type ReportRepositoryInterface interface {
	Insert(reportInput ReportCore) (ReportCore, error)
	SelectById(idReport string)(ReportCore, error)
}


type ReportServiceInterface interface {
	Create(reportInput ReportCore, userId string) (ReportCore, error)
}
