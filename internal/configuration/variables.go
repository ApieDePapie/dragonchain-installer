package configuration

// Version is the version of this tool (changes for each release, set when compiling with the Makefile)
var Version string

// DragonchainHelmVersion helm version of dragonchain to use
var DragonchainHelmVersion = "1.0.4"

// OpenfaasHelmVersion helm version of openfaas (faas-netes) to use
var OpenfaasHelmVersion = "5.2.2"

// RegistryHelmVersion helm version of docker container registry to use
var RegistryHelmVersion = "1.8.3"

// RegistryIP the clusterip to use for the docker registry deployment
var RegistryIP = "10.98.76.54"

// RegistryPort the port to use for the docker registry deployment
var RegistryPort = 5000

// MinikubeContext the name of the minikube profile to use, which is also the kubernetes context and VM name
var MinikubeContext = "dragonchain"

// KubernetesVersion the kubernetes version to use with the dragonchain's minikube cluster
var KubernetesVersion = "v1.15.5"

// MinikubeVMMemory amount of memory to give to the minikube VM (only applicable when creating new minikube cluster)
var MinikubeVMMemory = "4000mb"

// MinikubeCpus number of cpus to give to the minikube VM (only applicable when creating new minikube cluster)
var MinikubeCpus = 2

// LinuxVirtualboxLink direct link for linux virtualbox installer download
var LinuxVirtualboxLink = "https://download.virtualbox.org/virtualbox/6.0.14/VirtualBox-6.0.14-133895-Linux_amd64.run"

// MacosVirtualboxLink direct link for macos virtualbox installer download
var MacosVirtualboxLink = "https://download.virtualbox.org/virtualbox/6.0.14/VirtualBox-6.0.14-133895-OSX.dmg"

// WindowsVirtualboxLink direct link for windows virtualbox installer download
var WindowsVirtualboxLink = "https://download.virtualbox.org/virtualbox/6.0.14/VirtualBox-6.0.14-133895-Win.exe"

// LinuxMinikubeLink direct link for linux minikube executable
var LinuxMinikubeLink = "https://github.com/kubernetes/minikube/releases/download/v1.5.2/minikube-linux-amd64"

// MacosMinikubeLink direct link for macos minikube executable
var MacosMinikubeLink = "https://github.com/kubernetes/minikube/releases/download/v1.5.2/minikube-darwin-amd64"

// WindowsMinikubeLink direct link for windows minikube executable
var WindowsMinikubeLink = "https://github.com/kubernetes/minikube/releases/download/v1.5.2/minikube-windows-amd64.exe"

// LinuxHelmLink direct link for linux helm package
var LinuxHelmLink = "https://get.helm.sh/helm-v2.16.1-linux-amd64.tar.gz"

// MacosHelmLink direct link for macos helm package
var MacosHelmLink = "https://get.helm.sh/helm-v2.16.1-darwin-amd64.tar.gz"

// WindowsHelmLink direct link for windows helm package
var WindowsHelmLink = "https://get.helm.sh/helm-v2.16.1-windows-amd64.zip"

// LinuxKubectlLink direct link for linux kubectl executable
var LinuxKubectlLink = "https://storage.googleapis.com/kubernetes-release/release/v1.16.2/bin/linux/amd64/kubectl"

// MacosKubectlLink direct link for macos kubectl executable
var MacosKubectlLink = "https://storage.googleapis.com/kubernetes-release/release/v1.16.2/bin/darwin/amd64/kubectl"

// WindowsKubectlLink direct link for windows kubectl executable
var WindowsKubectlLink = "https://storage.googleapis.com/kubernetes-release/release/v1.16.2/bin/windows/amd64/kubectl.exe"

// SetDefaultCredentials indicates whether or not to set the default chain whe configuring the credentials ini file
var SetDefaultCredentials = true