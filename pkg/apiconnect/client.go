package apiconnect

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	agenterrors "github.com/Axway/agent-sdk/pkg/util/errors"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/config"

	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
)

const HealthCheckEndpoint = "health"

// Page describes the page query parameter
type Page struct {
	Offset   int
	PageSize int
}

type Client interface {
	ListAPIs() (*ListCatalogResponse, error)
	DownloadWsdl(apiId string) (*Wsdl, error)
	DownloadApiSpec(apiId string) ([]byte, error)
	GetDatapowerGatewayDetails(url string) (*DatapowerGateway, error)
	OnConfigChange(ibmApiConnectConfig *config.ApiConnectConfig)
}

// IbmApiConnectClient is the client for interacting with IBM APIM.
type IbmApiConnectClient struct {
	url              string
	organizationName string
	catalogName      string
	clientId         string
	clientSecret     string
	apiKey           string
	apiClient        coreapi.Client
	lifetime         time.Duration
	auth             Auth
}

type AuthClient interface {
	GetAccessToken() (*OauthToken, *Credential, error)
	RefreshToken(oauthToken *OauthToken) (*OauthToken, error)
}

// NewClient creates a new client for interacting with IBM Api connect.
func NewClient(apiconnectConfig *config.ApiConnectConfig, apiClient coreapi.Client, mock bool) *IbmApiConnectClient {
	client := &IbmApiConnectClient{}
	//client.apiClient = coreapi.NewClient(apiconnectConfig.TLS, apiconnectConfig.ProxyURL)
	client.apiClient = apiClient

	if mock {
		client.auth = &auth{
			token: OauthToken{
				AccessToken: "test",
			},
		}
	} else {
		client.OnConfigChange(apiconnectConfig)
	}

	//hc.RegisterHealthcheck("IBM API Gateway", HealthCheckEndpoint, client.healthcheck)
	return client
}

func (c *IbmApiConnectClient) OnConfigChange(apiconnectConfig *config.ApiConnectConfig) {
	if c.auth != nil {
		c.auth.Stop()
	}

	c.url = apiconnectConfig.IbmApiConnectUrl
	c.organizationName = apiconnectConfig.OrganizationName
	c.catalogName = apiconnectConfig.CatalogName
	c.clientId = apiconnectConfig.ClientId
	c.clientSecret = apiconnectConfig.ClientSecret
	c.apiKey = apiconnectConfig.ApiKey
	c.lifetime = apiconnectConfig.SessionLifetime
	var err error
	c.auth, err = NewAuth(c)
	if err != nil {
		log.Fatalf("Failed to authenticate with IBM API Connect: %s", err.Error())
	}
}

func (c *IbmApiConnectClient) healthcheck(name string) (status *hc.Status) {
	url := c.url + HealthCheckEndpoint
	fmt.Println(url)
	status = &hc.Status{
		Result: hc.OK,
	}

	request := coreapi.Request{
		Method: coreapi.GET,
		URL:    url,
	}
	response, err := c.apiClient.Send(request)

	if err != nil {
		status = &hc.Status{
			Result:  hc.FAIL,
			Details: fmt.Sprintf("%s Failed. Unable to connect to IBM API Connect, check IBM API Connect configuration. %s", name, err.Error()),
		}
	}

	if response.Code != http.StatusOK {
		status = &hc.Status{
			Result:  hc.FAIL,
			Details: fmt.Sprintf("%s Failed. Unable to connect to IBM API Connect, check IBM API Connect configuration.", name),
		}
	}

	return status
}

// ListAPIs lists the API.
func (c *IbmApiConnectClient) ListAPIs() (*ListCatalogResponse, error) {

	application := &ListCatalogResponse{}

	url := fmt.Sprintf("%s/catalogs/%s/%s/apis", c.url, c.organizationName, c.catalogName)

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + c.auth.GetToken(),
	}

	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     url,
		Headers: headers,
	}
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body, application)
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (c *IbmApiConnectClient) DownloadWsdl(apiId string) (*Wsdl, error) {
	wsdl := &Wsdl{}

	url := fmt.Sprintf("%s/catalogs/%s/%s/apis/%s/wsdl", c.url, c.organizationName, c.catalogName, apiId)
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + c.auth.GetToken(),
	}

	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     url,
		Headers: headers,
	}
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(response.Body, wsdl)
	if err != nil {
		return nil, err
	}
	return wsdl, nil

}

func (c *IbmApiConnectClient) DownloadApiSpec(apiId string) ([]byte, error) {
	url := fmt.Sprintf("%s/catalogs/%s/%s/apis/%s/document", c.url, c.organizationName, c.catalogName, apiId)
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + c.auth.GetToken(),
	}

	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     url,
		Headers: headers,
	}
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, err
	}

	return response.Body, nil

}

func (c *IbmApiConnectClient) GetDatapowerGatewayDetails(url string) (*DatapowerGateway, error) {
	datapowerGateway := &DatapowerGateway{}

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + c.auth.GetToken(),
	}
	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     url,
		Headers: headers,
	}
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(response.Body, datapowerGateway)
	if err != nil {
		return nil, err
	}
	return datapowerGateway, nil
}

func (c *IbmApiConnectClient) GetAccessToken() (*OauthToken, *Credential, error) {

	oauthToken := &OauthToken{}
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	authnRequest := AuthnRequest{
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		ApiKey:       c.apiKey,
		GrantType:    "api_key",
	}

	requestBody, err := json.Marshal(authnRequest)
	request := coreapi.Request{
		Method:  coreapi.POST,
		URL:     c.url + "/token",
		Headers: headers,
		Body:    requestBody,
	}

	fmt.Println(string(requestBody))

	log.Println("Create OAuth token")
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, nil, agenterrors.Wrap(ErrCommunicatingWithGateway, err.Error())
	}
	if response.Code != http.StatusOK {
		return nil, nil, ErrAuthentication
	}
	err = json.Unmarshal(response.Body, oauthToken)
	if err != nil {
		return nil, nil, err
	}
	user := Credential{
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		ApiKey:       c.apiKey,
	}
	return oauthToken, &user, nil
}

func (c *IbmApiConnectClient) RefreshToken(oauthToken *OauthToken) (*OauthToken, error) {

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	authnRequest := AuthnRequest{
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		RefreshToken: oauthToken.RefreshToken,
		GrantType:    "refresh_token",
	}

	requestBody, err := json.Marshal(authnRequest)
	request := coreapi.Request{
		Method:  coreapi.POST,
		URL:     c.url + "/token",
		Headers: headers,
		Body:    requestBody,
	}

	log.Println("Calling  Refersh Oauth token")
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, agenterrors.Wrap(ErrCommunicatingWithGateway, err.Error())
	}
	if response.Code != http.StatusOK {
		return nil, ErrAuthentication
	}
	err = json.Unmarshal(response.Body, oauthToken)
	if err != nil {
		return nil, err
	}
	return oauthToken, nil
}
