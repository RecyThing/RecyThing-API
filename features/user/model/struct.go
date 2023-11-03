package model

import "time"

type ResponProfileUsers struct {
  Username string
  Email string
  Nama_lengkap string
  No_telp int
  Alamat string
  Tanggal_lahir time.Time
  Badges string
}

type ResponManageUsers struct {
  No uint
  Username string
  Email string
  No_telp int
}

type ResponDetailManageUsers struct {
  Email string
  No_telp int
  Tanggal_lahir string
  Alamat string
  Total_point int
  Tujuan_penggunaan string
  Akun_terdaftar string
}

type ResponseFAQ struct {
  Title string
  Content string
}




