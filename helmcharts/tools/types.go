package tools

import (
	mlog "github.com/maxwell92/log"
)

var log = mlog.Log

// Charts struct for charts
type Charts struct {
	Name string `json:"name"`
}

// NewCharts gives a new instance
func NewCharts(name string) *Charts {
	return &Charts{
		Name: name,
	}
}

// Definations for Values genterator
// Image struct in values.yaml
type Image struct {
	Repository string
	Tag        string
	PullPolicy string
}

// Service struct in values.yaml
type Service struct {
	Type     string
	Port     string
	NodePort string
}

// Environment struct in values.yaml
type Environment struct {
	Name  string
	Value string
}

// Environments list in values.yaml
type Env []Environment

// LimitsRequests struct in values.yaml
type LimitsRequests struct {
	CPU    string
	Memory string
}

// Resources struct in values.yaml
type Resources struct {
	Limits   LimitsRequests
	Requests LimitsRequests
}

// Values struct in values.yaml
type Values struct {
	Namespace        string
	AppName          string
	ReplicaCount     int32
	Image            Image
	DebugImage       Image
	ImagePullSecrets string
	NameOverride     string
	FullnameOverride string
	Env              []Environment
	Resources        Resources
	Service          Service
}

// NewValues gives a new/empty instance
func NewValues() *Values {
	return &Values{
		Env: make([]Environment, 0),
	}
}
