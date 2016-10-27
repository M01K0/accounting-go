package postgresql

import "regexp"

// ValidateStringSQL valida que no se intente un SQL Injection,
// generalmente se usa en el Order By
func ValidateStringSQL(value string) string {
	valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
	if valid.MatchString(value) {
		return value
	}
	return ""
}
