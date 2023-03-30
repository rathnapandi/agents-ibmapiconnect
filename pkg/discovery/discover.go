package discovery

import (
	"fmt"
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

	response, err := d.client.ListAPIs()
	if err != nil {
		log.Error(err)
	}
	fmt.Println(response)

	// for _, deployedApi := range response.DeployedApi {
	// 	spec := deployedApi.Component.Definition

	// 	if err := xml.Unmarshal([]byte(spec), &component); err != nil {
	// 		panic(err)
	// 	}
	// 	fileContnet := component.PublicationInfo.SwaggerFile.FileContent
	// 	decodedString, err := base64.StdEncoding.DecodeString(fileContnet)
	// 	if err != nil {
	// 		panic(fmt.Sprintf("Error decoding file: %s", err))
	// 	}
	// 	reader, err := gzip.NewReader(bytes.NewReader(decodedString))
	// 	if err != nil {
	// 		panic(fmt.Sprintf("Error setting up gzip: %s", err))
	// 	}
	// 	result, err := ioutil.ReadAll(reader)
	// 	if err != nil {
	// 		panic(fmt.Sprintf("Unable to decompress : %s", err))
	// 	}
	// 	authPolicy := handleAuthPolicy(deployedApi)
	// 	url := handleEndpointType(deployedApi)
	// 	api := apiconnect.API{
	// 		ID:          deployedApi.ID,
	// 		Name:        component.Name,
	// 		Description: component.PublicationInfo.Description,
	// 		Version:     component.Version,
	// 		// append endpoint url
	// 		Url:           url,
	// 		Documentation: []byte(deployedApi.Metadata.Description),
	// 		ApiSpec:       result,
	// 		AuthPolicy:    authPolicy,
	// 		ComponentId:   component.Id,
	// 		EnvironmentId: deployedApi.Environment.ID,
	// 	}
	// 	apis = append(apis, api)
	// }
	// for _, api := range apis {
	// 	go func(api apiconnect.API) {
	// 		svcDetail := d.serviceHandler.ToServiceDetail(&api)
	// 		if svcDetail != nil {
	// 			d.apiChan <- svcDetail
	// 		}
	// 	}(api)
	// }
}
