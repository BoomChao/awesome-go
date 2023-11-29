package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	KubeadmJoinConf = `
	apiVersion: {{.apiVersion}}
	kind: JoinConfiguration
	discovery:
	  file:
		kubeConfigPath: {{.kubeConfigPath}}
	  tlsBootstrapToken: {{.tlsBootstrapToken}}
	nodeRegistration:
	  criSocket: {{.criSocket}}
	  name: {{.name}}
	  ignorePreflightErrors:
		- FileAvailable--etc-kubernetes-kubelet.conf
		- DirAvailable--etc-kubernetes-manifests
		{{- range $index, $value := .ignorePreflightErrors}}
		- {{$value}}
		{{- end}}
	  kubeletExtraArgs:
		rotate-certificates: "false"
		pod-infra-container-image: {{.podInfraContainerImage}}
		node-labels: {{.nodeLabels}}
		{{- if .networkPlugin}}
		network-plugin: {{.networkPlugin}}
		{{end}}
		{{- if .containerRuntime}}
		container-runtime: {{.containerRuntime}}
		{{end}}
		{{- if .containerRuntimeEndpoint}}
		container-runtime-endpoint: {{.containerRuntimeEndpoint}}
		{{end}}
	`
)

func TestTemlate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := map[string]interface{}{
			"kubeConfigPath":    "kubeconfigPaht",
			"tlsBootstrapToken": "token",
		}

		ctx["containerRuntime"] = "remote"

		s, err := SubsituteTemplate(KubeadmJoinConf, ctx)
		assert.NoError(t, err)
		t.Log(s)
	})
}
