package traceability

import (
	"github.com/hpcloud/tail"
)

const (
	healthCheckEndpoint = "ingestion"
	CacheKeyTimeStamp   = "LAST_RUN"
)

type Emitter interface {
	Start() error
}

// IbmEventEmitter - Gathers analytics data for publishing to Central.
type IbmEventEmitter struct {
	eventChannel chan string
	logFile      string
}

// NewIbmEventEmitter - Creates a client to poll for events.
func NewIbmEventEmitter(logFile string, eventChannel chan string) *IbmEventEmitter {
	me := &IbmEventEmitter{
		eventChannel: eventChannel,
		logFile:      logFile,
	}
	return me
}

// Start retrieves analytics data from anypoint and sends them on the event channel for processing.
func (me *IbmEventEmitter) Start() error {
	go me.tailFile()

	return nil

}

func (me IbmEventEmitter) tailFile() {
	t, _ := tail.TailFile(me.logFile, tail.Config{Follow: true})
	for line := range t.Lines {
		me.eventChannel <- line.Text
	}
}
