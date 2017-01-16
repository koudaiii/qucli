package quay

import (
	"encoding/json"
	"net/url"
	"os"
	"path"

	"github.com/koudaiii/dockerepos/utils"
)

func GetPermissions(namespace string, name string, accountType string) (QuayPermissions, error) {
	var resp QuayPermissionsResponse
	var permissions QuayPermissions

	u, err := url.Parse(QuayURLBase)

	if err != nil {
		return permissions, err
	}
	u.Path = path.Join(u.Path, "repository", namespace, name, "permissions", accountType) + "/"

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

func DeletePermission(namespace string, name string, accountType string, account string) error {
	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "repository", namespace, name, "permissions", accountType, account)

	_, err = utils.HttpDelete(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return err
	}

	return nil
}

func AddPermission(namespace string, name string, accountType string, account string, role string) (QuayPermission, error) {
	var repos QuayPermission
	var permission QuayPermission
	req, err := json.Marshal(QuayPermission{
		Role: role,
	})

	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return permission, err
	}
	u.Path = path.Join(u.Path, "repository", namespace, name, "permissions", accountType, account)

	body, err := utils.HttpPut(u.String(), os.Getenv("QUAY_API_TOKEN"), req)
	if err != nil {
		return permission, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return repos, err
	}

	return repos, nil
}
