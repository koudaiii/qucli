package quay

import (
	"encoding/json"
	"net/url"
	"os"
	"path"

	"github.com/koudaiii/dockerepos/utils"
)

func GetRepository(namespace string, name string) (QuayRepository, error) {
	var resp QuayRepository
	u, err := url.Parse(QuayURLBase)
	if err != nil {
		return resp, err
	}
	u.Path = path.Join(u.Path, namespace, name)

	body, err := utils.HttpGet(u.String(), os.Getenv("QUAY_API_TOKEN"))
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
