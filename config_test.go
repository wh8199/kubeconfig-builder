package kubeconfig

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGenerateTokenConfig(t *testing.T) {
	config, err := GenerateTokenConfig("test bearer token", "https://127.0.0.1:6443", "", "test-user", "test-cluster", true)
	if err != nil {
		t.Error(err)
	}

	k8sconfig := &K8SConfig{}

	t.Log(string(config))

	if err := yaml.Unmarshal(config, k8sconfig); err != nil {
		t.Error(err)
	}
}

func TestGenerateTokenConfigWithTLS(t *testing.T) {
	config, err := GenerateTokenConfig("test bearer token", "https://127.0.0.1:6443", "cacert", "test-user", "test-cluster", false)
	if err != nil {
		t.Error(err)
	}

	k8sconfig := &K8SConfig{}

	t.Log(string(config))

	if err := yaml.Unmarshal(config, k8sconfig); err != nil {
		t.Error(err)
	}
}

func TestGenerateCertConfig(t *testing.T) {
	config, err := GenerateCertConfig("https://127.0.0.1:6443", "", "clientcert", "clientkey", "test-user", "test-cluster", true)
	if err != nil {
		t.Error(err)
	}

	k8sconfig := &K8SConfig{}

	t.Log(string(config))

	if err := yaml.Unmarshal(config, k8sconfig); err != nil {
		t.Error(err)
	}
}

func TestGenerateCertConfigWithTLS(t *testing.T) {
	config, err := GenerateCertConfig("https://127.0.0.1:6443", "cacert", "clientcert", "clientkey", "test-user", "test-cluster", false)
	if err != nil {
		t.Error(err)
	}

	k8sconfig := &K8SConfig{}

	t.Log(string(config))

	if err := yaml.Unmarshal(config, k8sconfig); err != nil {
		t.Error(err)
	}
}

func TestGenerateBasicAuthConfig(t *testing.T) {
	config, err := GenerateBasciAuthConfig("https://127.0.0.1:6443", "", "admin", "admin-password", "test-cluster", true)
	if err != nil {
		t.Error(err)
	}

	k8sconfig := &K8SConfig{}

	t.Log(string(config))

	if err := yaml.Unmarshal(config, k8sconfig); err != nil {
		t.Error(err)
	}
}

func TestGenerateBasicAuthConfigWithTLS(t *testing.T) {
	config, err := GenerateBasciAuthConfig("https://127.0.0.1:6443", "cacert", "admin", "admin-password", "test-cluster", false)
	if err != nil {
		t.Error(err)
	}

	k8sconfig := &K8SConfig{}

	t.Log(string(config))

	if err := yaml.Unmarshal(config, k8sconfig); err != nil {
		t.Error(err)
	}
}
