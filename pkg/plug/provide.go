package plug

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type ProvideInit = func(db DB, logger Logger) ProviderInterface

func Provide(pinit ProvideInit) {
	logger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
	provider := pinit(DB{}, Logger{logger})

	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: provider,
			},
		},
		Logger: logger,
	}
	plugin.Serve(&config)
}