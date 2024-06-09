package migration

import (
	"fmt"
	"log"
	"simkes-go/config/database"
	"simkes-go/models/entity"
)

func RunMigration() {
	// Migrasi entity model user
	// Method AutoMigrate mengembalikan error, sehingga harus kita tampung
	err := database.DB.AutoMigrate(
		&entity.Users{},
		&entity.Siswa{},
		&entity.SiswaPelanggar{},
		&entity.Guru{},
		&entity.PasalPelanggaran{},
		&entity.Jurusan{},
		&entity.UserRole{},
		&entity.TahunAjar{},
		&entity.Sanksi{},
	)

	// Cek error
	if err != nil {
		// Tampilkan error
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
