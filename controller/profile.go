package controller

import (
	"errors"

	"github.com/alexyslozada/accounting-go/dao/configuration"
	"github.com/alexyslozada/accounting-go/dao/postgresql"
	"github.com/alexyslozada/accounting-go/models"
)

// ProfileInsert Inserta un perfil en la BD
func ProfileInsert(profile *models.Profile) error {
	switch configuration.Config.Engine {
	case "postgresql":
		return postgresql.ProfileInsert(profile)
	case "mysql":
		return errors.New("Aún no se encuentra soportado")
	default:
		return errors.New("Aún no se encuentra soportado")
	}
}
