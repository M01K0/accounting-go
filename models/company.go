package models

// Company Datos generales de la empresa
type Company struct {
	IdentificationType `json:"tipoIdentificacion"`
	Identification     string `json:"numeroIdentificacion"`
	VerificationDigit  string `json:"digitoVerificacion"`
	Name               string `json:"nombre"`
	Address            string `json:"direccion"`
	PhoneNumber        string `json:"telefono"`
	Department        `json:"departamento"`
	City               `json:"ciudad"`
	WebSite            string `json:"sitioWeb"`
	Email              string `json:"correo"`
	Activity           string `json:"actividad"`
	AutoRetainer       bool   `json:"autorretenedor"`
	PersonType         `json:"tipoPersona"`
	RegimeType         `json:"tipoRegimen"`
	Logo               string `json:"logo"`
	TaxpayerType       `json:"tipoContribuyente"`
}
