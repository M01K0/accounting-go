package models

// ThirdParties Terceros
type ThirdParties struct {
	ID                   int16 `json:"id"`
	IdentificationType   `json:"tipoIdentificacion"`
	IdentificationNumber string `json:"numeroIdentificacion"`
	VerificationDigit    string `json:"digitoVerificacion"`
	PersonType           `json:"tipoPersona"`
	BusinessName         string `json:"razonSocial"`
	Lastname             string `json:"primerApellido"`
	SecondSurname        string `json:"segundoApellido"`
	Firstname            string `json:"primerNombre"`
	MiddleName           string `json:"segundoNombre"`
	Address              string `json:"direccion"`
	PhoneNumber          string `json:"telefono"`
	Departament          `json:"departamento"`
	City                 City   `json:"ciudad"`
	Email                string `json:"correo"`
	TaxpayerType         `json:"contribuyente"`
}
