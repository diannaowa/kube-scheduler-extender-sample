/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

const (
	// APIServerAdvertiseAddress flag sets the IP address the API Server will advertise it's listening on. Specify '0.0.0.0' to use the address of the default network interface.
	APIServerAdvertiseAddress = "apiserver-advertise-address"

	// APIServerBindPort flag sets the port for the API Server to bind to.
	APIServerBindPort = "apiserver-bind-port"

	// APIServerCertSANs flag sets extra Subject Alternative Names (SANs) to use for the API Server serving certificate. Can be both IP addresses and DNS names.
	APIServerCertSANs = "apiserver-cert-extra-sans"

	// APIServerExtraArgs flag sets a extra flags to pass to the API Server or override default ones in form of <flagname>=<value>.
	APIServerExtraArgs = "apiserver-extra-args"

	// CertificatesDir flag sets the path where to save and read the certificates.
	CertificatesDir = "cert-dir"

	// CfgPath flag sets the path to kubeadm config file.
	CfgPath = "config"

	// ControllerManagerExtraArgs flag sets extra flags to pass to the Controller Manager or override default ones in form of <flagname>=<value>.
	ControllerManagerExtraArgs = "controller-manager-extra-args"

	// DryRun flag instruct kubeadm to don't apply any changes; just output what would be done.
	DryRun = "dry-run"

	// FeatureGatesString flag sets key=value pairs that describe feature gates for various features.
	FeatureGatesString = "feature-gates"

	// IgnorePreflightErrors sets the path a list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.
	IgnorePreflightErrors = "ignore-preflight-errors"

	// ImageRepository sets the container registry to pull control plane images from.
	ImageRepository = "image-repository"

	// KubeconfigDir flag sets the path where to save the kubeconfig file.
	KubeconfigDir = "kubeconfig-dir"

	// KubeconfigPath flag sets the kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations are searched for an existing KubeConfig file.
	KubeconfigPath = "kubeconfig"

	// KubernetesVersion flag sets the Kubernetes version for the control plane.
	KubernetesVersion = "kubernetes-version"

	// NetworkingDNSDomain flag sets the domain for services, e.g. "myorg.internal".
	NetworkingDNSDomain = "service-dns-domain"

	// NetworkingServiceSubnet flag sets the range of IP address for service VIPs.
	NetworkingServiceSubnet = "service-cidr"

	// NetworkingPodSubnet flag sets the range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.
	NetworkingPodSubnet = "pod-network-cidr"

	// NodeCRISocket flag sets the CRI socket to connect to.
	NodeCRISocket = "cri-socket"

	// NodeName flag sets the node name.
	NodeName = "node-name"

	// SchedulerExtraArgs flag sets extra flags to pass to the Scheduler or override default ones in form of <flagname>=<value>".
	SchedulerExtraArgs = "scheduler-extra-args"

	// SkipTokenPrint flag instruct kubeadm to skip printing of the default bootstrap token generated by 'kubeadm init'.
	SkipTokenPrint = "skip-token-print"

	// CSROnly flag instructs kubeadm to create CSRs instead of automatically creating or renewing certs
	CSROnly = "csr-only"

	// CSRDir flag sets the location for CSRs and flags to be output
	CSRDir = "csr-dir"

	// TokenStr flags sets both the discovery-token and the tls-bootstrap-token when those values are not provided
	TokenStr = "token"

	// TokenTTL flag sets the time to live for token
	TokenTTL = "token-ttl"

	// TokenUsages flag sets the usages of the token
	TokenUsages = "usages"

	// TokenGroups flag sets the authentication groups of the token
	TokenGroups = "groups"

	// TokenDescription flag sets the description of the token
	TokenDescription = "description"

	// TLSBootstrapToken flag sets the token used to temporarily authenticate with the Kubernetes Control Plane to submit a certificate signing request (CSR) for a locally created key pair
	TLSBootstrapToken = "tls-bootstrap-token"

	// TokenDiscovery flag sets the token used to validate cluster information fetched from the API server (for token-based discovery)
	TokenDiscovery = "discovery-token"

	// TokenDiscoveryCAHash flag instruct kubeadm to validate that the root CA public key matches this hash (for token-based discovery)
	TokenDiscoveryCAHash = "discovery-token-ca-cert-hash"

	// TokenDiscoverySkipCAHash flag instruct kubeadm to skip CA hash verification (for token-based discovery)
	TokenDiscoverySkipCAHash = "discovery-token-unsafe-skip-ca-verification"

	// FileDiscovery flag sets the file or URL from which to load cluster information (for file-based discovery)
	FileDiscovery = "discovery-file"

	// ControlPlane flag instruct kubeadm to create a new control plane instance on this node
	ControlPlane = "experimental-control-plane"

	// UploadCerts flag instruct kubeadm to upload certificates
	UploadCerts = "experimental-upload-certs"

	// CertificateKey flag sets the key used to encrypt and decrypt certificate secrets
	CertificateKey = "certificate-key"

	// SkipCertificateKeyPrint flag instruct kubeadm to skip printing certificate key used to encrypt certs by 'kubeadm init'.
	SkipCertificateKeyPrint = "skip-certificate-key-print"
)
