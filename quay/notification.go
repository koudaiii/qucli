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

	u.Path = path.Join(u.Path, "repository", namespace, name, "notification")

	body, err := utils.HttpGet(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return notifications, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return notifications, err
	}

	for _, item := range repos.Items {
		if item["title"] == nil {
			item["title"] = ""
		}
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

func DeleteRepositoryNotification(namespace string, name string, uuid string, hostname string) error {
	u := QuayURLParse(hostname)

	u.Path = path.Join(u.Path, "repository", namespace, name, "notification", uuid)

	_, err := utils.HttpDelete(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return err
	}

	return nil
}

func AddRepositoryNotification(namespace string, name string, request RequestRepositoryNotification, hostname string) (RepositoryNotification, error) {
	var repos RepositoryNotification
	req, err := json.Marshal(request)

	u := QuayURLParse(hostname)
	u.Path = path.Join(u.Path, "repository", namespace, name, "notification")

	body, err := utils.HttpPost(u.String()+"/", os.Getenv("QUAY_API_TOKEN"), req)
	if err != nil {
		return repos, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return repos, err
	}

	return repos, nil
}

func TestRepositoryNotification(namespace string, name string, uuid string, hostname string) error {
	u := QuayURLParse(hostname)

	u.Path = path.Join(u.Path, "repository", namespace, name, "notification", uuid, "test")

	_, err := utils.HttpPost(u.String(), os.Getenv("QUAY_API_TOKEN"), nil)
	if err != nil {
		return err
	}

	return nil
}
