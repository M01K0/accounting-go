package models

// Account cuenta puc
type Account struct {
	ID            int16  `json:"id"`
	Account       string `json:"cuenta"`
	Name          string `json:"nombre"`
	ParentAccount int16  `json:"idCuentaPadre"`
	AccountClass  `json:"claseCuenta"`
	AccountLevel  `json:"nivelCuenta"`
}
