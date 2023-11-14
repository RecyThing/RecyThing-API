package request

type RecybotRequest struct {
	Category string `json:"category" valid:"required~kategori tidak boleh kosong"`
	Question string `json:"question" valid:"required~pertanyaan tidak boleh kosong"`
}

