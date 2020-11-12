package kubeconfig

// information about cluster
type ClusterInfo struct {
	// cluster cacert
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`

	// the path of cluster cacert
	CertificateAuthority string `yaml:"certificate-authority,omitempty"`

	// skip cacert check
	InsecureSkipTLSVerify bool `yaml:"insecure-skip-tls-verify,omitempty"`

	// cluster server address
	Server string `yaml:"server"`
}

// cluster in k8s config
type Cluster struct {
	// cluster info
	Cluster ClusterInfo `yaml:"cluster"`

	// cluster name
	Name string `yaml:"name"`
}

// user info
type UserInfo struct {
	//client cert for tls check
	ClientCertificateData string `yaml:"client-certificate-data,omitempty"`

	//client cert for tls check
	ClientKeyData string `yaml:"client-key-data,omitempty"`

	// the client cert file path
	ClientCertificate string `yaml:"client-certificate,omitempty"`

	// the client key file path
	ClientKey string `yaml:"client-key,omitempty"`

	// bearer token, eg: service accout token
	Token string `yaml:"token,omitempty"`

	// basic auth username
	UserName string `yaml:"username,omitempty"`

	// basic auth password
	Password string `yaml:"password,omitempty"`
}

// user in k8s config
type User struct {
	// user info
	User UserInfo `yaml:"user"`

	// username
	Name string `yaml:"name"`
}

// k8s config context info, it contains cluster name and username
type ContextInfo struct {
	// clustername
	Cluster string `yaml:"cluster"`

	// context name
	User string `yaml:"user"`
}

// k8s context
type Context struct {
	// context info
	Context ContextInfo `yaml:"context"`

	//username
	Name string `yaml:"name"`

	// namespace in this context, if this value is empty, it will be set to 'default'
	Namespace string `yaml:"namespace"`
}

// the defination of k8s config, its style is same to K8S kind/group:version
type K8SConfig struct {
	//api version of config resource, its value is 'v1'
	APIVersion string `yaml:"apiVersion"`

	// kind of config result, its value is 'config'
	Kind string `yaml:"kind"`

	// the context being used
	CurrentContext string `yaml:"current-context"`

	// clusters
	Clusters []Cluster `yaml:"clusters"`

	// users
	Users []User `yaml:"users"`

	// contexts
	Contexts []Context `yaml:"contexts"`
}
