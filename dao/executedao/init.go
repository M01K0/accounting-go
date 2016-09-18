package executedao

import (
	"github.com/alexyslozada/accounting-go/dao/configuration"
	"github.com/alexyslozada/accounting-go/dao/interfaces"
	"github.com/alexyslozada/accounting-go/dao/postgresql"
	"log"
	"sync"
)

var (
	CityDAO               interfaces.CityDAO
	DepartmentDAO         interfaces.DepartmentDAO
	IdentificationTypeDAO interfaces.IdentificationTypeDAO
	ObjectDAO             interfaces.ObjectDAO
	ObjectProfileDAO      interfaces.ObjectProfileDAO
	PersonTypeDAO         interfaces.PersonTypeDAO
	ProfileDAO            interfaces.ProfileDAO
	RegimeTypeDAO         interfaces.RegimeTypeDAO
	TaxpayerTypeDAO       interfaces.TaxpayerTypeDAO
	ThirdPartyDAO         interfaces.ThirdPartyDAO
	UserDAO               interfaces.UserDAO
	once                  sync.Once
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
		CityDAO = postgresql.CityDAOPsql{}
		DepartmentDAO = postgresql.DepartmentDAOPsql{}
		IdentificationTypeDAO = postgresql.IdentificationTypeDAOPsql{}
		ObjectDAO = postgresql.ObjectDAOPsql{}
		ObjectProfileDAO = postgresql.ObjectProfileDAOPsql{}
		PersonTypeDAO = postgresql.PersonTypeDAOPsql{}
		ProfileDAO = postgresql.ProfileDAOPsql{}
		RegimeTypeDAO = postgresql.RegimeTypeDAOPsql{}
		TaxpayerTypeDAO = postgresql.TaxpayerTypeDAOPsql{}
		ThirdPartyDAO = postgresql.ThirdPartyDAOPsql{}
		UserDAO = postgresql.UserDAOPsql{}
	default:
		log.Fatal("No existe el motor de persistencia solicitado")
	}
}
