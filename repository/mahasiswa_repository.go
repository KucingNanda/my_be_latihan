package repository

import (
	"be_latihan/config"
	"be_latihan/model"
	"fmt"
)

// Ambil semua data mahasiswa
func GetAllMahasiswa() ([]model.Mahasiswa, error) {
	var data []model.Mahasiswa

	result := config.GetDB().Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

// Insert mahasiswa baru
func InsertMahasiswa(mhs *model.Mahasiswa) (*model.Mahasiswa, error) {
	result := config.GetDB().Create(mhs)
	if result.Error != nil {
		return nil, result.Error
	}

	return mhs, nil
}

// Ambil satu data mahasiswa berdasarkan NPM
func GetMahasiswaByNPM(npm int64) (model.Mahasiswa, error) {
	var mhs model.Mahasiswa

	result := config.GetDB().First(&mhs, "npm = ?", npm)
	if result.Error != nil {
		return mhs, result.Error
	}

	return mhs, nil
}

// Update data mahasiswa berdasarkan NPM
func UpdateMahasiswa(npm int64, newData *model.Mahasiswa) (*model.Mahasiswa, error) {
	var mhs model.Mahasiswa
	db := config.GetDB()

	// cek apakah data ada
	if err := db.First(&mhs, "npm = ?", npm).Error; err != nil {
		return nil, fmt.Errorf("data dengan NPM %d tidak ditemukan", npm)
	}

	// update hanya field yang dikirim (lebih aman)
	if err := db.Model(&mhs).Updates(newData).Error; err != nil {
		return nil, err
	}

	return &mhs, nil
}

// Hapus data mahasiswa berdasarkan NPM
func DeleteMahasiswa(npm int64) error {
	result := config.GetDB().Where("npm = ?", npm).Delete(&model.Mahasiswa{})

	// cek apakah data benar-benar terhapus
	if result.RowsAffected == 0 {
		return fmt.Errorf("data dengan NPM %d tidak ditemukan", npm)
	}

	return result.Error
}
