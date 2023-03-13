package mahasiswa

import "time"

type Mahasiswa struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	NIK         string    `json:"nik,omitempty"`
	Religion    string    `json:"religion,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	Address     string    `json:"address,omitempty"`
}
