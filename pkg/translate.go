package main

import (
	"net/http"

	"github.com/cdevents/github-translator/pkg/github"
	"github.com/cdevents/webhook-adapter/pkg/cdevents"
	"github.com/hashicorp/go-plugin"
)

type EventTranslator struct{}

// TranslateEvent Invoked from external application to translate Github event into CDEvent
func (EventTranslator) TranslateEvent(event string, headers http.Header) (string, error) {
	github.Log().Info("Serving from Github-translator plugin")
	cdEvent, err := github.HandleTranslateGithubEvent(event, headers)
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: cdevents.Handshake,
		Plugins: map[string]plugin.Plugin{
			"github-translator-cdevents": &cdevents.TranslatorGRPCPlugin{Impl: &EventTranslator{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
