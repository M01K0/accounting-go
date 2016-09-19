package executedao

import (
	"github.com/alexyslozada/accounting-go/dao/configuration"
	"github.com/alexyslozada/accounting-go/dao/interfaces"
	"github.com/alexyslozada/accounting-go/dao/postgresql"
	"log"
	"sync"
)

var (
	AccountClassDAO       interfaces.AccountClassDAO
	AccountingDocumentDAO interfaces.AccountingDocumentDAO
	AccountLevelDAO       interfaces.AccountLevelDAO
	AccountPeriodDAO      interfaces.AccountPeriodDAO
	AccountPUCDAO         interfaces.AccountPUCDAO
	CityDAO               interfaces.CityDAO
	CompanyDAO            interfaces.CompanyDAO
	CostCenterDAO         interfaces.CostCenterDAO
	DepartmentDAO         interfaces.DepartmentDAO
	FunctionaryCompanyDAO interfaces.FunctionaryCompanyDAO
	FunctionaryTypeDAO    interfaces.FunctionaryTypeDAO
	IdentificationTypeDAO interfaces.IdentificationTypeDAO
	ObjectDAO             interfaces.ObjectDAO
	ObjectProfileDAO      interfaces.ObjectProfileDAO
	PersonTypeDAO         interfaces.PersonTypeDAO
	ProfileDAO            interfaces.ProfileDAO
	RegimeTypeDAO         interfaces.RegimeTypeDAO
	ReportTypeDAO         interfaces.ReportTypeDAO
	TaxDAO                interfaces.TaxDAO
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
		AccountClassDAO = postgresql.AccountClassDAOPsql{}
		AccountingDocumentDAO = postgresql.AccountingDocumentDAOPsql{}
		AccountLevelDAO = postgresql.AccountLevelDAOPsql{}
		AccountPeriodDAO = postgresql.AccountPeriodDAOPsql{}
		AccountPUCDAO = postgresql.AccountPUCDAOPsql{}
		CityDAO = postgresql.CityDAOPsql{}
		CompanyDAO = postgresql.CompanyDAOPsql{}
		CostCenterDAO = postgresql.CostCenterDAOPsql{}
		DepartmentDAO = postgresql.DepartmentDAOPsql{}
		FunctionaryCompanyDAO = postgresql.FunctionaryCompanyDAOPsql{}
		FunctionaryTypeDAO = postgresql.FunctionaryTypeDAOPsql{}
		IdentificationTypeDAO = postgresql.IdentificationTypeDAOPsql{}
		ObjectDAO = postgresql.ObjectDAOPsql{}
		ObjectProfileDAO = postgresql.ObjectProfileDAOPsql{}
		PersonTypeDAO = postgresql.PersonTypeDAOPsql{}
		ProfileDAO = postgresql.ProfileDAOPsql{}
		RegimeTypeDAO = postgresql.RegimeTypeDAOPsql{}
		ReportTypeDAO = postgresql.ReportTypeDAOPsql{}
		TaxDAO = postgresql.TaxDAOPsql{}
		TaxpayerTypeDAO = postgresql.TaxpayerTypeDAOPsql{}
		ThirdPartyDAO = postgresql.ThirdPartyDAOPsql{}
		UserDAO = postgresql.UserDAOPsql{}
	default:
		log.Fatal("No existe el motor de persistencia solicitado")
	}
}
