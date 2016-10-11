package interfaces

type PermissionsDAO interface {
	GetScopes(id int16, method string) ([]string, error)
}
