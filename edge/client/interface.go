package client

import (
	"net/http"
	"time"

	portainer "github.com/portainer/portainer/api"

	"github.com/portainer/agent"
)

type EnvironmentStatusOptions struct {
	DoSnapshot bool
	DoCommand  bool
}

type PortainerClient interface {
	GetEnvironmentID() (portainer.EndpointID, error)
	GetEnvironmentStatus(options EnvironmentStatusOptions) (*PollStatusResponse, error)
	GetEdgeStackConfig(edgeStackID int) (*agent.EdgeStackConfig, error)
	SetEdgeStackStatus(edgeStackID, edgeStackStatus int, error string) error
	DeleteEdgeStackStatus(edgeStackID int) error
	SetEdgeJobStatus(edgeJobStatus agent.EdgeJobStatus) error
	SetTimeout(t time.Duration)
	SetLastCommandTimestamp(timestamp time.Time)
	EnqueueLogCollectionForStack(logCmd LogCommandData) error
}

type PollStatusResponse struct {
	Status          string           `json:"status"`
	Port            int              `json:"port"`
	Schedules       []agent.Schedule `json:"schedules"`
	CheckinInterval float64          `json:"checkin"`
	Credentials     string           `json:"credentials"`
	Stacks          []StackStatus    `json:"stacks"`
	VersionUpdate   VersionUpdate    `json:"versionUpdate"`

	// Async mode only
	EndpointID       int            `json:"endpointID"`
	PingInterval     time.Duration  `json:"pingInterval"`
	SnapshotInterval time.Duration  `json:"snapshotInterval"`
	CommandInterval  time.Duration  `json:"commandInterval"`
	AsyncCommands    []AsyncCommand `json:"commands"`
}

type StackStatus struct {
	ID               int
	Version          int
	Name             string // used in async mode
	FileContent      string // used in async mode
	CommandOperation string // used in async mode
}

type VersionUpdate struct {
	// Target version
	Version string
	// Scheduled time
	ScheduledTime int64
	// If need to update
	Active bool
	// Update schedule ID
	ScheduleID int
}

type setEndpointIDFn func(portainer.EndpointID)
type getEndpointIDFn func() portainer.EndpointID

// NewPortainerClient returns a pointer to a new PortainerClient instance
func NewPortainerClient(serverAddress string, setEIDFn setEndpointIDFn, getEIDFn getEndpointIDFn, edgeID string, edgeAsyncMode bool, agentPlatform agent.ContainerPlatform, httpClient *http.Client, updateScheduleID int) PortainerClient {
	if edgeAsyncMode {
		return NewPortainerAsyncClient(serverAddress, setEIDFn, getEIDFn, edgeID, agentPlatform, httpClient, updateScheduleID)
	}

	return NewPortainerEdgeClient(serverAddress, setEIDFn, getEIDFn, edgeID, agentPlatform, httpClient, updateScheduleID)
}
