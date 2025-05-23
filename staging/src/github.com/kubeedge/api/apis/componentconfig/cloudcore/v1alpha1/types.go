/*
Copyright 2019 The KubeEdge Authors.

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

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudCoreConfig indicates the config of cloudCore which get from cloudCore config file
type CloudCoreConfig struct {
	metav1.TypeMeta
	// CommonConfig indicates common config for all modules
	// +Required
	CommonConfig *CommonConfig `json:"commonConfig,omitempty"`
	// KubeAPIConfig indicates the kubernetes cluster info which cloudCore will connected
	// +Required
	KubeAPIConfig *KubeAPIConfig `json:"kubeAPIConfig,omitempty"`
	// Modules indicates cloudCore modules config
	// +Required
	Modules *Modules `json:"modules,omitempty"`
	// FeatureGates is a map of feature names to bools that enable or disable alpha/experimental features.
	FeatureGates map[string]bool `json:"featureGates,omitempty"`
}

// CommonConfig indicates common config for all modules
type CommonConfig struct {
	// TunnelPort indicates the port that the cloudcore tunnel listened
	TunnelPort int `json:"tunnelPort,omitempty"`

	// MonitorServer holds config that exposes prometheus metrics and pprof
	MonitorServer MonitorServer `json:"monitorServer,omitempty"`
}

// MonitorServer indicates MonitorServer config
type MonitorServer struct {
	// BindAddress is the IP address and port for the monitor server to serve on,
	// defaulting to 127.0.0.1:9091 (set to 0.0.0.0 for all interfaces)
	BindAddress string `json:"bindAddress,omitempty"`

	// EnableProfiling enables profiling via web interface on /debug/pprof handler.
	// Profiling handlers will be handled by monitor server.
	EnableProfiling bool `json:"enableProfiling,omitempty"`
}

// KubeAPIConfig indicates the configuration for interacting with k8s server
type KubeAPIConfig struct {
	// Master indicates the address of the Kubernetes API server (overrides any value in KubeConfig)
	// such as https://127.0.0.1:8443
	// default ""
	// Note: Can not use "omitempty" option,  It will affect the output of the default configuration file
	Master string `json:"master"`
	// ContentType indicates the ContentType of message transmission when interacting with k8s
	// default "application/vnd.kubernetes.protobuf"
	ContentType string `json:"contentType,omitempty"`
	// QPS to while talking with kubernetes apiserver
	// default 100
	QPS int32 `json:"qps,omitempty"`
	// Burst to use while talking with kubernetes apiserver
	// default 200
	Burst int32 `json:"burst,omitempty"`
	// KubeConfig indicates the path to kubeConfig file with authorization and master location information.
	// default "/root/.kube/config"
	// +Required
	KubeConfig string `json:"kubeConfig"`
}

// Modules indicates the modules of CloudCore will be use
type Modules struct {
	// CloudHub indicates CloudHub module config
	CloudHub *CloudHub `json:"cloudHub,omitempty"`
	// EdgeController indicates EdgeController module config
	EdgeController *EdgeController `json:"edgeController,omitempty"`
	// DeviceController indicates DeviceController module config
	DeviceController *DeviceController `json:"deviceController,omitempty"`
	// TaskManager indicates TaskManager module config
	TaskManager *TaskManager `json:"taskManager,omitempty"`
	// SyncController indicates SyncController module config
	SyncController *SyncController `json:"syncController,omitempty"`
	// DynamicController indicates DynamicController module config
	DynamicController *DynamicController `json:"dynamicController,omitempty"`
	// CloudStream indicates cloudstream module config
	CloudStream *CloudStream `json:"cloudStream,omitempty"`
	// Router indicates router module config
	Router *Router `json:"router,omitempty"`
	// IptablesManager indicates iptables module config
	IptablesManager *IptablesManager `json:"iptablesManager,omitempty"`
}

// CloudHub indicates the config of CloudHub module.
// CloudHub is a web socket or quic server responsible for watching changes at the cloud side,
// caching and sending messages to EdgeHub.
type CloudHub struct {
	// Enable indicates whether CloudHub is enabled, if set to false (for debugging etc.),
	// skip checking other CloudHub configs.
	// default true
	Enable bool `json:"enable"`
	// KeepaliveInterval indicates keep-alive interval (second)
	// default 30
	KeepaliveInterval int32 `json:"keepaliveInterval,omitempty"`
	// NodeLimit is a maximum number of edge node that can connect to the single CloudCore
	// default 1000
	NodeLimit int32 `json:"nodeLimit,omitempty"`
	// TLSCAFile indicates ca file path
	// default "/etc/kubeedge/ca/rootCA.crt"
	TLSCAFile string `json:"tlsCAFile,omitempty"`
	// TLSCAKeyFile indicates caKey file path
	// default "/etc/kubeedge/ca/rootCA.key"
	TLSCAKeyFile string `json:"tlsCAKeyFile,omitempty"`
	// TLSPrivateKeyFile indicates key file path
	// default "/etc/kubeedge/certs/server.crt"
	TLSCertFile string `json:"tlsCertFile,omitempty"`
	// TLSPrivateKeyFile indicates key file path
	// default "/etc/kubeedge/certs/server.key"
	TLSPrivateKeyFile string `json:"tlsPrivateKeyFile,omitempty"`
	// WriteTimeout indicates write time (second)
	// default 30
	WriteTimeout int32 `json:"writeTimeout,omitempty"`
	// Quic indicates quic server info
	Quic *CloudHubQUIC `json:"quic,omitempty"`
	// UnixSocket set unixsocket server info
	UnixSocket *CloudHubUnixSocket `json:"unixsocket,omitempty"`
	// WebSocket indicates websocket server info
	// +Required
	WebSocket *CloudHubWebSocket `json:"websocket,omitempty"`
	// HTTPS indicates https server info
	// +Required
	HTTPS *CloudHubHTTPS `json:"https,omitempty"`
	// AdvertiseAddress sets the IP address for the cloudcore to advertise.
	AdvertiseAddress []string `json:"advertiseAddress,omitempty"`
	// DNSNames sets the DNSNames for CloudCore.
	DNSNames []string `json:"dnsNames,omitempty"`
	// EdgeCertSigningDuration indicates the validity period of edge certificate
	// default 365d
	EdgeCertSigningDuration time.Duration `json:"edgeCertSigningDuration,omitempty"`
	// TokenRefreshDuration indicates the interval of cloudcore token refresh, unit is hour
	// default 12h
	TokenRefreshDuration time.Duration `json:"tokenRefreshDuration,omitempty"`
	// Authorization authz configurations
	Authorization *CloudHubAuthorization `json:"authorization,omitempty"`
}

// CloudHubQUIC indicates the quic server config
type CloudHubQUIC struct {
	// Enable indicates whether enable quic protocol
	// default false
	Enable bool `json:"enable"`
	// Address set server ip address
	// default 0.0.0.0
	Address string `json:"address,omitempty"`
	// Port set open port for quic server
	// default 10001
	Port uint32 `json:"port,omitempty"`
	// MaxIncomingStreams set the max incoming stream for quic server
	// default 10000
	MaxIncomingStreams int32 `json:"maxIncomingStreams,omitempty"`
}

// CloudHubUnixSocket indicates the unix socket config
type CloudHubUnixSocket struct {
	// Enable indicates whether enable unix domain socket protocol
	// default true
	Enable bool `json:"enable"`
	// Address indicates unix domain socket address
	// default "unix:///var/lib/kubeedge/kubeedge.sock"
	Address string `json:"address,omitempty"`
}

// CloudHubWebSocket indicates the websocket config of CloudHub
type CloudHubWebSocket struct {
	// Enable indicates whether enable websocket protocol
	// default true
	Enable bool `json:"enable"`
	// Address indicates server ip address
	// default 0.0.0.0
	Address string `json:"address,omitempty"`
	// Port indicates the open port for websocket server
	// default 10000
	Port uint32 `json:"port,omitempty"`
}

// CloudHubHTTPS indicates the http config of CloudHub
type CloudHubHTTPS struct {
	// Enable indicates whether enable Https protocol
	// default true
	Enable bool `json:"enable"`
	// Address indicates server ip address
	// default 0.0.0.0
	Address string `json:"address,omitempty"`
	// Port indicates the open port for HTTPS server
	// default 10002
	Port uint32 `json:"port,omitempty"`
}

// CloudHubAuthorization CloudHub authz configurations
type CloudHubAuthorization struct {
	// Enable indicates whether enable CloudHub Authorization
	// default false
	Enable bool `json:"enable"`
	// Debug only logs errors but always allow messages
	// default false
	Debug bool `json:"debug"`
	// Modes a list of authorization modes will be used
	// default node
	Modes []AuthorizationMode `json:"modes"`
}

// AuthorizationMode indicates an authorization mdoe
type AuthorizationMode struct {
	// Node node authorization
	Node *NodeAuthorization `json:"node,omitempty"`
}

// NodeAuthorization node authorization
type NodeAuthorization struct {
	// Enable enables node authorization
	// default true
	Enable bool `json:"enable"`
}

// EdgeController indicates the config of EdgeController module
type EdgeController struct {
	// Enable indicates whether EdgeController is enabled,
	// if set to false (for debugging etc.), skip checking other EdgeController configs.
	// default true
	Enable bool `json:"enable"`
	// NodeUpdateFrequency indicates node update frequency (second)
	// default 10
	NodeUpdateFrequency int32 `json:"nodeUpdateFrequency,omitempty"`
	// Buffer indicates k8s resource buffer
	Buffer *EdgeControllerBuffer `json:"buffer,omitempty"`
	// Load indicates EdgeController load
	Load *EdgeControllerLoad `json:"load,omitempty"`
}

// EdgeControllerBuffer indicates the EdgeController buffer
type EdgeControllerBuffer struct {
	// ProcessEvent indicates the buffer of process event
	// default 1024
	ProcessEvent int32 `json:"processEvent,omitempty"`
	// UpdatePodStatus indicates the buffer of update pod status
	// default 1024
	UpdatePodStatus int32 `json:"updatePodStatus,omitempty"`
	// UpdateNodeStatus indicates the buffer of update node status
	// default 1024
	UpdateNodeStatus int32 `json:"updateNodeStatus,omitempty"`
	// QueryConfigMap indicates the buffer of query configMap
	// default 1024
	QueryConfigMap int32 `json:"queryConfigMap,omitempty"`
	// QuerySecret indicates the buffer of query secret
	// default 1024
	QuerySecret int32 `json:"querySecret,omitempty"`
	// PodEvent indicates the buffer of pod event
	// default 1
	PodEvent int32 `json:"podEvent,omitempty"`
	// ConfigMapEvent indicates the buffer of configMap event
	// default 1
	ConfigMapEvent int32 `json:"configMapEvent,omitempty"`
	// SecretEvent indicates the buffer of secret event
	// default 1
	SecretEvent int32 `json:"secretEvent,omitempty"`
	// RulesEvent indicates the buffer of rule event
	// default 1
	RulesEvent int32 `json:"rulesEvent,omitempty"`
	// RuleEndpointsEvent indicates the buffer of endpoint event
	// default 1
	RuleEndpointsEvent int32 `json:"ruleEndpointsEvent,omitempty"`
	// QueryPersistentVolume indicates the buffer of query persistent volume
	// default 1024
	QueryPersistentVolume int32 `json:"queryPersistentVolume,omitempty"`
	// QueryPersistentVolumeClaim indicates the buffer of query persistent volume claim
	// default 1024
	QueryPersistentVolumeClaim int32 `json:"queryPersistentVolumeClaim,omitempty"`
	// QueryVolumeAttachment indicates the buffer of query volume attachment
	// default 1024
	QueryVolumeAttachment int32 `json:"queryVolumeAttachment,omitempty"`
	// CreateNode indicates the buffer of create node
	// default 1024
	CreateNode int32 `json:"createNode,omitempty"`
	// PatchNode indicates the buffer of patch node
	// default 1024
	PatchNode int32 `json:"patchNode,omitempty"`
	// QueryNode indicates the buffer of query node
	// default 1024
	QueryNode int32 `json:"queryNode,omitempty"`
	// UpdateNode indicates the buffer of update node
	// default 1024
	UpdateNode int32 `json:"updateNode,omitempty"`
	// PatchPod indicates the buffer of patch pod
	// default 1024
	PatchPod int32 `json:"patchPod,omitempty"`
	// DeletePod indicates the buffer of delete pod message from edge
	// default 1024
	DeletePod int32 `json:"deletePod,omitempty"`
	// CreateLease indicates the buffer of create lease message from edge
	// default 1024
	CreateLease int32 `json:"createLease,omitempty"`
	// QueryLease indicates the buffer of query lease message from edge
	// default 1024
	QueryLease int32 `json:"queryLease,omitempty"`
	// ServiceAccount indicates the buffer of service account token
	// default 1024
	ServiceAccountToken int32 `json:"serviceAccountToken,omitempty"`
	// CreatePod indicates the buffer of create pod
	// default 1024
	CreatePod int32 `json:"createPod,omitempty"`
	// CertificateSigningRequest indicates the buffer of certificatesSigningRequest
	// default 1024
	CertificateSigningRequest int32 `json:"certificateSigningRequest,omitempty"`
}

// EdgeControllerLoad indicates the EdgeController load
type EdgeControllerLoad struct {
	// ProcessEventWorkers indicates the load of process event workers
	// default 4
	ProcessEventWorkers int32 `json:"processEventWorkers,omitempty"`
	// UpdatePodStatusWorkers indicates the load of update pod status workers
	// default 1
	UpdatePodStatusWorkers int32 `json:"updatePodStatusWorkers,omitempty"`
	// UpdateNodeStatusWorkers indicates the load of update node status workers
	// default 1
	UpdateNodeStatusWorkers int32 `json:"updateNodeStatusWorkers,omitempty"`
	// QueryConfigMapWorkers indicates the load of query config map workers
	// default 4
	QueryConfigMapWorkers int32 `json:"queryConfigMapWorkers,omitempty"`
	// QuerySecretWorkers indicates the load of query secret workers
	// default 4
	QuerySecretWorkers int32 `json:"querySecretWorkers,omitempty"`
	// QueryPersistentVolumeWorkers indicates the load of query persistent volume workers
	// default 4
	QueryPersistentVolumeWorkers int32 `json:"queryPersistentVolumeWorkers,omitempty"`
	// QueryPersistentVolumeClaimWorkers indicates the load of query persistent volume claim workers
	// default 4
	QueryPersistentVolumeClaimWorkers int32 `json:"queryPersistentVolumeClaimWorkers,omitempty"`
	// QueryVolumeAttachmentWorkers indicates the load of query volume attachment workers
	// default 4
	QueryVolumeAttachmentWorkers int32 `json:"queryVolumeAttachmentWorkers,omitempty"`
	// CreateNodeWorkers indicates the load of create node workers
	// default 4
	CreateNodeWorkers int32 `json:"createNodeWorkers,omitempty"`
	// PatchNodeWorkers indicates the load of patch node workers
	// default 4
	PatchNodeWorkers int32 `json:"patchNodeWorkers,omitempty"`
	// QueryNodeWorkers indicates the load of query node workers
	// default 4
	QueryNodeWorkers int32 `json:"queryNodeWorkers,omitempty"`
	// UpdateNodeWorkers indicates the load of update node workers
	// default 4
	UpdateNodeWorkers int32 `json:"updateNodeWorkers,omitempty"`
	// PatchPodWorkers indicates the load of patch pod workers
	// default 4
	PatchPodWorkers int32 `json:"patchPodWorkers,omitempty"`
	// DeletePodWorkers indicates the load of delete pod workers
	// default 4
	DeletePodWorkers int32 `json:"deletePodWorkers,omitempty"`
	// CreateLeaseWorkers indicates the load of create lease workers
	// default 4
	CreateLeaseWorkers int32 `json:"createLeaseWorkers,omitempty"`
	// QueryLeaseWorkers indicates the load of query lease workers
	// default 4
	QueryLeaseWorkers int32 `json:"queryLeaseWorkers,omitempty"`
	// UpdateRuleStatusWorkers indicates the load of update rule status
	// default 4
	UpdateRuleStatusWorkers int32 `json:"UpdateRuleStatusWorkers,omitempty"`
	// ServiceAccountTokenWorkers indicates the load of service account token
	// default 4
	ServiceAccountTokenWorkers int32 `json:"ServiceAccountTokenWorkers,omitempty"`
	// CreatePodWorks indicates the load of create pod
	// default 4
	CreatePodWorks int32 `json:"CreatePodWorks,omitempty"`
	// CertificateSigningRequestWorkers indicates the load of CertificateSigningRequest
	// default 4
	CertificateSigningRequestWorkers int32 `json:"certificateSigningRequestWorkers,omitempty"`
}

// DeviceController indicates the device controller
type DeviceController struct {
	// Enable indicates whether deviceController is enabled,
	// if set to false (for debugging etc.), skip checking other deviceController configs.
	// default true
	Enable bool `json:"enable"`
	// Buffer indicates Device controller buffer
	Buffer *DeviceControllerBuffer `json:"buffer,omitempty"`
	// Load indicates DeviceController Load
	Load *DeviceControllerLoad `json:"load,omitempty"`
}

// DeviceControllerBuffer indicates deviceController buffer
type DeviceControllerBuffer struct {
	// UpdateDeviceTwins indicates the buffer of update device twins
	// default 1024
	UpdateDeviceTwins int32 `json:"updateDeviceTwins,omitempty"`
	// UpdateDeviceStates indicates the buffer of update device states
	// default 1024
	UpdateDeviceStates int32 `json:"updateDeviceStatus,omitempty"`
	// DeviceEvent indicates the buffer of device event
	// default 1
	DeviceEvent int32 `json:"deviceEvent,omitempty"`
	// DeviceModelEvent indicates the buffer of device model event
	// default 1
	DeviceModelEvent int32 `json:"deviceModelEvent,omitempty"`
}

// DeviceControllerLoad indicates the deviceController load
type DeviceControllerLoad struct {
	// UpdateDeviceStatusWorkers indicates the load of update device status workers
	// default 1
	UpdateDeviceStatusWorkers int32 `json:"updateDeviceStatusWorkers,omitempty"`
}

// TaskManager indicates the operations controller
type TaskManager struct {
	// Enable indicates whether TaskManager is enabled,
	// if set to false (for debugging etc.), skip checking other TaskManager configs.
	// default false
	Enable bool `json:"enable"`
	// Buffer indicates Operation Controller buffer
	Buffer *TaskManagerBuffer `json:"buffer,omitempty"`
	// Load indicates Operation Controller Load
	Load *TaskManagerLoad `json:"load,omitempty"`
}

// TaskManagerBuffer indicates TaskManager buffer
type TaskManagerBuffer struct {
	// TaskStatus indicates the buffer of update NodeUpgradeJob status
	// default 1024
	TaskStatus int32 `json:"taskStatus,omitempty"`
	// TaskEvent indicates the buffer of NodeUpgradeJob event
	// default 1
	TaskEvent int32 `json:"taskEvent,omitempty"`
}

// TaskManagerLoad indicates the TaskManager load
type TaskManagerLoad struct {
	// TaskWorkers indicates the load of update NodeUpgradeJob workers
	// default 1
	TaskWorkers int32 `json:"taskWorkers,omitempty"`
}

// ImagePrePullController indicates the operations controller
type ImagePrePullController struct {
	// Enable indicates whether ImagePrePullController is enabled,
	// if set to false (for debugging etc.), skip checking other ImagePrePullController configs.
	// default false
	Enable bool `json:"enable"`
	// Buffer indicates Operation Controller buffer
	Buffer *ImagePrePullControllerBuffer `json:"buffer,omitempty"`
	// Load indicates Operation Controller Load
	Load *ImagePrePullControllerLoad `json:"load,omitempty"`
}

// ImagePrePullControllerBuffer indicates ImagePrePullController buffer
type ImagePrePullControllerBuffer struct {
	// ImagePrePullJobStatus indicates the buffer of update ImagePrePullJob status
	// default 1024
	ImagePrePullJobStatus int32 `json:"imagePrePullJobStatus,omitempty"`
	// ImagePrePullJobEvent indicates the buffer of ImagePrePullJob event
	// default 1
	ImagePrePullJobEvent int32 `json:"imagePrePullJobEvent,omitempty"`
}

// ImagePrePullControllerLoad indicates the ImagePrePullController load
type ImagePrePullControllerLoad struct {
	// ImagePrePullJobWorkers indicates the load of update ImagePrePullJob workers
	// default 1
	ImagePrePullJobWorkers int32 `json:"imagePrePullJobWorkers,omitempty"`
}

// SyncController indicates the sync controller
type SyncController struct {
	// Enable indicates whether syncController is enabled,
	// if set to false (for debugging etc.), skip checking other syncController configs.
	// default true
	Enable bool `json:"enable"`
}

// DynamicController indicates the dynamic controller
type DynamicController struct {
	// Enable indicates whether dynamicController is enabled,
	// if set to false (for debugging etc.), skip checking other dynamicController configs.
	// default true
	Enable bool `json:"enable"`
}

// CloudStream indicates the stream controller
type CloudStream struct {
	// Enable indicates whether cloudstream is enabled, if set to false (for debugging etc.), skip checking other configs.
	// default true
	Enable bool `json:"enable"`

	// TLSTunnelCAFile indicates ca file path
	// default /etc/kubeedge/ca/rootCA.crt
	TLSTunnelCAFile string `json:"tlsTunnelCAFile,omitempty"`
	// TLSTunnelCertFile indicates cert file path
	// default /etc/kubeedge/certs/server.crt
	TLSTunnelCertFile string `json:"tlsTunnelCertFile,omitempty"`
	// TLSTunnelPrivateKeyFile indicates key file path
	// default /etc/kubeedge/certs/server.key
	TLSTunnelPrivateKeyFile string `json:"tlsTunnelPrivateKeyFile,omitempty"`
	// TunnelPort set open port for tunnel server
	// default 10004
	TunnelPort uint32 `json:"tunnelPort,omitempty"`

	// TLSStreamCAFile indicates kube-apiserver ca file path
	// default /etc/kubeedge/ca/streamCA.crt
	TLSStreamCAFile string `json:"tlsStreamCAFile,omitempty"`
	// TLSStreamCertFile indicates cert file path
	// default /etc/kubeedge/certs/stream.crt
	TLSStreamCertFile string `json:"tlsStreamCertFile,omitempty"`
	// TLSStreamPrivateKeyFile indicates key file path
	// default /etc/kubeedge/certs/stream.key
	TLSStreamPrivateKeyFile string `json:"tlsStreamPrivateKeyFile,omitempty"`
	// StreamPort set open port for stream server
	// default 10003
	StreamPort uint32 `json:"streamPort,omitempty"`
}

type Router struct {
	// default true
	Enable      bool   `json:"enable"`
	Address     string `json:"address,omitempty"`
	Port        uint32 `json:"port,omitempty"`
	RestTimeout uint32 `json:"restTimeout,omitempty"`
}

// IptablesManager indicates the config of Iptables
type IptablesManager struct {
	// Enable indicates whether enable IptablesManager
	// default true
	Enable bool `json:"enable"`
	// It indicates how the component is deployed, valid mode can use "internal" or "external".
	// The iptables manager component with the internal mode is always deployed inside the cloudcore, will share the host network, forward to the internal port of the tunnel port.
	// The iptables manager component with the external mode is always deployed outside the cloudcore, will share the host network, forward to the internal cloudcore service and port.
	// default internal.
	// +kubebuilder:validation:Enum=internal;external
	Mode IptablesMgrMode `json:"mode,omitempty"`
}
