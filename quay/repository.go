package quay

import (
	"encoding/json"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/koudaiii/qcli/utils"
)

func ListRepository(namespace string, public bool) (QuayRepositories, error) {
	var repos ResponseRepositories
	var repositories QuayRepositories
	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return repositories, err
	}

	values := url.Values{}
	values.Set("public", strconv.FormatBool(public))
	values.Add("namespace", namespace)

	u.Path = path.Join(u.Path, "repository")

	body, err := utils.HttpGet(u.String()+"?"+values.Encode(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return repositories, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return repositories, err
	}

	for _, item := range repos.Items {
		repositories.Items = append(repositories.Items,
			ResponseRepository{
				Namespace: item["namespace"].(string),
				Name:      item["name"].(string),
				IsPublic:  item["is_public"].(bool),
			})
	}
	return repositories, nil
}

func GetRepository(namespace string, name string) (ResponseRepository, error) {
	var repos ResponseRepository
	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return repos, err
	}
	u.Path = path.Join(u.Path, "repository", namespace, name)

	body, err := utils.HttpGet(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return repos, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return repos, err
	}

	return repos, nil
}

func DeleteRepository(namespace string, name string) error {
	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "repository", namespace, name)

	_, err = utils.HttpDelete(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return err
	}

	return nil
}

func CreateRepository(namespace string, name string, visibility string) (QuayRepository, error) {
	var repos QuayRepository
	req, err := json.Marshal(RequestRepository{
		Namespace:  namespace,
		Repository: name,
		Visibility: visibility,
	})

	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return repos, err
	}
	u.Path = path.Join(u.Path, "repository")

	body, err := utils.HttpPost(u.String(), os.Getenv("QUAY_API_TOKEN"), req)
	if err != nil {
		return repos, err
	}

	if err := json.Unmarshal([]byte(body), &repos); err != nil {
		return repos, err
	}

	return repos, nil
}
