package discovery

import (
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/v3/registry"
)

var (
	ErrDuplication = errors.New("register failed: instance duplicated: ")
	ErrServerError = errors.New("server error")
)

const (
	_registerURL = "http://%s/discovery/register"

	_cancelURL = "http://%s/discovery/cancel"
	_renewURL  = "http://%s/discovery/renew"
	_pollURL   = "http://%s/discovery/polls"

	_codeOK          = 0
	_codeNotFound    = -404
	_codeNotModified = -304

	_registerGap    = 30 * time.Second
	_statusUP       = "1"
	_discoveryAppID = "infra.discovery"
)

type Config struct {
	Nodes  []string
	Region string
	Zone   string
	Env    string
	Host   string
}

func fixConfig(c *Config) error { _ = "STUB: not implemented"; return nil }

type discoveryInstance struct {
	Region   string   `json:"region"`
	Zone     string   `json:"zone"`
	Env      string   `json:"env"`
	AppID    string   `json:"appid"`
	Hostname string   `json:"hostname"`
	Addrs    []string `json:"addrs"`
	Version  string   `json:"version"`
	LastTs   int64    `json:"latest_timestamp"`

	Metadata map[string]string `json:"metadata"`
	Status   int64             `json:"status"`
}

const _reservedInstanceIDKey = "kratos.v2.serviceinstance.id"

func fromServerInstance(ins *registry.ServiceInstance, config *Config) *discoveryInstance {
	_ = "STUB: not implemented"
	return nil
}

func toServiceInstance(ins *discoveryInstance) *registry.ServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

type disInstancesInfo struct {
	Instances map[string][]*discoveryInstance `json:"instances"`
	LastTs    int64                           `json:"latest_timestamp"`
	Scheduler *scheduler                      `json:"scheduler"`
}

type scheduler struct {
	Clients map[string]*zoneStrategy `json:"clients"`
}

type zoneStrategy struct {
	Zones map[string]*strategy `json:"zones"`
}

type strategy struct {
	Weight int64 `json:"weight"`
}

const (
	_paramKeyRegion   = "region"
	_paramKeyZone     = "zone"
	_paramKeyEnv      = "env"
	_paramKeyHostname = "hostname"
	_paramKeyAppID    = "appid"
	_paramKeyAddrs    = "addrs"
	_paramKeyVersion  = "version"
	_paramKeyStatus   = "status"
	_paramKeyMetadata = "metadata"
)

func newParams(c *Config) url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

type discoveryCommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type discoveryPollsResp struct {
	Code int                          `json:"code"`
	Data map[string]*disInstancesInfo `json:"data"`
}
