package search

import (
	"common-web-framework/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type MeiliSearchClient struct {
	uri     string
	headers http.Header
}

func NewMeiliSearchClient(host string, apiKey string) *MeiliSearchClient {
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	headers.Add("Content-Type", "application/json;charset=utf-8")

	return &MeiliSearchClient{
		uri:     host,
		headers: headers,
	}
}

func (c *MeiliSearchClient) SendRequest(method string, endpoint string, body string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.uri, endpoint)

	client := http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header = c.headers

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *MeiliSearchClient) CreateIndex(index string) error {
	var endpoint = "indexes"

	var payload = map[string]interface{}{
		"uid":        index,
		"primaryKey": "id",
	}

	var jsonStr = utils.ObjectToJson(payload)

	var _, err = c.SendRequest(http.MethodPost, endpoint, jsonStr)

	return err
}

func (c *MeiliSearchClient) DropIndex(index string) error {
	var endpoint = "indexes/" + index

	var _, err = c.SendRequest(http.MethodDelete, endpoint, "")

	return err
}

func (c *MeiliSearchClient) DeleteAllDocument(index string) error {
	var endpoint = fmt.Sprintf("indexes/%s/documents", index)

	var _, err = c.SendRequest(http.MethodDelete, endpoint, "")

	return err
}

func (c *MeiliSearchClient) SaveDocument(index string, jsonDocument string) error {
	var endpoint = fmt.Sprintf("indexes/%s/documents", index)

	var _, err = c.SendRequest(http.MethodPost, endpoint, jsonDocument)

	return err
}

func (c *MeiliSearchClient) SearchDocument(index string, req MeiliSearchRequest) MeiliSearchResponse {

	var endpoint = fmt.Sprintf("indexes/%s/search", index)

	var response, _ = c.SendRequest(http.MethodPost, endpoint, utils.ObjectToJson(req))

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		var result, _ = ioutil.ReadAll(response.Body)

		return utils.JsonToObject[MeiliSearchResponse](string(result))
	}

	return MeiliSearchResponse{}
}
