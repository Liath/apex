package clojure

import (
	"strings"

	"github.com/apex/apex/function"
	"path/filepath"
)

func init() {
	function.RegisterPlugin("clojure", &Plugin{})
}

const (
	// Runtime name used by Apex
	Runtime = "clojure"

	// RuntimeCanonical represents names used by AWS Lambda
	RuntimeCanonical = "java8"
)

// Plugin Does plugin things.
type Plugin struct{}

// Open adds the shim and golang defaults.
func (p *Plugin) Open(fn *function.Function) error {
	if !strings.HasPrefix(fn.Runtime, "clojure") {
		return nil
	}

	if fn.Hooks.Build == "" {
		fn.Hooks.Build = "lein uberjar && mv target/*-standalone.jar target/apex.jar"
	}

	if fn.Hooks.Clean == "" {
		fn.Hooks.Clean = "rm -fr target"
	}

	if fn.Config.BuildArtifact == "" {
		fn.Config.BuildArtifact = "target/apex.jar"
	}
	fn.Config.BuildArtifact = filepath.Join(fn.Path, fn.Config.BuildArtifact)

	return nil
}

// Deploy sets actual runtime to java.
func (p *Plugin) Deploy(fn *function.Function) error {
	if !strings.HasPrefix(fn.Runtime, "clojure") {
		return nil
	}
	fn.Runtime = RuntimeCanonical

	return nil
}
