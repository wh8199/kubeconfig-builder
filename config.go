package kubeconfig

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

/*
	description: generate k8s config with token
	params:
		token: k8s config token
		server: k8s server address
		cacert: the cacert, if skipCA is equal true, this value will be ignored
		username: the display username of k8s cluster
		clustername: the display clustername of k8s clutser
		skipCA: skip ca check or not, if this value is equal true, the cacert is ignored
*/
func GenerateTokenConfig(token, server, cacert, username, clustername, nsName string, skipCA bool) ([]byte, error) {
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

	if nsName != "" {
		config.Contexts[0].Namespace = nsName
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fail to generate k8s config: %s", err.Error())
	}

	return data, nil
}

/*
	description: generate k8s config with clientcert or key
	params:
		server: k8s server address
		cacert: the cacert, if skipCA is equal true, this value will be ignored
		clientcert: the clientcert of tls
		clientkey: the clientkey of tls
		username: the display username of k8s cluster
		clustername: the display clustername of k8s clutser
		skipCA: skip ca check or not, if this value is equal true, the cacert is ignored
*/
func GenerateCertConfig(server, cacert, clientcert, clientkey, username, clustername, nsName string, skipCA bool) ([]byte, error) {
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

	if nsName != "" {
		config.Contexts[0].Namespace = nsName
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fail to generate k8s config: %s", err.Error())
	}

	return data, nil
}

/*
	description: generate k8s config with clientcert or key
	params:
		server: k8s server address
		cacert: the cacert, if skipCA is equal true, this value will be ignored
		username: the username of basicauth
		password: the password of basicauth
		clustername: the display clustername of k8s clutser
		skipCA: skip ca check or not, if this value is equal true, the cacert is ignored
*/
func GenerateBasciAuthConfig(server, cacert, username, password, clustername, nsName string, skipCA bool) ([]byte, error) {
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

	if nsName != "" {
		config.Contexts[0].Namespace = nsName
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fail to generate k8s config: %s", err.Error())
	}

	return data, nil
}
