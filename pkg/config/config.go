package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Axway/agent-sdk/pkg/cmd/properties"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
)

var config *AgentConfig

const (
	pathPollInterval          = "ibm.pollInterval"
	pathLogFile               = "ibm.logFile"
	pathProcessOnInput        = "ibm.processOnInput"
	pathIbmApiConnectUrl      = "ibm.url"
	pathOrganizationName      = "ibm.organizationName"
	pathCatalogName           = "ibm.catalogName"
	pathApiKey                = "ibm.auth.api_key"
	pathClientId              = "ibm.auth.client_id"
	pathClientSecret          = "ibm.auth.client_secret"
	pathAuthLifetime          = "ibm.auth.lifetime"
	pathSSLNextProtos         = "ibm.ssl.nextProtos"
	pathSSLInsecureSkipVerify = "ibm.ssl.insecureSkipVerify"
	pathSSLCipherSuites       = "ibm.ssl.cipherSuites"
	pathSSLMinVersion         = "ibm.ssl.minVersion"
	pathSSLMaxVersion         = "ibm.ssl.maxVersion"
	pathProxyURL              = "ibm.proxyUrl"
	pathCachePath             = "ibm.cachePath"
)

// SetConfig sets the global AgentConfig reference.
func SetConfig(newConfig *AgentConfig) {
	config = newConfig
}

// GetConfig gets the AgentConfig
func GetConfig() *AgentConfig {
	return config
}

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralConfig    corecfg.CentralConfig `config:"central"`
	ApiconnectConfig *ApiconnectConfig     `config:"ibm"`
}

// ApiconnectConfig - represents the config for the IBM API connect gateway
type ApiconnectConfig struct {
	corecfg.IConfigValidator
	AgentType        corecfg.AgentType
	PollInterval     time.Duration     `config:"PollInterval"`
	LogFile          string            `config:"logFile"`
	ProcessOnInput   bool              `config:"processOnInput"`
	CachePath        string            `config:"cachePath"`
	IbmApiConnectUrl string            `config:"url"`
	OrganizationName string            `config:"organizationName"`
	CatalogName      string            `config:"catalogName"`
	ClientId         string            `config:"auth.client_id"`
	ClientSecret     string            `config:"auth.client_secret"`
	ApiKey           string            `config:"auth.api_key"`
	ProxyURL         string            `config:"proxyUrl"`
	SessionLifetime  time.Duration     `config:"auth.lifetime"`
	TLS              corecfg.TLSConfig `config:"ssl"`
}

// ValidateCfg - Validates the gateway config
func (c *ApiconnectConfig) ValidateCfg() (err error) {
	if c.IbmApiConnectUrl == "" {
		return fmt.Errorf("invalid IBM Api Connect configuration: atomSphereUrl is not configured")
	} else {
		_, err := url.ParseRequestURI(c.IbmApiConnectUrl)
		if err != nil {
			return fmt.Errorf("invalid IBM Api Connect Platform URL: %s", c.IbmApiConnectUrl)
		}
	}

	if c.OrganizationName == "" {
		return fmt.Errorf("invalid IBM Api Connect configuration: organizationName is not configured")
	}

	if c.CatalogName == "" {
		return fmt.Errorf("invalid IBM Api Connect configuration: catalogName is not configured")
	}

	if c.ClientId == "" {
		return fmt.Errorf("invalid IBM Api Connect configuration: client_id is not configured")
	}

	if c.ClientSecret == "" {
		return fmt.Errorf("invalid IBM Api Connect configuration: client_secret is not configured")
	}
	if c.ApiKey == "" {
		return fmt.Errorf("invalid IBM Api Connect configuration: api_key is not configured")
	}
	if c.PollInterval == 0 {
		return errors.New("invalid IBM Api Connect configuration: pollInterval is invalid")
	}

	if _, err := os.Stat(c.CachePath); os.IsNotExist(err) {
		return fmt.Errorf("invalid IBM Api Connect cache path: path does not exist: %s", c.CachePath)
	}
	c.CachePath = filepath.Clean(c.CachePath)

	if c.AgentType == corecfg.TraceabilityAgent && c.LogFile != "" {
		if _, err := os.Stat(c.LogFile); os.IsNotExist(err) {
			return fmt.Errorf("invalid IBM Api Connect log path: path does not exist: %s", c.LogFile)
		}
	}
	return
}

// AddConfigProperties - Adds the command properties needed for IBM Api Connect agent
func AddConfigProperties(props properties.Properties) {
	props.AddDurationProperty(pathPollInterval, 30*time.Second, "Poll interval for read spec discovery/traffic log")
	props.AddStringProperty(pathLogFile, "./logs/traffic.log", "Sample log file with traffic event from gateway")
	props.AddBoolProperty(pathProcessOnInput, true, "Flag to process received event on input or by output before publishing the event by transport")
	props.AddStringProperty(pathIbmApiConnectUrl, "https://apimserver.example.com/api", "IBM Api Connect URL.")
	props.AddStringProperty(pathOrganizationName, "", "IBM Api Connect Provider Organization name.")
	props.AddStringProperty(pathCatalogName, "", "IBM Api Connect Catalog name.")
	props.AddStringProperty(pathClientId, "", "IBM Api Connect client_id.")
	props.AddStringProperty(pathClientSecret, "", "IBM Api Connect client_secret")
	props.AddStringProperty(pathApiKey, "", "IBM Api Connect API Key.")
	props.AddDurationProperty(pathAuthLifetime, 5*time.Minute, "IBM Api Connect session lifetime.")
	props.AddStringProperty(pathCachePath, "/tmp", "IBM Api Connect Cache Path")
	// ssl properties and command flags
	props.AddStringSliceProperty(pathSSLNextProtos, []string{}, "List of supported application level protocols, comma separated.")
	props.AddBoolProperty(pathSSLInsecureSkipVerify, false, "Controls whether a client verifies the server's certificate chain and host name.")
	props.AddStringSliceProperty(pathSSLCipherSuites, corecfg.TLSDefaultCipherSuitesStringSlice(), "List of supported cipher suites, comma separated.")
	props.AddStringProperty(pathSSLMinVersion, corecfg.TLSDefaultMinVersionString(), "Minimum acceptable SSL/TLS protocol version.")
	props.AddStringProperty(pathSSLMaxVersion, "0", "Maximum acceptable SSL/TLS protocol version.")
}

// NewIbmApiconnectConfig - parse the props and create an IBM Api Connect Configuration structure
func NewIbmApiconnectConfig(props properties.Properties, agentType corecfg.AgentType) *ApiconnectConfig {
	return &ApiconnectConfig{
		AgentType:        agentType,
		PollInterval:     props.DurationPropertyValue(pathPollInterval),
		LogFile:          props.StringPropertyValue(pathLogFile),
		ProcessOnInput:   props.BoolPropertyValue(pathProcessOnInput),
		IbmApiConnectUrl: props.StringPropertyValue(pathIbmApiConnectUrl),
		OrganizationName: props.StringPropertyValue(pathOrganizationName),
		CatalogName:      props.StringPropertyValue(pathCatalogName),
		CachePath:        props.StringPropertyValue(pathCachePath),
		ClientId:         props.StringPropertyValue(pathClientId),
		ClientSecret:     props.StringPropertyValue(pathClientSecret),
		ApiKey:           props.StringPropertyValue(pathApiKey),
		ProxyURL:         props.StringPropertyValue(pathProxyURL),
		SessionLifetime:  props.DurationPropertyValue(pathAuthLifetime),
		TLS: &corecfg.TLSConfiguration{
			NextProtos:         props.StringSlicePropertyValue(pathSSLNextProtos),
			InsecureSkipVerify: props.BoolPropertyValue(pathSSLInsecureSkipVerify),
			CipherSuites:       corecfg.NewCipherArray(props.StringSlicePropertyValue(pathSSLCipherSuites)),
			MinVersion:         corecfg.TLSVersionAsValue(props.StringPropertyValue(pathSSLMinVersion)),
			MaxVersion:         corecfg.TLSVersionAsValue(props.StringPropertyValue(pathSSLMaxVersion)),
		},
	}
}
