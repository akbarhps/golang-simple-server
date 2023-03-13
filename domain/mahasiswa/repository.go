package mahasiswa

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, mahasiswa *Mahasiswa) error
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) Create(ctx context.Context, mhs *Mahasiswa) error {
	query := `INSERT INTO biodata_mahasiswa(id, nama, nik, agama, jenis_kelamin, tanggal_lahir, alamat)
			  VALUES(?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query,
		mhs.ID,
		mhs.Name,
		mhs.NIK,
		mhs.Religion,
		mhs.Gender,
		mhs.DateOfBirth,
		mhs.Address)

	return err
}
