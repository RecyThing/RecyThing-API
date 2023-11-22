package constanta

// Constanta For Role
const (
	USER       = "user"
	ADMIN      = "admin"
	SUPERADMIN = "super_admin"
)

// Constanta For Success
const (
	SUCCESS_LOGIN       = "berhasil melakukan login"
	SUCCESS_NULL        = "data belum tersedia"
	SUCCESS_CREATE_DATA = "berhasil membuat data"
	SUCCESS_DELETE_DATA = "berhasil menghapus data"
	SUCCESS_GET_DATA    = "berhasil mendapatkan data"
)

// Constanta For Utils
const (
	VERIFICATION_URL = "http://localhost:8080/verify-token?token="
	// VERIFICATION_URL   = "https://api.recything.my.id/verify-token?token="
	EMAIL_NOT_REGISTER = "email belum terdaftar"
)

// Constanta For Error
const (
	ERROR_TEMPLATE         = "gagal menguraikan template"
	ERROR_DATA_ID          = "id tidak ditemukan"
	ERROR_ID_INVALID       = "id salah"
	ERROR_DATA_EMAIL       = "email tidak ditemukan"
	ERROR_FORMAT_EMAIL     = "format email tidak valid"
	ERROR_EMAIL_EXIST      = "email sudah digunakan"
	ERROR_AKSES_ROLE       = "akses ditolak"
	ERROR_PASSWORD         = "password lama tidak sesuai"
	ERROR_CONFIRM_PASSWORD = "konfirmasi password tidak sesuai"
	ERROR_EXTRA_TOKEN      = "gagal ekstrak token"
	ERROR_ID_ROLE          = "id atau role tidak ditemukan"
	ERROR_GET_DATA         = "data tidak ditemukan"
	ERROR_EMPTY            = "harap lengkapi data dengan benar"
	ERROR_HASH_PASSWORD    = "error hash password"
	ERROR_DATA_NOT_FOUND   = "data tidak ditemukan"
	ERROR_DATA_EXIST       = "data sudah ada"
	ERROR_INVALID_INPUT    = "data yang diinput tidak sesuai"
	ERROR_NOT_FOUND        = "data tidak ditemukan"
)

var (
	Unit     = []string{"barang", "kilogram"}
	Category = []string{"sampah anorganik", "sampah organik", "informasi", "batasan"}
)
