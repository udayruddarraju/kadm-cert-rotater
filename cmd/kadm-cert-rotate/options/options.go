package options

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	PkiDir           string
	RenewalThreshold string
}

func NewRotaterConfig() *Config {
	return &Config{}
}

// AddFlags adds flags for a specific CMServer to the specified FlagSet
func (s *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.PkiDir, "pki-dir", s.PkiDir, "path to the pki directory that is used during kubeadm init.")
	fs.StringVar(&s.RenewalThreshold, "renewal-threshold", s.RenewalThreshold, "duration with which the rotater is expected to renew expiring certificates.")

	viper.BindPFlag("pki-dir", fs.Lookup("pki-dir"))
	viper.BindPFlag("renewal-threshold", fs.Lookup("renewal-threshold"))

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})
}
