package mahasiswa

type (
	createRequest struct {
		Name        string `json:"name"`
		NIK         string `json:"nik"`
		Religion    string `json:"religion"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"date_of_birth"`
		Address     string `json:"address"`
	}
)
