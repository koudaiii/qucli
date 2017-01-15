package quay

const QuayURLBase = "https://quay.io/api/v1/"

type QuayPermission struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type QuayPermissions struct {
	Items []QuayPermission
}

type QuayPermissionsResponse struct {
	Items map[string]interface{} `json:"permissions"`
}

type QuayRepository struct {
	Namespace   string `json:"namespace"`
	IsPublic    bool   `json:"is_public"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RequestRepository struct {
	Namespace   string `json:"namespace"`
	Visibility  string `json:"visibility"`
	Repository  string `json:"repository"`
	Description string `json:"description"`
}
