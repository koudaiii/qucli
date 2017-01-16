package quay

import (
	"encoding/json"
	"net/url"
	"os"
	"path"

	"github.com/koudaiii/dockerepos/utils"
)

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

func DeleteRepository(namespace string, name string) (ResponseRepository, error) {
	repos := ResponseRepository{
		Name:      name,
		Namespace: namespace,
	}
	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return repos, err
	}
	u.Path = path.Join(u.Path, "repository", namespace, name)

	_, err = utils.HttpDelete(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return repos, err
	}

	return repos, nil
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
