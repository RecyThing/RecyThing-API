package constanta

// Constanta For Role
const (
	USER             = "user"
	ADMIN            = "admin"
	SUPERADMIN       = "super_admin"
)

// Constanta For Success
const (
	SUCCESS_LOGIN      = "berhasil melakukan login"
	SUCCESS_NULL       = "data belum tersedia"
)

// Constanta For Utils
const (
	VERIFICATION_URL = "https://api.recything.my.id/verify-token?token="
	EMAIL_NOT_REGISTER = "email belum terdaftar"
)

// Constanta For Error
const (
	ERROR_TEMPLATE         = "gagal menguraikan template"
	ERROR_DATA_ID          = "id tidak ditemukan"
	ERROR_ID_INVALID       = "id salah"
	ERROR_DATA_EMAIL       = "email tidak ditemukan"
	ERROR_FORMAT_EMAIL     = "format email tidak valid"
	ERROR_AKSES_ROLE       = "akses ditolak"
	ERROR_PASSWORD         = "password lama tidak sesuai"
	ERROR_CONFIRM_PASSWORD = "konfirmasi password tidak sesuai"
	ERROR_EXTRA_TOKEN      = "gagal ekstrak token"
	ERROR_ID_ROLE          = "id atau role tidak ditemukan"
	ERROR_GET_DATA         = "data tidak ditemukan"
	ERROR_EMPTY            = "harap lengkapi data dengan benar"
	ERROR_HASH_PASSWORD    = "error hash password"
)
