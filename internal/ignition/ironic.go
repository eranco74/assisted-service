package ignition

import (
	"encoding/json"

	ignition_config_types_32 "github.com/coreos/ignition/v2/config/v3_2/types"
	aiv1beta1 "github.com/openshift/assisted-service/api/v1beta1"
	iccignition "github.com/openshift/image-customization-controller/pkg/ignition"
)

type IronicIgniotionBuilder interface {
	GenerateIronicConfig(ironicBaseURL string, infraEnv aiv1beta1.InfraEnv, ironicAgentImage string) ([]byte, error)
}

type IronicIgniotionConfig struct {
	// TODO: MGMT-10375	Get the ironic image from the default release image for the arch
	BaremetalIronicAgentImage string `envconfig:"IRONIC_AGENT_IMAGE" default:"registry.ci.openshift.org/openshift:ironic-agent:latest"`
}

type ironicIgniotionBuilder struct {
	config IronicIgniotionConfig
}

func NewIronicIgniotionBuilder(config IronicIgniotionConfig) IronicIgniotionBuilder {
	ib := ironicIgniotionBuilder{config}
	return &ib
}

func (r *ironicIgniotionBuilder) GenerateIronicConfig(ironicBaseURL string, infraEnv aiv1beta1.InfraEnv, ironicAgentImage string) ([]byte, error) {
	config := ignition_config_types_32.Config{}
	config.Ignition.Version = "3.2.0"

	httpProxy, httpsProxy, noProxy := getProxyConfigs(infraEnv.Spec.Proxy)
	if ironicAgentImage == "" {
		ironicAgentImage = r.config.BaremetalIronicAgentImage
	}
	// TODO: this should probably get the pullSecret as well
	ib, err := iccignition.New([]byte{}, []byte{}, ironicBaseURL, ironicAgentImage, "", "", "", httpProxy, httpsProxy, noProxy, "")
	if err != nil {
		return []byte{}, err
	}
	config.Storage.Files = []ignition_config_types_32.File{ib.IronicAgentConf()}
	// TODO: sort out the flags (authfile...) and copy network
	config.Systemd.Units = []ignition_config_types_32.Unit{ib.IronicAgentService(false)}
	return json.Marshal(config)
}

func getProxyConfigs(proxy *aiv1beta1.Proxy) (string, string, string) {
	if proxy == nil {
		return "", "", ""
	}
	return proxy.HTTPProxy, proxy.HTTPSProxy, proxy.NoProxy
}
