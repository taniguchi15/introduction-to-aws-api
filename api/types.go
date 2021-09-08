package api

type (
	Request interface{ Request() }

	User struct {
		Name         string `json:"name"`
		NameKana     string `json:"nameKana"`
		Email        string `json:"email"`
		Age          int    `json:"age"`
		ZipCode      int    `json:"zipCode"`
		Address      string `json:"address"`
		Tel          string `json:"tel"`
		MobileNumber string `json:"mobileNumber"`
		Hobby        string `json:"hobby"`
		WorkTel      string `json:"workTel"`
	}
)

func (u *User) Request() {}
