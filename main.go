package main

import (
	"fmt"

	"github.com/alexyslozada/accounting-go/models"
	"github.com/alexyslozada/accounting-go/dao/execute"
)

func main() {
	profile := models.Profile{Name: "Desde DAO"}
	err := execute.InsertProfile(&profile)
	if err != nil {
		fmt.Println("Error al insertar perfil", err)
		profile = models.Profile{}
	}
	fmt.Println(profile)


	/*profile.ID = 27
	err = execute.DeleteProfile(&profile)
	if err != nil {
		fmt.Println("Error al eliminar el perfil", err)
	}
	fmt.Println(profile)
	otherProfile, err := execute.GetProfileByID(12)
	if err != nil {
		fmt.Println("Error al consultar perfil", err)
	}
	fmt.Println(otherProfile)
	otherProfile.Name = "CoNtaBle"
	err = execute.UpdateProfile(otherProfile)
	if err != nil {
		fmt.Println("Error al actualizar perfil", err)
	}
	fmt.Println(otherProfile)
	profiles, err := execute.GetAllProfiles()
	if err != nil {
		fmt.Println("Error al listar los perfiles", err)
	}
	for _, p := range profiles {
		fmt.Println(p)
	}*/
}
