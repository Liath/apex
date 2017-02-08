package clojure

import (
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
  defaultZipPath = "target/apex.jar"
)

type Plugin struct{}

// Open adds the shim and golang defaults.
func (p *Plugin) Open(fn *function.Function) error {
  if fn.Runtime != Runtime {
    return nil
  }
  fn.Runtime = RuntimeCanonical

  if fn.Hooks.Build == "" {
    fn.Hooks.Build = "lein uberjar && mv target/*-standalone.jar target/apex.jar"
  }

  if fn.Hooks.Clean == "" {
    fn.Hooks.Clean = "rm -f target &> /dev/null"
  }

  if (fn.Config.ZipPath == "") {
    fn.Config.ZipPath = defaultZipPath
  }
  fn.Config.ZipPath = filepath.Join(fn.Path, fn.Config.ZipPath)

  return nil
}
