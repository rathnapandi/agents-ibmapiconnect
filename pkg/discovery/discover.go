package discovery

import (
	"encoding/base64"
	"time"

	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/apiconnect"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/config"

	"github.com/Axway/agent-sdk/pkg/cache"

	"github.com/sirupsen/logrus"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

// discovery implements the Repeater interface. Polls IBM API Connect for APIs.
type discovery struct {
	apiChan           chan *ServiceDetail
	cache             cache.Cache
	centralClient     apic.Client
	client            apiconnect.Client
	discoveryPageSize int
	pollInterval      time.Duration
	stopDiscovery     chan bool
	serviceHandler    ServiceHandler
}

func (d *discovery) Stop() {
	d.stopDiscovery <- true
}

func (d *discovery) OnConfigChange(cfg *config.ApiConnectConfig) {
	d.pollInterval = cfg.PollInterval
	d.serviceHandler.OnConfigChange(cfg)
}

// Loop Discovery event loop.
func (d *discovery) Loop() {
	go func() {
		// Instant fist "tick"
		d.discoverAPIs()
		logrus.Info("Starting poller for IBM API Connect APIs")
		ticker := time.NewTicker(d.pollInterval)
		for {
			select {
			case <-ticker.C:
				d.discoverAPIs()
			case <-d.stopDiscovery:
				log.Debug("stopping discovery loop")
				ticker.Stop()
			}
		}
	}()
}

// discoverAPIs Finds APIs from exchange
func (d *discovery) discoverAPIs() {

	apis, err := d.client.ListAPIs()
	if err != nil {
		log.Error(err)
		return
	}

	for _, api := range apis.Results {
		go func(api apiconnect.Results) {
			apiId := api.Id
			wsdl, err := d.client.DownloadWsdl(apiId)
			if err != nil {
				log.Error("Unable to Download WSDL : %s", err)
				return
			}
			var specification []byte
			if wsdl.ContentType == "" {
				specification, err = d.client.DownloadApiSpec(apiId)
				if err != nil {
					log.Error("Unable to download Api specification : %s", err)
					return
				}
			} else {
				specification, err = base64.StdEncoding.DecodeString(wsdl.Content)
				if err != nil {
					log.Error("Unable to DecodeWSDL base 64 content: %s", err)
					return
				}
			}

			dataPowerGateway, err := d.client.GetDatapowerGatewayDetails(api.GatewayServiceUrls[0])
			if err != nil {
				log.Error("Error reaching out Datapower gateway: %s", err)
				return
			}

			amplifyApi := apiconnect.API{
				ID:          api.Id,
				Name:        api.Name,
				Description: api.Title,
				Version:     api.Version,
				// append endpoint url
				Url:           dataPowerGateway.CatalogBase + api.BasePaths[0],
				Documentation: []byte(api.Title),
				ApiSpec:       specification,
				//	AuthPolicy:    authPolicy,
			}
			svcDetail := d.serviceHandler.ToServiceDetail(&amplifyApi)
			if svcDetail != nil {
				d.apiChan <- svcDetail
			}

		}(api)
	}
}
