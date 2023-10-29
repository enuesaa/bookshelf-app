package main

import (
	"net/rpc"
	"encoding/gob"

	"github.com/hashicorp/go-plugin"
	"github.com/enuesaa/teatime/pkg/plug"
)

type PinitResource struct {
	A string
}
func (r *PinitResource) Schema() plug.Schema {
	schema := plug.Schema {
		Name: "hey",
	}
	return schema
}

type PluginConnectorServer struct {}
func (g *PluginConnectorServer) Resource(args interface{}, resp *plug.ResourceWrapper) error {
	resp.Resource = &PinitResource{
		A: "a",
	}
	return nil
}

func (g *PluginConnectorServer) Info(args interface{}, resp *plug.Info) error {
	*resp = plug.Info{
		Name: "aa",
		Description: "bb",
	}
	return nil
}

type PluginConnector struct{}
func (p *PluginConnector) Server(*plugin.MuxBroker) (interface{}, error) {
	return &PluginConnectorServer{}, nil
}
func (PluginConnector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return nil, nil
}

func main() {
	gob.Register(&plug.ResourceWrapper{})
    gob.Register(&PinitResource{})
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &PluginConnector{},
		},
	})
}
