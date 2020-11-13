# kubeconfig-builder

## description

When generating the k8s config file, the format of the config file was yaml. The way I generated the config file before was to use a config template, but I was troubled by the tabs and spaces in the yaml file for a long time. If you are also troubled by this, you can also use this tool to help you quickly generate config files.

## usage

First, please import this package,

go get github.com/wh8199/kubeconfig-builder

Then, you can generate config with some parameters

### generate config with token

```go
     GenerateTokenConfig(token, server, cacert, username, clustername string, skipCA bool) ([]byte, error)
```

### generate config with client cert and key

```go
    GenerateCertConfig(server, cacert, clientcert, clientkey, username, clustername string, skipCA bool) ([]byte, error)
```

### generate config with basic auth

```go
    GenerateBasciAuthConfig(server, cacert, username, password, clustername string, skipCA bool) ([]byte, error)
```

## test

```makefile
    make test
```
