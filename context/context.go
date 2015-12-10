package context

import (
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/kardianos/service"

	"github.com/Graylog2/nxlog-sidecar/backends/nxlog"
	"github.com/Graylog2/nxlog-sidecar/daemon"
)

type Ctx struct {
	ServerUrl *url.URL
	NodeId    string
	NxPath    string
	Config    *daemon.Config
	Program   *daemon.Program
	Service   service.Service
	NxConfig  *nxlog.NxConfig
}

func NewContext(serverUrl string, nxPath string, nodeId string) *Ctx {
	dc := daemon.NewConfig(nxPath)
	dp := daemon.NewProgram(dc)
	nxc := nxlog.NewNxConfig(nxPath)

	url, err := url.Parse(serverUrl)
	if err != nil {
		logrus.Fatal("server-url is not valid", err)
	}

	if nodeId == "" {
		logrus.Fatal("please provide a valid node-id")
	}

	return &Ctx{
		ServerUrl: url,
		NodeId:    nodeId,
		NxPath:    nxPath,
		Config:    dc,
		Program:   dp,
		NxConfig:  nxc,
	}
}