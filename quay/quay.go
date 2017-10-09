package quay

import (
	"fmt"
	"net/url"
	"os"
)

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

type QuayRepositories struct {
	Items []ResponseRepository
}

type ResponseRepositories struct {
	Items []map[string]interface{} `json:"repositories"`
}

type RequestRepositoryNotification struct {
	Title       string `json:"title"`
	Event       string `json:"event"`
	Method      string `json:"method"`
	Config      []NotificationConfig
	EventConfig []NotificationEventConfig
}

type NotificationConfig struct {
	URL string `json:"url"`
}

type NotificationEventConfig struct {
	EventConfig []map[string]interface{} `json:"eventConfig"`
}

type ResponseRepositoryNotifications struct {
	Items []map[string]interface{} `json:"notifications"`
}

type RepositoryNotifications struct {
	Items []RepositoryNotification
}

type RepositoryNotification struct {
	Title            string                 `json:"title"`
	Event            string                 `json:"event"`
	Method           string                 `json:"method"`
	EventConfig      map[string]interface{} `json:"event_config"`
	UUID             string                 `json:"uuid"`
	NumberOfFailures float64                `json:"number_of_failures"`
	Config           map[string]interface{} `json:"config"`
}

type QuayRepository struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type ResponseRepository struct {
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

func QuayURLParse(hostname string) *url.URL {
	if os.Getenv("QUAY_HOSTNAME") != "" {
		hostname = os.Getenv("QUAY_HOSTNAME")
	}
	u, err := url.Parse("https://" + hostname + "/api/v1/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	return u
}
