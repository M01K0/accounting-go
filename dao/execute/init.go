package execute

import (
	"github.com/alexyslozada/accounting-go/dao/configuration"
	"github.com/alexyslozada/accounting-go/dao/interfaces"
	"github.com/alexyslozada/accounting-go/dao/postgresql"
	"log"
	"sync"
)

var (
	profiledao interfaces.ProfileDAO
	once       sync.Once
)

func init() {
	once.Do(func() {
		initDAO()
	})
}

// initDAO Inicia los dao dependiendo de la configuración de connection.json
func initDAO() {
	log.Println("Se ha llamado initDAO")
	switch configuration.Config.Engine {
	case "postgresql":
		profiledao = postgresql.ProfileDAOPsql{}
	default:
		log.Fatal("No existe el motor de persistencia solicitado")
	}
}
