package models

// AccountClass Clase de cuenta
type AccountClass struct {
	ID         int16  `json:"id"`
	Class      string `json:"clase"`
	ReportType `json:"tipoInforme"`
	Nature     string `json:"naturaleza"`
}
