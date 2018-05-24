/*
Copyright 2018 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Options struct {
	Apiserver                           string
	Kubeconfig                          string
	Help                                bool
	Port                                int
	Host                                string
	TelemetryPort                       int
	TelemetryHost                       string
	Collectors                          CollectorSet
	Namespaces                          NamespaceList
	Version                             bool
	DisablePodNonGenericResourceMetrics bool
	AllowedMetrics                      MetricList
	DisabledMetrics                     MetricList

	flags *pflag.FlagSet
}

func NewOptions() *Options {
	return &Options{
		Collectors: CollectorSet{},
		AllowedMetrics: MetricList{},
		DisabledMetrics: MetricList{},
	}
}

func (o *Options) AddFlags() {
	o.flags = pflag.NewFlagSet("", pflag.ExitOnError)
	// add glog flags
	o.flags.AddGoFlagSet(flag.CommandLine)
	o.flags.Lookup("logtostderr").Value.Set("true")
	o.flags.Lookup("logtostderr").DefValue = "true"
	o.flags.Lookup("logtostderr").NoOptDefVal = "true"

	o.flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		o.flags.PrintDefaults()
	}

	o.flags.StringVar(&o.Apiserver, "apiserver", "", `The URL of the apiserver to use as a master`)
	o.flags.StringVar(&o.Kubeconfig, "kubeconfig", "", "Absolute path to the kubeconfig file")
	o.flags.BoolVarP(&o.Help, "help", "h", false, "Print Help text")
	o.flags.IntVar(&o.Port, "port", 80, `Port to expose metrics on.`)
	o.flags.StringVar(&o.Host, "host", "0.0.0.0", `Host to expose metrics on.`)
	o.flags.IntVar(&o.TelemetryPort, "telemetry-port", 81, `Port to expose kube-state-metrics self metrics on.`)
	o.flags.StringVar(&o.TelemetryHost, "telemetry-host", "0.0.0.0", `Host to expose kube-state-metrics self metrics on.`)
	o.flags.Var(&o.Collectors, "collectors", fmt.Sprintf("Comma-separated list of collectors to be enabled. Defaults to %q", &DefaultCollectors))
	o.flags.Var(&o.Namespaces, "namespace", fmt.Sprintf("Comma-separated list of namespaces to be enabled. Defaults to %q", &DefaultNamespaces))
	o.flags.Var(&o.AllowedMetrics, "allow-metric", fmt.Sprintf("Comma-separated list of metrics to be exposed. Defaults to %q", &MetricList{}))
	o.flags.Var(&o.DisabledMetrics, "disable-metric", fmt.Sprintf("Comma-separated list of metrics to not be exposed. Defaults to %q", &MetricList{}))
	o.flags.BoolVarP(&o.Version, "version", "", false, "kube-state-metrics build version information")
	o.flags.BoolVarP(&o.DisablePodNonGenericResourceMetrics, "disable-pod-non-generic-resource-metrics", "", false, "Disable pod non generic resource request and limit metrics")
}

func (o *Options) Parse() error {
	err := o.flags.Parse(os.Args)
	return err
}

func (o *Options) Usage() {
	o.flags.Usage()
}
