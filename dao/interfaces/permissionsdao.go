package interfaces

type PermissionsDAO interface {
	GetScopes(id int16) (map[string][]string, error)
}
