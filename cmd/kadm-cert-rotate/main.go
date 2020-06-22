package main

import (
	"bytes"
	goflag "flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	gops "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/udayruddarraju/kadm-cert-rotater/cmd/kadm-cert-rotate/options"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta2 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta2"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/certs/renewal"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
)

var cmd = &cobra.Command{
	Use:   "kadm-cert-rotate",
	Short: "kadm-cert-rotate - rotates/renews control plane certificates if the expiration is lesser than the threshold.",
	Long:  "kadm-cert-rotate - rotates/renews control plane certificates if the expiration is lesser than the threshold.",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

// serverConfig is used to capture options passed through CLI and used to initialize the WebhookServer.
var rotaterConfig *options.Config

func init() {
	rotaterConfig = options.NewRotaterConfig()
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	rotaterConfig.AddFlags(pflag.CommandLine)
}

func Run() {
	internalcfg, err := getInternalCfg("", "", kubeadmapiv1beta2.ClusterConfiguration{}, os.Stdout, "kadm-cert-rotater: check-expiration")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to build internal config: %s\n", err.Error())
		os.Exit(1)
	}

	rm, err := renewal.NewManager(&internalcfg.ClusterConfiguration, "/etc/kubernetes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get certificate expiration details: %s\n", err.Error())
		os.Exit(1)
	}

	expiryThreshold, err := time.ParseDuration(rotaterConfig.RenewalThreshold)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid time duration format passed for threshold: %s\n", err.Error())
		os.Exit(1)
	}

	for _, handler := range rm.Certificates() {
		if ok, _ := rm.CertificateExists(handler.Name); ok {
			e, err := rm.GetCertificateExpirationInfo(handler.Name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to get certificate expiration details: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Fprintf(os.Stdout, "time to expiry for certificate %s: %s\n", handler.Name, e.ResidualTime().String())
			if e.ResidualTime() < expiryThreshold {
				renewed, err := rm.RenewUsingLocalCA(handler.Name)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Renewing ceritifcate %s failed: %s\n", handler.Name, err.Error())
					os.Exit(1)
				}
				if !renewed {
					fmt.Printf("Detected external %s, %s can't be renewed\n", handler.CABaseName, handler.LongName)
					os.Exit(1)
				}
				fmt.Printf("%s renewed\n", handler.LongName)
			}
		}
	}

	// find all processes
	processList, err := gops.Processes()
	for _, process := range processList {
		// restart control plane processes by ending a SIGHUP to the process id
		if isControlPlaneProcess(process.Executable()) {
			fmt.Fprintf(os.Stdout, "Restarting process %s with Pid %d\n", process.Executable(), process.Pid())
			cmd := exec.Command("kill", "-1", fmt.Sprintf("%d", process.Pid()))
			var outBytes, errBytes bytes.Buffer
			cmd.Stdout = &outBytes
			cmd.Stderr = &errBytes
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("Restarting process %s with Pid %d failed: %s\n", process.Executable(), process.Pid(), err.Error()))
				continue
			}
			fmt.Sprintf("successfully restarted process %s with pid %d\n", process.Executable(), process.Pid())
		}
	}
}

func main() {
	cmd.Execute()
}

func getInternalCfg(cfgPath string, kubeconfigPath string, cfg kubeadmapiv1beta2.ClusterConfiguration, out io.Writer, logPrefix string) (*kubeadmapi.InitConfiguration, error) {
	// In case the user is not providing a custom config, try to get current config from the cluster.
	// NB. this operation should not block, because we want to allow certificate renewal also in case of not-working clusters
	if cfgPath == "" {
		client, err := kubeconfigutil.ClientSetFromFile(kubeconfigPath)
		if err == nil {
			internalcfg, err := configutil.FetchInitConfigurationFromCluster(client, out, logPrefix, false)
			if err == nil {
				fmt.Println() // add empty line to separate the FetchInitConfigurationFromCluster output from the command output
				return internalcfg, nil
			}
			fmt.Printf("[%s] Error reading configuration from the Cluster. Falling back to default configuration\n\n", logPrefix)
		}
	}

	// Otherwise read config from --config if provided, otherwise use default configuration
	return configutil.LoadOrDefaultInitConfiguration(cfgPath, &kubeadmapiv1beta2.InitConfiguration{}, &cfg)
}

func isControlPlaneProcess(processName string) bool {
	controlPlaneProcessNames := []string{
		"kube-apiserver",
		"kube-controller-manager",
		"kube-scheduler",
		"etcd",
	}
	for _, controlPlaneProcessName := range controlPlaneProcessNames {
		if strings.Contains(controlPlaneProcessName, processName) {
			return true
		}
	}
	return false
}
