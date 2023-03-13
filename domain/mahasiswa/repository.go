package mahasiswa

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, mahasiswa *Mahasiswa) error
	GetByID(ctx context.Context, id string) (*Mahasiswa, error)
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

func (r *repositoryImpl) GetByID(ctx context.Context, id string) (*Mahasiswa, error) {
	query := `SELECT id, nama, nik, agama, jenis_kelamin, tanggal_lahir, alamat
			  FROM biodata_mahasiswa
			  WHERE id = ?`

	res, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var mhs *Mahasiswa
	for res.Next() {
		mhs = &Mahasiswa{}
		err = res.Scan(&mhs.ID, &mhs.Name, &mhs.NIK, &mhs.Religion, &mhs.Gender, &mhs.DateOfBirth, &mhs.Address)
		if err != nil {
			return nil, err
		}
	}

	return mhs, nil
}
