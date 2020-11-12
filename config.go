package kubeconfig

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func GenerateTokenConfig(token, server, cacert, username, clustername string, skipCA bool) ([]byte, error) {
	config := &K8SConfig{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters: []Cluster{
			Cluster{
				Cluster: ClusterInfo{
					InsecureSkipTLSVerify:    skipCA,
					Server:                   server,
					CertificateAuthorityData: cacert,
				},
				Name: clustername,
			},
		},
		Users: []User{
			User{
				User: UserInfo{
					Token: token,
				},
				Name: username,
			},
		},
		Contexts: []Context{
			Context{
				Context: ContextInfo{
					Cluster: clustername,
					User:    username,
				},
				Name: username + "-" + clustername + "-context",
			},
		},
		CurrentContext: username + "-" + clustername + "-context",
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fail to generate k8s config: %s", err.Error())
	}

	return data, nil
}

func GenerateCertConfig(server, cacert, clientcert, clientkey, username, clustername string, skipCA bool) ([]byte, error) {
	config := &K8SConfig{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters: []Cluster{
			Cluster{
				Cluster: ClusterInfo{
					InsecureSkipTLSVerify:    skipCA,
					Server:                   server,
					CertificateAuthorityData: cacert,
				},
				Name: clustername,
			},
		},
		Users: []User{
			User{
				User: UserInfo{
					ClientCertificateData: clientcert,
					ClientKeyData:         clientkey,
				},
				Name: username,
			},
		},
		Contexts: []Context{
			Context{
				Context: ContextInfo{
					Cluster: clustername,
					User:    username,
				},
				Name: username + "-" + clustername + "-context",
			},
		},
		CurrentContext: username + "-" + clustername + "-context",
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fail to generate k8s config: %s", err.Error())
	}

	return data, nil
}

func GenerateBasciAuthConfig(server, cacert, username, password, clustername string, skipCA bool) ([]byte, error) {
	config := &K8SConfig{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters: []Cluster{
			Cluster{
				Cluster: ClusterInfo{
					InsecureSkipTLSVerify:    skipCA,
					Server:                   server,
					CertificateAuthorityData: cacert,
				},
				Name: clustername,
			},
		},
		Users: []User{
			User{
				User: UserInfo{
					UserName: username,
					Password: password,
				},
				Name: username,
			},
		},
		Contexts: []Context{
			Context{
				Context: ContextInfo{
					Cluster: clustername,
					User:    username,
				},
				Name: username + "-" + clustername + "-context",
			},
		},
		CurrentContext: username + "-" + clustername + "-context",
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fail to generate k8s config: %s", err.Error())
	}

	return data, nil
}
