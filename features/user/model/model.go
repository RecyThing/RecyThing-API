package model

import "time"

type Users struct {

  NamaLengkap string
  Username string
  Email string
  KataSandi string
  NoTelp string
  Alamat string
  TanggalLahir string
  Badges string
  Tujuan string
  AkunTerdaftar time.Time
}

type Admins struct {

  NamaLengkap string
  Email string
  KataSandi string
  Status string //Aktif & Tidak Aktif

}

type Achievement struct {

  NamaAchivement string
  Badge string
  TotalTercapai int
  TargetPoin int

}

