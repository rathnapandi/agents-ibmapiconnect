package apiconnect

import (
	"context"
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
	GetAPI(id string) (*API, error)
	ListAPIs() (*ListCatalogResponse, error)
	OnConfigChange(ibmApiConnectConfig *config.IbmApiConnectConfig)
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
	GetAccessToken() (*OauthToken, *Credential, time.Duration, error)
}

// NewClient creates a new client for interacting with IBM Api connect.
func NewClient(apiconnectConfig *config.ApiconnectConfig) *IbmApiConnectClient {
	client := &IbmApiConnectClient{}
	client.apiClient = coreapi.NewClient(apiconnectConfig.TLS, apiconnectConfig.ProxyURL)
	client.OnConfigChange(apiconnectConfig)
	hc.RegisterHealthcheck("IBM API Gateway", HealthCheckEndpoint, client.healthcheck)
	return client
}

func (c *IbmApiConnectClient) OnConfigChange(apiconnectConfig *config.ApiconnectConfig) {
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
	url := c.url + "/graphql/health"
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
func (c *IbmApiConnectClient) ListAPIs(ctx context.Context) (*ListCatalogResponse, error) {

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

func (a *IbmApiConnectClient) getSpec(specFile string) (string, []byte, error) {
	return "", nil, nil
}

// GetAPI gets a single api by id
func (c *IbmApiConnectClient) GetAPI(id string) (*API, error) {

	return nil, nil
}

func (c *IbmApiConnectClient) GetAccessToken() (*OauthToken, *Credential, time.Duration, error) {

	oauthToken := &OauthToken{}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	request := coreapi.Request{
		Method:  coreapi.POST,
		URL:     c.url + "/token",
		Headers: headers,
	}

	log.Println("Calling get JWT token")
	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, nil, 0, agenterrors.Wrap(ErrCommunicatingWithGateway, err.Error())
	}
	if response.Code != http.StatusOK {
		return nil, nil, 0, ErrAuthentication
	}
	err = json.Unmarshal(response.Body, oauthToken)
	if err != nil {
		return nil, nil, 0, err
	}

	user := Credential{
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		ApiKey:       c.apiKey,
	}
	return oauthToken, &user, c.lifetime, nil
}
