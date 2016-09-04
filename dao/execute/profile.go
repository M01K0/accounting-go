package execute

import "github.com/alexyslozada/accounting-go/models"

// InsertProfile inserta un registro en la BD
func InsertProfile(profile *models.Profile) error {
	return profiledao.InsertProfile(profile)
}

// UpdateProfile actualiza un registro en la BD
func UpdateProfile(profile *models.Profile) error {
	return profiledao.UpdateProfile(profile)
}

// DeleteProfile borra un registro de la BD
func DeleteProfile(profile *models.Profile) error {
	return profiledao.DeleteProfile(profile)
}

// GetProfileByID devuelve un registro de la BD por el ID
func GetProfileByID(id int16) (*models.Profile, error) {
	return profiledao.GetProfileByID(id)
}

// GetAllProfiles devuelve varios registros de la BD
func GetAllProfiles() ([]models.Profile, error) {
	return profiledao.GetAllProfiles()
}
