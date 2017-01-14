package quay

import (
	"encoding/json"
	"net/url"
	"os"
	"path"

	"github.com/koudaiii/dockerepos/utils"
)

func GetUserPermissions(namespace string, name string) (QuayPermissions, error) {
	var resp QuayPermissionsResponse
	var permissions QuayPermissions

	u, err := url.Parse(QuayURLBase)

	if err != nil {
		return permissions, err
	}
	u.Path = path.Join(u.Path, namespace, name, "permissions", "user") + "/"

	body, err := utils.HttpGet(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return permissions, err
	}
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return permissions, err
	}
	for _, item := range resp.Items {
		permissions.Items = append(permissions.Items,
			QuayPermission{
				Name: item.(map[string]interface{})["name"].(string),
				Role: item.(map[string]interface{})["role"].(string),
			})
	}
	return permissions, nil
}

func GetTeamPermissions(namespace string, name string) (QuayPermissions, error) {
	var resp QuayPermissionsResponse
	var permissions QuayPermissions

	u, err := url.Parse(QuayURLBase)

	if err != nil {
		return permissions, err
	}
	u.Path = path.Join(u.Path, namespace, name, "permissions", "team") + "/"

	body, err := utils.HttpGet(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return permissions, err
	}
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return permissions, err
	}
	for _, item := range resp.Items {
		permissions.Items = append(permissions.Items,
			QuayPermission{
				Name: item.(map[string]interface{})["name"].(string),
				Role: item.(map[string]interface{})["role"].(string),
			})
	}
	return permissions, nil
}
