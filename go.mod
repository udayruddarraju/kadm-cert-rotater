module github.com/udayruddarraju/kadm-cert-rotater

replace github.com/udayruddarraju/kadm-cert-rotater => ./

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/aws/aws-sdk-go v1.28.2 // indirect
	github.com/checkpoint-restore/go-criu v0.0.0-20190109184317-bdb7599cd87b // indirect
	github.com/cilium/ebpf v0.0.0-20191025125908-95b36a581eed // indirect
	github.com/coredns/corefile-migration v1.0.6 // indirect
	github.com/docker/libnetwork v0.8.0-dev.2.0.20190624125649-f0e46a78ea34 // indirect
	github.com/elazarl/goproxy v0.0.0-20180725130230-947c36da3153 // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mitchellh/go-ps v1.0.0
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/opencontainers/runc v1.0.0-rc9 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/spf13/afero v1.3.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	gonum.org/v1/gonum v0.6.2 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	k8s.io/api v0.18.4 // indirect
	k8s.io/apimachinery v0.18.4 // indirect
	k8s.io/client-go v0.17.2 // indirect
	k8s.io/cluster-bootstrap v0.18.4 // indirect
	k8s.io/component-base v0.18.4 // indirect
	k8s.io/gengo v0.0.0-20200413195148-3a45101e95ac // indirect
	k8s.io/klog/v2 v2.2.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6 // indirect
	k8s.io/kube-proxy v0.18.4 // indirect
	k8s.io/kubelet v0.18.4 // indirect
	k8s.io/kubernetes v1.17.2
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace k8s.io/api => k8s.io/api v0.17.2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.2

replace k8s.io/apimachinery => k8s.io/apimachinery v0.17.3-beta.0

replace k8s.io/apiserver => k8s.io/apiserver v0.17.2

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.2

replace k8s.io/client-go => k8s.io/client-go v0.17.2

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.2

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.2

replace k8s.io/code-generator => k8s.io/code-generator v0.17.3-beta.0

replace k8s.io/component-base => k8s.io/component-base v0.17.2

replace k8s.io/cri-api => k8s.io/cri-api v0.17.3-beta.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.2

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.2

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.2

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.2

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.2

replace k8s.io/kubectl => k8s.io/kubectl v0.17.2

replace k8s.io/kubelet => k8s.io/kubelet v0.17.2

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.2

replace k8s.io/metrics => k8s.io/metrics v0.17.2

replace k8s.io/node-api => k8s.io/node-api v0.17.2

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.2

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.17.2

replace k8s.io/sample-controller => k8s.io/sample-controller v0.17.2
