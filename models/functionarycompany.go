package models

// FunctionaryCompany Funcionario de la empresa.
type FunctionaryCompany struct {
	ID                 int16 `json:"id"`
	IdentificationType `json:"tipoIdentificacion"`
	Identification     string `json:"numeroIdentificacion"`
	VerificationDigit  string `json:"digitoVerificacion"`
	Name               string `json:"nombre"`
	FunctionaryType    `json:"tipoFuncionario"`
}
