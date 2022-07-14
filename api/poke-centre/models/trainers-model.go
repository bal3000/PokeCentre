package models

type TrainerAddModel struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	NhsNumber string `json:"nhsNumber"`
}
