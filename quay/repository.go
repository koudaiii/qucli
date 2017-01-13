package quay

import (
	"encoding/json"
	"net/url"
	"os"
	"path"

	"github.com/koudaiii/dockerepos/utils"
)

const QuayURLBase = "https://quay.io/api/v1/repository/"

type QuayRepository struct {
	Namespace   string `json:"namespace"`
	Visibility  string `json:"name"`
	Name        string `json;"repository`
	Description string `json:"description`
}

type QuayRepositoryResponse struct {
	Namespace   string `json:"namespace"`
	Name        string `json;"repository`
	Description string `json:"description`
}

func GetRepository(namespace string, name string) (QuayRepositoryResponse, error) {
	var resp QuayRepositoryResponse
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
