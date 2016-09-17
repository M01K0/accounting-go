package executedao

import (
	"github.com/alexyslozada/accounting-go/dao/configuration"
	"github.com/alexyslozada/accounting-go/dao/interfaces"
	"github.com/alexyslozada/accounting-go/dao/postgresql"
	"log"
	"sync"
)

var (
	ObjectDAO interfaces.ObjectDAO
	ObjectProfileDAO interfaces.ObjectProfileDAO
	ProfileDAO interfaces.ProfileDAO
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
		ObjectDAO = postgresql.ObjectDAOPsql{}
		ObjectProfileDAO = postgresql.ObjectProfileDAOPsql{}
		ProfileDAO = postgresql.ProfileDAOPsql{}
	default:
		log.Fatal("No existe el motor de persistencia solicitado")
	}
}
