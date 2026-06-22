package repository_test

import (
	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/repository"
	"fmt"
	"testing"
	"time"
)

func setupTest(t *testing.T) {
	config.InitDB()

	err := config.GetDB().AutoMigrate(&model.Mahasiswa{})
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}
}

// helper untuk buat data dummy
func createTestMahasiswa() model.Mahasiswa {
	npm := time.Now().UnixNano()

	return model.Mahasiswa{
		NPM:    npm,
		Nama:   "Test User",
		Prodi:  "Informatika",
		Alamat: "Bandung",
		NoHP:   "08123456789",
		Hobi:   []string{"Coding"},
	}
}

func TestInsertMahasiswa(t *testing.T) {
	setupTest(t)

	mhs := createTestMahasiswa()

	_, err := repository.InsertMahasiswa(&mhs)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}

	fmt.Printf("INSERTED NPM: %d\n", mhs.NPM)
}

func TestGetAllMahasiswa(t *testing.T) {
	setupTest(t)

	data, err := repository.GetAllMahasiswa()
	if err != nil {
		t.Errorf("GetAll failed: %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Expected data, got empty")
	}

	fmt.Printf("DATA DI TABLE: %+v\n", data)
}

func TestGetMahasiswaByNPM(t *testing.T) {
	setupTest(t)

	mhs := createTestMahasiswa()
	repository.InsertMahasiswa(&mhs)

	result, err := repository.GetMahasiswaByNPM(mhs.NPM)
	if err != nil {
		t.Errorf("GetByNPM failed: %v", err)
	}

	if result.NPM != mhs.NPM {
		t.Errorf("Expected %d, got %d", mhs.NPM, result.NPM)
	}
}

func TestUpdateMahasiswa(t *testing.T) {
	setupTest(t)

	mhs := createTestMahasiswa()
	repository.InsertMahasiswa(&mhs)

	updated, err := repository.UpdateMahasiswa(mhs.NPM, &model.Mahasiswa{
		Nama:   "Updated Name",
		Prodi:  "SI",
		Alamat: "Jakarta",
		Hobi:   []string{"Gaming"},
	})

	if err != nil {
		t.Errorf("Update failed: %v", err)
	}

	if updated.Nama != "Updated Name" {
		t.Errorf("Update tidak berhasil")
	}
}

func TestDeleteMahasiswa(t *testing.T) {
	setupTest(t)

	mhs := createTestMahasiswa()
	repository.InsertMahasiswa(&mhs)

	err := repository.DeleteMahasiswa(mhs.NPM)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	_, err = repository.GetMahasiswaByNPM(mhs.NPM)
	if err == nil {
		t.Errorf("Data masih ada setelah delete")
	}
}
