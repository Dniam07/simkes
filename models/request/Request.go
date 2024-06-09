package request

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id         uint      `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	UserroleId int       `json:"userrole_id"`
	SiswaId    int       `json:"siswa_id"`
	GuruId     int       `json:"guru_id"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// Enables soft deletion by using the deleted_at column to mark records as deleted without permanently removing them.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

type Siswa struct {
	ID            uint      `json:"id"`
	Nama          string    `json:"nama"`
	NISN          string    `json:"nisn"`
	Kelas         string    `json:"kelas"`
	TahunAjaranID int       `json:"tahun_ajaran_id"`
	JurusanID     int       `json:"jurusan_id"`
	NamaOrtu      string    `json:"nama_ortu"`
	NoTelpOrtu    string    `json:"no_telp_ortu"`
	NoTelpSiswa   string    `json:"no_telp_siswa"`
	Point         int       `json:"point"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	// Enables soft deletion by using the deleted_at column to mark records as deleted without permanently removing them.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
