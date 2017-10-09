package quay

import (
	"encoding/json"
	"os"
	"path"

	"github.com/koudaiii/qucli/utils"
)

func ListRepositoryNotifications(namespace string, name string, hostname string) (RepositoryNotifications, error) {
	var repos ResponseRepositoryNotifications
	var notifications RepositoryNotifications

	u := QuayURLParse(hostname)

	u.Path = path.Join(u.Path, "repository")
	u.Path = path.Join(u.Path, namespace)
	u.Path = path.Join(u.Path, name)
	u.Path = path.Join(u.Path, "notification")

	body, err := utils.HttpGet(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return notifications, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return notifications, err
	}

	for _, item := range repos.Items {
		notifications.Items = append(notifications.Items,
			RepositoryNotification{
				Title:            item["title"].(string),
				Event:            item["event"].(string),
				Method:           item["method"].(string),
				EventConfig:      item["event_config"].(map[string]interface{}),
				UUID:             item["uuid"].(string),
				NumberOfFailures: item["number_of_failures"].(float64),
				Config:           item["config"].(map[string]interface{}),
			})
	}
	return notifications, nil
}
