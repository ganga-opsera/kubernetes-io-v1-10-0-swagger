package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Iok8skubernetespkgapiv1Capabilities represents the Iok8skubernetespkgapiv1Capabilities schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Capabilities struct {
	Add []string `json:"add,omitempty"` // Added capabilities
	Drop []string `json:"drop,omitempty"` // Removed capabilities
}

// Iok8skubernetespkgapisrbacv1beta1Subject represents the Iok8skubernetespkgapisrbacv1beta1Subject schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1Subject struct {
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	Kind string `json:"kind"` // Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Name string `json:"name"` // Name of the object being referenced.
	Namespace string `json:"namespace,omitempty"` // Namespace of the referenced object. If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
}

// Iok8sapinetworkingv1NetworkPolicyEgressRule represents the Iok8sapinetworkingv1NetworkPolicyEgressRule schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyEgressRule struct {
	Ports []Iok8sapinetworkingv1NetworkPolicyPort `json:"ports,omitempty"` // List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
	To []Iok8sapinetworkingv1NetworkPolicyPeer `json:"to,omitempty"` // List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.
}

// Iok8sapicorev1Pod represents the Iok8sapicorev1Pod schema from the OpenAPI specification
type Iok8sapicorev1Pod struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PodSpec `json:"spec,omitempty"` // PodSpec is a description of a pod.
	Status Iok8sapicorev1PodStatus `json:"status,omitempty"` // PodStatus represents information about the status of a pod. Status may trail the actual state of a system.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta1DeploymentCondition represents the Iok8sapiappsv1beta1DeploymentCondition schema from the OpenAPI specification
type Iok8sapiappsv1beta1DeploymentCondition struct {
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
}

// Iok8skubernetespkgapiv1AzureFileVolumeSource represents the Iok8skubernetespkgapiv1AzureFileVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1AzureFileVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretname string `json:"secretName"` // the name of secret that contains Azure Storage Account Name and Key
	Sharename string `json:"shareName"` // Share Name
}

// Iok8skubernetespkgapiv1EndpointAddress represents the Iok8skubernetespkgapiv1EndpointAddress schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EndpointAddress struct {
	Ip string `json:"ip"` // The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready.
	Nodename string `json:"nodeName,omitempty"` // Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
	Targetref Iok8sapicorev1ObjectReference `json:"targetRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Hostname string `json:"hostname,omitempty"` // The Hostname of this endpoint
}

// Iok8skubernetespkgapisstoragev1beta1StorageClass represents the Iok8skubernetespkgapisstoragev1beta1StorageClass schema from the OpenAPI specification
type Iok8skubernetespkgapisstoragev1beta1StorageClass struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Volumebindingmode string `json:"volumeBindingMode,omitempty"` // VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is alpha-level and is only honored by servers that enable the VolumeScheduling feature.
	Allowvolumeexpansion bool `json:"allowVolumeExpansion,omitempty"` // AllowVolumeExpansion shows whether the storage class allow volume expand
	Mountoptions []string `json:"mountOptions,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	Parameters map[string]interface{} `json:"parameters,omitempty"` // Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Provisioner string `json:"provisioner"` // Provisioner indicates the type of the provisioner.
	Reclaimpolicy string `json:"reclaimPolicy,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapirbacv1alpha1AggregationRule represents the Iok8sapirbacv1alpha1AggregationRule schema from the OpenAPI specification
type Iok8sapirbacv1alpha1AggregationRule struct {
	Clusterroleselectors []Iok8sapimachinerypkgapismetav1LabelSelector `json:"clusterRoleSelectors,omitempty"` // ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
}

// Iok8sapirbacv1Subject represents the Iok8sapirbacv1Subject schema from the OpenAPI specification
type Iok8sapirbacv1Subject struct {
	Namespace string `json:"namespace,omitempty"` // Namespace of the referenced object. If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	Kind string `json:"kind"` // Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Name string `json:"name"` // Name of the object being referenced.
}

// Iok8sapiappsv1beta2ControllerRevisionList represents the Iok8sapiappsv1beta2ControllerRevisionList schema from the OpenAPI specification
type Iok8sapiappsv1beta2ControllerRevisionList struct {
	Items []Iok8sapiappsv1beta2ControllerRevision `json:"items"` // Items is the list of ControllerRevisions
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Statusreplicaspath string `json:"statusReplicasPath"` // StatusReplicasPath defines the JSON path inside of a CustomResource that corresponds to Scale.Status.Replicas. Only JSON paths without the array notation are allowed. Must be a JSON Path under .status. If there is no value under the given path in the CustomResource, the status replica value in the /scale subresource will default to 0.
	Labelselectorpath string `json:"labelSelectorPath,omitempty"` // LabelSelectorPath defines the JSON path inside of a CustomResource that corresponds to Scale.Status.Selector. Only JSON paths without the array notation are allowed. Must be a JSON Path under .status. Must be set to work with HPA. If there is no value under the given path in the CustomResource, the status label selector value in the /scale subresource will default to the empty string.
	Specreplicaspath string `json:"specReplicasPath"` // SpecReplicasPath defines the JSON path inside of a CustomResource that corresponds to Scale.Spec.Replicas. Only JSON paths without the array notation are allowed. Must be a JSON Path under .spec. If there is no value under the given path in the CustomResource, the /scale subresource will return an error on GET.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown.
	TypeField string `json:"type"` // Type is the type of the condition.
}

// Iok8sapipolicyv1beta1PodDisruptionBudget represents the Iok8sapipolicyv1beta1PodDisruptionBudget schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodDisruptionBudget struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapipolicyv1beta1PodDisruptionBudgetSpec `json:"spec,omitempty"` // PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
	Status Iok8sapipolicyv1beta1PodDisruptionBudgetStatus `json:"status,omitempty"` // PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapibatchv1beta1CronJobStatus represents the Iok8sapibatchv1beta1CronJobStatus schema from the OpenAPI specification
type Iok8sapibatchv1beta1CronJobStatus struct {
	Lastscheduletime string `json:"lastScheduleTime,omitempty"`
	Active []Iok8sapicorev1ObjectReference `json:"active,omitempty"` // A list of pointers to currently running jobs.
}

// Iok8skubernetespkgapiv1LimitRangeItem represents the Iok8skubernetespkgapiv1LimitRangeItem schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LimitRangeItem struct {
	TypeField string `json:"type,omitempty"` // Type of resource that this limit applies to.
	DefaultField map[string]interface{} `json:"default,omitempty"` // Default resource requirement limit value by resource name if resource limit is omitted.
	Defaultrequest map[string]interface{} `json:"defaultRequest,omitempty"` // DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
	Max map[string]interface{} `json:"max,omitempty"` // Max usage constraints on this kind by resource name.
	Maxlimitrequestratio map[string]interface{} `json:"maxLimitRequestRatio,omitempty"` // MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
	Min map[string]interface{} `json:"min,omitempty"` // Min usage constraints on this kind by resource name.
}

// Iok8sapicorev1SecretProjection represents the Iok8sapicorev1SecretProjection schema from the OpenAPI specification
type Iok8sapicorev1SecretProjection struct {
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

// Iok8sapipolicyv1beta1IDRange represents the Iok8sapipolicyv1beta1IDRange schema from the OpenAPI specification
type Iok8sapipolicyv1beta1IDRange struct {
	Max int64 `json:"max"` // Max is the end of the range, inclusive.
	Min int64 `json:"min"` // Min is the start of the range, inclusive.
}

// Iok8sapirbacv1beta1ClusterRoleBindingList represents the Iok8sapirbacv1beta1ClusterRoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1beta1ClusterRoleBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1ClusterRoleBinding `json:"items"` // Items is a list of ClusterRoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisappsv1beta1ScaleSpec represents the Iok8skubernetespkgapisappsv1beta1ScaleSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8skubernetespkgapiv1PodCondition represents the Iok8skubernetespkgapiv1PodCondition schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	TypeField string `json:"type"` // Type is the type of the condition. Currently only Ready. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Lastprobetime string `json:"lastProbeTime,omitempty"`
}

// Iok8sapiauthenticationv1beta1TokenReviewSpec represents the Iok8sapiauthenticationv1beta1TokenReviewSpec schema from the OpenAPI specification
type Iok8sapiauthenticationv1beta1TokenReviewSpec struct {
	Token string `json:"token,omitempty"` // Token is the opaque bearer token.
}

// Iok8sapicorev1Toleration represents the Iok8sapicorev1Toleration schema from the OpenAPI specification
type Iok8sapicorev1Toleration struct {
	Tolerationseconds int64 `json:"tolerationSeconds,omitempty"` // TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	Value string `json:"value,omitempty"` // Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
	Effect string `json:"effect,omitempty"` // Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	Key string `json:"key,omitempty"` // Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	Operator string `json:"operator,omitempty"` // Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
}

// Iok8skubernetespkgapiv1RBDVolumeSource represents the Iok8skubernetespkgapiv1RBDVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1RBDVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	Image string `json:"image"` // The rados image name. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"` // Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Monitors []string `json:"monitors"` // A collection of Ceph monitors. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"` // The rados pool name. Default is rbd. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	User string `json:"user,omitempty"` // The rados user name. Default is admin. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
}

// Iok8skubernetespkgapisextensionsv1beta1RunAsUserStrategyOptions represents the Iok8skubernetespkgapisextensionsv1beta1RunAsUserStrategyOptions schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1RunAsUserStrategyOptions struct {
	Ranges []Iok8sapiextensionsv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of uids that may be used.
	Rule string `json:"rule"` // Rule is the strategy that will dictate the allowable RunAsUser values that may be set.
}

// Iok8skubernetespkgapisadmissionregistrationv1alpha1InitializerConfigurationList represents the Iok8skubernetespkgapisadmissionregistrationv1alpha1InitializerConfigurationList schema from the OpenAPI specification
type Iok8skubernetespkgapisadmissionregistrationv1alpha1InitializerConfigurationList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1alpha1InitializerConfiguration `json:"items"` // List of InitializerConfiguration.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisauthorizationv1SelfSubjectAccessReviewSpec represents the Iok8skubernetespkgapisauthorizationv1SelfSubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1SelfSubjectAccessReviewSpec struct {
	Nonresourceattributes Iok8sapiauthorizationv1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
}

// Iok8sapiappsv1DaemonSet represents the Iok8sapiappsv1DaemonSet schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1DaemonSetSpec `json:"spec,omitempty"` // DaemonSetSpec is the specification of a daemon set.
	Status Iok8sapiappsv1DaemonSetStatus `json:"status,omitempty"` // DaemonSetStatus represents the current status of a daemon set.
}

// Iok8sapiautoscalingv1Scale represents the Iok8sapiautoscalingv1Scale schema from the OpenAPI specification
type Iok8sapiautoscalingv1Scale struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiautoscalingv1ScaleSpec `json:"spec,omitempty"` // ScaleSpec describes the attributes of a scale subresource.
	Status Iok8sapiautoscalingv1ScaleStatus `json:"status,omitempty"` // ScaleStatus represents the current status of a scale subresource.
}

// Iok8sapiappsv1beta1DeploymentStrategy represents the Iok8sapiappsv1beta1DeploymentStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta1DeploymentStrategy struct {
	Rollingupdate Iok8sapiappsv1beta1RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
}

// Iok8sapistoragev1alpha1VolumeAttachmentList represents the Iok8sapistoragev1alpha1VolumeAttachmentList schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttachmentList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapistoragev1alpha1VolumeAttachment `json:"items"` // Items is the list of VolumeAttachments
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapibatchv2alpha1CronJobSpec represents the Iok8sapibatchv2alpha1CronJobSpec schema from the OpenAPI specification
type Iok8sapibatchv2alpha1CronJobSpec struct {
	Successfuljobshistorylimit int `json:"successfulJobsHistoryLimit,omitempty"` // The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	Suspend bool `json:"suspend,omitempty"` // This flag tells the controller to suspend subsequent executions, it does not apply to already started executions. Defaults to false.
	Concurrencypolicy string `json:"concurrencyPolicy,omitempty"` // Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	Failedjobshistorylimit int `json:"failedJobsHistoryLimit,omitempty"` // The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	Jobtemplate Iok8sapibatchv2alpha1JobTemplateSpec `json:"jobTemplate"` // JobTemplateSpec describes the data a Job should have when created from a template
	Schedule string `json:"schedule"` // The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Startingdeadlineseconds int64 `json:"startingDeadlineSeconds,omitempty"` // Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.
}

// Iok8skubernetespkgapisextensionsv1beta1Ingress represents the Iok8skubernetespkgapisextensionsv1beta1Ingress schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1Ingress struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1IngressSpec `json:"spec,omitempty"` // IngressSpec describes the Ingress the user wishes to exist.
	Status Iok8sapiextensionsv1beta1IngressStatus `json:"status,omitempty"` // IngressStatus describe the current state of the Ingress.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiadmissionregistrationv1alpha1InitializerConfiguration represents the Iok8sapiadmissionregistrationv1alpha1InitializerConfiguration schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1InitializerConfiguration struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Initializers []Iok8sapiadmissionregistrationv1alpha1Initializer `json:"initializers,omitempty"` // Initializers is a list of resources and their default initializers Order-sensitive. When merging multiple InitializerConfigurations, we sort the initializers from different InitializerConfigurations by the name of the InitializerConfigurations; the order of the initializers from the same InitializerConfiguration is preserved.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapiv1Node represents the Iok8skubernetespkgapiv1Node schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Node struct {
	Status Iok8sapicorev1NodeStatus `json:"status,omitempty"` // NodeStatus is information about the current status of a node.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1NodeSpec `json:"spec,omitempty"` // NodeSpec describes the attributes that a node is created with.
}

// Iok8skubernetespkgapisextensionsv1beta1ScaleSpec represents the Iok8skubernetespkgapisextensionsv1beta1ScaleSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8skubernetespkgapiv1ObjectFieldSelector represents the Iok8skubernetespkgapiv1ObjectFieldSelector schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ObjectFieldSelector struct {
	Apiversion string `json:"apiVersion,omitempty"` // Version of the schema the FieldPath is written in terms of, defaults to "v1".
	Fieldpath string `json:"fieldPath"` // Path of the field to select in the specified API version.
}

// Iok8skubernetespkgapiv1ReplicationController represents the Iok8skubernetespkgapiv1ReplicationController schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ReplicationController struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ReplicationControllerSpec `json:"spec,omitempty"` // ReplicationControllerSpec is the specification of a replication controller.
	Status Iok8sapicorev1ReplicationControllerStatus `json:"status,omitempty"` // ReplicationControllerStatus represents the current status of a replication controller.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1NodeList represents the Iok8skubernetespkgapiv1NodeList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeList struct {
	Items []Iok8sapicorev1Node `json:"items"` // List of nodes
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1Endpoints represents the Iok8sapicorev1Endpoints schema from the OpenAPI specification
type Iok8sapicorev1Endpoints struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Subsets []Iok8sapicorev1EndpointSubset `json:"subsets,omitempty"` // The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1Taint represents the Iok8skubernetespkgapiv1Taint schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Taint struct {
	Effect string `json:"effect"` // Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Key string `json:"key"` // Required. The taint key to be applied to a node.
	Timeadded string `json:"timeAdded,omitempty"`
	Value string `json:"value,omitempty"` // Required. The taint value corresponding to the taint key.
}

// Iok8sapiappsv1DeploymentStatus represents the Iok8sapiappsv1DeploymentStatus schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentStatus struct {
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Conditions []Iok8sapiappsv1DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of ready pods targeted by this deployment.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
}

// Iok8skubernetespkgapiv1ExecAction represents the Iok8skubernetespkgapiv1ExecAction schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ExecAction struct {
	Command []string `json:"command,omitempty"` // Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
}

// Iok8skubernetespkgapisauthorizationv1SubjectAccessReviewStatus represents the Iok8skubernetespkgapisauthorizationv1SubjectAccessReviewStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1SubjectAccessReviewStatus struct {
	Allowed bool `json:"allowed"` // Allowed is required. True if the action would be allowed, false otherwise.
	Denied bool `json:"denied,omitempty"` // Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	Reason string `json:"reason,omitempty"` // Reason is optional. It indicates why a request was allowed or denied.
}

// Iok8sapischedulingv1alpha1PriorityClass represents the Iok8sapischedulingv1alpha1PriorityClass schema from the OpenAPI specification
type Iok8sapischedulingv1alpha1PriorityClass struct {
	Globaldefault bool `json:"globalDefault,omitempty"` // globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Value int `json:"value"` // The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Description string `json:"description,omitempty"` // description is an arbitrary string that usually provides guidelines on when this priority class should be used.
}

// Iok8sapicorev1PodTemplateSpec represents the Iok8sapicorev1PodTemplateSpec schema from the OpenAPI specification
type Iok8sapicorev1PodTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PodSpec `json:"spec,omitempty"` // PodSpec is a description of a pod.
}

// Iok8sapicorev1SecretReference represents the Iok8sapicorev1SecretReference schema from the OpenAPI specification
type Iok8sapicorev1SecretReference struct {
	Name string `json:"name,omitempty"` // Name is unique within a namespace to reference a secret resource.
	Namespace string `json:"namespace,omitempty"` // Namespace defines the space within which the secret name must be unique.
}

// Iok8sapibatchv1JobCondition represents the Iok8sapibatchv1JobCondition schema from the OpenAPI specification
type Iok8sapibatchv1JobCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // (brief) reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of job condition, Complete or Failed.
	Lastprobetime string `json:"lastProbeTime,omitempty"`
}

// Iok8skubernetespkgapisbatchv2alpha1JobTemplateSpec represents the Iok8skubernetespkgapisbatchv2alpha1JobTemplateSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv2alpha1JobTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
}

// Iok8skubernetespkgapisappsv1beta1StatefulSetUpdateStrategy represents the Iok8skubernetespkgapisappsv1beta1StatefulSetUpdateStrategy schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1StatefulSetUpdateStrategy struct {
	TypeField string `json:"type,omitempty"` // Type indicates the type of the StatefulSetUpdateStrategy.
	Rollingupdate Iok8sapiappsv1beta1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"` // RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
}

// Iok8sapiauthorizationv1NonResourceRule represents the Iok8sapiauthorizationv1NonResourceRule schema from the OpenAPI specification
type Iok8sapiauthorizationv1NonResourceRule struct {
	Verbs []string `json:"verbs"` // Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options. "*" means all.
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path. "*" means all.
}

// Iok8skubernetespkgapiv1GitRepoVolumeSource represents the Iok8skubernetespkgapiv1GitRepoVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1GitRepoVolumeSource struct {
	Revision string `json:"revision,omitempty"` // Commit hash for the specified revision.
	Directory string `json:"directory,omitempty"` // Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	Repository string `json:"repository"` // Repository URL
}

// Iok8sapiappsv1DaemonSetUpdateStrategy represents the Iok8sapiappsv1DaemonSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetUpdateStrategy struct {
	Rollingupdate Iok8sapiappsv1RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of daemon set rolling update.
	TypeField string `json:"type,omitempty"` // Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
}

// Iok8skubernetespkgapisbatchv1Job represents the Iok8skubernetespkgapisbatchv1Job schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv1Job struct {
	Status Iok8sapibatchv1JobStatus `json:"status,omitempty"` // JobStatus represents the current state of a Job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
}

// Iok8sapicorev1ISCSIPersistentVolumeSource represents the Iok8sapicorev1ISCSIPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ISCSIPersistentVolumeSource struct {
	Iscsiinterface string `json:"iscsiInterface,omitempty"` // iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Lun int `json:"lun"` // iSCSI Target Lun number.
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	Targetportal string `json:"targetPortal"` // iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Chapauthdiscovery bool `json:"chapAuthDiscovery,omitempty"` // whether support iSCSI Discovery CHAP authentication
	Chapauthsession bool `json:"chapAuthSession,omitempty"` // whether support iSCSI Session CHAP authentication
	Initiatorname string `json:"initiatorName,omitempty"` // Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	Portals []string `json:"portals,omitempty"` // iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	Iqn string `json:"iqn"` // Target iSCSI Qualified Name.
}

// Iok8sapicorev1PodSecurityContext represents the Iok8sapicorev1PodSecurityContext schema from the OpenAPI specification
type Iok8sapicorev1PodSecurityContext struct {
	Runasnonroot bool `json:"runAsNonRoot,omitempty"` // Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Runasuser int64 `json:"runAsUser,omitempty"` // The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
	Supplementalgroups []int64 `json:"supplementalGroups,omitempty"` // A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container.
	Fsgroup int64 `json:"fsGroup,omitempty"` // A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.
	Runasgroup int64 `json:"runAsGroup,omitempty"` // The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
}

// Iok8sapisettingsv1alpha1PodPresetSpec represents the Iok8sapisettingsv1alpha1PodPresetSpec schema from the OpenAPI specification
type Iok8sapisettingsv1alpha1PodPresetSpec struct {
	Volumes []Iok8sapicorev1Volume `json:"volumes,omitempty"` // Volumes defines the collection of Volume to inject into the pod.
	Env []Iok8sapicorev1EnvVar `json:"env,omitempty"` // Env defines the collection of EnvVar to inject into containers.
	Envfrom []Iok8sapicorev1EnvFromSource `json:"envFrom,omitempty"` // EnvFrom defines the collection of EnvFromSource to inject into containers.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Volumemounts []Iok8sapicorev1VolumeMount `json:"volumeMounts,omitempty"` // VolumeMounts defines the collection of VolumeMount to inject into containers.
}

// Iok8sapiauthorizationv1beta1LocalSubjectAccessReview represents the Iok8sapiauthorizationv1beta1LocalSubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1LocalSubjectAccessReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
}

// Iok8sapistoragev1alpha1VolumeError represents the Iok8sapistoragev1alpha1VolumeError schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeError struct {
	Message string `json:"message,omitempty"` // String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Time string `json:"time,omitempty"`
}

// Iok8sapicorev1PersistentVolumeClaim represents the Iok8sapicorev1PersistentVolumeClaim schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaim struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PersistentVolumeClaimSpec `json:"spec,omitempty"` // PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
	Status Iok8sapicorev1PersistentVolumeClaimStatus `json:"status,omitempty"` // PersistentVolumeClaimStatus is the current status of a persistent volume claim.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1NodeStatus represents the Iok8skubernetespkgapiv1NodeStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeStatus struct {
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Images []Iok8sapicorev1ContainerImage `json:"images,omitempty"` // List of container images on this node
	Volumesinuse []string `json:"volumesInUse,omitempty"` // List of attachable volumes in use (mounted) by the node.
	Daemonendpoints Iok8sapicorev1NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"` // NodeDaemonEndpoints lists ports opened by daemons running on the Node.
	Nodeinfo Iok8sapicorev1NodeSystemInfo `json:"nodeInfo,omitempty"` // NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
	Allocatable map[string]interface{} `json:"allocatable,omitempty"` // Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
	Volumesattached []Iok8sapicorev1AttachedVolume `json:"volumesAttached,omitempty"` // List of volumes that are attached to the node.
	Addresses []Iok8sapicorev1NodeAddress `json:"addresses,omitempty"` // List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	Conditions []Iok8sapicorev1NodeCondition `json:"conditions,omitempty"` // Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
	Phase string `json:"phase,omitempty"` // NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
}

// Iok8sapiextensionsv1beta1DaemonSetSpec represents the Iok8sapiextensionsv1beta1DaemonSetSpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DaemonSetSpec struct {
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Templategeneration int64 `json:"templateGeneration,omitempty"` // DEPRECATED. A sequence number representing a specific generation of the template. Populated by the system. It can be set only during the creation.
	Updatestrategy Iok8sapiextensionsv1beta1DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
}

// Iok8sapirbacv1RoleBinding represents the Iok8sapirbacv1RoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1RoleBinding struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Roleref Iok8sapirbacv1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1NodeConfigSource represents the Iok8sapicorev1NodeConfigSource schema from the OpenAPI specification
type Iok8sapicorev1NodeConfigSource struct {
	Configmapref Iok8sapicorev1ObjectReference `json:"configMapRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1DeploymentStatus represents the Iok8sapiextensionsv1beta1DeploymentStatus schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DeploymentStatus struct {
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of ready pods targeted by this deployment.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Conditions []Iok8sapiextensionsv1beta1DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
}

// Iok8sapiappsv1beta1StatefulSetUpdateStrategy represents the Iok8sapiappsv1beta1StatefulSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta1StatefulSetUpdateStrategy struct {
	TypeField string `json:"type,omitempty"` // Type indicates the type of the StatefulSetUpdateStrategy.
	Rollingupdate Iok8sapiappsv1beta1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"` // RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
}

// Iok8sapiappsv1StatefulSetStatus represents the Iok8sapiappsv1StatefulSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
	Readyreplicas int `json:"readyReplicas,omitempty"` // readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	Replicas int `json:"replicas"` // replicas is the number of Pods created by the StatefulSet controller.
	Updaterevision string `json:"updateRevision,omitempty"` // updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	Currentrevision string `json:"currentRevision,omitempty"` // currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
	Collisioncount int `json:"collisionCount,omitempty"` // collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Conditions []Iok8sapiappsv1StatefulSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a statefulset's current state.
	Currentreplicas int `json:"currentReplicas,omitempty"` // currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
}

// Iok8sapicorev1ConfigMap represents the Iok8sapicorev1ConfigMap schema from the OpenAPI specification
type Iok8sapicorev1ConfigMap struct {
	Data map[string]interface{} `json:"data,omitempty"` // Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Binarydata map[string]interface{} `json:"binaryData,omitempty"` // BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
}

// Iok8skubernetespkgapisextensionsv1beta1DaemonSetUpdateStrategy represents the Iok8skubernetespkgapisextensionsv1beta1DaemonSetUpdateStrategy schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DaemonSetUpdateStrategy struct {
	Rollingupdate Iok8sapiextensionsv1beta1RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of daemon set rolling update.
	TypeField string `json:"type,omitempty"` // Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is OnDelete.
}

// Iok8sapicertificatesv1beta1CertificateSigningRequestList represents the Iok8sapicertificatesv1beta1CertificateSigningRequestList schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1CertificateSigningRequestList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicertificatesv1beta1CertificateSigningRequest `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfiguration represents the Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfiguration schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfiguration struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Webhooks []Iok8sapiadmissionregistrationv1beta1Webhook `json:"webhooks,omitempty"` // Webhooks is a list of webhooks and the affected resources and operations.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyPeer represents the Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyPeer schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyPeer struct {
	Ipblock Iok8sapiextensionsv1beta1IPBlock `json:"ipBlock,omitempty"` // DEPRECATED 1.9 - This group version of IPBlock is deprecated by networking/v1/IPBlock. IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8skubernetespkgapiv1ConfigMapProjection represents the Iok8skubernetespkgapiv1ConfigMapProjection schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ConfigMapProjection struct {
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or it's keys must be defined
}

// Iok8sapiextensionsv1beta1RollingUpdateDaemonSet represents the Iok8sapiextensionsv1beta1RollingUpdateDaemonSet schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1RollingUpdateDaemonSet struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8sapiauthenticationv1TokenReview represents the Iok8sapiauthenticationv1TokenReview schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenReview struct {
	Spec Iok8sapiauthenticationv1TokenReviewSpec `json:"spec"` // TokenReviewSpec is a description of the token authentication request.
	Status Iok8sapiauthenticationv1TokenReviewStatus `json:"status,omitempty"` // TokenReviewStatus is the result of the token authentication request.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ResourceRequirements represents the Iok8sapicorev1ResourceRequirements schema from the OpenAPI specification
type Iok8sapicorev1ResourceRequirements struct {
	Requests map[string]interface{} `json:"requests,omitempty"` // Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Limits map[string]interface{} `json:"limits,omitempty"` // Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
}

// Iok8sapiextensionsv1beta1ReplicaSet represents the Iok8sapiextensionsv1beta1ReplicaSet schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ReplicaSet struct {
	Spec Iok8sapiextensionsv1beta1ReplicaSetSpec `json:"spec,omitempty"` // ReplicaSetSpec is the specification of a ReplicaSet.
	Status Iok8sapiextensionsv1beta1ReplicaSetStatus `json:"status,omitempty"` // ReplicaSetStatus represents the current status of a ReplicaSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ResourceQuotaSpec represents the Iok8sapicorev1ResourceQuotaSpec schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuotaSpec struct {
	Hard map[string]interface{} `json:"hard,omitempty"` // Hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Scopes []string `json:"scopes,omitempty"` // A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
}

// Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerStatus represents the Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerStatus struct {
	Conditions []Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerCondition `json:"conditions"` // conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
	Currentmetrics []Iok8sapiautoscalingv2beta1MetricStatus `json:"currentMetrics"` // currentMetrics is the last read state of the metrics used by this autoscaler.
	Currentreplicas int `json:"currentReplicas"` // currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	Desiredreplicas int `json:"desiredReplicas"` // desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
	Lastscaletime string `json:"lastScaleTime,omitempty"`
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed by this autoscaler.
}

// Iok8sapirbacv1beta1ClusterRole represents the Iok8sapirbacv1beta1ClusterRole schema from the OpenAPI specification
type Iok8sapirbacv1beta1ClusterRole struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1beta1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this ClusterRole
	Aggregationrule Iok8sapirbacv1beta1AggregationRule `json:"aggregationRule,omitempty"` // AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
}

// Iok8skubernetespkgapisrbacv1alpha1ClusterRoleBinding represents the Iok8skubernetespkgapisrbacv1alpha1ClusterRoleBinding schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1ClusterRoleBinding struct {
	Roleref Iok8sapirbacv1alpha1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1alpha1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1NamespaceList represents the Iok8sapicorev1NamespaceList schema from the OpenAPI specification
type Iok8sapicorev1NamespaceList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Namespace `json:"items"` // Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1StorageOSVolumeSource represents the Iok8sapicorev1StorageOSVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1StorageOSVolumeSource struct {
	Volumename string `json:"volumeName,omitempty"` // VolumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.
	Volumenamespace string `json:"volumeNamespace,omitempty"` // VolumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
}

// Iok8skubernetespkgapiv1AWSElasticBlockStoreVolumeSource represents the Iok8skubernetespkgapiv1AWSElasticBlockStoreVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1AWSElasticBlockStoreVolumeSource struct {
	Volumeid string `json:"volumeID"` // Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	Partition int `json:"partition,omitempty"` // The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Readonly bool `json:"readOnly,omitempty"` // Specify "true" to force and set the ReadOnly property in VolumeMounts to "true". If omitted, the default is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
}

// Iok8sapiextensionsv1beta1HTTPIngressPath represents the Iok8sapiextensionsv1beta1HTTPIngressPath schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1HTTPIngressPath struct {
	Path string `json:"path,omitempty"` // Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
	Backend Iok8sapiextensionsv1beta1IngressBackend `json:"backend"` // IngressBackend describes all endpoints for a given service and port.
}

// Iok8sapiextensionsv1beta1DaemonSet represents the Iok8sapiextensionsv1beta1DaemonSet schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DaemonSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1DaemonSetSpec `json:"spec,omitempty"` // DaemonSetSpec is the specification of a daemon set.
	Status Iok8sapiextensionsv1beta1DaemonSetStatus `json:"status,omitempty"` // DaemonSetStatus represents the current status of a daemon set.
}

// Iok8skubernetespkgapisappsv1beta1DeploymentStatus represents the Iok8skubernetespkgapisappsv1beta1DeploymentStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1DeploymentStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of ready pods targeted by this deployment.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Conditions []Iok8sapiappsv1beta1DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
}

// Iok8skubernetespkgapisbatchv1JobStatus represents the Iok8skubernetespkgapisbatchv1JobStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv1JobStatus struct {
	Succeeded int `json:"succeeded,omitempty"` // The number of pods which reached phase Succeeded.
	Active int `json:"active,omitempty"` // The number of actively running pods.
	Completiontime string `json:"completionTime,omitempty"`
	Conditions []Iok8sapibatchv1JobCondition `json:"conditions,omitempty"` // The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Failed int `json:"failed,omitempty"` // The number of pods which reached phase Failed.
	Starttime string `json:"startTime,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"` // Type is the type of the condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown.
}

// Iok8sapistoragev1beta1StorageClassList represents the Iok8sapistoragev1beta1StorageClassList schema from the OpenAPI specification
type Iok8sapistoragev1beta1StorageClassList struct {
	Items []Iok8sapistoragev1beta1StorageClass `json:"items"` // Items is the list of StorageClasses
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus represents the Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus struct {
	Allowed bool `json:"allowed"` // Allowed is required. True if the action would be allowed, false otherwise.
	Denied bool `json:"denied,omitempty"` // Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	Reason string `json:"reason,omitempty"` // Reason is optional. It indicates why a request was allowed or denied.
}

// Iok8sapistoragev1beta1VolumeAttachmentStatus represents the Iok8sapistoragev1beta1VolumeAttachmentStatus schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttachmentStatus struct {
	Attachmentmetadata map[string]interface{} `json:"attachmentMetadata,omitempty"` // Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Detacherror Iok8sapistoragev1beta1VolumeError `json:"detachError,omitempty"` // VolumeError captures an error encountered during a volume operation.
	Attacherror Iok8sapistoragev1beta1VolumeError `json:"attachError,omitempty"` // VolumeError captures an error encountered during a volume operation.
	Attached bool `json:"attached"` // Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
}

// Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyList represents the Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyList schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1NetworkPolicy `json:"items"` // Items is a list of schema objects.
}

// Iok8sapiadmissionregistrationv1alpha1Rule represents the Iok8sapiadmissionregistrationv1alpha1Rule schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1Rule struct {
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
	Apiversions []string `json:"apiVersions,omitempty"` // APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.
}

// Iok8skubernetespkgapisextensionsv1beta1IngressStatus represents the Iok8skubernetespkgapisextensionsv1beta1IngressStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IngressStatus struct {
	Loadbalancer Iok8sapicorev1LoadBalancerStatus `json:"loadBalancer,omitempty"` // LoadBalancerStatus represents the status of a load-balancer.
}

// Iok8sapicorev1ConfigMapProjection represents the Iok8sapicorev1ConfigMapProjection schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapProjection struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or it's keys must be defined
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
}

// Iok8skubernetespkgapiv1Handler represents the Iok8skubernetespkgapiv1Handler schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Handler struct {
	Exec Iok8sapicorev1ExecAction `json:"exec,omitempty"` // ExecAction describes a "run in container" action.
	Httpget Iok8sapicorev1HTTPGetAction `json:"httpGet,omitempty"` // HTTPGetAction describes an action based on HTTP Get requests.
	Tcpsocket Iok8sapicorev1TCPSocketAction `json:"tcpSocket,omitempty"` // TCPSocketAction describes an action based on opening a socket
}

// Iok8skubernetespkgapiv1GlusterfsVolumeSource represents the Iok8skubernetespkgapiv1GlusterfsVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1GlusterfsVolumeSource struct {
	Endpoints string `json:"endpoints"` // EndpointsName is the endpoint name that details Glusterfs topology. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path"` // Path is the Glusterfs volume path. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
}

// Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerSpec represents the Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerSpec struct {
	Scaletargetref Iok8sapiautoscalingv1CrossVersionObjectReference `json:"scaleTargetRef"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
	Targetcpuutilizationpercentage int `json:"targetCPUUtilizationPercentage,omitempty"` // target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
	Maxreplicas int `json:"maxReplicas"` // upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	Minreplicas int `json:"minReplicas,omitempty"` // lower limit for the number of pods that can be set by the autoscaler, default 1.
}

// Iok8skubernetespkgapiv1CinderVolumeSource represents the Iok8skubernetespkgapiv1CinderVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1CinderVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	Volumeid string `json:"volumeID"` // volume id used to identify the volume in cinder More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
}

// Iok8sapicorev1Probe represents the Iok8sapicorev1Probe schema from the OpenAPI specification
type Iok8sapicorev1Probe struct {
	Failurethreshold int `json:"failureThreshold,omitempty"` // Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
	Httpget Iok8sapicorev1HTTPGetAction `json:"httpGet,omitempty"` // HTTPGetAction describes an action based on HTTP Get requests.
	Initialdelayseconds int `json:"initialDelaySeconds,omitempty"` // Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	Periodseconds int `json:"periodSeconds,omitempty"` // How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
	Successthreshold int `json:"successThreshold,omitempty"` // Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	Tcpsocket Iok8sapicorev1TCPSocketAction `json:"tcpSocket,omitempty"` // TCPSocketAction describes an action based on opening a socket
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	Exec Iok8sapicorev1ExecAction `json:"exec,omitempty"` // ExecAction describes a "run in container" action.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status GeneratedType `json:"status,omitempty"` // APIServiceStatus contains derived information about an API server
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec GeneratedType `json:"spec,omitempty"` // APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
}

// Iok8sapicorev1ConfigMapKeySelector represents the Iok8sapicorev1ConfigMapKeySelector schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapKeySelector struct {
	Key string `json:"key"` // The key to select.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or it's key must be defined
}

// Iok8skubernetespkgapiv1SecretList represents the Iok8skubernetespkgapiv1SecretList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SecretList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Secret `json:"items"` // Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiextensionsv1beta1AllowedFlexVolume represents the Iok8sapiextensionsv1beta1AllowedFlexVolume schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1AllowedFlexVolume struct {
	Driver string `json:"driver"` // Driver is the name of the Flexvolume driver.
}

// Iok8sapicorev1PersistentVolumeClaimCondition represents the Iok8sapicorev1PersistentVolumeClaimCondition schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimCondition struct {
	TypeField string `json:"type"`
	Lastprobetime string `json:"lastProbeTime,omitempty"`
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "ResizeStarted" that means the underlying persistent volume is being resized.
	Status string `json:"status"`
}

// Iok8sapieventsv1beta1Event represents the Iok8sapieventsv1beta1Event schema from the OpenAPI specification
type Iok8sapieventsv1beta1Event struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Eventtime string `json:"eventTime"`
	Deprecatedcount int `json:"deprecatedCount,omitempty"` // Deprecated field assuring backward compatibility with core.v1 Event type
	Deprecatedsource Iok8sapicorev1EventSource `json:"deprecatedSource,omitempty"` // EventSource contains information for an event.
	Action string `json:"action,omitempty"` // What action was taken/failed regarding to the regarding object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Deprecatedfirsttimestamp string `json:"deprecatedFirstTimestamp,omitempty"`
	Deprecatedlasttimestamp string `json:"deprecatedLastTimestamp,omitempty"`
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Note string `json:"note,omitempty"` // Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Reason string `json:"reason,omitempty"` // Why the action was taken.
	Regarding Iok8sapicorev1ObjectReference `json:"regarding,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Reportingcontroller string `json:"reportingController,omitempty"` // Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	Reportinginstance string `json:"reportingInstance,omitempty"` // ID of the controller instance, e.g. `kubelet-xyzf`.
	Series Iok8sapieventsv1beta1EventSeries `json:"series,omitempty"` // EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
	TypeField string `json:"type,omitempty"` // Type of this event (Normal, Warning), new types could be added in the future.
	Related Iok8sapicorev1ObjectReference `json:"related,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
}

// Iok8sapirbacv1beta1RoleList represents the Iok8sapirbacv1beta1RoleList schema from the OpenAPI specification
type Iok8sapirbacv1beta1RoleList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1Role `json:"items"` // Items is a list of Roles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapibatchv2alpha1JobTemplateSpec represents the Iok8sapibatchv2alpha1JobTemplateSpec schema from the OpenAPI specification
type Iok8sapibatchv2alpha1JobTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
}

// Iok8sapimachinerypkgapismetav1LabelSelector represents the Iok8sapimachinerypkgapismetav1LabelSelector schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1LabelSelector struct {
	Matchlabels map[string]interface{} `json:"matchLabels,omitempty"` // matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	Matchexpressions []Iok8sapimachinerypkgapismetav1LabelSelectorRequirement `json:"matchExpressions,omitempty"` // matchExpressions is a list of label selector requirements. The requirements are ANDed.
}

// Iok8sapicorev1VolumeDevice represents the Iok8sapicorev1VolumeDevice schema from the OpenAPI specification
type Iok8sapicorev1VolumeDevice struct {
	Devicepath string `json:"devicePath"` // devicePath is the path inside of the container that the device will be mapped to.
	Name string `json:"name"` // name must match the name of a persistentVolumeClaim in the pod
}

// Iok8sapiauthorizationv1beta1SubjectRulesReviewStatus represents the Iok8sapiauthorizationv1beta1SubjectRulesReviewStatus schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SubjectRulesReviewStatus struct {
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
	Incomplete bool `json:"incomplete"` // Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.
	Nonresourcerules []Iok8sapiauthorizationv1beta1NonResourceRule `json:"nonResourceRules"` // NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	Resourcerules []Iok8sapiauthorizationv1beta1ResourceRule `json:"resourceRules"` // ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
}

// Iok8sapiextensionsv1beta1NetworkPolicyEgressRule represents the Iok8sapiextensionsv1beta1NetworkPolicyEgressRule schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicyEgressRule struct {
	Ports []Iok8sapiextensionsv1beta1NetworkPolicyPort `json:"ports,omitempty"` // List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
	To []Iok8sapiextensionsv1beta1NetworkPolicyPeer `json:"to,omitempty"` // List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.
}

// Iok8sapicorev1NodeDaemonEndpoints represents the Iok8sapicorev1NodeDaemonEndpoints schema from the OpenAPI specification
type Iok8sapicorev1NodeDaemonEndpoints struct {
	Kubeletendpoint Iok8sapicorev1DaemonEndpoint `json:"kubeletEndpoint,omitempty"` // DaemonEndpoint contains information about a single Daemon endpoint.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Raw string `json:"Raw"`
}

// Iok8sapistoragev1beta1StorageClass represents the Iok8sapistoragev1beta1StorageClass schema from the OpenAPI specification
type Iok8sapistoragev1beta1StorageClass struct {
	Provisioner string `json:"provisioner"` // Provisioner indicates the type of the provisioner.
	Reclaimpolicy string `json:"reclaimPolicy,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Volumebindingmode string `json:"volumeBindingMode,omitempty"` // VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is alpha-level and is only honored by servers that enable the VolumeScheduling feature.
	Allowvolumeexpansion bool `json:"allowVolumeExpansion,omitempty"` // AllowVolumeExpansion shows whether the storage class allow volume expand
	Mountoptions []string `json:"mountOptions,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	Parameters map[string]interface{} `json:"parameters,omitempty"` // Parameters holds the parameters for the provisioner that should create volumes of this storage class.
}

// Iok8sapirbacv1ClusterRole represents the Iok8sapirbacv1ClusterRole schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRole struct {
	Rules []Iok8sapirbacv1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this ClusterRole
	Aggregationrule Iok8sapirbacv1AggregationRule `json:"aggregationRule,omitempty"` // AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1NodeStatus represents the Iok8sapicorev1NodeStatus schema from the OpenAPI specification
type Iok8sapicorev1NodeStatus struct {
	Daemonendpoints Iok8sapicorev1NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"` // NodeDaemonEndpoints lists ports opened by daemons running on the Node.
	Nodeinfo Iok8sapicorev1NodeSystemInfo `json:"nodeInfo,omitempty"` // NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
	Allocatable map[string]interface{} `json:"allocatable,omitempty"` // Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
	Volumesattached []Iok8sapicorev1AttachedVolume `json:"volumesAttached,omitempty"` // List of volumes that are attached to the node.
	Addresses []Iok8sapicorev1NodeAddress `json:"addresses,omitempty"` // List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	Conditions []Iok8sapicorev1NodeCondition `json:"conditions,omitempty"` // Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
	Phase string `json:"phase,omitempty"` // NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Images []Iok8sapicorev1ContainerImage `json:"images,omitempty"` // List of container images on this node
	Volumesinuse []string `json:"volumesInUse,omitempty"` // List of attachable volumes in use (mounted) by the node.
}

// Iok8sapiappsv1beta2ControllerRevision represents the Iok8sapiappsv1beta2ControllerRevision schema from the OpenAPI specification
type Iok8sapiappsv1beta2ControllerRevision struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Revision int64 `json:"revision"` // Revision indicates the revision of the state represented by Data.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.Object `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // External package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.RawExtension `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // On the wire, the JSON will look something like this: { 	"kind":"MyAPIObject", 	"apiVersion":"v1", 	"myPlugin": { 		"kind":"PluginA", 		"aOption":"foo", 	}, } So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
}

// Iok8skubernetespkgapisstoragev1StorageClass represents the Iok8skubernetespkgapisstoragev1StorageClass schema from the OpenAPI specification
type Iok8skubernetespkgapisstoragev1StorageClass struct {
	Allowvolumeexpansion bool `json:"allowVolumeExpansion,omitempty"` // AllowVolumeExpansion shows whether the storage class allow volume expand
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Provisioner string `json:"provisioner"` // Provisioner indicates the type of the provisioner.
	Volumebindingmode string `json:"volumeBindingMode,omitempty"` // VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is alpha-level and is only honored by servers that enable the VolumeScheduling feature.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Mountoptions []string `json:"mountOptions,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	Parameters map[string]interface{} `json:"parameters,omitempty"` // Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Reclaimpolicy string `json:"reclaimPolicy,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapicorev1LimitRangeItem represents the Iok8sapicorev1LimitRangeItem schema from the OpenAPI specification
type Iok8sapicorev1LimitRangeItem struct {
	DefaultField map[string]interface{} `json:"default,omitempty"` // Default resource requirement limit value by resource name if resource limit is omitted.
	Defaultrequest map[string]interface{} `json:"defaultRequest,omitempty"` // DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
	Max map[string]interface{} `json:"max,omitempty"` // Max usage constraints on this kind by resource name.
	Maxlimitrequestratio map[string]interface{} `json:"maxLimitRequestRatio,omitempty"` // MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
	Min map[string]interface{} `json:"min,omitempty"` // Min usage constraints on this kind by resource name.
	TypeField string `json:"type,omitempty"` // Type of resource that this limit applies to.
}

// Iok8sapicorev1ReplicationControllerStatus represents the Iok8sapicorev1ReplicationControllerStatus schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerStatus struct {
	Conditions []Iok8sapicorev1ReplicationControllerCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replication controller's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replication controller.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed replication controller.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replication controller.
	Replicas int `json:"replicas"` // Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replication controller.
}

// Iok8sapiappsv1beta2ScaleSpec represents the Iok8sapiappsv1beta2ScaleSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta2ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8skubernetespkgapiv1AttachedVolume represents the Iok8skubernetespkgapiv1AttachedVolume schema from the OpenAPI specification
type Iok8skubernetespkgapiv1AttachedVolume struct {
	Devicepath string `json:"devicePath"` // DevicePath represents the device path where the volume should be available
	Name string `json:"name"` // Name of the attached volume
}

// Iok8sapiauthenticationv1beta1TokenReviewStatus represents the Iok8sapiauthenticationv1beta1TokenReviewStatus schema from the OpenAPI specification
type Iok8sapiauthenticationv1beta1TokenReviewStatus struct {
	User Iok8sapiauthenticationv1beta1UserInfo `json:"user,omitempty"` // UserInfo holds the information about the user needed to implement the user.Info interface.
	Authenticated bool `json:"authenticated,omitempty"` // Authenticated indicates that the token was associated with a known user.
	ErrorField string `json:"error,omitempty"` // Error indicates that the token couldn't be checked
}

// Iok8sapiappsv1beta1StatefulSetList represents the Iok8sapiappsv1beta1StatefulSetList schema from the OpenAPI specification
type Iok8sapiappsv1beta1StatefulSetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta1StatefulSet `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1RBDVolumeSource represents the Iok8sapicorev1RBDVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1RBDVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	User string `json:"user,omitempty"` // The rados user name. Default is admin. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	Image string `json:"image"` // The rados image name. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"` // Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Monitors []string `json:"monitors"` // A collection of Ceph monitors. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"` // The rados pool name. Default is rbd. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
}

// Iok8skubernetespkgapisrbacv1alpha1ClusterRoleBindingList represents the Iok8skubernetespkgapisrbacv1alpha1ClusterRoleBindingList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1ClusterRoleBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1ClusterRoleBinding `json:"items"` // Items is a list of ClusterRoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisappsv1beta1Scale represents the Iok8skubernetespkgapisappsv1beta1Scale schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1Scale struct {
	Status Iok8sapiappsv1beta1ScaleStatus `json:"status,omitempty"` // ScaleStatus represents the current status of a scale subresource.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta1ScaleSpec `json:"spec,omitempty"` // ScaleSpec describes the attributes of a scale subresource
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // Name is the name of the service
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the service
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	DefaultField GeneratedType `json:"default,omitempty"` // JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
	Format string `json:"format,omitempty"`
	Schema string `json:"$schema,omitempty"`
	Additionalitems GeneratedType `json:"additionalItems,omitempty"` // JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property.
	Maxitems int64 `json:"maxItems,omitempty"`
	Oneof []GeneratedType `json:"oneOf,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Patternproperties map[string]interface{} `json:"patternProperties,omitempty"`
	Maxlength int64 `json:"maxLength,omitempty"`
	Enum []GeneratedType `json:"enum,omitempty"`
	Exclusivemaximum bool `json:"exclusiveMaximum,omitempty"`
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
	Minlength int64 `json:"minLength,omitempty"`
	Id string `json:"id,omitempty"`
	Definitions map[string]interface{} `json:"definitions,omitempty"`
	Multipleof float64 `json:"multipleOf,omitempty"`
	Not GeneratedType `json:"not,omitempty"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
	Minproperties int64 `json:"minProperties,omitempty"`
	Maxproperties int64 `json:"maxProperties,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
	Uniqueitems bool `json:"uniqueItems,omitempty"`
	Ref string `json:"$ref,omitempty"`
	Minimum float64 `json:"minimum,omitempty"`
	Title string `json:"title,omitempty"`
	Items GeneratedType `json:"items,omitempty"` // JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps or an array of JSONSchemaProps. Mainly here for serialization purposes.
	Description string `json:"description,omitempty"`
	TypeField string `json:"type,omitempty"`
	Example GeneratedType `json:"example,omitempty"` // JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
	Exclusiveminimum bool `json:"exclusiveMinimum,omitempty"`
	Additionalproperties GeneratedType `json:"additionalProperties,omitempty"` // JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property.
	Required []string `json:"required,omitempty"`
	Anyof []GeneratedType `json:"anyOf,omitempty"`
	Externaldocs GeneratedType `json:"externalDocs,omitempty"` // ExternalDocumentation allows referencing an external resource for extended documentation.
	Properties map[string]interface{} `json:"properties,omitempty"`
	Allof []GeneratedType `json:"allOf,omitempty"`
	Minitems int64 `json:"minItems,omitempty"`
}

// Iok8sapiextensionsv1beta1AllowedHostPath represents the Iok8sapiextensionsv1beta1AllowedHostPath schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1AllowedHostPath struct {
	Pathprefix string `json:"pathPrefix,omitempty"` // is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path. Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`
}

// Iok8sapiappsv1beta2DaemonSetSpec represents the Iok8sapiappsv1beta2DaemonSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta2DaemonSetSpec struct {
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Updatestrategy Iok8sapiappsv1beta2DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"` // DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiextensionsv1beta1IngressRule represents the Iok8sapiextensionsv1beta1IngressRule schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IngressRule struct {
	Host string `json:"host,omitempty"` // Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the 	 IP in the Spec of the parent Ingress. 2. The `:` delimiter is not respected because ports are not allowed. 	 Currently the port of an Ingress is implicitly :80 for http and 	 :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
	Http Iok8sapiextensionsv1beta1HTTPIngressRuleValue `json:"http,omitempty"` // HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
}

// Iok8skubernetespkgapisauthorizationv1beta1ResourceAttributes represents the Iok8skubernetespkgapisauthorizationv1beta1ResourceAttributes schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1ResourceAttributes struct {
	Subresource string `json:"subresource,omitempty"` // Subresource is one of the existing resource types. "" means none.
	Verb string `json:"verb,omitempty"` // Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Version string `json:"version,omitempty"` // Version is the API Version of the Resource. "*" means all.
	Group string `json:"group,omitempty"` // Group is the API Group of the Resource. "*" means all.
	Name string `json:"name,omitempty"` // Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the action being requested. Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Resource string `json:"resource,omitempty"` // Resource is one of the existing resource types. "*" means all.
}

// Iok8skubernetespkgapiv1Secret represents the Iok8skubernetespkgapiv1Secret schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Secret struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Stringdata map[string]interface{} `json:"stringData,omitempty"` // stringData allows specifying non-binary secret data in string form. It is provided as a write-only convenience method. All keys and values are merged into the data field on write, overwriting any existing values. It is never output when reading from the API.
	TypeField string `json:"type,omitempty"` // Used to facilitate programmatic handling of secret data.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Data map[string]interface{} `json:"data,omitempty"` // Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta1StatefulSetSpec represents the Iok8sapiappsv1beta1StatefulSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta1StatefulSetSpec struct {
	Podmanagementpolicy string `json:"podManagementPolicy,omitempty"` // podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	Replicas int `json:"replicas,omitempty"` // replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Servicename string `json:"serviceName"` // serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Updatestrategy Iok8sapiappsv1beta1StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"` // StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
	Volumeclaimtemplates []Iok8sapicorev1PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"` // volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
}

// Iok8skubernetespkgapiv1NFSVolumeSource represents the Iok8skubernetespkgapiv1NFSVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NFSVolumeSource struct {
	Path string `json:"path"` // Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server string `json:"server"` // Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
}

// Iok8sapicorev1ServiceAccountList represents the Iok8sapicorev1ServiceAccountList schema from the OpenAPI specification
type Iok8sapicorev1ServiceAccountList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ServiceAccount `json:"items"` // List of ServiceAccounts. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapirbacv1Role represents the Iok8sapirbacv1Role schema from the OpenAPI specification
type Iok8sapirbacv1Role struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this Role
}

// Iok8sapicorev1AttachedVolume represents the Iok8sapicorev1AttachedVolume schema from the OpenAPI specification
type Iok8sapicorev1AttachedVolume struct {
	Name string `json:"name"` // Name of the attached volume
	Devicepath string `json:"devicePath"` // DevicePath represents the device path where the volume should be available
}

// Iok8sapiappsv1Deployment represents the Iok8sapiappsv1Deployment schema from the OpenAPI specification
type Iok8sapiappsv1Deployment struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
	Status Iok8sapiappsv1DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapicorev1SecretVolumeSource represents the Iok8sapicorev1SecretVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1SecretVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or it's keys must be defined
	Secretname string `json:"secretName,omitempty"` // Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
}

// Iok8skubernetespkgapisbatchv2alpha1CronJobSpec represents the Iok8skubernetespkgapisbatchv2alpha1CronJobSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv2alpha1CronJobSpec struct {
	Schedule string `json:"schedule"` // The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Startingdeadlineseconds int64 `json:"startingDeadlineSeconds,omitempty"` // Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.
	Successfuljobshistorylimit int `json:"successfulJobsHistoryLimit,omitempty"` // The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	Suspend bool `json:"suspend,omitempty"` // This flag tells the controller to suspend subsequent executions, it does not apply to already started executions. Defaults to false.
	Concurrencypolicy string `json:"concurrencyPolicy,omitempty"` // Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	Failedjobshistorylimit int `json:"failedJobsHistoryLimit,omitempty"` // The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	Jobtemplate Iok8sapibatchv2alpha1JobTemplateSpec `json:"jobTemplate"` // JobTemplateSpec describes the data a Job should have when created from a template
}

// Iok8skubernetespkgapiv1Event represents the Iok8skubernetespkgapiv1Event schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Event struct {
	Firsttimestamp string `json:"firstTimestamp,omitempty"`
	Reason string `json:"reason,omitempty"` // This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	TypeField string `json:"type,omitempty"` // Type of this event (Normal, Warning), new types could be added in the future
	Eventtime string `json:"eventTime,omitempty"`
	Source Iok8sapicorev1EventSource `json:"source,omitempty"` // EventSource contains information for an event.
	Count int `json:"count,omitempty"` // The number of times this event has occurred.
	Lasttimestamp string `json:"lastTimestamp,omitempty"`
	Message string `json:"message,omitempty"` // A human-readable description of the status of this operation.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Series Iok8sapicorev1EventSeries `json:"series,omitempty"` // EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
	Reportinginstance string `json:"reportingInstance,omitempty"` // ID of the controller instance, e.g. `kubelet-xyzf`.
	Action string `json:"action,omitempty"` // What action was taken/failed regarding to the Regarding object.
	Involvedobject Iok8sapicorev1ObjectReference `json:"involvedObject"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Reportingcomponent string `json:"reportingComponent,omitempty"` // Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	Related Iok8sapicorev1ObjectReference `json:"related,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
}

// Iok8skubernetespkgapiv1WeightedPodAffinityTerm represents the Iok8skubernetespkgapiv1WeightedPodAffinityTerm schema from the OpenAPI specification
type Iok8skubernetespkgapiv1WeightedPodAffinityTerm struct {
	Podaffinityterm Iok8sapicorev1PodAffinityTerm `json:"podAffinityTerm"` // Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
	Weight int `json:"weight"` // weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
}

// Iok8sapipolicyv1beta1SupplementalGroupsStrategyOptions represents the Iok8sapipolicyv1beta1SupplementalGroupsStrategyOptions schema from the OpenAPI specification
type Iok8sapipolicyv1beta1SupplementalGroupsStrategyOptions struct {
	Rule string `json:"rule,omitempty"` // Rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	Ranges []Iok8sapipolicyv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of supplemental groups. If you would like to force a single supplemental group then supply a single range with the same start and end.
}

// Iok8sapicorev1VolumeNodeAffinity represents the Iok8sapicorev1VolumeNodeAffinity schema from the OpenAPI specification
type Iok8sapicorev1VolumeNodeAffinity struct {
	Required Iok8sapicorev1NodeSelector `json:"required,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8skubernetespkgapiv1VolumeMount represents the Iok8skubernetespkgapiv1VolumeMount schema from the OpenAPI specification
type Iok8skubernetespkgapiv1VolumeMount struct {
	Subpath string `json:"subPath,omitempty"` // Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	Mountpath string `json:"mountPath"` // Path within the container at which the volume should be mounted. Must not contain ':'.
	Mountpropagation string `json:"mountPropagation,omitempty"` // mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationHostToContainer is used. This field is beta in 1.10.
	Name string `json:"name"` // This must match the Name of a Volume.
	Readonly bool `json:"readOnly,omitempty"` // Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
}

// Iok8sapiauthorizationv1SubjectAccessReviewSpec represents the Iok8sapiauthorizationv1SubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectAccessReviewSpec struct {
	Uid string `json:"uid,omitempty"` // UID information about the requesting user.
	User string `json:"user,omitempty"` // User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra corresponds to the user.Info.GetExtra() method from the authenticator. Since that is input to the authorizer it needs a reflection here.
	Groups []string `json:"groups,omitempty"` // Groups is the groups you're testing for.
	Nonresourceattributes Iok8sapiauthorizationv1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
}

// Iok8sapipolicyv1beta1AllowedHostPath represents the Iok8sapipolicyv1beta1AllowedHostPath schema from the OpenAPI specification
type Iok8sapipolicyv1beta1AllowedHostPath struct {
	Pathprefix string `json:"pathPrefix,omitempty"` // is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path. Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`
}

// Iok8skubernetespkgapisauthorizationv1beta1SelfSubjectAccessReviewSpec represents the Iok8skubernetespkgapisauthorizationv1beta1SelfSubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1SelfSubjectAccessReviewSpec struct {
	Nonresourceattributes Iok8sapiauthorizationv1beta1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1beta1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
}

// Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyIngressRule represents the Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyIngressRule schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyIngressRule struct {
	From []Iok8sapiextensionsv1beta1NetworkPolicyPeer `json:"from,omitempty"` // List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.
	Ports []Iok8sapiextensionsv1beta1NetworkPolicyPort `json:"ports,omitempty"` // List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
}

// Iok8skubernetespkgapiv1EndpointPort represents the Iok8skubernetespkgapiv1EndpointPort schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EndpointPort struct {
	Protocol string `json:"protocol,omitempty"` // The IP protocol for this port. Must be UDP or TCP. Default is TCP.
	Name string `json:"name,omitempty"` // The name of this port (corresponds to ServicePort.Name). Must be a DNS_LABEL. Optional only if one port is defined.
	Port int `json:"port"` // The port number of the endpoint.
}

// Iok8sapiappsv1StatefulSet represents the Iok8sapiappsv1StatefulSet schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1StatefulSetSpec `json:"spec,omitempty"` // A StatefulSetSpec is the specification of a StatefulSet.
	Status Iok8sapiappsv1StatefulSetStatus `json:"status,omitempty"` // StatefulSetStatus represents the current state of a StatefulSet.
}

// Iok8skubernetespkgapiv1SecurityContext represents the Iok8skubernetespkgapiv1SecurityContext schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SecurityContext struct {
	Readonlyrootfilesystem bool `json:"readOnlyRootFilesystem,omitempty"` // Whether this container has a read-only root filesystem. Default is false.
	Runasgroup int64 `json:"runAsGroup,omitempty"` // The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Runasnonroot bool `json:"runAsNonRoot,omitempty"` // Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Runasuser int64 `json:"runAsUser,omitempty"` // The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
	Allowprivilegeescalation bool `json:"allowPrivilegeEscalation,omitempty"` // AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	Capabilities Iok8sapicorev1Capabilities `json:"capabilities,omitempty"` // Adds and removes POSIX capabilities from running containers.
	Privileged bool `json:"privileged,omitempty"` // Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
}

// Iok8skubernetespkgapiv1PodAffinityTerm represents the Iok8skubernetespkgapiv1PodAffinityTerm schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodAffinityTerm struct {
	Topologykey string `json:"topologyKey"` // This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	Labelselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"labelSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Namespaces []string `json:"namespaces,omitempty"` // namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this pod's namespace"
}

// Iok8sapicorev1NodeSelector represents the Iok8sapicorev1NodeSelector schema from the OpenAPI specification
type Iok8sapicorev1NodeSelector struct {
	Nodeselectorterms []Iok8sapicorev1NodeSelectorTerm `json:"nodeSelectorTerms"` // Required. A list of node selector terms. The terms are ORed.
}

// Iok8skubernetespkgapisextensionsv1beta1NetworkPolicy represents the Iok8skubernetespkgapisextensionsv1beta1NetworkPolicy schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1NetworkPolicy struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1NetworkPolicySpec `json:"spec,omitempty"` // DEPRECATED 1.9 - This group version of NetworkPolicySpec is deprecated by networking/v1/NetworkPolicySpec.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta2StatefulSetStatus represents the Iok8sapiappsv1beta2StatefulSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta2StatefulSetStatus struct {
	Conditions []Iok8sapiappsv1beta2StatefulSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a statefulset's current state.
	Replicas int `json:"replicas"` // replicas is the number of Pods created by the StatefulSet controller.
	Updaterevision string `json:"updateRevision,omitempty"` // updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
	Currentreplicas int `json:"currentReplicas,omitempty"` // currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
	Readyreplicas int `json:"readyReplicas,omitempty"` // readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	Currentrevision string `json:"currentRevision,omitempty"` // currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	Collisioncount int `json:"collisionCount,omitempty"` // collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
}

// Iok8skubernetespkgapisnetworkingv1NetworkPolicySpec represents the Iok8skubernetespkgapisnetworkingv1NetworkPolicySpec schema from the OpenAPI specification
type Iok8skubernetespkgapisnetworkingv1NetworkPolicySpec struct {
	Egress []Iok8sapinetworkingv1NetworkPolicyEgressRule `json:"egress,omitempty"` // List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
	Ingress []Iok8sapinetworkingv1NetworkPolicyIngressRule `json:"ingress,omitempty"` // List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Policytypes []string `json:"policyTypes,omitempty"` // List of rule types that the NetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
}

// Iok8sapicorev1HTTPHeader represents the Iok8sapicorev1HTTPHeader schema from the OpenAPI specification
type Iok8sapicorev1HTTPHeader struct {
	Name string `json:"name"` // The header field name
	Value string `json:"value"` // The header field value
}

// Iok8skubernetespkgapisrbacv1alpha1Subject represents the Iok8skubernetespkgapisrbacv1alpha1Subject schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1Subject struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion holds the API group and version of the referenced subject. Defaults to "v1" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io/v1alpha1" for User and Group subjects.
	Kind string `json:"kind"` // Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Name string `json:"name"` // Name of the object being referenced.
	Namespace string `json:"namespace,omitempty"` // Namespace of the referenced object. If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
}

// Iok8sapimachinerypkgapismetav1ObjectMeta represents the Iok8sapimachinerypkgapismetav1ObjectMeta schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ObjectMeta struct {
	Clustername string `json:"clusterName,omitempty"` // The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.
	Labels map[string]interface{} `json:"labels,omitempty"` // Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	Deletiontimestamp string `json:"deletionTimestamp,omitempty"`
	Initializers Iok8sapimachinerypkgapismetav1Initializers `json:"initializers,omitempty"` // Initializers tracks the progress of initialization.
	Namespace string `json:"namespace,omitempty"` // Namespace defines the space within each name must be unique. An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty. Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
	Name string `json:"name,omitempty"` // Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Uid string `json:"uid,omitempty"` // UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	Selflink string `json:"selfLink,omitempty"` // SelfLink is a URL representing this object. Populated by the system. Read-only.
	Creationtimestamp string `json:"creationTimestamp,omitempty"`
	Deletiongraceperiodseconds int64 `json:"deletionGracePeriodSeconds,omitempty"` // Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.
	Generation int64 `json:"generation,omitempty"` // A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Ownerreferences []Iok8sapimachinerypkgapismetav1OwnerReference `json:"ownerReferences,omitempty"` // List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.
	Generatename string `json:"generateName,omitempty"` // GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server. If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header). Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency
	Finalizers []string `json:"finalizers,omitempty"` // Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.
	Resourceversion string `json:"resourceVersion,omitempty"` // An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources. Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
}

// Iok8sapicorev1CephFSVolumeSource represents the Iok8sapicorev1CephFSVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CephFSVolumeSource struct {
	Monitors []string `json:"monitors"` // Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Path string `json:"path,omitempty"` // Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Secretfile string `json:"secretFile,omitempty"` // Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	User string `json:"user,omitempty"` // Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
}

// Iok8sapicorev1ObjectFieldSelector represents the Iok8sapicorev1ObjectFieldSelector schema from the OpenAPI specification
type Iok8sapicorev1ObjectFieldSelector struct {
	Apiversion string `json:"apiVersion,omitempty"` // Version of the schema the FieldPath is written in terms of, defaults to "v1".
	Fieldpath string `json:"fieldPath"` // Path of the field to select in the specified API version.
}

// Iok8sapiextensionsv1beta1ReplicaSetSpec represents the Iok8sapiextensionsv1beta1ReplicaSetSpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ReplicaSetSpec struct {
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
}

// Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerList represents the Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerList schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiautoscalingv2beta1HorizontalPodAutoscaler `json:"items"` // items is the list of horizontal pod autoscaler objects.
}

// Iok8skubernetespkgapiv1Volume represents the Iok8skubernetespkgapiv1Volume schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Volume struct {
	Azurefile Iok8sapicorev1AzureFileVolumeSource `json:"azureFile,omitempty"` // AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	Downwardapi Iok8sapicorev1DownwardAPIVolumeSource `json:"downwardAPI,omitempty"` // DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
	Cephfs Iok8sapicorev1CephFSVolumeSource `json:"cephfs,omitempty"` // Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Quobyte Iok8sapicorev1QuobyteVolumeSource `json:"quobyte,omitempty"` // Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Cinder Iok8sapicorev1CinderVolumeSource `json:"cinder,omitempty"` // Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Hostpath Iok8sapicorev1HostPathVolumeSource `json:"hostPath,omitempty"` // Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	Name string `json:"name"` // Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Photonpersistentdisk Iok8sapicorev1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"` // Represents a Photon Controller persistent disk resource.
	Awselasticblockstore Iok8sapicorev1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"` // Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	Projected Iok8sapicorev1ProjectedVolumeSource `json:"projected,omitempty"` // Represents a projected volume source
	Scaleio Iok8sapicorev1ScaleIOVolumeSource `json:"scaleIO,omitempty"` // ScaleIOVolumeSource represents a persistent ScaleIO volume
	Configmap Iok8sapicorev1ConfigMapVolumeSource `json:"configMap,omitempty"` // Adapts a ConfigMap into a volume. The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.
	Iscsi Iok8sapicorev1ISCSIVolumeSource `json:"iscsi,omitempty"` // Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Flocker Iok8sapicorev1FlockerVolumeSource `json:"flocker,omitempty"` // Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Azuredisk Iok8sapicorev1AzureDiskVolumeSource `json:"azureDisk,omitempty"` // AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	Glusterfs Iok8sapicorev1GlusterfsVolumeSource `json:"glusterfs,omitempty"` // Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Gcepersistentdisk Iok8sapicorev1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"` // Represents a Persistent Disk resource in Google Compute Engine. A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	Portworxvolume Iok8sapicorev1PortworxVolumeSource `json:"portworxVolume,omitempty"` // PortworxVolumeSource represents a Portworx volume resource.
	Secret Iok8sapicorev1SecretVolumeSource `json:"secret,omitempty"` // Adapts a Secret into a volume. The contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.
	Flexvolume Iok8sapicorev1FlexVolumeSource `json:"flexVolume,omitempty"` // FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	Nfs Iok8sapicorev1NFSVolumeSource `json:"nfs,omitempty"` // Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Rbd Iok8sapicorev1RBDVolumeSource `json:"rbd,omitempty"` // Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Gitrepo Iok8sapicorev1GitRepoVolumeSource `json:"gitRepo,omitempty"` // Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.
	Storageos Iok8sapicorev1StorageOSVolumeSource `json:"storageos,omitempty"` // Represents a StorageOS persistent volume resource.
	Vspherevolume Iok8sapicorev1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"` // Represents a vSphere volume resource.
	Persistentvolumeclaim Iok8sapicorev1PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"` // PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
	Emptydir Iok8sapicorev1EmptyDirVolumeSource `json:"emptyDir,omitempty"` // Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
	Fc Iok8sapicorev1FCVolumeSource `json:"fc,omitempty"` // Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
}

// Iok8skubernetespkgapiv1NamespaceList represents the Iok8skubernetespkgapiv1NamespaceList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NamespaceList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Namespace `json:"items"` // Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapimachinerypkgversionInfo represents the Iok8sapimachinerypkgversionInfo schema from the OpenAPI specification
type Iok8sapimachinerypkgversionInfo struct {
	Minor string `json:"minor"`
	Builddate string `json:"buildDate"`
	Gitcommit string `json:"gitCommit"`
	Gitversion string `json:"gitVersion"`
	Goversion string `json:"goVersion"`
	Platform string `json:"platform"`
	Gittreestate string `json:"gitTreeState"`
	Compiler string `json:"compiler"`
	Major string `json:"major"`
}

// Iok8skubernetespkgapiv1EndpointSubset represents the Iok8skubernetespkgapiv1EndpointSubset schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EndpointSubset struct {
	Notreadyaddresses []Iok8sapicorev1EndpointAddress `json:"notReadyAddresses,omitempty"` // IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
	Ports []Iok8sapicorev1EndpointPort `json:"ports,omitempty"` // Port numbers available on the related IP addresses.
	Addresses []Iok8sapicorev1EndpointAddress `json:"addresses,omitempty"` // IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
}

// Iok8sapirbacv1alpha1Subject represents the Iok8sapirbacv1alpha1Subject schema from the OpenAPI specification
type Iok8sapirbacv1alpha1Subject struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion holds the API group and version of the referenced subject. Defaults to "v1" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io/v1alpha1" for User and Group subjects.
	Kind string `json:"kind"` // Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Name string `json:"name"` // Name of the object being referenced.
	Namespace string `json:"namespace,omitempty"` // Namespace of the referenced object. If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
}

// Iok8skubernetespkgapissettingsv1alpha1PodPresetSpec represents the Iok8skubernetespkgapissettingsv1alpha1PodPresetSpec schema from the OpenAPI specification
type Iok8skubernetespkgapissettingsv1alpha1PodPresetSpec struct {
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Volumemounts []Iok8sapicorev1VolumeMount `json:"volumeMounts,omitempty"` // VolumeMounts defines the collection of VolumeMount to inject into containers.
	Volumes []Iok8sapicorev1Volume `json:"volumes,omitempty"` // Volumes defines the collection of Volume to inject into the pod.
	Env []Iok8sapicorev1EnvVar `json:"env,omitempty"` // Env defines the collection of EnvVar to inject into containers.
	Envfrom []Iok8sapicorev1EnvFromSource `json:"envFrom,omitempty"` // EnvFrom defines the collection of EnvFromSource to inject into containers.
}

// Iok8sapicorev1ReplicationControllerSpec represents the Iok8sapicorev1ReplicationControllerSpec schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerSpec struct {
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Selector map[string]interface{} `json:"selector,omitempty"` // Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
}

// Iok8sapiappsv1beta1RollingUpdateStatefulSetStrategy represents the Iok8sapiappsv1beta1RollingUpdateStatefulSetStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta1RollingUpdateStatefulSetStrategy struct {
	Partition int `json:"partition,omitempty"` // Partition indicates the ordinal at which the StatefulSet should be partitioned.
}

// Iok8skubernetespkgapiv1ContainerStateWaiting represents the Iok8skubernetespkgapiv1ContainerStateWaiting schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerStateWaiting struct {
	Message string `json:"message,omitempty"` // Message regarding why the container is not yet running.
	Reason string `json:"reason,omitempty"` // (brief) reason the container is not yet running.
}

// Iok8sapinetworkingv1NetworkPolicyIngressRule represents the Iok8sapinetworkingv1NetworkPolicyIngressRule schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyIngressRule struct {
	From []Iok8sapinetworkingv1NetworkPolicyPeer `json:"from,omitempty"` // List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.
	Ports []Iok8sapinetworkingv1NetworkPolicyPort `json:"ports,omitempty"` // List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
}

// Iok8sapistoragev1alpha1VolumeAttachmentSource represents the Iok8sapistoragev1alpha1VolumeAttachmentSource schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttachmentSource struct {
	Persistentvolumename string `json:"persistentVolumeName,omitempty"` // Name of the persistent volume to attach.
}

// Iok8sapiauthorizationv1beta1SelfSubjectAccessReviewSpec represents the Iok8sapiauthorizationv1beta1SelfSubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SelfSubjectAccessReviewSpec struct {
	Nonresourceattributes Iok8sapiauthorizationv1beta1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1beta1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
}

// Iok8skubernetespkgapisrbacv1beta1ClusterRoleBinding represents the Iok8skubernetespkgapisrbacv1beta1ClusterRoleBinding schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1ClusterRoleBinding struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Roleref Iok8sapirbacv1beta1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1beta1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapicorev1ReplicationController represents the Iok8sapicorev1ReplicationController schema from the OpenAPI specification
type Iok8sapicorev1ReplicationController struct {
	Spec Iok8sapicorev1ReplicationControllerSpec `json:"spec,omitempty"` // ReplicationControllerSpec is the specification of a replication controller.
	Status Iok8sapicorev1ReplicationControllerStatus `json:"status,omitempty"` // ReplicationControllerStatus represents the current status of a replication controller.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiextensionsv1beta1DeploymentRollback represents the Iok8sapiextensionsv1beta1DeploymentRollback schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DeploymentRollback struct {
	Name string `json:"name"` // Required: This must match the Name of a deployment.
	Rollbackto Iok8sapiextensionsv1beta1RollbackConfig `json:"rollbackTo"` // DEPRECATED.
	Updatedannotations map[string]interface{} `json:"updatedAnnotations,omitempty"` // The annotations to be updated to a deployment
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1EnvVarSource represents the Iok8skubernetespkgapiv1EnvVarSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EnvVarSource struct {
	Configmapkeyref Iok8sapicorev1ConfigMapKeySelector `json:"configMapKeyRef,omitempty"` // Selects a key from a ConfigMap.
	Fieldref Iok8sapicorev1ObjectFieldSelector `json:"fieldRef,omitempty"` // ObjectFieldSelector selects an APIVersioned field of an object.
	Resourcefieldref Iok8sapicorev1ResourceFieldSelector `json:"resourceFieldRef,omitempty"` // ResourceFieldSelector represents container resources (cpu, memory) and their output format
	Secretkeyref Iok8sapicorev1SecretKeySelector `json:"secretKeyRef,omitempty"` // SecretKeySelector selects a key of a Secret.
}

// Iok8skubernetespkgapiv1ServiceList represents the Iok8skubernetespkgapiv1ServiceList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ServiceList struct {
	Items []Iok8sapicorev1Service `json:"items"` // List of services
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1VolumeProjection represents the Iok8skubernetespkgapiv1VolumeProjection schema from the OpenAPI specification
type Iok8skubernetespkgapiv1VolumeProjection struct {
	Configmap Iok8sapicorev1ConfigMapProjection `json:"configMap,omitempty"` // Adapts a ConfigMap into a projected volume. The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.
	Downwardapi Iok8sapicorev1DownwardAPIProjection `json:"downwardAPI,omitempty"` // Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
	Secret Iok8sapicorev1SecretProjection `json:"secret,omitempty"` // Adapts a secret into a projected volume. The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.
}

// Iok8skubernetespkgapisextensionsv1beta1RollingUpdateDaemonSet represents the Iok8skubernetespkgapisextensionsv1beta1RollingUpdateDaemonSet schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1RollingUpdateDaemonSet struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8skubernetespkgapisappsv1beta1StatefulSetStatus represents the Iok8skubernetespkgapisappsv1beta1StatefulSetStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1StatefulSetStatus struct {
	Replicas int `json:"replicas"` // replicas is the number of Pods created by the StatefulSet controller.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	Collisioncount int `json:"collisionCount,omitempty"` // collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Conditions []Iok8sapiappsv1beta1StatefulSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a statefulset's current state.
	Currentreplicas int `json:"currentReplicas,omitempty"` // currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
	Currentrevision string `json:"currentRevision,omitempty"` // currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
	Readyreplicas int `json:"readyReplicas,omitempty"` // readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	Updaterevision string `json:"updateRevision,omitempty"` // updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
}

// Iok8sapicorev1VsphereVirtualDiskVolumeSource represents the Iok8sapicorev1VsphereVirtualDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1VsphereVirtualDiskVolumeSource struct {
	Volumepath string `json:"volumePath"` // Path that identifies vSphere volume vmdk
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Storagepolicyid string `json:"storagePolicyID,omitempty"` // Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	Storagepolicyname string `json:"storagePolicyName,omitempty"` // Storage Policy Based Management (SPBM) profile name.
}

// Iok8sapicorev1DaemonEndpoint represents the Iok8sapicorev1DaemonEndpoint schema from the OpenAPI specification
type Iok8sapicorev1DaemonEndpoint struct {
	Port int `json:"Port"` // Port number of the given endpoint.
}

// Iok8skubernetespkgapiv1PersistentVolumeClaimSpec represents the Iok8skubernetespkgapiv1PersistentVolumeClaimSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeClaimSpec struct {
	Volumename string `json:"volumeName,omitempty"` // VolumeName is the binding reference to the PersistentVolume backing this claim.
	Accessmodes []string `json:"accessModes,omitempty"` // AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Storageclassname string `json:"storageClassName,omitempty"` // Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	Volumemode string `json:"volumeMode,omitempty"` // volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. This is an alpha feature and may change in the future.
}

// Iok8skubernetespkgapisautoscalingv1Scale represents the Iok8skubernetespkgapisautoscalingv1Scale schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1Scale struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiautoscalingv1ScaleSpec `json:"spec,omitempty"` // ScaleSpec describes the attributes of a scale subresource.
	Status Iok8sapiautoscalingv1ScaleStatus `json:"status,omitempty"` // ScaleStatus represents the current status of a scale subresource.
}

// Iok8sapiextensionsv1beta1NetworkPolicyPort represents the Iok8sapiextensionsv1beta1NetworkPolicyPort schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicyPort struct {
	Port string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"` // Optional. The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.
}

// Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerList represents the Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerList schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiautoscalingv1HorizontalPodAutoscaler `json:"items"` // list of horizontal pod autoscaler objects.
}

// Iok8skubernetespkgapisrbacv1beta1RoleList represents the Iok8skubernetespkgapisrbacv1beta1RoleList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1RoleList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1Role `json:"items"` // Items is a list of Roles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1StatefulSetSpec represents the Iok8sapiappsv1StatefulSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetSpec struct {
	Updatestrategy Iok8sapiappsv1StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"` // StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
	Volumeclaimtemplates []Iok8sapicorev1PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"` // volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	Podmanagementpolicy string `json:"podManagementPolicy,omitempty"` // podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	Replicas int `json:"replicas,omitempty"` // replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Servicename string `json:"serviceName"` // serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// Iok8sapimachinerypkgapismetav1Preconditions represents the Iok8sapimachinerypkgapismetav1Preconditions schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Preconditions struct {
	Uid string `json:"uid,omitempty"` // Specifies the target UID.
}

// Iok8sapimachinerypkgapismetav1Status represents the Iok8sapimachinerypkgapismetav1Status schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Status struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Reason string `json:"reason,omitempty"` // A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Status string `json:"status,omitempty"` // Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Code int `json:"code,omitempty"` // Suggested HTTP return code for this status, 0 if not set.
	Details Iok8sapimachinerypkgapismetav1StatusDetails `json:"details,omitempty"` // StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Message string `json:"message,omitempty"` // A human-readable description of the status of this operation.
}

// Iok8sapiappsv1StatefulSetUpdateStrategy represents the Iok8sapiappsv1StatefulSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetUpdateStrategy struct {
	Rollingupdate Iok8sapiappsv1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"` // RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
	TypeField string `json:"type,omitempty"` // Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
}

// Iok8sapicorev1PersistentVolumeSpec represents the Iok8sapicorev1PersistentVolumeSpec schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeSpec struct {
	Portworxvolume Iok8sapicorev1PortworxVolumeSource `json:"portworxVolume,omitempty"` // PortworxVolumeSource represents a Portworx volume resource.
	Storageclassname string `json:"storageClassName,omitempty"` // Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	Quobyte Iok8sapicorev1QuobyteVolumeSource `json:"quobyte,omitempty"` // Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Iscsi Iok8sapicorev1ISCSIPersistentVolumeSource `json:"iscsi,omitempty"` // ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Photonpersistentdisk Iok8sapicorev1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"` // Represents a Photon Controller persistent disk resource.
	Cinder Iok8sapicorev1CinderVolumeSource `json:"cinder,omitempty"` // Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Vspherevolume Iok8sapicorev1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"` // Represents a vSphere volume resource.
	Azurefile Iok8sapicorev1AzureFilePersistentVolumeSource `json:"azureFile,omitempty"` // AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	Storageos Iok8sapicorev1StorageOSPersistentVolumeSource `json:"storageos,omitempty"` // Represents a StorageOS persistent volume resource.
	Local Iok8sapicorev1LocalVolumeSource `json:"local,omitempty"` // Local represents directly-attached storage with node affinity
	Mountoptions []string `json:"mountOptions,omitempty"` // A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	Gcepersistentdisk Iok8sapicorev1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"` // Represents a Persistent Disk resource in Google Compute Engine. A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	Hostpath Iok8sapicorev1HostPathVolumeSource `json:"hostPath,omitempty"` // Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	Csi Iok8sapicorev1CSIPersistentVolumeSource `json:"csi,omitempty"` // Represents storage that is managed by an external CSI volume driver (Beta feature)
	Scaleio Iok8sapicorev1ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty"` // ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
	Capacity map[string]interface{} `json:"capacity,omitempty"` // A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Nfs Iok8sapicorev1NFSVolumeSource `json:"nfs,omitempty"` // Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Glusterfs Iok8sapicorev1GlusterfsVolumeSource `json:"glusterfs,omitempty"` // Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Accessmodes []string `json:"accessModes,omitempty"` // AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	Nodeaffinity Iok8sapicorev1VolumeNodeAffinity `json:"nodeAffinity,omitempty"` // VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
	Claimref Iok8sapicorev1ObjectReference `json:"claimRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Azuredisk Iok8sapicorev1AzureDiskVolumeSource `json:"azureDisk,omitempty"` // AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	Flexvolume Iok8sapicorev1FlexPersistentVolumeSource `json:"flexVolume,omitempty"` // FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
	Fc Iok8sapicorev1FCVolumeSource `json:"fc,omitempty"` // Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Flocker Iok8sapicorev1FlockerVolumeSource `json:"flocker,omitempty"` // Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Cephfs Iok8sapicorev1CephFSPersistentVolumeSource `json:"cephfs,omitempty"` // Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Awselasticblockstore Iok8sapicorev1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"` // Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	Persistentvolumereclaimpolicy string `json:"persistentVolumeReclaimPolicy,omitempty"` // What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	Rbd Iok8sapicorev1RBDPersistentVolumeSource `json:"rbd,omitempty"` // Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Volumemode string `json:"volumeMode,omitempty"` // volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec. This is an alpha feature and may change in the future.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Shortnames []string `json:"shortNames,omitempty"` // ShortNames are short names for the resource. It must be all lowercase.
	Singular string `json:"singular,omitempty"` // Singular is the singular name of the resource. It must be all lowercase Defaults to lowercased <kind>
	Categories []string `json:"categories,omitempty"` // Categories is a list of grouped resources custom resources belong to (e.g. 'all')
	Kind string `json:"kind"` // Kind is the serialized kind of the resource. It is normally CamelCase and singular.
	Listkind string `json:"listKind,omitempty"` // ListKind is the serialized kind of the list for this resource. Defaults to <kind>List.
	Plural string `json:"plural"` // Plural is the plural name of the resource to serve. It must match the name of the CustomResourceDefinition-registration too: plural.group and it must be all lowercase.
}

// Iok8sapicorev1LocalVolumeSource represents the Iok8sapicorev1LocalVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1LocalVolumeSource struct {
	Path string `json:"path"` // The full path to the volume on the node For alpha, this path must be a directory Once block as a source is supported, then this path can point to a block device
}

// Iok8sapibatchv1JobSpec represents the Iok8sapibatchv1JobSpec schema from the OpenAPI specification
type Iok8sapibatchv1JobSpec struct {
	Parallelism int `json:"parallelism,omitempty"` // Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Activedeadlineseconds int64 `json:"activeDeadlineSeconds,omitempty"` // Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
	Backofflimit int `json:"backoffLimit,omitempty"` // Specifies the number of retries before marking this job failed. Defaults to 6
	Completions int `json:"completions,omitempty"` // Specifies the desired number of successfully finished pods the job should be run with. Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Manualselector bool `json:"manualSelector,omitempty"` // manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
}

// Iok8sapirbacv1alpha1RoleList represents the Iok8sapirbacv1alpha1RoleList schema from the OpenAPI specification
type Iok8sapirbacv1alpha1RoleList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1Role `json:"items"` // Items is a list of Roles
}

// Iok8sapiadmissionregistrationv1beta1ServiceReference represents the Iok8sapiadmissionregistrationv1beta1ServiceReference schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1ServiceReference struct {
	Name string `json:"name"` // `name` is the name of the service. Required
	Namespace string `json:"namespace"` // `namespace` is the namespace of the service. Required
	Path string `json:"path,omitempty"` // `path` is an optional URL path which will be sent in any request to this service.
}

// Iok8sapicorev1ClientIPConfig represents the Iok8sapicorev1ClientIPConfig schema from the OpenAPI specification
type Iok8sapicorev1ClientIPConfig struct {
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
}

// Iok8sapiappsv1beta1ScaleStatus represents the Iok8sapiappsv1beta1ScaleStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta1ScaleStatus struct {
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
	Selector map[string]interface{} `json:"selector,omitempty"` // label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Targetselector string `json:"targetSelector,omitempty"` // label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
}

// Iok8sapiappsv1beta2DeploymentSpec represents the Iok8sapiappsv1beta2DeploymentSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta2DeploymentSpec struct {
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Strategy Iok8sapiappsv1beta2DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused.
}

// Iok8sapiappsv1beta1RollbackConfig represents the Iok8sapiappsv1beta1RollbackConfig schema from the OpenAPI specification
type Iok8sapiappsv1beta1RollbackConfig struct {
	Revision int64 `json:"revision,omitempty"` // The revision to rollback to. If set to 0, rollback to the last revision.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allows bool `json:"Allows"`
	Schema GeneratedType `json:"Schema"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
}

// Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR represents the Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR struct {
	Clientcidr string `json:"clientCIDR"` // The CIDR with which clients can match their IP to figure out the server address that they should use.
	Serveraddress string `json:"serverAddress"` // Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.
}

// Iok8sapiappsv1beta1ScaleSpec represents the Iok8sapiappsv1beta1ScaleSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8sapicorev1HTTPGetAction represents the Iok8sapicorev1HTTPGetAction schema from the OpenAPI specification
type Iok8sapicorev1HTTPGetAction struct {
	Httpheaders []Iok8sapicorev1HTTPHeader `json:"httpHeaders,omitempty"` // Custom headers to set in the request. HTTP allows repeated headers.
	Path string `json:"path,omitempty"` // Path to access on the HTTP server.
	Port string `json:"port"`
	Scheme string `json:"scheme,omitempty"` // Scheme to use for connecting to the host. Defaults to HTTP.
	Host string `json:"host,omitempty"` // Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
}

// Iok8sapiextensionsv1beta1IPBlock represents the Iok8sapiextensionsv1beta1IPBlock schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IPBlock struct {
	Cidr string `json:"cidr"` // CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24"
	Except []string `json:"except,omitempty"` // Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" Except values will be rejected if they are outside the CIDR range
}

// Iok8sapiappsv1RollingUpdateStatefulSetStrategy represents the Iok8sapiappsv1RollingUpdateStatefulSetStrategy schema from the OpenAPI specification
type Iok8sapiappsv1RollingUpdateStatefulSetStrategy struct {
	Partition int `json:"partition,omitempty"` // Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
}

// Iok8skubernetespkgapiv1PersistentVolumeList represents the Iok8skubernetespkgapiv1PersistentVolumeList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeList struct {
	Items []Iok8sapicorev1PersistentVolume `json:"items"` // List of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1EventList represents the Iok8skubernetespkgapiv1EventList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EventList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Event `json:"items"` // List of events
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiextensionsv1beta1DaemonSetStatus represents the Iok8sapiextensionsv1beta1DaemonSetStatus schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DaemonSetStatus struct {
	Updatednumberscheduled int `json:"updatedNumberScheduled,omitempty"` // The total number of nodes that are running updated daemon pod
	Conditions []Iok8sapiextensionsv1beta1DaemonSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a DaemonSet's current state.
	Currentnumberscheduled int `json:"currentNumberScheduled"` // The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Desirednumberscheduled int `json:"desiredNumberScheduled"` // The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Numberavailable int `json:"numberAvailable,omitempty"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numberready int `json:"numberReady"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The most recent generation observed by the daemon set controller.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Numberunavailable int `json:"numberUnavailable,omitempty"` // The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numbermisscheduled int `json:"numberMisscheduled"` // The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
}

// Iok8skubernetespkgapiv1Pod represents the Iok8skubernetespkgapiv1Pod schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Pod struct {
	Status Iok8sapicorev1PodStatus `json:"status,omitempty"` // PodStatus represents information about the status of a pod. Status may trail the actual state of a system.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PodSpec `json:"spec,omitempty"` // PodSpec is a description of a pod.
}

// Iok8sapicorev1Volume represents the Iok8sapicorev1Volume schema from the OpenAPI specification
type Iok8sapicorev1Volume struct {
	Downwardapi Iok8sapicorev1DownwardAPIVolumeSource `json:"downwardAPI,omitempty"` // DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
	Cephfs Iok8sapicorev1CephFSVolumeSource `json:"cephfs,omitempty"` // Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Quobyte Iok8sapicorev1QuobyteVolumeSource `json:"quobyte,omitempty"` // Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Cinder Iok8sapicorev1CinderVolumeSource `json:"cinder,omitempty"` // Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Hostpath Iok8sapicorev1HostPathVolumeSource `json:"hostPath,omitempty"` // Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	Name string `json:"name"` // Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Photonpersistentdisk Iok8sapicorev1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"` // Represents a Photon Controller persistent disk resource.
	Awselasticblockstore Iok8sapicorev1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"` // Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	Projected Iok8sapicorev1ProjectedVolumeSource `json:"projected,omitempty"` // Represents a projected volume source
	Scaleio Iok8sapicorev1ScaleIOVolumeSource `json:"scaleIO,omitempty"` // ScaleIOVolumeSource represents a persistent ScaleIO volume
	Configmap Iok8sapicorev1ConfigMapVolumeSource `json:"configMap,omitempty"` // Adapts a ConfigMap into a volume. The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.
	Iscsi Iok8sapicorev1ISCSIVolumeSource `json:"iscsi,omitempty"` // Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Flocker Iok8sapicorev1FlockerVolumeSource `json:"flocker,omitempty"` // Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Azuredisk Iok8sapicorev1AzureDiskVolumeSource `json:"azureDisk,omitempty"` // AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	Glusterfs Iok8sapicorev1GlusterfsVolumeSource `json:"glusterfs,omitempty"` // Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Gcepersistentdisk Iok8sapicorev1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"` // Represents a Persistent Disk resource in Google Compute Engine. A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	Portworxvolume Iok8sapicorev1PortworxVolumeSource `json:"portworxVolume,omitempty"` // PortworxVolumeSource represents a Portworx volume resource.
	Secret Iok8sapicorev1SecretVolumeSource `json:"secret,omitempty"` // Adapts a Secret into a volume. The contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.
	Flexvolume Iok8sapicorev1FlexVolumeSource `json:"flexVolume,omitempty"` // FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	Nfs Iok8sapicorev1NFSVolumeSource `json:"nfs,omitempty"` // Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Rbd Iok8sapicorev1RBDVolumeSource `json:"rbd,omitempty"` // Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Gitrepo Iok8sapicorev1GitRepoVolumeSource `json:"gitRepo,omitempty"` // Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.
	Storageos Iok8sapicorev1StorageOSVolumeSource `json:"storageos,omitempty"` // Represents a StorageOS persistent volume resource.
	Vspherevolume Iok8sapicorev1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"` // Represents a vSphere volume resource.
	Persistentvolumeclaim Iok8sapicorev1PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"` // PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
	Emptydir Iok8sapicorev1EmptyDirVolumeSource `json:"emptyDir,omitempty"` // Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
	Fc Iok8sapicorev1FCVolumeSource `json:"fc,omitempty"` // Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Azurefile Iok8sapicorev1AzureFileVolumeSource `json:"azureFile,omitempty"` // AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
}

// Iok8skubernetespkgapiv1ConfigMapKeySelector represents the Iok8skubernetespkgapiv1ConfigMapKeySelector schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ConfigMapKeySelector struct {
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or it's key must be defined
	Key string `json:"key"` // The key to select.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8skubernetespkgapiv1ContainerStateTerminated represents the Iok8skubernetespkgapiv1ContainerStateTerminated schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerStateTerminated struct {
	Reason string `json:"reason,omitempty"` // (brief) reason from the last termination of the container
	Signal int `json:"signal,omitempty"` // Signal from the last termination of the container
	Startedat string `json:"startedAt,omitempty"`
	Containerid string `json:"containerID,omitempty"` // Container's ID in the format 'docker://<container_id>'
	Exitcode int `json:"exitCode"` // Exit status from the last termination of the container
	Finishedat string `json:"finishedAt,omitempty"`
	Message string `json:"message,omitempty"` // Message regarding the last termination of the container
}

// Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequest represents the Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequest schema from the OpenAPI specification
type Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequest struct {
	Status Iok8sapicertificatesv1beta1CertificateSigningRequestStatus `json:"status,omitempty"`
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicertificatesv1beta1CertificateSigningRequestSpec `json:"spec,omitempty"` // This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
}

// Iok8skubernetespkgapisautoscalingv1ScaleSpec represents the Iok8skubernetespkgapisautoscalingv1ScaleSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8sapicorev1Lifecycle represents the Iok8sapicorev1Lifecycle schema from the OpenAPI specification
type Iok8sapicorev1Lifecycle struct {
	Poststart Iok8sapicorev1Handler `json:"postStart,omitempty"` // Handler defines a specific action that should be taken
	Prestop Iok8sapicorev1Handler `json:"preStop,omitempty"` // Handler defines a specific action that should be taken
}

// Iok8skubernetespkgapiv1Probe represents the Iok8skubernetespkgapiv1Probe schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Probe struct {
	Exec Iok8sapicorev1ExecAction `json:"exec,omitempty"` // ExecAction describes a "run in container" action.
	Failurethreshold int `json:"failureThreshold,omitempty"` // Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
	Httpget Iok8sapicorev1HTTPGetAction `json:"httpGet,omitempty"` // HTTPGetAction describes an action based on HTTP Get requests.
	Initialdelayseconds int `json:"initialDelaySeconds,omitempty"` // Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	Periodseconds int `json:"periodSeconds,omitempty"` // How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
	Successthreshold int `json:"successThreshold,omitempty"` // Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	Tcpsocket Iok8sapicorev1TCPSocketAction `json:"tcpSocket,omitempty"` // TCPSocketAction describes an action based on opening a socket
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
}

// Iok8skubernetespkgapiv1SecretEnvSource represents the Iok8skubernetespkgapiv1SecretEnvSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SecretEnvSource struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret must be defined
}

// Iok8skubernetespkgapisextensionsv1beta1IngressTLS represents the Iok8skubernetespkgapisextensionsv1beta1IngressTLS schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IngressTLS struct {
	Hosts []string `json:"hosts,omitempty"` // Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Secretname string `json:"secretName,omitempty"` // SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
}

// Iok8skubernetespkgapiv1FCVolumeSource represents the Iok8skubernetespkgapiv1FCVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1FCVolumeSource struct {
	Targetwwns []string `json:"targetWWNs,omitempty"` // Optional: FC target worldwide names (WWNs)
	Wwids []string `json:"wwids,omitempty"` // Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Lun int `json:"lun,omitempty"` // Optional: FC target lun number
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8skubernetespkgapiv1TCPSocketAction represents the Iok8skubernetespkgapiv1TCPSocketAction schema from the OpenAPI specification
type Iok8skubernetespkgapiv1TCPSocketAction struct {
	Port string `json:"port"`
	Host string `json:"host,omitempty"` // Optional: Host name to connect to, defaults to the pod IP.
}

// Iok8skubernetespkgapisextensionsv1beta1DaemonSetList represents the Iok8skubernetespkgapisextensionsv1beta1DaemonSetList schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DaemonSetList struct {
	Items []Iok8sapiextensionsv1beta1DaemonSet `json:"items"` // A list of daemon sets.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1Scale represents the Iok8sapiextensionsv1beta1Scale schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1Scale struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1ScaleSpec `json:"spec,omitempty"` // describes the attributes of a scale subresource
	Status Iok8sapiextensionsv1beta1ScaleStatus `json:"status,omitempty"` // represents the current status of a scale subresource.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisauthorizationv1beta1SelfSubjectAccessReview represents the Iok8skubernetespkgapisauthorizationv1beta1SelfSubjectAccessReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1SelfSubjectAccessReview struct {
	Status Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1beta1SelfSubjectAccessReviewSpec `json:"spec"` // SelfSubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
}

// Iok8sapicorev1ContainerStateRunning represents the Iok8sapicorev1ContainerStateRunning schema from the OpenAPI specification
type Iok8sapicorev1ContainerStateRunning struct {
	Startedat string `json:"startedAt,omitempty"`
}

// Iok8skubernetespkgapisstoragev1StorageClassList represents the Iok8skubernetespkgapisstoragev1StorageClassList schema from the OpenAPI specification
type Iok8skubernetespkgapisstoragev1StorageClassList struct {
	Items []Iok8sapistoragev1StorageClass `json:"items"` // Items is the list of StorageClasses
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1PodList represents the Iok8sapicorev1PodList schema from the OpenAPI specification
type Iok8sapicorev1PodList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Pod `json:"items"` // List of pods. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapisadmissionregistrationv1alpha1InitializerConfiguration represents the Iok8skubernetespkgapisadmissionregistrationv1alpha1InitializerConfiguration schema from the OpenAPI specification
type Iok8skubernetespkgapisadmissionregistrationv1alpha1InitializerConfiguration struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Initializers []Iok8sapiadmissionregistrationv1alpha1Initializer `json:"initializers,omitempty"` // Initializers is a list of resources and their default initializers Order-sensitive. When merging multiple InitializerConfigurations, we sort the initializers from different InitializerConfigurations by the name of the InitializerConfigurations; the order of the initializers from the same InitializerConfiguration is preserved.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta1StatefulSetCondition represents the Iok8sapiappsv1beta1StatefulSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1beta1StatefulSetCondition struct {
	TypeField string `json:"type"` // Type of statefulset condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
}

// Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicyList represents the Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicyList schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicyList struct {
	Items []Iok8sapiextensionsv1beta1PodSecurityPolicy `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapistoragev1StorageClass represents the Iok8sapistoragev1StorageClass schema from the OpenAPI specification
type Iok8sapistoragev1StorageClass struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Mountoptions []string `json:"mountOptions,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	Parameters map[string]interface{} `json:"parameters,omitempty"` // Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Reclaimpolicy string `json:"reclaimPolicy,omitempty"` // Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Allowvolumeexpansion bool `json:"allowVolumeExpansion,omitempty"` // AllowVolumeExpansion shows whether the storage class allow volume expand
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Provisioner string `json:"provisioner"` // Provisioner indicates the type of the provisioner.
	Volumebindingmode string `json:"volumeBindingMode,omitempty"` // VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is alpha-level and is only honored by servers that enable the VolumeScheduling feature.
}

// Iok8skubernetespkgapisextensionsv1beta1DeploymentRollback represents the Iok8skubernetespkgapisextensionsv1beta1DeploymentRollback schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DeploymentRollback struct {
	Name string `json:"name"` // Required: This must match the Name of a deployment.
	Rollbackto Iok8sapiextensionsv1beta1RollbackConfig `json:"rollbackTo"` // DEPRECATED.
	Updatedannotations map[string]interface{} `json:"updatedAnnotations,omitempty"` // The annotations to be updated to a deployment
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1StorageOSVolumeSource represents the Iok8skubernetespkgapiv1StorageOSVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1StorageOSVolumeSource struct {
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Volumename string `json:"volumeName,omitempty"` // VolumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.
	Volumenamespace string `json:"volumeNamespace,omitempty"` // VolumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8skubernetespkgapisautoscalingv1CrossVersionObjectReference represents the Iok8skubernetespkgapisautoscalingv1CrossVersionObjectReference schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1CrossVersionObjectReference struct {
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent
	Kind string `json:"kind"` // Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
	Name string `json:"name"` // Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
}

// Iok8sapiextensionsv1beta1DeploymentSpec represents the Iok8sapiextensionsv1beta1DeploymentSpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DeploymentSpec struct {
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused and will not be processed by the deployment controller.
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. This is not set by default.
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Strategy Iok8sapiextensionsv1beta1DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Rollbackto Iok8sapiextensionsv1beta1RollbackConfig `json:"rollbackTo,omitempty"` // DEPRECATED.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// Iok8skubernetespkgapiv1LocalVolumeSource represents the Iok8skubernetespkgapiv1LocalVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LocalVolumeSource struct {
	Path string `json:"path"` // The full path to the volume on the node For alpha, this path must be a directory Once block as a source is supported, then this path can point to a block device
}

// Iok8sapiauthorizationv1SubjectRulesReviewStatus represents the Iok8sapiauthorizationv1SubjectRulesReviewStatus schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectRulesReviewStatus struct {
	Nonresourcerules []Iok8sapiauthorizationv1NonResourceRule `json:"nonResourceRules"` // NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	Resourcerules []Iok8sapiauthorizationv1ResourceRule `json:"resourceRules"` // ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
	Incomplete bool `json:"incomplete"` // Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.
}

// Iok8sapiappsv1beta2ReplicaSetStatus represents the Iok8sapiappsv1beta2ReplicaSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta2ReplicaSetStatus struct {
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replica set.
	Replicas int `json:"replicas"` // Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replica set.
	Conditions []Iok8sapiappsv1beta2ReplicaSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replica set's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replicaset.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
}

// Iok8skubernetespkgapisbatchv2alpha1CronJobStatus represents the Iok8skubernetespkgapisbatchv2alpha1CronJobStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv2alpha1CronJobStatus struct {
	Active []Iok8sapicorev1ObjectReference `json:"active,omitempty"` // A list of pointers to currently running jobs.
	Lastscheduletime string `json:"lastScheduleTime,omitempty"`
}

// Iok8sapicorev1EventSeries represents the Iok8sapicorev1EventSeries schema from the OpenAPI specification
type Iok8sapicorev1EventSeries struct {
	Count int `json:"count,omitempty"` // Number of occurrences in this series up to the last heartbeat time
	Lastobservedtime string `json:"lastObservedTime,omitempty"`
	State string `json:"state,omitempty"` // State of this Series: Ongoing or Finished
}

// Iok8sapicorev1LoadBalancerStatus represents the Iok8sapicorev1LoadBalancerStatus schema from the OpenAPI specification
type Iok8sapicorev1LoadBalancerStatus struct {
	Ingress []Iok8sapicorev1LoadBalancerIngress `json:"ingress,omitempty"` // Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
}

// Iok8sapicorev1ContainerImage represents the Iok8sapicorev1ContainerImage schema from the OpenAPI specification
type Iok8sapicorev1ContainerImage struct {
	Sizebytes int64 `json:"sizeBytes,omitempty"` // The size of the image in bytes.
	Names []string `json:"names"` // Names by which this image is known. e.g. ["k8s.gcr.io/hyperkube:v1.0.7", "dockerhub.io/google_containers/hyperkube:v1.0.7"]
}

// Iok8sapicorev1ResourceQuota represents the Iok8sapicorev1ResourceQuota schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuota struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ResourceQuotaSpec `json:"spec,omitempty"` // ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
	Status Iok8sapicorev1ResourceQuotaStatus `json:"status,omitempty"` // ResourceQuotaStatus defines the enforced hard limits and observed use.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapirbacv1RoleList represents the Iok8sapirbacv1RoleList schema from the OpenAPI specification
type Iok8sapirbacv1RoleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1Role `json:"items"` // Items is a list of Roles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapieventsv1beta1EventSeries represents the Iok8sapieventsv1beta1EventSeries schema from the OpenAPI specification
type Iok8sapieventsv1beta1EventSeries struct {
	Count int `json:"count"` // Number of occurrences in this series up to the last heartbeat time
	Lastobservedtime string `json:"lastObservedTime"`
	State string `json:"state"` // Information whether this series is ongoing or finished.
}

// Iok8sapipolicyv1beta1PodDisruptionBudgetList represents the Iok8sapipolicyv1beta1PodDisruptionBudgetList schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodDisruptionBudgetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapipolicyv1beta1PodDisruptionBudget `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisrbacv1beta1ClusterRole represents the Iok8skubernetespkgapisrbacv1beta1ClusterRole schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1ClusterRole struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1beta1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this ClusterRole
	Aggregationrule Iok8sapirbacv1beta1AggregationRule `json:"aggregationRule,omitempty"` // AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiautoscalingv2beta1ResourceMetricSource represents the Iok8sapiautoscalingv2beta1ResourceMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1ResourceMetricSource struct {
	Name string `json:"name"` // name is the name of the resource in question.
	Targetaverageutilization int `json:"targetAverageUtilization,omitempty"` // targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	Targetaveragevalue string `json:"targetAverageValue,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Schema GeneratedType `json:"Schema"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
	Property []string `json:"Property"`
}

// Iok8sapiextensionsv1beta1DaemonSetCondition represents the Iok8sapiextensionsv1beta1DaemonSetCondition schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DaemonSetCondition struct {
	TypeField string `json:"type"` // Type of DaemonSet condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
}

// Iok8sapibatchv1beta1CronJobList represents the Iok8sapibatchv1beta1CronJobList schema from the OpenAPI specification
type Iok8sapibatchv1beta1CronJobList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapibatchv1beta1CronJob `json:"items"` // items is the list of CronJobs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec GeneratedType `json:"spec,omitempty"` // CustomResourceDefinitionSpec describes how a user wants their resource to appear
	Status GeneratedType `json:"status,omitempty"` // CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
}

// Iok8sapicorev1ReplicationControllerList represents the Iok8sapicorev1ReplicationControllerList schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ReplicationController `json:"items"` // List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiauthorizationv1beta1SelfSubjectRulesReview represents the Iok8sapiauthorizationv1beta1SelfSubjectRulesReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SelfSubjectRulesReview struct {
	Spec Iok8sapiauthorizationv1beta1SelfSubjectRulesReviewSpec `json:"spec"`
	Status Iok8sapiauthorizationv1beta1SubjectRulesReviewStatus `json:"status,omitempty"` // SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapiv1ContainerPort represents the Iok8skubernetespkgapiv1ContainerPort schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerPort struct {
	Containerport int `json:"containerPort"` // Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
	Hostip string `json:"hostIP,omitempty"` // What host IP to bind the external port to.
	Hostport int `json:"hostPort,omitempty"` // Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	Name string `json:"name,omitempty"` // If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
	Protocol string `json:"protocol,omitempty"` // Protocol for port. Must be UDP or TCP. Defaults to "TCP".
}

// Iok8sapirbacv1beta1RoleBinding represents the Iok8sapirbacv1beta1RoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1beta1RoleBinding struct {
	Roleref Iok8sapirbacv1beta1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1beta1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapisappsv1beta1StatefulSetSpec represents the Iok8skubernetespkgapisappsv1beta1StatefulSetSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1StatefulSetSpec struct {
	Servicename string `json:"serviceName"` // serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Updatestrategy Iok8sapiappsv1beta1StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"` // StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
	Volumeclaimtemplates []Iok8sapicorev1PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"` // volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	Podmanagementpolicy string `json:"podManagementPolicy,omitempty"` // podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	Replicas int `json:"replicas,omitempty"` // replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapistoragev1beta1VolumeAttachmentSource represents the Iok8sapistoragev1beta1VolumeAttachmentSource schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttachmentSource struct {
	Persistentvolumename string `json:"persistentVolumeName,omitempty"` // Name of the persistent volume to attach.
}

// Iok8sapimachinerypkgapismetav1DeleteOptions represents the Iok8sapimachinerypkgapismetav1DeleteOptions schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1DeleteOptions struct {
	Preconditions Iok8sapimachinerypkgapismetav1Preconditions `json:"preconditions,omitempty"` // Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
	Propagationpolicy string `json:"propagationPolicy,omitempty"` // Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Graceperiodseconds int64 `json:"gracePeriodSeconds,omitempty"` // The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Orphandependents bool `json:"orphanDependents,omitempty"` // Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.
}

// Iok8sapimachinerypkgapismetav1Initializers represents the Iok8sapimachinerypkgapismetav1Initializers schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Initializers struct {
	Pending []Iok8sapimachinerypkgapismetav1Initializer `json:"pending"` // Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.
	Result Iok8sapimachinerypkgapismetav1Status `json:"result,omitempty"` // Status is a return value for calls that don't return other objects.
}

// Iok8skubernetespkgapiv1SecretKeySelector represents the Iok8skubernetespkgapiv1SecretKeySelector schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SecretKeySelector struct {
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or it's key must be defined
	Key string `json:"key"` // The key of the secret to select from. Must be a valid secret key.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8sapibatchv1beta1CronJob represents the Iok8sapibatchv1beta1CronJob schema from the OpenAPI specification
type Iok8sapibatchv1beta1CronJob struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1beta1CronJobSpec `json:"spec,omitempty"` // CronJobSpec describes how the job execution will look like and when it will actually run.
	Status Iok8sapibatchv1beta1CronJobStatus `json:"status,omitempty"` // CronJobStatus represents the current state of a cron job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1Affinity represents the Iok8skubernetespkgapiv1Affinity schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Affinity struct {
	Nodeaffinity Iok8sapicorev1NodeAffinity `json:"nodeAffinity,omitempty"` // Node affinity is a group of node affinity scheduling rules.
	Podaffinity Iok8sapicorev1PodAffinity `json:"podAffinity,omitempty"` // Pod affinity is a group of inter pod affinity scheduling rules.
	Podantiaffinity Iok8sapicorev1PodAntiAffinity `json:"podAntiAffinity,omitempty"` // Pod anti affinity is a group of inter pod anti affinity scheduling rules.
}

// Iok8sapiauthenticationv1TokenReviewSpec represents the Iok8sapiauthenticationv1TokenReviewSpec schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenReviewSpec struct {
	Token string `json:"token,omitempty"` // Token is the opaque bearer token.
}

// Iok8skubernetespkgapiv1NodeSystemInfo represents the Iok8skubernetespkgapiv1NodeSystemInfo schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeSystemInfo struct {
	Kubeproxyversion string `json:"kubeProxyVersion"` // KubeProxy Version reported by the node.
	Osimage string `json:"osImage"` // OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	Architecture string `json:"architecture"` // The Architecture reported by the node
	Kernelversion string `json:"kernelVersion"` // Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	Operatingsystem string `json:"operatingSystem"` // The Operating System reported by the node
	Containerruntimeversion string `json:"containerRuntimeVersion"` // ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	Systemuuid string `json:"systemUUID"` // SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	Bootid string `json:"bootID"` // Boot ID reported by the node.
	Kubeletversion string `json:"kubeletVersion"` // Kubelet Version reported by the node.
	Machineid string `json:"machineID"` // MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
}

// Iok8skubernetespkgapiv1NodeSpec represents the Iok8skubernetespkgapiv1NodeSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeSpec struct {
	Taints []Iok8sapicorev1Taint `json:"taints,omitempty"` // If specified, the node's taints.
	Unschedulable bool `json:"unschedulable,omitempty"` // Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	Configsource Iok8sapicorev1NodeConfigSource `json:"configSource,omitempty"` // NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil.
	Externalid string `json:"externalID,omitempty"` // External ID of the node assigned by some machine database (e.g. a cloud provider). Deprecated.
	Podcidr string `json:"podCIDR,omitempty"` // PodCIDR represents the pod IP range assigned to the node.
	Providerid string `json:"providerID,omitempty"` // ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
}

// Iok8sapicorev1ComponentStatusList represents the Iok8sapicorev1ComponentStatusList schema from the OpenAPI specification
type Iok8sapicorev1ComponentStatusList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ComponentStatus `json:"items"` // List of ComponentStatus objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1beta2ScaleStatus represents the Iok8sapiappsv1beta2ScaleStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta2ScaleStatus struct {
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
	Selector map[string]interface{} `json:"selector,omitempty"` // label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Targetselector string `json:"targetSelector,omitempty"` // label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
}

// Iok8sapicorev1SessionAffinityConfig represents the Iok8sapicorev1SessionAffinityConfig schema from the OpenAPI specification
type Iok8sapicorev1SessionAffinityConfig struct {
	Clientip Iok8sapicorev1ClientIPConfig `json:"clientIP,omitempty"` // ClientIPConfig represents the configurations of Client IP based session affinity.
}

// Iok8sapiextensionsv1beta1IngressSpec represents the Iok8sapiextensionsv1beta1IngressSpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IngressSpec struct {
	Rules []Iok8sapiextensionsv1beta1IngressRule `json:"rules,omitempty"` // A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Tls []Iok8sapiextensionsv1beta1IngressTLS `json:"tls,omitempty"` // TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	Backend Iok8sapiextensionsv1beta1IngressBackend `json:"backend,omitempty"` // IngressBackend describes all endpoints for a given service and port.
}

// Iok8skubernetespkgapiv1PodSecurityContext represents the Iok8skubernetespkgapiv1PodSecurityContext schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodSecurityContext struct {
	Supplementalgroups []int64 `json:"supplementalGroups,omitempty"` // A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container.
	Fsgroup int64 `json:"fsGroup,omitempty"` // A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.
	Runasgroup int64 `json:"runAsGroup,omitempty"` // The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	Runasnonroot bool `json:"runAsNonRoot,omitempty"` // Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Runasuser int64 `json:"runAsUser,omitempty"` // The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
}

// Iok8sapiappsv1beta1Scale represents the Iok8sapiappsv1beta1Scale schema from the OpenAPI specification
type Iok8sapiappsv1beta1Scale struct {
	Status Iok8sapiappsv1beta1ScaleStatus `json:"status,omitempty"` // ScaleStatus represents the current status of a scale subresource.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta1ScaleSpec `json:"spec,omitempty"` // ScaleSpec describes the attributes of a scale subresource
}

// Iok8sapicorev1PhotonPersistentDiskVolumeSource represents the Iok8sapicorev1PhotonPersistentDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1PhotonPersistentDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Pdid string `json:"pdID"` // ID that identifies Photon Controller persistent disk
}

// Iok8skubernetespkgapisextensionsv1beta1IngressSpec represents the Iok8skubernetespkgapisextensionsv1beta1IngressSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IngressSpec struct {
	Rules []Iok8sapiextensionsv1beta1IngressRule `json:"rules,omitempty"` // A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Tls []Iok8sapiextensionsv1beta1IngressTLS `json:"tls,omitempty"` // TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	Backend Iok8sapiextensionsv1beta1IngressBackend `json:"backend,omitempty"` // IngressBackend describes all endpoints for a given service and port.
}

// Iok8sapicertificatesv1beta1CertificateSigningRequestCondition represents the Iok8sapicertificatesv1beta1CertificateSigningRequestCondition schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1CertificateSigningRequestCondition struct {
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
	Message string `json:"message,omitempty"` // human readable message with details about the request state
	Reason string `json:"reason,omitempty"` // brief reason for the request state
	TypeField string `json:"type"` // request approval state, currently Approved or Denied.
}

// Iok8skubernetespkgapisappsv1beta1StatefulSet represents the Iok8skubernetespkgapisappsv1beta1StatefulSet schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1StatefulSet struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta1StatefulSetSpec `json:"spec,omitempty"` // A StatefulSetSpec is the specification of a StatefulSet.
	Status Iok8sapiappsv1beta1StatefulSetStatus `json:"status,omitempty"` // StatefulSetStatus represents the current state of a StatefulSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1EnvFromSource represents the Iok8skubernetespkgapiv1EnvFromSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EnvFromSource struct {
	Configmapref Iok8sapicorev1ConfigMapEnvSource `json:"configMapRef,omitempty"` // ConfigMapEnvSource selects a ConfigMap to populate the environment variables with. The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
	Prefix string `json:"prefix,omitempty"` // An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Secretref Iok8sapicorev1SecretEnvSource `json:"secretRef,omitempty"` // SecretEnvSource selects a Secret to populate the environment variables with. The contents of the target Secret's Data field will represent the key-value pairs as environment variables.
}

// Iok8sapimachinerypkgapismetav1StatusCause represents the Iok8sapimachinerypkgapismetav1StatusCause schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1StatusCause struct {
	Field string `json:"field,omitempty"` // The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed. Fields may appear more than once in an array of causes due to fields having multiple errors. Optional. Examples: "name" - the field "name" on the current resource "items[0].name" - the field "name" on the first array entry in "items"
	Message string `json:"message,omitempty"` // A human-readable description of the cause of the error. This field may be presented as-is to a reader.
	Reason string `json:"reason,omitempty"` // A machine-readable description of the cause of the error. If this value is empty there is no information available.
}

// Iok8sapipolicyv1beta1SELinuxStrategyOptions represents the Iok8sapipolicyv1beta1SELinuxStrategyOptions schema from the OpenAPI specification
type Iok8sapipolicyv1beta1SELinuxStrategyOptions struct {
	Rule string `json:"rule"` // type is the strategy that will dictate the allowable labels that may be set.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
}

// Iok8sapiauthorizationv1SelfSubjectAccessReview represents the Iok8sapiauthorizationv1SelfSubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectAccessReview struct {
	Spec Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec `json:"spec"` // SelfSubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerCondition represents the Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerCondition schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // message is a human-readable explanation containing details about the transition
	Reason string `json:"reason,omitempty"` // reason is the reason for the condition's last transition.
	Status string `json:"status"` // status is the status of the condition (True, False, Unknown)
	TypeField string `json:"type"` // type describes the current condition
}

// Iok8sapipolicyv1beta1PodSecurityPolicyList represents the Iok8sapipolicyv1beta1PodSecurityPolicyList schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodSecurityPolicyList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapipolicyv1beta1PodSecurityPolicy `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapiv1EventSource represents the Iok8skubernetespkgapiv1EventSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EventSource struct {
	Component string `json:"component,omitempty"` // Component from which the event is generated.
	Host string `json:"host,omitempty"` // Node name on which the event is generated.
}

// Iok8sapibatchv1JobList represents the Iok8sapibatchv1JobList schema from the OpenAPI specification
type Iok8sapibatchv1JobList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapibatchv1Job `json:"items"` // items is the list of Jobs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1PersistentVolumeClaimSpec represents the Iok8sapicorev1PersistentVolumeClaimSpec schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimSpec struct {
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Storageclassname string `json:"storageClassName,omitempty"` // Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	Volumemode string `json:"volumeMode,omitempty"` // volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. This is an alpha feature and may change in the future.
	Volumename string `json:"volumeName,omitempty"` // VolumeName is the binding reference to the PersistentVolume backing this claim.
	Accessmodes []string `json:"accessModes,omitempty"` // AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
}

// Iok8sapiauthenticationv1UserInfo represents the Iok8sapiauthenticationv1UserInfo schema from the OpenAPI specification
type Iok8sapiauthenticationv1UserInfo struct {
	Uid string `json:"uid,omitempty"` // A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	Username string `json:"username,omitempty"` // The name that uniquely identifies this user among all active users.
	Extra map[string]interface{} `json:"extra,omitempty"` // Any additional information provided by the authenticator.
	Groups []string `json:"groups,omitempty"` // The names of groups this user is a part of.
}

// Iok8sapipolicyv1beta1HostPortRange represents the Iok8sapipolicyv1beta1HostPortRange schema from the OpenAPI specification
type Iok8sapipolicyv1beta1HostPortRange struct {
	Max int `json:"max"` // max is the end of the range, inclusive.
	Min int `json:"min"` // min is the start of the range, inclusive.
}

// Iok8skubernetespkgapisautoscalingv1ScaleStatus represents the Iok8skubernetespkgapisautoscalingv1ScaleStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1ScaleStatus struct {
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
	Selector string `json:"selector,omitempty"` // label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors
}

// Iok8sapicorev1PodSpec represents the Iok8sapicorev1PodSpec schema from the OpenAPI specification
type Iok8sapicorev1PodSpec struct {
	Dnsconfig Iok8sapicorev1PodDNSConfig `json:"dnsConfig,omitempty"` // PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
	Affinity Iok8sapicorev1Affinity `json:"affinity,omitempty"` // Affinity is a group of affinity scheduling rules.
	Containers []Iok8sapicorev1Container `json:"containers"` // List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.
	Terminationgraceperiodseconds int64 `json:"terminationGracePeriodSeconds,omitempty"` // Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
	Volumes []Iok8sapicorev1Volume `json:"volumes,omitempty"` // List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes
	Schedulername string `json:"schedulerName,omitempty"` // If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
	Restartpolicy string `json:"restartPolicy,omitempty"` // Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	Serviceaccountname string `json:"serviceAccountName,omitempty"` // ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Subdomain string `json:"subdomain,omitempty"` // If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Tolerations []Iok8sapicorev1Toleration `json:"tolerations,omitempty"` // If specified, the pod's tolerations.
	Hostpid bool `json:"hostPID,omitempty"` // Use the host's pid namespace. Optional: Default to false.
	Hostname string `json:"hostname,omitempty"` // Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
	Imagepullsecrets []Iok8sapicorev1LocalObjectReference `json:"imagePullSecrets,omitempty"` // ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	Dnspolicy string `json:"dnsPolicy,omitempty"` // Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	Securitycontext Iok8sapicorev1PodSecurityContext `json:"securityContext,omitempty"` // PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext. Field values of container.securityContext take precedence over field values of PodSecurityContext.
	Priorityclassname string `json:"priorityClassName,omitempty"` // If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	Shareprocessnamespace bool `json:"shareProcessNamespace,omitempty"` // Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. This field is alpha-level and is honored only by servers that enable the PodShareProcessNamespace feature.
	Automountserviceaccounttoken bool `json:"automountServiceAccountToken,omitempty"` // AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	Initcontainers []Iok8sapicorev1Container `json:"initContainers,omitempty"` // List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	Nodeselector map[string]interface{} `json:"nodeSelector,omitempty"` // NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	Hostipc bool `json:"hostIPC,omitempty"` // Use the host's ipc namespace. Optional: Default to false.
	Hostnetwork bool `json:"hostNetwork,omitempty"` // Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.
	Nodename string `json:"nodeName,omitempty"` // NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	Priority int `json:"priority,omitempty"` // The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
	Activedeadlineseconds int64 `json:"activeDeadlineSeconds,omitempty"` // Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
	Hostaliases []Iok8sapicorev1HostAlias `json:"hostAliases,omitempty"` // HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.
	Serviceaccount string `json:"serviceAccount,omitempty"` // DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
}

// Iok8sapimachinerypkgapismetav1StatusDetails represents the Iok8sapimachinerypkgapismetav1StatusDetails schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1StatusDetails struct {
	Causes []Iok8sapimachinerypkgapismetav1StatusCause `json:"causes,omitempty"` // The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Group string `json:"group,omitempty"` // The group attribute of the resource associated with the status StatusReason.
	Kind string `json:"kind,omitempty"` // The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name,omitempty"` // The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Retryafterseconds int `json:"retryAfterSeconds,omitempty"` // If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	Uid string `json:"uid,omitempty"` // UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
}

// Iok8sapiappsv1beta1DeploymentList represents the Iok8sapiappsv1beta1DeploymentList schema from the OpenAPI specification
type Iok8sapiappsv1beta1DeploymentList struct {
	Items []Iok8sapiappsv1beta1Deployment `json:"items"` // Items is the list of Deployments.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisextensionsv1beta1HTTPIngressPath represents the Iok8skubernetespkgapisextensionsv1beta1HTTPIngressPath schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1HTTPIngressPath struct {
	Backend Iok8sapiextensionsv1beta1IngressBackend `json:"backend"` // IngressBackend describes all endpoints for a given service and port.
	Path string `json:"path,omitempty"` // Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
}

// Iok8skubernetespkgapisextensionsv1beta1Scale represents the Iok8skubernetespkgapisextensionsv1beta1Scale schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1Scale struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1ScaleSpec `json:"spec,omitempty"` // describes the attributes of a scale subresource
	Status Iok8sapiextensionsv1beta1ScaleStatus `json:"status,omitempty"` // represents the current status of a scale subresource.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapistoragev1beta1VolumeAttachmentList represents the Iok8sapistoragev1beta1VolumeAttachmentList schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttachmentList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapistoragev1beta1VolumeAttachment `json:"items"` // Items is the list of VolumeAttachments
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapinetworkingv1NetworkPolicyPort represents the Iok8sapinetworkingv1NetworkPolicyPort schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyPort struct {
	Port string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"` // The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.
}

// Iok8skubernetespkgapiv1HTTPHeader represents the Iok8skubernetespkgapiv1HTTPHeader schema from the OpenAPI specification
type Iok8skubernetespkgapiv1HTTPHeader struct {
	Name string `json:"name"` // The header field name
	Value string `json:"value"` // The header field value
}

// Iok8skubernetespkgapisextensionsv1beta1RollingUpdateDeployment represents the Iok8skubernetespkgapisextensionsv1beta1RollingUpdateDeployment schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1RollingUpdateDeployment struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
	Maxsurge string `json:"maxSurge,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Openapiv3schema GeneratedType `json:"openAPIV3Schema,omitempty"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
}

// Iok8skubernetespkgapisappsv1beta1ControllerRevisionList represents the Iok8skubernetespkgapisappsv1beta1ControllerRevisionList schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1ControllerRevisionList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta1ControllerRevision `json:"items"` // Items is the list of ControllerRevisions
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapipolicyv1beta1Eviction represents the Iok8sapipolicyv1beta1Eviction schema from the OpenAPI specification
type Iok8sapipolicyv1beta1Eviction struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Deleteoptions Iok8sapimachinerypkgapismetav1DeleteOptions `json:"deleteOptions,omitempty"` // DeleteOptions may be provided when deleting an API object.
}

// Iok8sapiappsv1beta2RollingUpdateStatefulSetStrategy represents the Iok8sapiappsv1beta2RollingUpdateStatefulSetStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta2RollingUpdateStatefulSetStrategy struct {
	Partition int `json:"partition,omitempty"` // Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
}

// Iok8sapirbacv1alpha1RoleBinding represents the Iok8sapirbacv1alpha1RoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1alpha1RoleBinding struct {
	Roleref Iok8sapirbacv1alpha1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1alpha1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapirbacv1beta1RoleBindingList represents the Iok8sapirbacv1beta1RoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1beta1RoleBindingList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1RoleBinding `json:"items"` // Items is a list of RoleBindings
}

// Iok8sapieventsv1beta1EventList represents the Iok8sapieventsv1beta1EventList schema from the OpenAPI specification
type Iok8sapieventsv1beta1EventList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapieventsv1beta1Event `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta2DeploymentCondition represents the Iok8sapiappsv1beta2DeploymentCondition schema from the OpenAPI specification
type Iok8sapiappsv1beta2DeploymentCondition struct {
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
}

// Iok8sapirbacv1PolicyRule represents the Iok8sapirbacv1PolicyRule schema from the OpenAPI specification
type Iok8sapirbacv1PolicyRule struct {
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. ResourceAll represents all resources.
	Verbs []string `json:"verbs"` // Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []GeneratedType `json:"items"` // Items individual CustomResourceDefinitions
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisauthorizationv1SubjectAccessReview represents the Iok8skubernetespkgapisauthorizationv1SubjectAccessReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1SubjectAccessReview struct {
	Spec Iok8sapiauthorizationv1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapirbacv1alpha1ClusterRoleList represents the Iok8sapirbacv1alpha1ClusterRoleList schema from the OpenAPI specification
type Iok8sapirbacv1alpha1ClusterRoleList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1ClusterRole `json:"items"` // Items is a list of ClusterRoles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1Binding represents the Iok8skubernetespkgapiv1Binding schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Binding struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Target Iok8sapicorev1ObjectReference `json:"target"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisauthenticationv1beta1TokenReview represents the Iok8skubernetespkgapisauthenticationv1beta1TokenReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1beta1TokenReview struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthenticationv1beta1TokenReviewSpec `json:"spec"` // TokenReviewSpec is a description of the token authentication request.
	Status Iok8sapiauthenticationv1beta1TokenReviewStatus `json:"status,omitempty"` // TokenReviewStatus is the result of the token authentication request.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Conditions []GeneratedType `json:"conditions"` // Conditions indicate state for particular aspects of a CustomResourceDefinition
	Acceptednames GeneratedType `json:"acceptedNames"` // CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
}

// Iok8sapicorev1TCPSocketAction represents the Iok8sapicorev1TCPSocketAction schema from the OpenAPI specification
type Iok8sapicorev1TCPSocketAction struct {
	Host string `json:"host,omitempty"` // Optional: Host name to connect to, defaults to the pod IP.
	Port string `json:"port"`
}

// Iok8skubernetespkgapiv1ComponentStatus represents the Iok8skubernetespkgapiv1ComponentStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ComponentStatus struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Conditions []Iok8sapicorev1ComponentCondition `json:"conditions,omitempty"` // List of component conditions observed
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiextensionsv1beta1Ingress represents the Iok8sapiextensionsv1beta1Ingress schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1Ingress struct {
	Spec Iok8sapiextensionsv1beta1IngressSpec `json:"spec,omitempty"` // IngressSpec describes the Ingress the user wishes to exist.
	Status Iok8sapiextensionsv1beta1IngressStatus `json:"status,omitempty"` // IngressStatus describe the current state of the Ingress.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1NamespaceSpec represents the Iok8sapicorev1NamespaceSpec schema from the OpenAPI specification
type Iok8sapicorev1NamespaceSpec struct {
	Finalizers []string `json:"finalizers,omitempty"` // Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
}

// Iok8sapicorev1FlexPersistentVolumeSource represents the Iok8sapicorev1FlexPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FlexPersistentVolumeSource struct {
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Driver string `json:"driver"` // Driver is the name of the driver to use for this volume.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	Options map[string]interface{} `json:"options,omitempty"` // Optional: Extra command options if any.
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapicorev1PersistentVolumeClaimVolumeSource represents the Iok8sapicorev1PersistentVolumeClaimVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // Will force the ReadOnly setting in VolumeMounts. Default false.
	Claimname string `json:"claimName"` // ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
}

// Iok8sapicorev1GitRepoVolumeSource represents the Iok8sapicorev1GitRepoVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GitRepoVolumeSource struct {
	Revision string `json:"revision,omitempty"` // Commit hash for the specified revision.
	Directory string `json:"directory,omitempty"` // Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	Repository string `json:"repository"` // Repository URL
}

// Iok8skubernetespkgapiv1ServiceAccount represents the Iok8skubernetespkgapiv1ServiceAccount schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ServiceAccount struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Secrets []Iok8sapicorev1ObjectReference `json:"secrets,omitempty"` // Secrets is the list of secrets allowed to be used by pods running using this ServiceAccount. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Automountserviceaccounttoken bool `json:"automountServiceAccountToken,omitempty"` // AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	Imagepullsecrets []Iok8sapicorev1LocalObjectReference `json:"imagePullSecrets,omitempty"` // ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
}

// Iok8skubernetespkgapisstoragev1beta1StorageClassList represents the Iok8skubernetespkgapisstoragev1beta1StorageClassList schema from the OpenAPI specification
type Iok8skubernetespkgapisstoragev1beta1StorageClassList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapistoragev1beta1StorageClass `json:"items"` // Items is the list of StorageClasses
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisextensionsv1beta1DeploymentStatus represents the Iok8skubernetespkgapisextensionsv1beta1DeploymentStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DeploymentStatus struct {
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of ready pods targeted by this deployment.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Conditions []Iok8sapiextensionsv1beta1DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
}

// Iok8skubernetespkgapiv1PreferredSchedulingTerm represents the Iok8skubernetespkgapiv1PreferredSchedulingTerm schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PreferredSchedulingTerm struct {
	Preference Iok8sapicorev1NodeSelectorTerm `json:"preference"` // A null or empty node selector term matches no objects.
	Weight int `json:"weight"` // Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
}

// Iok8skubernetespkgapiv1ReplicationControllerStatus represents the Iok8skubernetespkgapiv1ReplicationControllerStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ReplicationControllerStatus struct {
	Conditions []Iok8sapicorev1ReplicationControllerCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replication controller's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replication controller.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed replication controller.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replication controller.
	Replicas int `json:"replicas"` // Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replication controller.
}

// Iok8sapistoragev1StorageClassList represents the Iok8sapistoragev1StorageClassList schema from the OpenAPI specification
type Iok8sapistoragev1StorageClassList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapistoragev1StorageClass `json:"items"` // Items is the list of StorageClasses
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapimachinerypkgapismetav1APIResourceList represents the Iok8sapimachinerypkgapismetav1APIResourceList schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIResourceList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Groupversion string `json:"groupVersion"` // groupVersion is the group and version this APIResourceList is for.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Resources []Iok8sapimachinerypkgapismetav1APIResource `json:"resources"` // resources contains the name of the resources and if they are namespaced.
}

// Iok8sapiextensionsv1beta1ReplicaSetStatus represents the Iok8sapiextensionsv1beta1ReplicaSetStatus schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ReplicaSetStatus struct {
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replicaset.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replica set.
	Replicas int `json:"replicas"` // Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replica set.
	Conditions []Iok8sapiextensionsv1beta1ReplicaSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replica set's current state.
}

// Iok8sapicorev1Affinity represents the Iok8sapicorev1Affinity schema from the OpenAPI specification
type Iok8sapicorev1Affinity struct {
	Podantiaffinity Iok8sapicorev1PodAntiAffinity `json:"podAntiAffinity,omitempty"` // Pod anti affinity is a group of inter pod anti affinity scheduling rules.
	Nodeaffinity Iok8sapicorev1NodeAffinity `json:"nodeAffinity,omitempty"` // Node affinity is a group of node affinity scheduling rules.
	Podaffinity Iok8sapicorev1PodAffinity `json:"podAffinity,omitempty"` // Pod affinity is a group of inter pod affinity scheduling rules.
}

// Iok8skubernetespkgapiv1StorageOSPersistentVolumeSource represents the Iok8skubernetespkgapiv1StorageOSPersistentVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1StorageOSPersistentVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1ObjectReference `json:"secretRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Volumename string `json:"volumeName,omitempty"` // VolumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.
	Volumenamespace string `json:"volumeNamespace,omitempty"` // VolumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
}

// Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerSpec represents the Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerSpec struct {
	Maxreplicas int `json:"maxReplicas"` // maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
	Metrics []Iok8sapiautoscalingv2beta1MetricSpec `json:"metrics,omitempty"` // metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used). The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods. Ergo, metrics used must decrease as the pod count is increased, and vice-versa. See the individual metric source types for more information about how each type of metric must respond.
	Minreplicas int `json:"minReplicas,omitempty"` // minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod.
	Scaletargetref Iok8sapiautoscalingv2beta1CrossVersionObjectReference `json:"scaleTargetRef"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
}

// Iok8sapiappsv1StatefulSetList represents the Iok8sapiappsv1StatefulSetList schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1StatefulSet `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapistoragev1beta1VolumeError represents the Iok8sapistoragev1beta1VolumeError schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeError struct {
	Message string `json:"message,omitempty"` // String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Time string `json:"time,omitempty"`
}

// Iok8sapicorev1Event represents the Iok8sapicorev1Event schema from the OpenAPI specification
type Iok8sapicorev1Event struct {
	Action string `json:"action,omitempty"` // What action was taken/failed regarding to the Regarding object.
	Involvedobject Iok8sapicorev1ObjectReference `json:"involvedObject"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Reportingcomponent string `json:"reportingComponent,omitempty"` // Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	Related Iok8sapicorev1ObjectReference `json:"related,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Firsttimestamp string `json:"firstTimestamp,omitempty"`
	Reason string `json:"reason,omitempty"` // This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	TypeField string `json:"type,omitempty"` // Type of this event (Normal, Warning), new types could be added in the future
	Eventtime string `json:"eventTime,omitempty"`
	Source Iok8sapicorev1EventSource `json:"source,omitempty"` // EventSource contains information for an event.
	Count int `json:"count,omitempty"` // The number of times this event has occurred.
	Lasttimestamp string `json:"lastTimestamp,omitempty"`
	Message string `json:"message,omitempty"` // A human-readable description of the status of this operation.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Series Iok8sapicorev1EventSeries `json:"series,omitempty"` // EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
	Reportinginstance string `json:"reportingInstance,omitempty"` // ID of the controller instance, e.g. `kubelet-xyzf`.
}

// Iok8skubernetespkgapiv1PodTemplateList represents the Iok8skubernetespkgapiv1PodTemplateList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodTemplateList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1PodTemplate `json:"items"` // List of pod templates
}

// Iok8sapiextensionsv1beta1PodSecurityPolicy represents the Iok8sapiextensionsv1beta1PodSecurityPolicy schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1PodSecurityPolicy struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1PodSecurityPolicySpec `json:"spec,omitempty"` // Pod Security Policy Spec defines the policy enforced.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapisauthenticationv1UserInfo represents the Iok8skubernetespkgapisauthenticationv1UserInfo schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1UserInfo struct {
	Uid string `json:"uid,omitempty"` // A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	Username string `json:"username,omitempty"` // The name that uniquely identifies this user among all active users.
	Extra map[string]interface{} `json:"extra,omitempty"` // Any additional information provided by the authenticator.
	Groups []string `json:"groups,omitempty"` // The names of groups this user is a part of.
}

// Iok8sapiappsv1beta2DaemonSetList represents the Iok8sapiappsv1beta2DaemonSetList schema from the OpenAPI specification
type Iok8sapiappsv1beta2DaemonSetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta2DaemonSet `json:"items"` // A list of daemon sets.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiautoscalingv2beta1ObjectMetricStatus represents the Iok8sapiautoscalingv2beta1ObjectMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1ObjectMetricStatus struct {
	Target Iok8sapiautoscalingv2beta1CrossVersionObjectReference `json:"target"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
	Currentvalue string `json:"currentValue"`
	Metricname string `json:"metricName"` // metricName is the name of the metric in question.
}

// Iok8skubernetespkgapisauthenticationv1TokenReviewStatus represents the Iok8skubernetespkgapisauthenticationv1TokenReviewStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1TokenReviewStatus struct {
	Authenticated bool `json:"authenticated,omitempty"` // Authenticated indicates that the token was associated with a known user.
	ErrorField string `json:"error,omitempty"` // Error indicates that the token couldn't be checked
	User Iok8sapiauthenticationv1UserInfo `json:"user,omitempty"` // UserInfo holds the information about the user needed to implement the user.Info interface.
}

// Iok8sapimachinerypkgruntimeRawExtension represents the Iok8sapimachinerypkgruntimeRawExtension schema from the OpenAPI specification
type Iok8sapimachinerypkgruntimeRawExtension struct {
	Raw string `json:"Raw"` // Raw is the underlying serialization of this object.
}

// Iok8sapicorev1NFSVolumeSource represents the Iok8sapicorev1NFSVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1NFSVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server string `json:"server"` // Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Path string `json:"path"` // Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
}

// Iok8skubernetespkgapiv1NamespaceStatus represents the Iok8skubernetespkgapiv1NamespaceStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NamespaceStatus struct {
	Phase string `json:"phase,omitempty"` // Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
}

// Iok8sapinetworkingv1IPBlock represents the Iok8sapinetworkingv1IPBlock schema from the OpenAPI specification
type Iok8sapinetworkingv1IPBlock struct {
	Cidr string `json:"cidr"` // CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24"
	Except []string `json:"except,omitempty"` // Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" Except values will be rejected if they are outside the CIDR range
}

// Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerStatus represents the Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscalerStatus struct {
	Currentcpuutilizationpercentage int `json:"currentCPUUtilizationPercentage,omitempty"` // current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
	Currentreplicas int `json:"currentReplicas"` // current number of replicas of pods managed by this autoscaler.
	Desiredreplicas int `json:"desiredReplicas"` // desired number of replicas of pods managed by this autoscaler.
	Lastscaletime string `json:"lastScaleTime,omitempty"`
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // most recent generation observed by this autoscaler.
}

// Iok8sapinetworkingv1NetworkPolicy represents the Iok8sapinetworkingv1NetworkPolicy schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicy struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1NetworkPolicySpec `json:"spec,omitempty"` // NetworkPolicySpec provides the specification of a NetworkPolicy
}

// Iok8sapicorev1Service represents the Iok8sapicorev1Service schema from the OpenAPI specification
type Iok8sapicorev1Service struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ServiceSpec `json:"spec,omitempty"` // ServiceSpec describes the attributes that a user creates on a service.
	Status Iok8sapicorev1ServiceStatus `json:"status,omitempty"` // ServiceStatus represents the current status of a service.
}

// Iok8sapicorev1EnvFromSource represents the Iok8sapicorev1EnvFromSource schema from the OpenAPI specification
type Iok8sapicorev1EnvFromSource struct {
	Configmapref Iok8sapicorev1ConfigMapEnvSource `json:"configMapRef,omitempty"` // ConfigMapEnvSource selects a ConfigMap to populate the environment variables with. The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
	Prefix string `json:"prefix,omitempty"` // An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Secretref Iok8sapicorev1SecretEnvSource `json:"secretRef,omitempty"` // SecretEnvSource selects a Secret to populate the environment variables with. The contents of the target Secret's Data field will represent the key-value pairs as environment variables.
}

// Iok8sapiappsv1beta2Scale represents the Iok8sapiappsv1beta2Scale schema from the OpenAPI specification
type Iok8sapiappsv1beta2Scale struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta2ScaleSpec `json:"spec,omitempty"` // ScaleSpec describes the attributes of a scale subresource
	Status Iok8sapiappsv1beta2ScaleStatus `json:"status,omitempty"` // ScaleStatus represents the current status of a scale subresource.
}

// Iok8sapirbacv1beta1ClusterRoleList represents the Iok8sapirbacv1beta1ClusterRoleList schema from the OpenAPI specification
type Iok8sapirbacv1beta1ClusterRoleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1ClusterRole `json:"items"` // Items is a list of ClusterRoles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1PodStatus represents the Iok8sapicorev1PodStatus schema from the OpenAPI specification
type Iok8sapicorev1PodStatus struct {
	Phase string `json:"phase,omitempty"` // Current condition of the pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
	Podip string `json:"podIP,omitempty"` // IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
	Qosclass string `json:"qosClass,omitempty"` // The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md
	Reason string `json:"reason,omitempty"` // A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'
	Containerstatuses []Iok8sapicorev1ContainerStatus `json:"containerStatuses,omitempty"` // The list has one entry per container in the manifest. Each entry is currently the output of `docker inspect`. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	Initcontainerstatuses []Iok8sapicorev1ContainerStatus `json:"initContainerStatuses,omitempty"` // The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	Nominatednodename string `json:"nominatedNodeName,omitempty"` // nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
	Starttime string `json:"startTime,omitempty"`
	Hostip string `json:"hostIP,omitempty"` // IP address of the host to which the pod is assigned. Empty if not yet scheduled.
	Conditions []Iok8sapicorev1PodCondition `json:"conditions,omitempty"` // Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Message string `json:"message,omitempty"` // A human readable message indicating details about why the pod is in this condition.
}

// Iok8sapiauthorizationv1beta1ResourceRule represents the Iok8sapiauthorizationv1beta1ResourceRule schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1ResourceRule struct {
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed. "*" means all.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. "*" means all in the specified apiGroups. "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
	Verbs []string `json:"verbs"` // Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed. "*" means all.
}

// Iok8skubernetespkgapisbatchv2alpha1CronJobList represents the Iok8skubernetespkgapisbatchv2alpha1CronJobList schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv2alpha1CronJobList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapibatchv2alpha1CronJob `json:"items"` // items is the list of CronJobs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisextensionsv1beta1FSGroupStrategyOptions represents the Iok8skubernetespkgapisextensionsv1beta1FSGroupStrategyOptions schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1FSGroupStrategyOptions struct {
	Ranges []Iok8sapiextensionsv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of fs groups. If you would like to force a single fs group then supply a single range with the same start and end.
	Rule string `json:"rule,omitempty"` // Rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
}

// Iok8skubernetespkgapiv1ResourceQuotaSpec represents the Iok8skubernetespkgapiv1ResourceQuotaSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ResourceQuotaSpec struct {
	Hard map[string]interface{} `json:"hard,omitempty"` // Hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Scopes []string `json:"scopes,omitempty"` // A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
}

// Iok8sapipolicyv1beta1PodDisruptionBudgetStatus represents the Iok8sapipolicyv1beta1PodDisruptionBudgetStatus schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodDisruptionBudgetStatus struct {
	Currenthealthy int `json:"currentHealthy"` // current number of healthy pods
	Desiredhealthy int `json:"desiredHealthy"` // minimum desired number of healthy pods
	Disruptedpods map[string]interface{} `json:"disruptedPods"` // DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
	Disruptionsallowed int `json:"disruptionsAllowed"` // Number of pod disruptions that are currently allowed.
	Expectedpods int `json:"expectedPods"` // total number of pods counted by this disruption budget
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // Most recent generation observed when updating this PDB status. PodDisruptionsAllowed and other status informatio is valid only if observedGeneration equals to PDB's object generation.
}

// Iok8skubernetespkgapiv1PortworxVolumeSource represents the Iok8skubernetespkgapiv1PortworxVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PortworxVolumeSource struct {
	Volumeid string `json:"volumeID"` // VolumeID uniquely identifies a Portworx volume
	Fstype string `json:"fsType,omitempty"` // FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Scale GeneratedType `json:"scale,omitempty"` // CustomResourceSubresourceScale defines how to serve the scale subresource for CustomResources.
	Status GeneratedType `json:"status,omitempty"` // CustomResourceSubresourceStatus defines how to serve the status subresource for CustomResources. Status is represented by the `.status` JSON path inside of a CustomResource. When set, * exposes a /status subresource for the custom resource * PUT requests to the /status subresource take a custom resource object, and ignore changes to anything except the status stanza * PUT/POST/PATCH requests to the custom resource ignore changes to the status stanza
}

// Iok8sapirbacv1beta1ClusterRoleBinding represents the Iok8sapirbacv1beta1ClusterRoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1beta1ClusterRoleBinding struct {
	Subjects []Iok8sapirbacv1beta1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Roleref Iok8sapirbacv1beta1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
}

// Iok8skubernetespkgapiv1NodeCondition represents the Iok8skubernetespkgapiv1NodeCondition schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // (brief) reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of node condition.
	Lastheartbeattime string `json:"lastHeartbeatTime,omitempty"`
}

// Iok8sapiextensionsv1beta1NetworkPolicyPeer represents the Iok8sapiextensionsv1beta1NetworkPolicyPeer schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicyPeer struct {
	Ipblock Iok8sapiextensionsv1beta1IPBlock `json:"ipBlock,omitempty"` // DEPRECATED 1.9 - This group version of IPBlock is deprecated by networking/v1/IPBlock. IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8skubernetespkgapisappsv1beta1RollbackConfig represents the Iok8skubernetespkgapisappsv1beta1RollbackConfig schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1RollbackConfig struct {
	Revision int64 `json:"revision,omitempty"` // The revision to rollback to. If set to 0, rollback to the last revision.
}

// Iok8sapiappsv1ReplicaSetSpec represents the Iok8sapiappsv1ReplicaSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetSpec struct {
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
}

// Iok8sapirbacv1ClusterRoleBinding represents the Iok8sapirbacv1ClusterRoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRoleBinding struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Roleref Iok8sapirbacv1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1FlexVolumeSource represents the Iok8skubernetespkgapiv1FlexVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1FlexVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Driver string `json:"driver"` // Driver is the name of the driver to use for this volume.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	Options map[string]interface{} `json:"options,omitempty"` // Optional: Extra command options if any.
}

// Iok8sapicorev1EndpointSubset represents the Iok8sapicorev1EndpointSubset schema from the OpenAPI specification
type Iok8sapicorev1EndpointSubset struct {
	Notreadyaddresses []Iok8sapicorev1EndpointAddress `json:"notReadyAddresses,omitempty"` // IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
	Ports []Iok8sapicorev1EndpointPort `json:"ports,omitempty"` // Port numbers available on the related IP addresses.
	Addresses []Iok8sapicorev1EndpointAddress `json:"addresses,omitempty"` // IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
}

// Iok8skubernetespkgapiv1Lifecycle represents the Iok8skubernetespkgapiv1Lifecycle schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Lifecycle struct {
	Poststart Iok8sapicorev1Handler `json:"postStart,omitempty"` // Handler defines a specific action that should be taken
	Prestop Iok8sapicorev1Handler `json:"preStop,omitempty"` // Handler defines a specific action that should be taken
}

// Iok8sapipolicyv1beta1AllowedFlexVolume represents the Iok8sapipolicyv1beta1AllowedFlexVolume schema from the OpenAPI specification
type Iok8sapipolicyv1beta1AllowedFlexVolume struct {
	Driver string `json:"driver"` // Driver is the name of the Flexvolume driver.
}

// Iok8sapimachinerypkgapismetav1WatchEvent represents the Iok8sapimachinerypkgapismetav1WatchEvent schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1WatchEvent struct {
	Object Iok8sapimachinerypkgruntimeRawExtension `json:"object"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.Object `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // External package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.RawExtension `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // On the wire, the JSON will look something like this: { 	"kind":"MyAPIObject", 	"apiVersion":"v1", 	"myPlugin": { 		"kind":"PluginA", 		"aOption":"foo", 	}, } So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	TypeField string `json:"type"`
}

// Iok8skubernetespkgapissettingsv1alpha1PodPresetList represents the Iok8skubernetespkgapissettingsv1alpha1PodPresetList schema from the OpenAPI specification
type Iok8skubernetespkgapissettingsv1alpha1PodPresetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapisettingsv1alpha1PodPreset `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1NodeAddress represents the Iok8sapicorev1NodeAddress schema from the OpenAPI specification
type Iok8sapicorev1NodeAddress struct {
	TypeField string `json:"type"` // Node address type, one of Hostname, ExternalIP or InternalIP.
	Address string `json:"address"` // The node address.
}

// Iok8sapicorev1ProjectedVolumeSource represents the Iok8sapicorev1ProjectedVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ProjectedVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Sources []Iok8sapicorev1VolumeProjection `json:"sources"` // list of volume projections
}

// Iok8sapinetworkingv1NetworkPolicyList represents the Iok8sapinetworkingv1NetworkPolicyList schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapinetworkingv1NetworkPolicy `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapiv1NodeSelectorRequirement represents the Iok8skubernetespkgapiv1NodeSelectorRequirement schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeSelectorRequirement struct {
	Key string `json:"key"` // The label key that the selector applies to.
	Operator string `json:"operator"` // Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Values []string `json:"values,omitempty"` // An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
}

// Iok8sapinetworkingv1NetworkPolicyPeer represents the Iok8sapinetworkingv1NetworkPolicyPeer schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyPeer struct {
	Ipblock Iok8sapinetworkingv1IPBlock `json:"ipBlock,omitempty"` // IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiautoscalingv1CrossVersionObjectReference represents the Iok8sapiautoscalingv1CrossVersionObjectReference schema from the OpenAPI specification
type Iok8sapiautoscalingv1CrossVersionObjectReference struct {
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent
	Kind string `json:"kind"` // Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
	Name string `json:"name"` // Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
}

// Iok8sapicorev1FlockerVolumeSource represents the Iok8sapicorev1FlockerVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FlockerVolumeSource struct {
	Datasetname string `json:"datasetName,omitempty"` // Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	Datasetuuid string `json:"datasetUUID,omitempty"` // UUID of the dataset. This is unique identifier of a Flocker dataset
}

// Iok8skubernetespkgapisextensionsv1beta1SELinuxStrategyOptions represents the Iok8skubernetespkgapisextensionsv1beta1SELinuxStrategyOptions schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1SELinuxStrategyOptions struct {
	Rule string `json:"rule"` // type is the strategy that will dictate the allowable labels that may be set.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Conditions []GeneratedType `json:"conditions,omitempty"` // Current service state of apiService.
}

// Iok8sapicorev1PersistentVolumeClaimList represents the Iok8sapicorev1PersistentVolumeClaimList schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1PersistentVolumeClaim `json:"items"` // A list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiauthorizationv1beta1SelfSubjectAccessReview represents the Iok8sapiauthorizationv1beta1SelfSubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SelfSubjectAccessReview struct {
	Status Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1beta1SelfSubjectAccessReviewSpec `json:"spec"` // SelfSubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
}

// Iok8skubernetespkgapiv1PersistentVolume represents the Iok8skubernetespkgapiv1PersistentVolume schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolume struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PersistentVolumeSpec `json:"spec,omitempty"` // PersistentVolumeSpec is the specification of a persistent volume.
	Status Iok8sapicorev1PersistentVolumeStatus `json:"status,omitempty"` // PersistentVolumeStatus is the current status of a persistent volume.
}

// Iok8sapicorev1PodAffinityTerm represents the Iok8sapicorev1PodAffinityTerm schema from the OpenAPI specification
type Iok8sapicorev1PodAffinityTerm struct {
	Topologykey string `json:"topologyKey"` // This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	Labelselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"labelSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Namespaces []string `json:"namespaces,omitempty"` // namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this pod's namespace"
}

// Iok8sapirbacv1alpha1Role represents the Iok8sapirbacv1alpha1Role schema from the OpenAPI specification
type Iok8sapirbacv1alpha1Role struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1alpha1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this Role
}

// Iok8skubernetespkgapiv1KeyToPath represents the Iok8skubernetespkgapiv1KeyToPath schema from the OpenAPI specification
type Iok8skubernetespkgapiv1KeyToPath struct {
	Path string `json:"path"` // The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	Key string `json:"key"` // The key to project.
	Mode int `json:"mode,omitempty"` // Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8sapiautoscalingv2beta1HorizontalPodAutoscaler represents the Iok8sapiautoscalingv2beta1HorizontalPodAutoscaler schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1HorizontalPodAutoscaler struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerSpec `json:"spec,omitempty"` // HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.
	Status Iok8sapiautoscalingv2beta1HorizontalPodAutoscalerStatus `json:"status,omitempty"` // HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiadmissionregistrationv1alpha1InitializerConfigurationList represents the Iok8sapiadmissionregistrationv1alpha1InitializerConfigurationList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1InitializerConfigurationList struct {
	Items []Iok8sapiadmissionregistrationv1alpha1InitializerConfiguration `json:"items"` // List of InitializerConfiguration.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1HostAlias represents the Iok8skubernetespkgapiv1HostAlias schema from the OpenAPI specification
type Iok8skubernetespkgapiv1HostAlias struct {
	Ip string `json:"ip,omitempty"` // IP address of the host file entry.
	Hostnames []string `json:"hostnames,omitempty"` // Hostnames for the above IP address.
}

// Iok8sapicorev1GCEPersistentDiskVolumeSource represents the Iok8sapicorev1GCEPersistentDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GCEPersistentDiskVolumeSource struct {
	Pdname string `json:"pdName"` // Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Partition int `json:"partition,omitempty"` // The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
}

// Iok8skubernetespkgapisauthenticationv1TokenReview represents the Iok8skubernetespkgapisauthenticationv1TokenReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1TokenReview struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthenticationv1TokenReviewSpec `json:"spec"` // TokenReviewSpec is a description of the token authentication request.
	Status Iok8sapiauthenticationv1TokenReviewStatus `json:"status,omitempty"` // TokenReviewStatus is the result of the token authentication request.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1EnvVar represents the Iok8skubernetespkgapiv1EnvVar schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EnvVar struct {
	Name string `json:"name"` // Name of the environment variable. Must be a C_IDENTIFIER.
	Value string `json:"value,omitempty"` // Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Valuefrom Iok8sapicorev1EnvVarSource `json:"valueFrom,omitempty"` // EnvVarSource represents a source for the value of an EnvVar.
}

// Iok8sapiextensionsv1beta1ReplicaSetList represents the Iok8sapiextensionsv1beta1ReplicaSetList schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ReplicaSetList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1ReplicaSet `json:"items"` // List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
}

// Iok8skubernetespkgapiv1ServiceAccountList represents the Iok8skubernetespkgapiv1ServiceAccountList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ServiceAccountList struct {
	Items []Iok8sapicorev1ServiceAccount `json:"items"` // List of ServiceAccounts. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1SELinuxStrategyOptions represents the Iok8sapiextensionsv1beta1SELinuxStrategyOptions schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1SELinuxStrategyOptions struct {
	Rule string `json:"rule"` // type is the strategy that will dictate the allowable labels that may be set.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
}

// Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetSpec represents the Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetSpec schema from the OpenAPI specification
type Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetSpec struct {
	Minavailable string `json:"minAvailable,omitempty"`
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8sapicorev1ContainerState represents the Iok8sapicorev1ContainerState schema from the OpenAPI specification
type Iok8sapicorev1ContainerState struct {
	Running Iok8sapicorev1ContainerStateRunning `json:"running,omitempty"` // ContainerStateRunning is a running state of a container.
	Terminated Iok8sapicorev1ContainerStateTerminated `json:"terminated,omitempty"` // ContainerStateTerminated is a terminated state of a container.
	Waiting Iok8sapicorev1ContainerStateWaiting `json:"waiting,omitempty"` // ContainerStateWaiting is a waiting state of a container.
}

// Iok8skubernetespkgapiv1ResourceQuotaStatus represents the Iok8skubernetespkgapiv1ResourceQuotaStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ResourceQuotaStatus struct {
	Used map[string]interface{} `json:"used,omitempty"` // Used is the current observed total usage of the resource in the namespace.
	Hard map[string]interface{} `json:"hard,omitempty"` // Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
}

// Iok8skubernetespkgapisappsv1beta1Deployment represents the Iok8skubernetespkgapisappsv1beta1Deployment schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1Deployment struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta1DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
	Status Iok8sapiappsv1beta1DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1EventSource represents the Iok8sapicorev1EventSource schema from the OpenAPI specification
type Iok8sapicorev1EventSource struct {
	Component string `json:"component,omitempty"` // Component from which the event is generated.
	Host string `json:"host,omitempty"` // Node name on which the event is generated.
}

// Iok8skubernetespkgapisadmissionregistrationv1alpha1Initializer represents the Iok8skubernetespkgapisadmissionregistrationv1alpha1Initializer schema from the OpenAPI specification
type Iok8skubernetespkgapisadmissionregistrationv1alpha1Initializer struct {
	Name string `json:"name"` // Name is the identifier of the initializer. It will be added to the object that needs to be initialized. Name should be fully qualified, e.g., alwayspullimages.kubernetes.io, where "alwayspullimages" is the name of the webhook, and kubernetes.io is the name of the organization. Required
	Rules []Iok8sapiadmissionregistrationv1alpha1Rule `json:"rules,omitempty"` // Rules describes what resources/subresources the initializer cares about. The initializer cares about an operation if it matches _any_ Rule. Rule.Resources must not include subresources.
}

// Iok8sapicertificatesv1beta1CertificateSigningRequest represents the Iok8sapicertificatesv1beta1CertificateSigningRequest schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1CertificateSigningRequest struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicertificatesv1beta1CertificateSigningRequestSpec `json:"spec,omitempty"` // This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
	Status Iok8sapicertificatesv1beta1CertificateSigningRequestStatus `json:"status,omitempty"`
}

// Iok8sapiappsv1beta2DeploymentStrategy represents the Iok8sapiappsv1beta2DeploymentStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta2DeploymentStrategy struct {
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	Rollingupdate Iok8sapiappsv1beta2RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
}

// Iok8sapicorev1EmptyDirVolumeSource represents the Iok8sapicorev1EmptyDirVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1EmptyDirVolumeSource struct {
	Sizelimit string `json:"sizeLimit,omitempty"`
	Medium string `json:"medium,omitempty"` // What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
}

// Iok8sapistoragev1beta1VolumeAttachment represents the Iok8sapistoragev1beta1VolumeAttachment schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttachment struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapistoragev1beta1VolumeAttachmentSpec `json:"spec"` // VolumeAttachmentSpec is the specification of a VolumeAttachment request.
	Status Iok8sapistoragev1beta1VolumeAttachmentStatus `json:"status,omitempty"` // VolumeAttachmentStatus is the status of a VolumeAttachment request.
}

// Iok8skubernetespkgapiv1ConfigMapList represents the Iok8skubernetespkgapiv1ConfigMapList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ConfigMapList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ConfigMap `json:"items"` // Items is the list of ConfigMaps.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiautoscalingv2beta1PodsMetricSource represents the Iok8sapiautoscalingv2beta1PodsMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1PodsMetricSource struct {
	Metricname string `json:"metricName"` // metricName is the name of the metric in question
	Targetaveragevalue string `json:"targetAverageValue"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec GeneratedType `json:"spec,omitempty"` // APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
	Status GeneratedType `json:"status,omitempty"` // APIServiceStatus contains derived information about an API server
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1Secret represents the Iok8sapicorev1Secret schema from the OpenAPI specification
type Iok8sapicorev1Secret struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Stringdata map[string]interface{} `json:"stringData,omitempty"` // stringData allows specifying non-binary secret data in string form. It is provided as a write-only convenience method. All keys and values are merged into the data field on write, overwriting any existing values. It is never output when reading from the API.
	TypeField string `json:"type,omitempty"` // Used to facilitate programmatic handling of secret data.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Data map[string]interface{} `json:"data,omitempty"` // Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
}

// Iok8skubernetespkgapisrbacv1alpha1PolicyRule represents the Iok8skubernetespkgapisrbacv1alpha1PolicyRule schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1PolicyRule struct {
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path This name is intentionally different than the internal type so that the DefaultConvert works nicely and because the ordering may be different. Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. ResourceAll represents all resources.
	Verbs []string `json:"verbs"` // Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.
}

// Iok8skubernetespkgapiv1PhotonPersistentDiskVolumeSource represents the Iok8skubernetespkgapiv1PhotonPersistentDiskVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PhotonPersistentDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Pdid string `json:"pdID"` // ID that identifies Photon Controller persistent disk
}

// Iok8sapiautoscalingv2beta1CrossVersionObjectReference represents the Iok8sapiautoscalingv2beta1CrossVersionObjectReference schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1CrossVersionObjectReference struct {
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent
	Kind string `json:"kind"` // Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
	Name string `json:"name"` // Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
}

// Iok8skubernetespkgapiv1HostPathVolumeSource represents the Iok8skubernetespkgapiv1HostPathVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1HostPathVolumeSource struct {
	Path string `json:"path"` // Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	TypeField string `json:"type,omitempty"` // Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
}

// Iok8sapiadmissionregistrationv1beta1Webhook represents the Iok8sapiadmissionregistrationv1beta1Webhook schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1Webhook struct {
	Clientconfig Iok8sapiadmissionregistrationv1beta1WebhookClientConfig `json:"clientConfig"` // WebhookClientConfig contains the information to make a TLS connection with the webhook
	Failurepolicy string `json:"failurePolicy,omitempty"` // FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Ignore.
	Name string `json:"name"` // The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is the name of the organization. Required.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Rules []Iok8sapiadmissionregistrationv1beta1RuleWithOperations `json:"rules,omitempty"` // Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
}

// Iok8sapicorev1Namespace represents the Iok8sapicorev1Namespace schema from the OpenAPI specification
type Iok8sapicorev1Namespace struct {
	Status Iok8sapicorev1NamespaceStatus `json:"status,omitempty"` // NamespaceStatus is information about the current status of a Namespace.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1NamespaceSpec `json:"spec,omitempty"` // NamespaceSpec describes the attributes on a Namespace.
}

// Iok8skubernetespkgapisrbacv1beta1RoleBindingList represents the Iok8skubernetespkgapisrbacv1beta1RoleBindingList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1RoleBindingList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1RoleBinding `json:"items"` // Items is a list of RoleBindings
}

// Iok8sapiextensionsv1beta1DaemonSetUpdateStrategy represents the Iok8sapiextensionsv1beta1DaemonSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DaemonSetUpdateStrategy struct {
	Rollingupdate Iok8sapiextensionsv1beta1RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of daemon set rolling update.
	TypeField string `json:"type,omitempty"` // Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is OnDelete.
}

// Iok8skubernetespkgapisrbacv1alpha1RoleBinding represents the Iok8skubernetespkgapisrbacv1alpha1RoleBinding schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1RoleBinding struct {
	Roleref Iok8sapirbacv1alpha1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1alpha1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiappsv1DeploymentList represents the Iok8sapiappsv1DeploymentList schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1Deployment `json:"items"` // Items is the list of Deployments.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1StorageOSPersistentVolumeSource represents the Iok8sapicorev1StorageOSPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1StorageOSPersistentVolumeSource struct {
	Secretref Iok8sapicorev1ObjectReference `json:"secretRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Volumename string `json:"volumeName,omitempty"` // VolumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.
	Volumenamespace string `json:"volumeNamespace,omitempty"` // VolumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapiappsv1beta2StatefulSetUpdateStrategy represents the Iok8sapiappsv1beta2StatefulSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta2StatefulSetUpdateStrategy struct {
	Rollingupdate Iok8sapiappsv1beta2RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"` // RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
	TypeField string `json:"type,omitempty"` // Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
}

// Iok8skubernetespkgapisnetworkingv1NetworkPolicyList represents the Iok8skubernetespkgapisnetworkingv1NetworkPolicyList schema from the OpenAPI specification
type Iok8skubernetespkgapisnetworkingv1NetworkPolicyList struct {
	Items []Iok8sapinetworkingv1NetworkPolicy `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1ComponentStatusList represents the Iok8skubernetespkgapiv1ComponentStatusList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ComponentStatusList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ComponentStatus `json:"items"` // List of ComponentStatus objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiextensionsv1beta1IngressTLS represents the Iok8sapiextensionsv1beta1IngressTLS schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IngressTLS struct {
	Hosts []string `json:"hosts,omitempty"` // Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Secretname string `json:"secretName,omitempty"` // SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
}

// Iok8sapicorev1ServiceList represents the Iok8sapicorev1ServiceList schema from the OpenAPI specification
type Iok8sapicorev1ServiceList struct {
	Items []Iok8sapicorev1Service `json:"items"` // List of services
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1ServiceStatus represents the Iok8sapicorev1ServiceStatus schema from the OpenAPI specification
type Iok8sapicorev1ServiceStatus struct {
	Loadbalancer Iok8sapicorev1LoadBalancerStatus `json:"loadBalancer,omitempty"` // LoadBalancerStatus represents the status of a load-balancer.
}

// Iok8sapiextensionsv1beta1DeploymentStrategy represents the Iok8sapiextensionsv1beta1DeploymentStrategy schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DeploymentStrategy struct {
	Rollingupdate Iok8sapiextensionsv1beta1RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
}

// Iok8skubernetespkgapisauthorizationv1ResourceAttributes represents the Iok8skubernetespkgapisauthorizationv1ResourceAttributes schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1ResourceAttributes struct {
	Group string `json:"group,omitempty"` // Group is the API Group of the Resource. "*" means all.
	Name string `json:"name,omitempty"` // Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the action being requested. Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Resource string `json:"resource,omitempty"` // Resource is one of the existing resource types. "*" means all.
	Subresource string `json:"subresource,omitempty"` // Subresource is one of the existing resource types. "" means none.
	Verb string `json:"verb,omitempty"` // Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Version string `json:"version,omitempty"` // Version is the API Version of the Resource. "*" means all.
}

// Iok8skubernetespkgapiv1PodTemplate represents the Iok8skubernetespkgapiv1PodTemplate schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodTemplate struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// Iok8skubernetespkgapisrbacv1beta1RoleBinding represents the Iok8skubernetespkgapisrbacv1beta1RoleBinding schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1RoleBinding struct {
	Roleref Iok8sapirbacv1beta1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1beta1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfigurationList represents the Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfigurationList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfigurationList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfiguration `json:"items"` // List of MutatingWebhookConfiguration.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiextensionsv1beta1RunAsUserStrategyOptions represents the Iok8sapiextensionsv1beta1RunAsUserStrategyOptions schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1RunAsUserStrategyOptions struct {
	Ranges []Iok8sapiextensionsv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of uids that may be used.
	Rule string `json:"rule"` // Rule is the strategy that will dictate the allowable RunAsUser values that may be set.
}

// Iok8sapicorev1EndpointPort represents the Iok8sapicorev1EndpointPort schema from the OpenAPI specification
type Iok8sapicorev1EndpointPort struct {
	Protocol string `json:"protocol,omitempty"` // The IP protocol for this port. Must be UDP or TCP. Default is TCP.
	Name string `json:"name,omitempty"` // The name of this port (corresponds to ServicePort.Name). Must be a DNS_LABEL. Optional only if one port is defined.
	Port int `json:"port"` // The port number of the endpoint.
}

// Iok8sapicorev1QuobyteVolumeSource represents the Iok8sapicorev1QuobyteVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1QuobyteVolumeSource struct {
	Group string `json:"group,omitempty"` // Group to map volume access to Default is no group
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	Registry string `json:"registry"` // Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	User string `json:"user,omitempty"` // User to map volume access to Defaults to serivceaccount user
	Volume string `json:"volume"` // Volume is a string that references an already created Quobyte volume by name.
}

// Iok8sapinetworkingv1NetworkPolicySpec represents the Iok8sapinetworkingv1NetworkPolicySpec schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicySpec struct {
	Policytypes []string `json:"policyTypes,omitempty"` // List of rule types that the NetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	Egress []Iok8sapinetworkingv1NetworkPolicyEgressRule `json:"egress,omitempty"` // List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
	Ingress []Iok8sapinetworkingv1NetworkPolicyIngressRule `json:"ingress,omitempty"` // List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8skubernetespkgapiv1PersistentVolumeClaimList represents the Iok8skubernetespkgapiv1PersistentVolumeClaimList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeClaimList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1PersistentVolumeClaim `json:"items"` // A list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1beta2ReplicaSetCondition represents the Iok8sapiappsv1beta2ReplicaSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1beta2ReplicaSetCondition struct {
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of replica set condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
}

// Iok8sapiappsv1beta2ReplicaSetList represents the Iok8sapiappsv1beta2ReplicaSetList schema from the OpenAPI specification
type Iok8sapiappsv1beta2ReplicaSetList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta2ReplicaSet `json:"items"` // List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
}

// Iok8sapiappsv1beta2DeploymentList represents the Iok8sapiappsv1beta2DeploymentList schema from the OpenAPI specification
type Iok8sapiappsv1beta2DeploymentList struct {
	Items []Iok8sapiappsv1beta2Deployment `json:"items"` // Items is the list of Deployments.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicertificatesv1beta1CertificateSigningRequestStatus represents the Iok8sapicertificatesv1beta1CertificateSigningRequestStatus schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1CertificateSigningRequestStatus struct {
	Conditions []Iok8sapicertificatesv1beta1CertificateSigningRequestCondition `json:"conditions,omitempty"` // Conditions applied to the request, such as approval or denial.
	Certificate string `json:"certificate,omitempty"` // If request was approved, the controller will place the issued certificate here.
}

// Iok8skubernetespkgapiv1QuobyteVolumeSource represents the Iok8skubernetespkgapiv1QuobyteVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1QuobyteVolumeSource struct {
	Group string `json:"group,omitempty"` // Group to map volume access to Default is no group
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	Registry string `json:"registry"` // Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	User string `json:"user,omitempty"` // User to map volume access to Defaults to serivceaccount user
	Volume string `json:"volume"` // Volume is a string that references an already created Quobyte volume by name.
}

// Iok8sapiappsv1beta1StatefulSetStatus represents the Iok8sapiappsv1beta1StatefulSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta1StatefulSetStatus struct {
	Replicas int `json:"replicas"` // replicas is the number of Pods created by the StatefulSet controller.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	Collisioncount int `json:"collisionCount,omitempty"` // collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Conditions []Iok8sapiappsv1beta1StatefulSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a statefulset's current state.
	Currentreplicas int `json:"currentReplicas,omitempty"` // currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
	Currentrevision string `json:"currentRevision,omitempty"` // currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
	Readyreplicas int `json:"readyReplicas,omitempty"` // readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	Updaterevision string `json:"updateRevision,omitempty"` // updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
}

// Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery represents the Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery struct {
	Version string `json:"version"` // version specifies the version in the form of "version". This is to save the clients the trouble of splitting the GroupVersion.
	Groupversion string `json:"groupVersion"` // groupVersion specifies the API group and version in the form "group/version"
}

// Iok8sapicorev1NamespaceStatus represents the Iok8sapicorev1NamespaceStatus schema from the OpenAPI specification
type Iok8sapicorev1NamespaceStatus struct {
	Phase string `json:"phase,omitempty"` // Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
}

// Iok8skubernetespkgapisextensionsv1beta1DeploymentStrategy represents the Iok8skubernetespkgapisextensionsv1beta1DeploymentStrategy schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DeploymentStrategy struct {
	Rollingupdate Iok8sapiextensionsv1beta1RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
}

// Iok8skubernetespkgapisappsv1beta1RollingUpdateDeployment represents the Iok8skubernetespkgapisappsv1beta1RollingUpdateDeployment schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1RollingUpdateDeployment struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
	Maxsurge string `json:"maxSurge,omitempty"`
}

// Iok8sapirbacv1alpha1PolicyRule represents the Iok8sapirbacv1alpha1PolicyRule schema from the OpenAPI specification
type Iok8sapirbacv1alpha1PolicyRule struct {
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path This name is intentionally different than the internal type so that the DefaultConvert works nicely and because the ordering may be different. Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. ResourceAll represents all resources.
	Verbs []string `json:"verbs"` // Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
}

// Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyPort represents the Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyPort schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1NetworkPolicyPort struct {
	Port string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"` // Optional. The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.
}

// Iok8sapiextensionsv1beta1IngressList represents the Iok8sapiextensionsv1beta1IngressList schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IngressList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1Ingress `json:"items"` // Items is the list of Ingress.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiauthorizationv1beta1NonResourceRule represents the Iok8sapiauthorizationv1beta1NonResourceRule schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1NonResourceRule struct {
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path. "*" means all.
	Verbs []string `json:"verbs"` // Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options. "*" means all.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Jsonschemas []GeneratedType `json:"JSONSchemas"`
	Schema GeneratedType `json:"Schema"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
}

// Iok8sapicorev1ScaleIOPersistentVolumeSource represents the Iok8sapicorev1ScaleIOPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ScaleIOPersistentVolumeSource struct {
	System string `json:"system"` // The name of the storage system as configured in ScaleIO.
	Volumename string `json:"volumeName,omitempty"` // The name of a volume already created in the ScaleIO system that is associated with this volume source.
	Sslenabled bool `json:"sslEnabled,omitempty"` // Flag to enable/disable SSL communication with Gateway, default false
	Storagepool string `json:"storagePool,omitempty"` // The ScaleIO Storage Pool associated with the protection domain.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1SecretReference `json:"secretRef"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Storagemode string `json:"storageMode,omitempty"` // Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Gateway string `json:"gateway"` // The host address of the ScaleIO API Gateway.
	Protectiondomain string `json:"protectionDomain,omitempty"` // The name of the ScaleIO Protection Domain for the configured storage.
}

// Iok8sapiappsv1beta2RollingUpdateDeployment represents the Iok8sapiappsv1beta2RollingUpdateDeployment schema from the OpenAPI specification
type Iok8sapiappsv1beta2RollingUpdateDeployment struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
	Maxsurge string `json:"maxSurge,omitempty"`
}

// Iok8sapibatchv1Job represents the Iok8sapibatchv1Job schema from the OpenAPI specification
type Iok8sapibatchv1Job struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
	Status Iok8sapibatchv1JobStatus `json:"status,omitempty"` // JobStatus represents the current state of a Job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1ConfigMapVolumeSource represents the Iok8skubernetespkgapiv1ConfigMapVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ConfigMapVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or it's keys must be defined
}

// Iok8sapiappsv1DaemonSetStatus represents the Iok8sapiappsv1DaemonSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetStatus struct {
	Numbermisscheduled int `json:"numberMisscheduled"` // The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Numberavailable int `json:"numberAvailable,omitempty"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numberunavailable int `json:"numberUnavailable,omitempty"` // The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Currentnumberscheduled int `json:"currentNumberScheduled"` // The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Desirednumberscheduled int `json:"desiredNumberScheduled"` // The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Numberready int `json:"numberReady"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The most recent generation observed by the daemon set controller.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Conditions []Iok8sapiappsv1DaemonSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a DaemonSet's current state.
	Updatednumberscheduled int `json:"updatedNumberScheduled,omitempty"` // The total number of nodes that are running updated daemon pod
}

// Iok8sapirbacv1alpha1ClusterRoleBinding represents the Iok8sapirbacv1alpha1ClusterRoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1alpha1ClusterRoleBinding struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Roleref Iok8sapirbacv1alpha1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1alpha1Subject `json:"subjects"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1FCVolumeSource represents the Iok8sapicorev1FCVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FCVolumeSource struct {
	Wwids []string `json:"wwids,omitempty"` // Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Lun int `json:"lun,omitempty"` // Optional: FC target lun number
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Targetwwns []string `json:"targetWWNs,omitempty"` // Optional: FC target worldwide names (WWNs)
}

// Iok8skubernetespkgapisappsv1beta1DeploymentList represents the Iok8skubernetespkgapisappsv1beta1DeploymentList schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1DeploymentList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta1Deployment `json:"items"` // Items is the list of Deployments.
}

// Iok8sapimachinerypkgapismetav1OwnerReference represents the Iok8sapimachinerypkgapismetav1OwnerReference schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1OwnerReference struct {
	Apiversion string `json:"apiVersion"` // API version of the referent.
	Blockownerdeletion bool `json:"blockOwnerDeletion,omitempty"` // If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
	Controller bool `json:"controller,omitempty"` // If true, this reference points to the managing controller.
	Kind string `json:"kind"` // Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name"` // Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Uid string `json:"uid"` // UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
}

// Iok8sapiextensionsv1beta1DeploymentList represents the Iok8sapiextensionsv1beta1DeploymentList schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DeploymentList struct {
	Items []Iok8sapiextensionsv1beta1Deployment `json:"items"` // Items is the list of Deployments.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1ReplicaSetCondition represents the Iok8sapiextensionsv1beta1ReplicaSetCondition schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ReplicaSetCondition struct {
	TypeField string `json:"type"` // Type of replica set condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
}

// Iok8sapicorev1WeightedPodAffinityTerm represents the Iok8sapicorev1WeightedPodAffinityTerm schema from the OpenAPI specification
type Iok8sapicorev1WeightedPodAffinityTerm struct {
	Podaffinityterm Iok8sapicorev1PodAffinityTerm `json:"podAffinityTerm"` // Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
	Weight int `json:"weight"` // weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
}

// Iok8skubernetespkgapisauthenticationv1beta1TokenReviewStatus represents the Iok8skubernetespkgapisauthenticationv1beta1TokenReviewStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1beta1TokenReviewStatus struct {
	User Iok8sapiauthenticationv1beta1UserInfo `json:"user,omitempty"` // UserInfo holds the information about the user needed to implement the user.Info interface.
	Authenticated bool `json:"authenticated,omitempty"` // Authenticated indicates that the token was associated with a known user.
	ErrorField string `json:"error,omitempty"` // Error indicates that the token couldn't be checked
}

// Iok8skubernetespkgapisbatchv1JobList represents the Iok8skubernetespkgapisbatchv1JobList schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv1JobList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapibatchv1Job `json:"items"` // items is the list of Jobs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1beta2StatefulSetCondition represents the Iok8sapiappsv1beta2StatefulSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1beta2StatefulSetCondition struct {
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of statefulset condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
}

// Iok8sapimachinerypkgapismetav1LabelSelectorRequirement represents the Iok8sapimachinerypkgapismetav1LabelSelectorRequirement schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1LabelSelectorRequirement struct {
	Key string `json:"key"` // key is the label key that the selector applies to.
	Operator string `json:"operator"` // operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Values []string `json:"values,omitempty"` // values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
}

// Iok8sapiautoscalingv2beta1MetricStatus represents the Iok8sapiautoscalingv2beta1MetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1MetricStatus struct {
	Resource Iok8sapiautoscalingv2beta1ResourceMetricStatus `json:"resource,omitempty"` // ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	TypeField string `json:"type"` // type is the type of metric source. It will be one of "Object", "Pods" or "Resource", each corresponds to a matching field in the object.
	External Iok8sapiautoscalingv2beta1ExternalMetricStatus `json:"external,omitempty"` // ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
	Object Iok8sapiautoscalingv2beta1ObjectMetricStatus `json:"object,omitempty"` // ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
	Pods Iok8sapiautoscalingv2beta1PodsMetricStatus `json:"pods,omitempty"` // PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
}

// Iok8skubernetespkgapissettingsv1alpha1PodPreset represents the Iok8skubernetespkgapissettingsv1alpha1PodPreset schema from the OpenAPI specification
type Iok8skubernetespkgapissettingsv1alpha1PodPreset struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapisettingsv1alpha1PodPresetSpec `json:"spec,omitempty"` // PodPresetSpec is a description of a pod preset.
}

// Iok8sapiappsv1beta1DeploymentSpec represents the Iok8sapiappsv1beta1DeploymentSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta1DeploymentSpec struct {
	Strategy Iok8sapiappsv1beta1DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 2.
	Rollbackto Iok8sapiappsv1beta1RollbackConfig `json:"rollbackTo,omitempty"` // DEPRECATED.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused.
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiappsv1DeploymentSpec represents the Iok8sapiappsv1DeploymentSpec schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentSpec struct {
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused.
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Strategy Iok8sapiappsv1DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// Iok8skubernetespkgapisrbacv1alpha1Role represents the Iok8skubernetespkgapisrbacv1alpha1Role schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1Role struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1alpha1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this Role
}

// Iok8skubernetespkgapiv1ContainerState represents the Iok8skubernetespkgapiv1ContainerState schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerState struct {
	Running Iok8sapicorev1ContainerStateRunning `json:"running,omitempty"` // ContainerStateRunning is a running state of a container.
	Terminated Iok8sapicorev1ContainerStateTerminated `json:"terminated,omitempty"` // ContainerStateTerminated is a terminated state of a container.
	Waiting Iok8sapicorev1ContainerStateWaiting `json:"waiting,omitempty"` // ContainerStateWaiting is a waiting state of a container.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version string `json:"version,omitempty"` // Version is the API version this server hosts. For example, "v1"
	Versionpriority int `json:"versionPriority"` // VersionPriority controls the ordering of this API version inside of its group. Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) Since it's inside of a group, the number can be small, probably in the 10s.
	Cabundle string `json:"caBundle"` // CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate.
	Group string `json:"group,omitempty"` // Group is the API group name this server hosts
	Grouppriorityminimum int `json:"groupPriorityMinimum"` // GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
	Insecureskiptlsverify bool `json:"insecureSkipTLSVerify,omitempty"` // InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged. You should use the CABundle instead.
	Service GeneratedType `json:"service"` // ServiceReference holds a reference to Service.legacy.k8s.io
}

// Iok8sapiappsv1beta2RollingUpdateDaemonSet represents the Iok8sapiappsv1beta2RollingUpdateDaemonSet schema from the OpenAPI specification
type Iok8sapiappsv1beta2RollingUpdateDaemonSet struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8skubernetespkgapisnetworkingv1NetworkPolicy represents the Iok8skubernetespkgapisnetworkingv1NetworkPolicy schema from the OpenAPI specification
type Iok8skubernetespkgapisnetworkingv1NetworkPolicy struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1NetworkPolicySpec `json:"spec,omitempty"` // NetworkPolicySpec provides the specification of a NetworkPolicy
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version string `json:"version,omitempty"` // Version is the API version this server hosts. For example, "v1"
	Versionpriority int `json:"versionPriority"` // VersionPriority controls the ordering of this API version inside of its group. Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) Since it's inside of a group, the number can be small, probably in the 10s.
	Cabundle string `json:"caBundle"` // CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate.
	Group string `json:"group,omitempty"` // Group is the API group name this server hosts
	Grouppriorityminimum int `json:"groupPriorityMinimum"` // GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
	Insecureskiptlsverify bool `json:"insecureSkipTLSVerify,omitempty"` // InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged. You should use the CABundle instead.
	Service GeneratedType `json:"service"` // ServiceReference holds a reference to Service.legacy.k8s.io
}

// Iok8sapiautoscalingv2beta1ObjectMetricSource represents the Iok8sapiautoscalingv2beta1ObjectMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1ObjectMetricSource struct {
	Targetvalue string `json:"targetValue"`
	Metricname string `json:"metricName"` // metricName is the name of the metric in question.
	Target Iok8sapiautoscalingv2beta1CrossVersionObjectReference `json:"target"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
}

// Iok8skubernetespkgapisextensionsv1beta1IDRange represents the Iok8skubernetespkgapisextensionsv1beta1IDRange schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IDRange struct {
	Max int64 `json:"max"` // Max is the end of the range, inclusive.
	Min int64 `json:"min"` // Min is the start of the range, inclusive.
}

// Iok8skubernetespkgapiv1ISCSIVolumeSource represents the Iok8skubernetespkgapiv1ISCSIVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ISCSIVolumeSource struct {
	Chapauthsession bool `json:"chapAuthSession,omitempty"` // whether support iSCSI Session CHAP authentication
	Iscsiinterface string `json:"iscsiInterface,omitempty"` // iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	Initiatorname string `json:"initiatorName,omitempty"` // Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Chapauthdiscovery bool `json:"chapAuthDiscovery,omitempty"` // whether support iSCSI Discovery CHAP authentication
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	Iqn string `json:"iqn"` // Target iSCSI Qualified Name.
	Lun int `json:"lun"` // iSCSI Target Lun number.
	Portals []string `json:"portals,omitempty"` // iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Targetportal string `json:"targetPortal"` // iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown.
	TypeField string `json:"type"` // Type is the type of the condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
}

// Iok8sapicorev1PersistentVolumeStatus represents the Iok8sapicorev1PersistentVolumeStatus schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeStatus struct {
	Reason string `json:"reason,omitempty"` // Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.
	Message string `json:"message,omitempty"` // A human-readable message indicating details about why the volume is in this state.
	Phase string `json:"phase,omitempty"` // Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
}

// Iok8sapicorev1ConfigMapVolumeSource represents the Iok8sapicorev1ConfigMapVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapVolumeSource struct {
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or it's keys must be defined
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8skubernetespkgapisrbacv1alpha1RoleList represents the Iok8skubernetespkgapisrbacv1alpha1RoleList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1RoleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1Role `json:"items"` // Items is a list of Roles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapiv1CephFSVolumeSource represents the Iok8skubernetespkgapiv1CephFSVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1CephFSVolumeSource struct {
	User string `json:"user,omitempty"` // Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors"` // Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Path string `json:"path,omitempty"` // Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Secretfile string `json:"secretFile,omitempty"` // Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
}

// Iok8sapiappsv1DaemonSetList represents the Iok8sapiappsv1DaemonSetList schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetList struct {
	Items []Iok8sapiappsv1DaemonSet `json:"items"` // A list of daemon sets.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1NetworkPolicy represents the Iok8sapiextensionsv1beta1NetworkPolicy schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicy struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1NetworkPolicySpec `json:"spec,omitempty"` // DEPRECATED 1.9 - This group version of NetworkPolicySpec is deprecated by networking/v1/NetworkPolicySpec.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1ServicePort represents the Iok8sapicorev1ServicePort schema from the OpenAPI specification
type Iok8sapicorev1ServicePort struct {
	Name string `json:"name,omitempty"` // The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. This maps to the 'Name' field in EndpointPort objects. Optional if only one ServicePort is defined on this service.
	Nodeport int `json:"nodePort,omitempty"` // The port on each node on which this service is exposed when type=NodePort or LoadBalancer. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ServiceType of this Service requires one. More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	Port int `json:"port"` // The port that will be exposed by this service.
	Protocol string `json:"protocol,omitempty"` // The IP protocol for this port. Supports "TCP" and "UDP". Default is TCP.
	Targetport string `json:"targetPort,omitempty"`
}

// Iok8skubernetespkgapiv1ReplicationControllerSpec represents the Iok8skubernetespkgapiv1ReplicationControllerSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ReplicationControllerSpec struct {
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Selector map[string]interface{} `json:"selector,omitempty"` // Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
}

// Iok8skubernetespkgapiv1DownwardAPIVolumeSource represents the Iok8skubernetespkgapiv1DownwardAPIVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1DownwardAPIVolumeSource struct {
	Items []Iok8sapicorev1DownwardAPIVolumeFile `json:"items,omitempty"` // Items is a list of downward API volume file
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8skubernetespkgapispolicyv1beta1Eviction represents the Iok8skubernetespkgapispolicyv1beta1Eviction schema from the OpenAPI specification
type Iok8skubernetespkgapispolicyv1beta1Eviction struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Deleteoptions Iok8sapimachinerypkgapismetav1DeleteOptions `json:"deleteOptions,omitempty"` // DeleteOptions may be provided when deleting an API object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapistoragev1alpha1VolumeAttachment represents the Iok8sapistoragev1alpha1VolumeAttachment schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttachment struct {
	Spec Iok8sapistoragev1alpha1VolumeAttachmentSpec `json:"spec"` // VolumeAttachmentSpec is the specification of a VolumeAttachment request.
	Status Iok8sapistoragev1alpha1VolumeAttachmentStatus `json:"status,omitempty"` // VolumeAttachmentStatus is the status of a VolumeAttachment request.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapiv1PodAntiAffinity represents the Iok8skubernetespkgapiv1PodAntiAffinity schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodAntiAffinity struct {
	Requiredduringschedulingignoredduringexecution []Iok8sapicorev1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
}

// Iok8sapiextensionsv1beta1HostPortRange represents the Iok8sapiextensionsv1beta1HostPortRange schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1HostPortRange struct {
	Max int `json:"max"` // max is the end of the range, inclusive.
	Min int `json:"min"` // min is the start of the range, inclusive.
}

// Iok8skubernetespkgapiv1ResourceQuotaList represents the Iok8skubernetespkgapiv1ResourceQuotaList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ResourceQuotaList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ResourceQuota `json:"items"` // Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1ConfigMapEnvSource represents the Iok8sapicorev1ConfigMapEnvSource schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapEnvSource struct {
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap must be defined
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8sapiauthorizationv1beta1ResourceAttributes represents the Iok8sapiauthorizationv1beta1ResourceAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1ResourceAttributes struct {
	Verb string `json:"verb,omitempty"` // Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Version string `json:"version,omitempty"` // Version is the API Version of the Resource. "*" means all.
	Group string `json:"group,omitempty"` // Group is the API Group of the Resource. "*" means all.
	Name string `json:"name,omitempty"` // Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the action being requested. Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Resource string `json:"resource,omitempty"` // Resource is one of the existing resource types. "*" means all.
	Subresource string `json:"subresource,omitempty"` // Subresource is one of the existing resource types. "" means none.
}

// Iok8sapiauthorizationv1beta1SubjectAccessReview represents the Iok8sapiauthorizationv1beta1SubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SubjectAccessReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
}

// Iok8skubernetespkgapisextensionsv1beta1DaemonSetSpec represents the Iok8skubernetespkgapisextensionsv1beta1DaemonSetSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DaemonSetSpec struct {
	Updatestrategy Iok8sapiextensionsv1beta1DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Templategeneration int64 `json:"templateGeneration,omitempty"` // DEPRECATED. A sequence number representing a specific generation of the template. Populated by the system. It can be set only during the creation.
}

// Iok8skubernetespkgapiv1ContainerImage represents the Iok8skubernetespkgapiv1ContainerImage schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerImage struct {
	Names []string `json:"names"` // Names by which this image is known. e.g. ["k8s.gcr.io/hyperkube:v1.0.7", "dockerhub.io/google_containers/hyperkube:v1.0.7"]
	Sizebytes int64 `json:"sizeBytes,omitempty"` // The size of the image in bytes.
}

// Iok8sapicorev1Container represents the Iok8sapicorev1Container schema from the OpenAPI specification
type Iok8sapicorev1Container struct {
	Command []string `json:"command,omitempty"` // Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Readinessprobe Iok8sapicorev1Probe `json:"readinessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Stdinonce bool `json:"stdinOnce,omitempty"` // Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	Envfrom []Iok8sapicorev1EnvFromSource `json:"envFrom,omitempty"` // List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	Workingdir string `json:"workingDir,omitempty"` // Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	Volumedevices []Iok8sapicorev1VolumeDevice `json:"volumeDevices,omitempty"` // volumeDevices is the list of block devices to be used by the container. This is an alpha feature and may change in the future.
	Lifecycle Iok8sapicorev1Lifecycle `json:"lifecycle,omitempty"` // Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Imagepullpolicy string `json:"imagePullPolicy,omitempty"` // Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	Name string `json:"name"` // Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	Terminationmessagepolicy string `json:"terminationMessagePolicy,omitempty"` // Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	Ports []Iok8sapicorev1ContainerPort `json:"ports,omitempty"` // List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.
	Args []string `json:"args,omitempty"` // Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Terminationmessagepath string `json:"terminationMessagePath,omitempty"` // Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	Tty bool `json:"tty,omitempty"` // Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Livenessprobe Iok8sapicorev1Probe `json:"livenessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Volumemounts []Iok8sapicorev1VolumeMount `json:"volumeMounts,omitempty"` // Pod volumes to mount into the container's filesystem. Cannot be updated.
	Securitycontext Iok8sapicorev1SecurityContext `json:"securityContext,omitempty"` // SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext. When both are set, the values in SecurityContext take precedence.
	Image string `json:"image,omitempty"` // Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
	Stdin bool `json:"stdin,omitempty"` // Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Env []Iok8sapicorev1EnvVar `json:"env,omitempty"` // List of environment variables to set in the container. Cannot be updated.
}

// Iok8sapirbacv1beta1RoleRef represents the Iok8sapirbacv1beta1RoleRef schema from the OpenAPI specification
type Iok8sapirbacv1beta1RoleRef struct {
	Apigroup string `json:"apiGroup"` // APIGroup is the group for the resource being referenced
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
}

// Iok8sapicorev1SecretEnvSource represents the Iok8sapicorev1SecretEnvSource schema from the OpenAPI specification
type Iok8sapicorev1SecretEnvSource struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret must be defined
}

// Iok8skubernetespkgapiv1ServiceSpec represents the Iok8skubernetespkgapiv1ServiceSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ServiceSpec struct {
	Publishnotreadyaddresses bool `json:"publishNotReadyAddresses,omitempty"` // publishNotReadyAddresses, when set to true, indicates that DNS implementations must publish the notReadyAddresses of subsets for the Endpoints associated with the Service. The default value is false. The primary use case for setting this field is to use a StatefulSet's Headless Service to propagate SRV records for its Pods without respect to their readiness for purpose of peer discovery. This field will replace the service.alpha.kubernetes.io/tolerate-unready-endpoints when that annotation is deprecated and all clients have been converted to use this field.
	Sessionaffinity string `json:"sessionAffinity,omitempty"` // Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Healthchecknodeport int `json:"healthCheckNodePort,omitempty"` // healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified, HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use user-specified nodePort value if specified by the client. Only effects when Type is set to LoadBalancer and ExternalTrafficPolicy is set to Local.
	Ports []Iok8sapicorev1ServicePort `json:"ports,omitempty"` // The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externaltrafficpolicy string `json:"externalTrafficPolicy,omitempty"` // externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. "Local" preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. "Cluster" obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.
	Loadbalancersourceranges []string `json:"loadBalancerSourceRanges,omitempty"` // If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature." More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	Selector map[string]interface{} `json:"selector,omitempty"` // Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
	Sessionaffinityconfig Iok8sapicorev1SessionAffinityConfig `json:"sessionAffinityConfig,omitempty"` // SessionAffinityConfig represents the configurations of session affinity.
	TypeField string `json:"type,omitempty"` // type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ExternalName" maps to the specified externalName. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a stable IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services---service-types
	Clusterip string `json:"clusterIP,omitempty"` // clusterIP is the IP address of the service and is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. This field can not be changed through updates. Valid values are "None", empty string (""), or a valid IP address. "None" can be specified for headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externalips []string `json:"externalIPs,omitempty"` // externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.
	Externalname string `json:"externalName,omitempty"` // externalName is the external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.
	Loadbalancerip string `json:"loadBalancerIP,omitempty"` // Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.
}

// Iok8sapistoragev1alpha1VolumeAttachmentStatus represents the Iok8sapistoragev1alpha1VolumeAttachmentStatus schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttachmentStatus struct {
	Attacherror Iok8sapistoragev1alpha1VolumeError `json:"attachError,omitempty"` // VolumeError captures an error encountered during a volume operation.
	Attached bool `json:"attached"` // Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attachmentmetadata map[string]interface{} `json:"attachmentMetadata,omitempty"` // Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Detacherror Iok8sapistoragev1alpha1VolumeError `json:"detachError,omitempty"` // VolumeError captures an error encountered during a volume operation.
}

// Iok8skubernetespkgapisbatchv1JobSpec represents the Iok8skubernetespkgapisbatchv1JobSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv1JobSpec struct {
	Parallelism int `json:"parallelism,omitempty"` // Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Activedeadlineseconds int64 `json:"activeDeadlineSeconds,omitempty"` // Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
	Backofflimit int `json:"backoffLimit,omitempty"` // Specifies the number of retries before marking this job failed. Defaults to 6
	Completions int `json:"completions,omitempty"` // Specifies the desired number of successfully finished pods the job should be run with. Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Manualselector bool `json:"manualSelector,omitempty"` // manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
}

// Iok8skubernetespkgapisrbacv1beta1ClusterRoleList represents the Iok8skubernetespkgapisrbacv1beta1ClusterRoleList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1ClusterRoleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1ClusterRole `json:"items"` // Items is a list of ClusterRoles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapiv1PersistentVolumeClaimStatus represents the Iok8skubernetespkgapiv1PersistentVolumeClaimStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeClaimStatus struct {
	Conditions []Iok8sapicorev1PersistentVolumeClaimCondition `json:"conditions,omitempty"` // Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
	Phase string `json:"phase,omitempty"` // Phase represents the current phase of PersistentVolumeClaim.
	Accessmodes []string `json:"accessModes,omitempty"` // AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Represents the actual resources of the underlying volume.
}

// Iok8sapiauthorizationv1SubjectAccessReviewStatus represents the Iok8sapiauthorizationv1SubjectAccessReviewStatus schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectAccessReviewStatus struct {
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	Reason string `json:"reason,omitempty"` // Reason is optional. It indicates why a request was allowed or denied.
	Allowed bool `json:"allowed"` // Allowed is required. True if the action would be allowed, false otherwise.
	Denied bool `json:"denied,omitempty"` // Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
}

// Iok8skubernetespkgapiv1ServiceStatus represents the Iok8skubernetespkgapiv1ServiceStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ServiceStatus struct {
	Loadbalancer Iok8sapicorev1LoadBalancerStatus `json:"loadBalancer,omitempty"` // LoadBalancerStatus represents the status of a load-balancer.
}

// Iok8sapiadmissionregistrationv1beta1RuleWithOperations represents the Iok8sapiadmissionregistrationv1beta1RuleWithOperations schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1RuleWithOperations struct {
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
	Apiversions []string `json:"apiVersions,omitempty"` // APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
	Operations []string `json:"operations,omitempty"` // Operations is the operations the admission hook cares about - CREATE, UPDATE, or * for all operations. If '*' is present, the length of the slice must be one. Required.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.
}

// Iok8sapistoragev1beta1VolumeAttachmentSpec represents the Iok8sapistoragev1beta1VolumeAttachmentSpec schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttachmentSpec struct {
	Attacher string `json:"attacher"` // Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Nodename string `json:"nodeName"` // The node that the volume should be attached to.
	Source Iok8sapistoragev1beta1VolumeAttachmentSource `json:"source"` // VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
}

// Iok8skubernetespkgapiv1DaemonEndpoint represents the Iok8skubernetespkgapiv1DaemonEndpoint schema from the OpenAPI specification
type Iok8skubernetespkgapiv1DaemonEndpoint struct {
	Port int `json:"Port"` // Port number of the given endpoint.
}

// Iok8sapicorev1ContainerStateTerminated represents the Iok8sapicorev1ContainerStateTerminated schema from the OpenAPI specification
type Iok8sapicorev1ContainerStateTerminated struct {
	Signal int `json:"signal,omitempty"` // Signal from the last termination of the container
	Startedat string `json:"startedAt,omitempty"`
	Containerid string `json:"containerID,omitempty"` // Container's ID in the format 'docker://<container_id>'
	Exitcode int `json:"exitCode"` // Exit status from the last termination of the container
	Finishedat string `json:"finishedAt,omitempty"`
	Message string `json:"message,omitempty"` // Message regarding the last termination of the container
	Reason string `json:"reason,omitempty"` // (brief) reason from the last termination of the container
}

// Iok8sapirbacv1beta1AggregationRule represents the Iok8sapirbacv1beta1AggregationRule schema from the OpenAPI specification
type Iok8sapirbacv1beta1AggregationRule struct {
	Clusterroleselectors []Iok8sapimachinerypkgapismetav1LabelSelector `json:"clusterRoleSelectors,omitempty"` // ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
}

// Iok8skubernetespkgapisrbacv1alpha1RoleRef represents the Iok8skubernetespkgapisrbacv1alpha1RoleRef schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1RoleRef struct {
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
	Apigroup string `json:"apiGroup"` // APIGroup is the group for the resource being referenced
}

// Iok8sapiappsv1beta1Deployment represents the Iok8sapiappsv1beta1Deployment schema from the OpenAPI specification
type Iok8sapiappsv1beta1Deployment struct {
	Spec Iok8sapiappsv1beta1DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
	Status Iok8sapiappsv1beta1DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReviewStatus represents the Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReviewStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReviewStatus struct {
	Denied bool `json:"denied,omitempty"` // Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	Reason string `json:"reason,omitempty"` // Reason is optional. It indicates why a request was allowed or denied.
	Allowed bool `json:"allowed"` // Allowed is required. True if the action would be allowed, false otherwise.
}

// Iok8sapibatchv2alpha1CronJobStatus represents the Iok8sapibatchv2alpha1CronJobStatus schema from the OpenAPI specification
type Iok8sapibatchv2alpha1CronJobStatus struct {
	Active []Iok8sapicorev1ObjectReference `json:"active,omitempty"` // A list of pointers to currently running jobs.
	Lastscheduletime string `json:"lastScheduleTime,omitempty"`
}

// Iok8skubernetespkgapisappsv1beta1StatefulSetList represents the Iok8skubernetespkgapisappsv1beta1StatefulSetList schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1StatefulSetList struct {
	Items []Iok8sapiappsv1beta1StatefulSet `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1HTTPIngressRuleValue represents the Iok8sapiextensionsv1beta1HTTPIngressRuleValue schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1HTTPIngressRuleValue struct {
	Paths []Iok8sapiextensionsv1beta1HTTPIngressPath `json:"paths"` // A collection of paths that map requests to backends.
}

// Iok8skubernetespkgapiv1ConfigMap represents the Iok8skubernetespkgapiv1ConfigMap schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ConfigMap struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Binarydata map[string]interface{} `json:"binaryData,omitempty"` // BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	Data map[string]interface{} `json:"data,omitempty"` // Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapirbacv1RoleRef represents the Iok8sapirbacv1RoleRef schema from the OpenAPI specification
type Iok8sapirbacv1RoleRef struct {
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
	Apigroup string `json:"apiGroup"` // APIGroup is the group for the resource being referenced
}

// Iok8sapicorev1ISCSIVolumeSource represents the Iok8sapicorev1ISCSIVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ISCSIVolumeSource struct {
	Initiatorname string `json:"initiatorName,omitempty"` // Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Chapauthdiscovery bool `json:"chapAuthDiscovery,omitempty"` // whether support iSCSI Discovery CHAP authentication
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	Iqn string `json:"iqn"` // Target iSCSI Qualified Name.
	Lun int `json:"lun"` // iSCSI Target Lun number.
	Portals []string `json:"portals,omitempty"` // iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Targetportal string `json:"targetPortal"` // iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Chapauthsession bool `json:"chapAuthSession,omitempty"` // whether support iSCSI Session CHAP authentication
	Iscsiinterface string `json:"iscsiInterface,omitempty"` // iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
}

// Iok8skubernetespkgapiv1ServicePort represents the Iok8skubernetespkgapiv1ServicePort schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ServicePort struct {
	Protocol string `json:"protocol,omitempty"` // The IP protocol for this port. Supports "TCP" and "UDP". Default is TCP.
	Targetport string `json:"targetPort,omitempty"`
	Name string `json:"name,omitempty"` // The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. This maps to the 'Name' field in EndpointPort objects. Optional if only one ServicePort is defined on this service.
	Nodeport int `json:"nodePort,omitempty"` // The port on each node on which this service is exposed when type=NodePort or LoadBalancer. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ServiceType of this Service requires one. More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	Port int `json:"port"` // The port that will be exposed by this service.
}

// Iok8skubernetespkgapisextensionsv1beta1NetworkPolicySpec represents the Iok8skubernetespkgapisextensionsv1beta1NetworkPolicySpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1NetworkPolicySpec struct {
	Policytypes []string `json:"policyTypes,omitempty"` // List of rule types that the NetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	Egress []Iok8sapiextensionsv1beta1NetworkPolicyEgressRule `json:"egress,omitempty"` // List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
	Ingress []Iok8sapiextensionsv1beta1NetworkPolicyIngressRule `json:"ingress,omitempty"` // List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default).
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8skubernetespkgapiv1GCEPersistentDiskVolumeSource represents the Iok8skubernetespkgapiv1GCEPersistentDiskVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1GCEPersistentDiskVolumeSource struct {
	Pdname string `json:"pdName"` // Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Partition int `json:"partition,omitempty"` // The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
}

// Iok8sapicorev1PodDNSConfig represents the Iok8sapicorev1PodDNSConfig schema from the OpenAPI specification
type Iok8sapicorev1PodDNSConfig struct {
	Options []Iok8sapicorev1PodDNSConfigOption `json:"options,omitempty"` // A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.
	Searches []string `json:"searches,omitempty"` // A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
	Nameservers []string `json:"nameservers,omitempty"` // A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
}

// Iok8sapirbacv1AggregationRule represents the Iok8sapirbacv1AggregationRule schema from the OpenAPI specification
type Iok8sapirbacv1AggregationRule struct {
	Clusterroleselectors []Iok8sapimachinerypkgapismetav1LabelSelector `json:"clusterRoleSelectors,omitempty"` // ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
}

// Iok8sapiappsv1beta2ReplicaSetSpec represents the Iok8sapiappsv1beta2ReplicaSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta2ReplicaSetSpec struct {
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
}

// Iok8skubernetespkgapiv1LimitRange represents the Iok8skubernetespkgapiv1LimitRange schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LimitRange struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1LimitRangeSpec `json:"spec,omitempty"` // LimitRangeSpec defines a min/max usage limit for resources that match on kind.
}

// Iok8sapicorev1EventList represents the Iok8sapicorev1EventList schema from the OpenAPI specification
type Iok8sapicorev1EventList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Event `json:"items"` // List of events
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1RollingUpdateDaemonSet represents the Iok8sapiappsv1RollingUpdateDaemonSet schema from the OpenAPI specification
type Iok8sapiappsv1RollingUpdateDaemonSet struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8sapisettingsv1alpha1PodPreset represents the Iok8sapisettingsv1alpha1PodPreset schema from the OpenAPI specification
type Iok8sapisettingsv1alpha1PodPreset struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapisettingsv1alpha1PodPresetSpec `json:"spec,omitempty"` // PodPresetSpec is a description of a pod preset.
}

// Iok8skubernetespkgapiv1PersistentVolumeClaimVolumeSource represents the Iok8skubernetespkgapiv1PersistentVolumeClaimVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeClaimVolumeSource struct {
	Claimname string `json:"claimName"` // ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Readonly bool `json:"readOnly,omitempty"` // Will force the ReadOnly setting in VolumeMounts. Default false.
}

// Iok8sapiextensionsv1beta1IngressStatus represents the Iok8sapiextensionsv1beta1IngressStatus schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IngressStatus struct {
	Loadbalancer Iok8sapicorev1LoadBalancerStatus `json:"loadBalancer,omitempty"` // LoadBalancerStatus represents the status of a load-balancer.
}

// Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestStatus represents the Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestStatus struct {
	Certificate string `json:"certificate,omitempty"` // If request was approved, the controller will place the issued certificate here.
	Conditions []Iok8sapicertificatesv1beta1CertificateSigningRequestCondition `json:"conditions,omitempty"` // Conditions applied to the request, such as approval or denial.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec represents the Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec struct {
	Scaletargetref Iok8sapiautoscalingv1CrossVersionObjectReference `json:"scaleTargetRef"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
	Targetcpuutilizationpercentage int `json:"targetCPUUtilizationPercentage,omitempty"` // target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
	Maxreplicas int `json:"maxReplicas"` // upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	Minreplicas int `json:"minReplicas,omitempty"` // lower limit for the number of pods that can be set by the autoscaler, default 1.
}

// Iok8skubernetespkgapisappsv1beta1RollingUpdateStatefulSetStrategy represents the Iok8skubernetespkgapisappsv1beta1RollingUpdateStatefulSetStrategy schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1RollingUpdateStatefulSetStrategy struct {
	Partition int `json:"partition,omitempty"` // Partition indicates the ordinal at which the StatefulSet should be partitioned.
}

// Iok8sapicorev1PodTemplateList represents the Iok8sapicorev1PodTemplateList schema from the OpenAPI specification
type Iok8sapicorev1PodTemplateList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1PodTemplate `json:"items"` // List of pod templates
}

// Iok8sapicorev1ResourceFieldSelector represents the Iok8sapicorev1ResourceFieldSelector schema from the OpenAPI specification
type Iok8sapicorev1ResourceFieldSelector struct {
	Containername string `json:"containerName,omitempty"` // Container name: required for volumes, optional for env vars
	Divisor string `json:"divisor,omitempty"`
	Resource string `json:"resource"` // Required: resource to select
}

// Iok8skubernetespkgapisrbacv1alpha1ClusterRoleList represents the Iok8skubernetespkgapisrbacv1alpha1ClusterRoleList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1ClusterRoleList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1ClusterRole `json:"items"` // Items is a list of ClusterRoles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta1DeploymentRollback represents the Iok8sapiappsv1beta1DeploymentRollback schema from the OpenAPI specification
type Iok8sapiappsv1beta1DeploymentRollback struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name"` // Required: This must match the Name of a deployment.
	Rollbackto Iok8sapiappsv1beta1RollbackConfig `json:"rollbackTo"` // DEPRECATED.
	Updatedannotations map[string]interface{} `json:"updatedAnnotations,omitempty"` // The annotations to be updated to a deployment
}

// Iok8sapicorev1PreferredSchedulingTerm represents the Iok8sapicorev1PreferredSchedulingTerm schema from the OpenAPI specification
type Iok8sapicorev1PreferredSchedulingTerm struct {
	Preference Iok8sapicorev1NodeSelectorTerm `json:"preference"` // A null or empty node selector term matches no objects.
	Weight int `json:"weight"` // Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
}

// Iok8skubernetespkgapiv1ResourceQuota represents the Iok8skubernetespkgapiv1ResourceQuota schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ResourceQuota struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ResourceQuotaSpec `json:"spec,omitempty"` // ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
	Status Iok8sapicorev1ResourceQuotaStatus `json:"status,omitempty"` // ResourceQuotaStatus defines the enforced hard limits and observed use.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1NetworkPolicyIngressRule represents the Iok8sapiextensionsv1beta1NetworkPolicyIngressRule schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicyIngressRule struct {
	From []Iok8sapiextensionsv1beta1NetworkPolicyPeer `json:"from,omitempty"` // List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.
	Ports []Iok8sapiextensionsv1beta1NetworkPolicyPort `json:"ports,omitempty"` // List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
}

// Iok8sapiappsv1beta2DaemonSetStatus represents the Iok8sapiappsv1beta2DaemonSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta2DaemonSetStatus struct {
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Numberready int `json:"numberReady"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	Conditions []Iok8sapiappsv1beta2DaemonSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a DaemonSet's current state.
	Desirednumberscheduled int `json:"desiredNumberScheduled"` // The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Currentnumberscheduled int `json:"currentNumberScheduled"` // The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Numberunavailable int `json:"numberUnavailable,omitempty"` // The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The most recent generation observed by the daemon set controller.
	Numberavailable int `json:"numberAvailable,omitempty"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numbermisscheduled int `json:"numberMisscheduled"` // The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Updatednumberscheduled int `json:"updatedNumberScheduled,omitempty"` // The total number of nodes that are running updated daemon pod
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Items []GeneratedType `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1PodStatus represents the Iok8skubernetespkgapiv1PodStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodStatus struct {
	Qosclass string `json:"qosClass,omitempty"` // The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md
	Reason string `json:"reason,omitempty"` // A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'
	Containerstatuses []Iok8sapicorev1ContainerStatus `json:"containerStatuses,omitempty"` // The list has one entry per container in the manifest. Each entry is currently the output of `docker inspect`. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	Initcontainerstatuses []Iok8sapicorev1ContainerStatus `json:"initContainerStatuses,omitempty"` // The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	Nominatednodename string `json:"nominatedNodeName,omitempty"` // nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
	Starttime string `json:"startTime,omitempty"`
	Hostip string `json:"hostIP,omitempty"` // IP address of the host to which the pod is assigned. Empty if not yet scheduled.
	Conditions []Iok8sapicorev1PodCondition `json:"conditions,omitempty"` // Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Message string `json:"message,omitempty"` // A human readable message indicating details about why the pod is in this condition.
	Phase string `json:"phase,omitempty"` // Current condition of the pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
	Podip string `json:"podIP,omitempty"` // IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
}

// Iok8sapiauthenticationv1beta1UserInfo represents the Iok8sapiauthenticationv1beta1UserInfo schema from the OpenAPI specification
type Iok8sapiauthenticationv1beta1UserInfo struct {
	Groups []string `json:"groups,omitempty"` // The names of groups this user is a part of.
	Uid string `json:"uid,omitempty"` // A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	Username string `json:"username,omitempty"` // The name that uniquely identifies this user among all active users.
	Extra map[string]interface{} `json:"extra,omitempty"` // Any additional information provided by the authenticator.
}

// Iok8sapiappsv1DeploymentCondition represents the Iok8sapiappsv1DeploymentCondition schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentCondition struct {
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
}

// Iok8sapiextensionsv1beta1FSGroupStrategyOptions represents the Iok8sapiextensionsv1beta1FSGroupStrategyOptions schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1FSGroupStrategyOptions struct {
	Ranges []Iok8sapiextensionsv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of fs groups. If you would like to force a single fs group then supply a single range with the same start and end.
	Rule string `json:"rule,omitempty"` // Rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
}

// Iok8sapiappsv1beta2DaemonSet represents the Iok8sapiappsv1beta2DaemonSet schema from the OpenAPI specification
type Iok8sapiappsv1beta2DaemonSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta2DaemonSetSpec `json:"spec,omitempty"` // DaemonSetSpec is the specification of a daemon set.
	Status Iok8sapiappsv1beta2DaemonSetStatus `json:"status,omitempty"` // DaemonSetStatus represents the current status of a daemon set.
}

// Iok8sapicorev1PersistentVolumeList represents the Iok8sapicorev1PersistentVolumeList schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1PersistentVolume `json:"items"` // List of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapicorev1GlusterfsVolumeSource represents the Iok8sapicorev1GlusterfsVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GlusterfsVolumeSource struct {
	Endpoints string `json:"endpoints"` // EndpointsName is the endpoint name that details Glusterfs topology. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path"` // Path is the Glusterfs volume path. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
}

// Iok8sapiextensionsv1beta1ScaleSpec represents the Iok8sapiextensionsv1beta1ScaleSpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8skubernetespkgapisauthorizationv1SelfSubjectAccessReview represents the Iok8skubernetespkgapisauthorizationv1SelfSubjectAccessReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1SelfSubjectAccessReview struct {
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec `json:"spec"` // SelfSubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
}

// Iok8skubernetespkgapiv1PodAffinity represents the Iok8skubernetespkgapiv1PodAffinity schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution []Iok8sapicorev1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
}

// Iok8skubernetespkgapisextensionsv1beta1Deployment represents the Iok8skubernetespkgapisextensionsv1beta1Deployment schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1Deployment struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
	Status Iok8sapiextensionsv1beta1DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1ReplicaSetCondition represents the Iok8sapiappsv1ReplicaSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetCondition struct {
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of replica set condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
}

// Iok8sapisettingsv1alpha1PodPresetList represents the Iok8sapisettingsv1alpha1PodPresetList schema from the OpenAPI specification
type Iok8sapisettingsv1alpha1PodPresetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapisettingsv1alpha1PodPreset `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1PersistentVolume represents the Iok8sapicorev1PersistentVolume schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolume struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PersistentVolumeSpec `json:"spec,omitempty"` // PersistentVolumeSpec is the specification of a persistent volume.
	Status Iok8sapicorev1PersistentVolumeStatus `json:"status,omitempty"` // PersistentVolumeStatus is the current status of a persistent volume.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapisbatchv1JobCondition represents the Iok8skubernetespkgapisbatchv1JobCondition schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv1JobCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // (brief) reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of job condition, Complete or Failed.
	Lastprobetime string `json:"lastProbeTime,omitempty"`
}

// Iok8sapicorev1VolumeProjection represents the Iok8sapicorev1VolumeProjection schema from the OpenAPI specification
type Iok8sapicorev1VolumeProjection struct {
	Configmap Iok8sapicorev1ConfigMapProjection `json:"configMap,omitempty"` // Adapts a ConfigMap into a projected volume. The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.
	Downwardapi Iok8sapicorev1DownwardAPIProjection `json:"downwardAPI,omitempty"` // Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
	Secret Iok8sapicorev1SecretProjection `json:"secret,omitempty"` // Adapts a secret into a projected volume. The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.
}

// Iok8skubernetespkgapisextensionsv1beta1RollbackConfig represents the Iok8skubernetespkgapisextensionsv1beta1RollbackConfig schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1RollbackConfig struct {
	Revision int64 `json:"revision,omitempty"` // The revision to rollback to. If set to 0, rollback to the last revision.
}

// Iok8skubernetespkgapiv1NodeAddress represents the Iok8skubernetespkgapiv1NodeAddress schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeAddress struct {
	Address string `json:"address"` // The node address.
	TypeField string `json:"type"` // Node address type, one of Hostname, ExternalIP or InternalIP.
}

// Iok8sapicertificatesv1beta1CertificateSigningRequestSpec represents the Iok8sapicertificatesv1beta1CertificateSigningRequestSpec schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1CertificateSigningRequestSpec struct {
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra information about the requesting user. See user.Info interface for details.
	Groups []string `json:"groups,omitempty"` // Group information about the requesting user. See user.Info interface for details.
	Request string `json:"request"` // Base64-encoded PKCS#10 CSR data
	Uid string `json:"uid,omitempty"` // UID information about the requesting user. See user.Info interface for details.
	Usages []string `json:"usages,omitempty"` // allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3 https://tools.ietf.org/html/rfc5280#section-4.2.1.12
	Username string `json:"username,omitempty"` // Information about the requesting user. See user.Info interface for details.
}

// Iok8sapicorev1NodeAffinity represents the Iok8sapicorev1NodeAffinity schema from the OpenAPI specification
type Iok8sapicorev1NodeAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution Iok8sapicorev1NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8sapicorev1PodTemplate represents the Iok8sapicorev1PodTemplate schema from the OpenAPI specification
type Iok8sapicorev1PodTemplate struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// Iok8sapicorev1ContainerStatus represents the Iok8sapicorev1ContainerStatus schema from the OpenAPI specification
type Iok8sapicorev1ContainerStatus struct {
	State Iok8sapicorev1ContainerState `json:"state,omitempty"` // ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	Containerid string `json:"containerID,omitempty"` // Container's ID in the format 'docker://<container_id>'.
	Image string `json:"image"` // The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images
	Imageid string `json:"imageID"` // ImageID of the container's image.
	Laststate Iok8sapicorev1ContainerState `json:"lastState,omitempty"` // ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	Name string `json:"name"` // This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
	Ready bool `json:"ready"` // Specifies whether the container has passed its readiness probe.
	Restartcount int `json:"restartCount"` // The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC.
}

// Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec represents the Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec struct {
	Resourceattributes Iok8sapiauthorizationv1beta1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
	Uid string `json:"uid,omitempty"` // UID information about the requesting user.
	User string `json:"user,omitempty"` // User is the user you're testing for. If you specify "User" but not "Group", then is it interpreted as "What if User were not a member of any groups
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra corresponds to the user.Info.GetExtra() method from the authenticator. Since that is input to the authorizer it needs a reflection here.
	Group []string `json:"group,omitempty"` // Groups is the groups you're testing for.
	Nonresourceattributes Iok8sapiauthorizationv1beta1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
}

// Iok8sapicorev1Node represents the Iok8sapicorev1Node schema from the OpenAPI specification
type Iok8sapicorev1Node struct {
	Spec Iok8sapicorev1NodeSpec `json:"spec,omitempty"` // NodeSpec describes the attributes that a node is created with.
	Status Iok8sapicorev1NodeStatus `json:"status,omitempty"` // NodeStatus is information about the current status of a node.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiappsv1ReplicaSetStatus represents the Iok8sapiappsv1ReplicaSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetStatus struct {
	Replicas int `json:"replicas"` // Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replica set.
	Conditions []Iok8sapiappsv1ReplicaSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replica set's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replicaset.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replica set.
}

// Iok8skubernetespkgapiv1FlockerVolumeSource represents the Iok8skubernetespkgapiv1FlockerVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1FlockerVolumeSource struct {
	Datasetname string `json:"datasetName,omitempty"` // Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	Datasetuuid string `json:"datasetUUID,omitempty"` // UUID of the dataset. This is unique identifier of a Flocker dataset
}

// Iok8sapicorev1NodeSystemInfo represents the Iok8sapicorev1NodeSystemInfo schema from the OpenAPI specification
type Iok8sapicorev1NodeSystemInfo struct {
	Architecture string `json:"architecture"` // The Architecture reported by the node
	Kernelversion string `json:"kernelVersion"` // Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	Operatingsystem string `json:"operatingSystem"` // The Operating System reported by the node
	Containerruntimeversion string `json:"containerRuntimeVersion"` // ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	Systemuuid string `json:"systemUUID"` // SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	Bootid string `json:"bootID"` // Boot ID reported by the node.
	Kubeletversion string `json:"kubeletVersion"` // Kubelet Version reported by the node.
	Machineid string `json:"machineID"` // MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	Kubeproxyversion string `json:"kubeProxyVersion"` // KubeProxy Version reported by the node.
	Osimage string `json:"osImage"` // OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
}

// Iok8skubernetespkgapiv1LoadBalancerStatus represents the Iok8skubernetespkgapiv1LoadBalancerStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LoadBalancerStatus struct {
	Ingress []Iok8sapicorev1LoadBalancerIngress `json:"ingress,omitempty"` // Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
}

// Iok8skubernetespkgapiv1NodeSelectorTerm represents the Iok8skubernetespkgapiv1NodeSelectorTerm schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeSelectorTerm struct {
	Matchexpressions []Iok8sapicorev1NodeSelectorRequirement `json:"matchExpressions"` // Required. A list of node selector requirements. The requirements are ANDed.
}

// Iok8sapipolicyv1beta1FSGroupStrategyOptions represents the Iok8sapipolicyv1beta1FSGroupStrategyOptions schema from the OpenAPI specification
type Iok8sapipolicyv1beta1FSGroupStrategyOptions struct {
	Ranges []Iok8sapipolicyv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of fs groups. If you would like to force a single fs group then supply a single range with the same start and end.
	Rule string `json:"rule,omitempty"` // Rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
}

// Iok8sapimachinerypkgapismetav1Initializer represents the Iok8sapimachinerypkgapismetav1Initializer schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Initializer struct {
	Name string `json:"name"` // name of the process that is responsible for initializing this object.
}

// Iok8skubernetespkgapisextensionsv1beta1ScaleStatus represents the Iok8skubernetespkgapisextensionsv1beta1ScaleStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ScaleStatus struct {
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
	Selector map[string]interface{} `json:"selector,omitempty"` // label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Targetselector string `json:"targetSelector,omitempty"` // label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
}

// Iok8sapicorev1SecretList represents the Iok8sapicorev1SecretList schema from the OpenAPI specification
type Iok8sapicorev1SecretList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Secret `json:"items"` // Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapirbacv1alpha1RoleRef represents the Iok8sapirbacv1alpha1RoleRef schema from the OpenAPI specification
type Iok8sapirbacv1alpha1RoleRef struct {
	Apigroup string `json:"apiGroup"` // APIGroup is the group for the resource being referenced
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
}

// Iok8skubernetespkgapisrbacv1beta1PolicyRule represents the Iok8skubernetespkgapisrbacv1beta1PolicyRule schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1PolicyRule struct {
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. '*' represents all resources in the specified apiGroups. '*/foo' represents the subresource 'foo' for all resources in the specified apiGroups.
	Verbs []string `json:"verbs"` // Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
}

// Iok8sapicorev1ReplicationControllerCondition represents the Iok8sapicorev1ReplicationControllerCondition schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of replication controller condition.
}

// Iok8skubernetespkgapisextensionsv1beta1DeploymentSpec represents the Iok8skubernetespkgapisextensionsv1beta1DeploymentSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DeploymentSpec struct {
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. This is not set by default.
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Strategy Iok8sapiextensionsv1beta1DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Rollbackto Iok8sapiextensionsv1beta1RollbackConfig `json:"rollbackTo,omitempty"` // DEPRECATED.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused and will not be processed by the deployment controller.
}

// Iok8skubernetespkgapisadmissionregistrationv1alpha1Rule represents the Iok8skubernetespkgapisadmissionregistrationv1alpha1Rule schema from the OpenAPI specification
type Iok8skubernetespkgapisadmissionregistrationv1alpha1Rule struct {
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
	Apiversions []string `json:"apiVersions,omitempty"` // APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.
}

// Iok8sapirbacv1beta1Subject represents the Iok8sapirbacv1beta1Subject schema from the OpenAPI specification
type Iok8sapirbacv1beta1Subject struct {
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	Kind string `json:"kind"` // Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Name string `json:"name"` // Name of the object being referenced.
	Namespace string `json:"namespace,omitempty"` // Namespace of the referenced object. If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
}

// Iok8sapiauthenticationv1TokenReviewStatus represents the Iok8sapiauthenticationv1TokenReviewStatus schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenReviewStatus struct {
	Authenticated bool `json:"authenticated,omitempty"` // Authenticated indicates that the token was associated with a known user.
	ErrorField string `json:"error,omitempty"` // Error indicates that the token couldn't be checked
	User Iok8sapiauthenticationv1UserInfo `json:"user,omitempty"` // UserInfo holds the information about the user needed to implement the user.Info interface.
}

// Iok8sapicorev1VolumeMount represents the Iok8sapicorev1VolumeMount schema from the OpenAPI specification
type Iok8sapicorev1VolumeMount struct {
	Subpath string `json:"subPath,omitempty"` // Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	Mountpath string `json:"mountPath"` // Path within the container at which the volume should be mounted. Must not contain ':'.
	Mountpropagation string `json:"mountPropagation,omitempty"` // mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationHostToContainer is used. This field is beta in 1.10.
	Name string `json:"name"` // This must match the Name of a Volume.
	Readonly bool `json:"readOnly,omitempty"` // Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Conditions []GeneratedType `json:"conditions,omitempty"` // Current service state of apiService.
}

// Iok8sapicorev1EndpointAddress represents the Iok8sapicorev1EndpointAddress schema from the OpenAPI specification
type Iok8sapicorev1EndpointAddress struct {
	Nodename string `json:"nodeName,omitempty"` // Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
	Targetref Iok8sapicorev1ObjectReference `json:"targetRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Hostname string `json:"hostname,omitempty"` // The Hostname of this endpoint
	Ip string `json:"ip"` // The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready.
}

// Iok8sapiadmissionregistrationv1beta1WebhookClientConfig represents the Iok8sapiadmissionregistrationv1beta1WebhookClientConfig schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1WebhookClientConfig struct {
	Cabundle string `json:"caBundle"` // `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. Required.
	Service Iok8sapiadmissionregistrationv1beta1ServiceReference `json:"service,omitempty"` // ServiceReference holds a reference to Service.legacy.k8s.io
	Url string `json:"url,omitempty"` // `url` gives the location of the webhook, in standard URL form (`[scheme://]host:port/path`). Exactly one of `url` or `service` must be specified. The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address. Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster. The scheme must be "https"; the URL must begin with "https://". A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier. Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
}

// Iok8sapibatchv1beta1JobTemplateSpec represents the Iok8sapibatchv1beta1JobTemplateSpec schema from the OpenAPI specification
type Iok8sapibatchv1beta1JobTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
}

// Iok8skubernetespkgapiv1ContainerStateRunning represents the Iok8skubernetespkgapiv1ContainerStateRunning schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerStateRunning struct {
	Startedat string `json:"startedAt,omitempty"`
}

// Iok8skubernetespkgapiv1ContainerStatus represents the Iok8skubernetespkgapiv1ContainerStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ContainerStatus struct {
	Name string `json:"name"` // This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
	Ready bool `json:"ready"` // Specifies whether the container has passed its readiness probe.
	Restartcount int `json:"restartCount"` // The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC.
	State Iok8sapicorev1ContainerState `json:"state,omitempty"` // ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	Containerid string `json:"containerID,omitempty"` // Container's ID in the format 'docker://<container_id>'.
	Image string `json:"image"` // The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images
	Imageid string `json:"imageID"` // ImageID of the container's image.
	Laststate Iok8sapicorev1ContainerState `json:"lastState,omitempty"` // ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
}

// Iok8sapirbacv1alpha1ClusterRole represents the Iok8sapirbacv1alpha1ClusterRole schema from the OpenAPI specification
type Iok8sapirbacv1alpha1ClusterRole struct {
	Rules []Iok8sapirbacv1alpha1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this ClusterRole
	Aggregationrule Iok8sapirbacv1alpha1AggregationRule `json:"aggregationRule,omitempty"` // AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiextensionsv1beta1IDRange represents the Iok8sapiextensionsv1beta1IDRange schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IDRange struct {
	Max int64 `json:"max"` // Max is the end of the range, inclusive.
	Min int64 `json:"min"` // Min is the start of the range, inclusive.
}

// Iok8sapiappsv1beta2StatefulSet represents the Iok8sapiappsv1beta2StatefulSet schema from the OpenAPI specification
type Iok8sapiappsv1beta2StatefulSet struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta2StatefulSetSpec `json:"spec,omitempty"` // A StatefulSetSpec is the specification of a StatefulSet.
	Status Iok8sapiappsv1beta2StatefulSetStatus `json:"status,omitempty"` // StatefulSetStatus represents the current state of a StatefulSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1NamespaceSpec represents the Iok8skubernetespkgapiv1NamespaceSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NamespaceSpec struct {
	Finalizers []string `json:"finalizers,omitempty"` // Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
}

// Iok8sapicorev1AzureFilePersistentVolumeSource represents the Iok8sapicorev1AzureFilePersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AzureFilePersistentVolumeSource struct {
	Secretnamespace string `json:"secretNamespace,omitempty"` // the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
	Sharename string `json:"shareName"` // Share Name
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretname string `json:"secretName"` // the name of secret that contains Azure Storage Account Name and Key
}

// Iok8sapicorev1AWSElasticBlockStoreVolumeSource represents the Iok8sapicorev1AWSElasticBlockStoreVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AWSElasticBlockStoreVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	Partition int `json:"partition,omitempty"` // The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Readonly bool `json:"readOnly,omitempty"` // Specify "true" to force and set the ReadOnly property in VolumeMounts to "true". If omitted, the default is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	Volumeid string `json:"volumeID"` // Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
}

// Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestSpec represents the Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestSpec struct {
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra information about the requesting user. See user.Info interface for details.
	Groups []string `json:"groups,omitempty"` // Group information about the requesting user. See user.Info interface for details.
	Request string `json:"request"` // Base64-encoded PKCS#10 CSR data
	Uid string `json:"uid,omitempty"` // UID information about the requesting user. See user.Info interface for details.
	Usages []string `json:"usages,omitempty"` // allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3 https://tools.ietf.org/html/rfc5280#section-4.2.1.12
	Username string `json:"username,omitempty"` // Information about the requesting user. See user.Info interface for details.
}

// Iok8sapicorev1NodeCondition represents the Iok8sapicorev1NodeCondition schema from the OpenAPI specification
type Iok8sapicorev1NodeCondition struct {
	Reason string `json:"reason,omitempty"` // (brief) reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of node condition.
	Lastheartbeattime string `json:"lastHeartbeatTime,omitempty"`
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human readable message indicating details about last transition.
}

// Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicy represents the Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicy schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicy struct {
	Spec Iok8sapiextensionsv1beta1PodSecurityPolicySpec `json:"spec,omitempty"` // Pod Security Policy Spec defines the policy enforced.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiappsv1beta2StatefulSetList represents the Iok8sapiappsv1beta2StatefulSetList schema from the OpenAPI specification
type Iok8sapiappsv1beta2StatefulSetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta2StatefulSet `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiauthorizationv1NonResourceAttributes represents the Iok8sapiauthorizationv1NonResourceAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1NonResourceAttributes struct {
	Path string `json:"path,omitempty"` // Path is the URL path of the request
	Verb string `json:"verb,omitempty"` // Verb is the standard HTTP verb
}

// Iok8sapicorev1EnvVarSource represents the Iok8sapicorev1EnvVarSource schema from the OpenAPI specification
type Iok8sapicorev1EnvVarSource struct {
	Secretkeyref Iok8sapicorev1SecretKeySelector `json:"secretKeyRef,omitempty"` // SecretKeySelector selects a key of a Secret.
	Configmapkeyref Iok8sapicorev1ConfigMapKeySelector `json:"configMapKeyRef,omitempty"` // Selects a key from a ConfigMap.
	Fieldref Iok8sapicorev1ObjectFieldSelector `json:"fieldRef,omitempty"` // ObjectFieldSelector selects an APIVersioned field of an object.
	Resourcefieldref Iok8sapicorev1ResourceFieldSelector `json:"resourceFieldRef,omitempty"` // ResourceFieldSelector represents container resources (cpu, memory) and their output format
}

// Iok8sapibatchv2alpha1CronJob represents the Iok8sapibatchv2alpha1CronJob schema from the OpenAPI specification
type Iok8sapibatchv2alpha1CronJob struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv2alpha1CronJobSpec `json:"spec,omitempty"` // CronJobSpec describes how the job execution will look like and when it will actually run.
	Status Iok8sapibatchv2alpha1CronJobStatus `json:"status,omitempty"` // CronJobStatus represents the current state of a cron job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1ObjectReference represents the Iok8sapicorev1ObjectReference schema from the OpenAPI specification
type Iok8sapicorev1ObjectReference struct {
	Kind string `json:"kind,omitempty"` // Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Namespace string `json:"namespace,omitempty"` // Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Resourceversion string `json:"resourceVersion,omitempty"` // Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	Uid string `json:"uid,omitempty"` // UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent.
	Fieldpath string `json:"fieldPath,omitempty"` // If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
}

// Iok8sapiextensionsv1beta1DeploymentCondition represents the Iok8sapiextensionsv1beta1DeploymentCondition schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DeploymentCondition struct {
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
}

// Iok8skubernetespkgapiv1SecretProjection represents the Iok8skubernetespkgapiv1SecretProjection schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SecretProjection struct {
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8skubernetespkgapisappsv1beta1DeploymentStrategy represents the Iok8skubernetespkgapisappsv1beta1DeploymentStrategy schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1DeploymentStrategy struct {
	Rollingupdate Iok8sapiappsv1beta1RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
}

// Iok8skubernetespkgapiv1LoadBalancerIngress represents the Iok8skubernetespkgapiv1LoadBalancerIngress schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LoadBalancerIngress struct {
	Ip string `json:"ip,omitempty"` // IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
	Hostname string `json:"hostname,omitempty"` // Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
}

// Iok8sapiauthorizationv1SelfSubjectRulesReview represents the Iok8sapiauthorizationv1SelfSubjectRulesReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectRulesReview struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec `json:"spec"`
	Status Iok8sapiauthorizationv1SubjectRulesReviewStatus `json:"status,omitempty"` // SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1FlexVolumeSource represents the Iok8sapicorev1FlexVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FlexVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Driver string `json:"driver"` // Driver is the name of the driver to use for this volume.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	Options map[string]interface{} `json:"options,omitempty"` // Optional: Extra command options if any.
}

// Iok8skubernetespkgapiv1ScaleIOVolumeSource represents the Iok8skubernetespkgapiv1ScaleIOVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ScaleIOVolumeSource struct {
	Protectiondomain string `json:"protectionDomain,omitempty"` // The name of the ScaleIO Protection Domain for the configured storage.
	Sslenabled bool `json:"sslEnabled,omitempty"` // Flag to enable/disable SSL communication with Gateway, default false
	Volumename string `json:"volumeName,omitempty"` // The name of a volume already created in the ScaleIO system that is associated with this volume source.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Gateway string `json:"gateway"` // The host address of the ScaleIO API Gateway.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	System string `json:"system"` // The name of the storage system as configured in ScaleIO.
	Storagemode string `json:"storageMode,omitempty"` // Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	Storagepool string `json:"storagePool,omitempty"` // The ScaleIO Storage Pool associated with the protection domain.
}

// Iok8sapiappsv1ControllerRevisionList represents the Iok8sapiappsv1ControllerRevisionList schema from the OpenAPI specification
type Iok8sapiappsv1ControllerRevisionList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1ControllerRevision `json:"items"` // Items is the list of ControllerRevisions
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapipolicyv1beta1PodSecurityPolicy represents the Iok8sapipolicyv1beta1PodSecurityPolicy schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodSecurityPolicy struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapipolicyv1beta1PodSecurityPolicySpec `json:"spec,omitempty"` // Pod Security Policy Spec defines the policy enforced.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapisextensionsv1beta1HTTPIngressRuleValue represents the Iok8skubernetespkgapisextensionsv1beta1HTTPIngressRuleValue schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1HTTPIngressRuleValue struct {
	Paths []Iok8sapiextensionsv1beta1HTTPIngressPath `json:"paths"` // A collection of paths that map requests to backends.
}

// Iok8sapicorev1HostPathVolumeSource represents the Iok8sapicorev1HostPathVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1HostPathVolumeSource struct {
	Path string `json:"path"` // Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	TypeField string `json:"type,omitempty"` // Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
}

// Iok8sapiauthorizationv1beta1SelfSubjectRulesReviewSpec represents the Iok8sapiauthorizationv1beta1SelfSubjectRulesReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1SelfSubjectRulesReviewSpec struct {
	Namespace string `json:"namespace,omitempty"` // Namespace to evaluate rules for. Required.
}

// Iok8sapiappsv1beta2StatefulSetSpec represents the Iok8sapiappsv1beta2StatefulSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1beta2StatefulSetSpec struct {
	Replicas int `json:"replicas,omitempty"` // replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Servicename string `json:"serviceName"` // serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Updatestrategy Iok8sapiappsv1beta2StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"` // StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
	Volumeclaimtemplates []Iok8sapicorev1PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"` // volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	Podmanagementpolicy string `json:"podManagementPolicy,omitempty"` // podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
}

// Iok8sapimachinerypkgapismetav1APIResource represents the Iok8sapimachinerypkgapismetav1APIResource schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIResource struct {
	Categories []string `json:"categories,omitempty"` // categories is a list of the grouped resources this resource belongs to (e.g. 'all')
	Kind string `json:"kind"` // kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
	Namespaced bool `json:"namespaced"` // namespaced indicates if a resource is namespaced or not.
	Shortnames []string `json:"shortNames,omitempty"` // shortNames is a list of suggested short names of the resource.
	Version string `json:"version,omitempty"` // version is the preferred version of the resource. Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)".
	Verbs []string `json:"verbs"` // verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)
	Group string `json:"group,omitempty"` // group is the preferred group of the resource. Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale".
	Name string `json:"name"` // name is the plural name of the resource.
	Singularname string `json:"singularName"` // singularName is the singular name of the resource. This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.
}

// Iok8skubernetespkgapiv1DownwardAPIProjection represents the Iok8skubernetespkgapiv1DownwardAPIProjection schema from the OpenAPI specification
type Iok8skubernetespkgapiv1DownwardAPIProjection struct {
	Items []Iok8sapicorev1DownwardAPIVolumeFile `json:"items,omitempty"` // Items is a list of DownwardAPIVolume file
}

// Iok8skubernetespkgapisextensionsv1beta1IngressList represents the Iok8skubernetespkgapisextensionsv1beta1IngressList schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IngressList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1Ingress `json:"items"` // Items is the list of Ingress.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiappsv1beta2DaemonSetUpdateStrategy represents the Iok8sapiappsv1beta2DaemonSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1beta2DaemonSetUpdateStrategy struct {
	TypeField string `json:"type,omitempty"` // Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
	Rollingupdate Iok8sapiappsv1beta2RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of daemon set rolling update.
}

// Iok8sapicorev1Taint represents the Iok8sapicorev1Taint schema from the OpenAPI specification
type Iok8sapicorev1Taint struct {
	Value string `json:"value,omitempty"` // Required. The taint value corresponding to the taint key.
	Effect string `json:"effect"` // Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Key string `json:"key"` // Required. The taint key to be applied to a node.
	Timeadded string `json:"timeAdded,omitempty"`
}

// Iok8sapimachinerypkgapismetav1APIVersions represents the Iok8sapimachinerypkgapismetav1APIVersions schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIVersions struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Serveraddressbyclientcidrs []Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs"` // a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.
	Versions []string `json:"versions"` // versions are the api versions that are available.
}

// Iok8sapiautoscalingv2beta1ExternalMetricSource represents the Iok8sapiautoscalingv2beta1ExternalMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1ExternalMetricSource struct {
	Metricselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"metricSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Targetaveragevalue string `json:"targetAverageValue,omitempty"`
	Targetvalue string `json:"targetValue,omitempty"`
	Metricname string `json:"metricName"` // metricName is the name of the metric in question.
}

// Iok8skubernetespkgapisauthenticationv1beta1UserInfo represents the Iok8skubernetespkgapisauthenticationv1beta1UserInfo schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1beta1UserInfo struct {
	Uid string `json:"uid,omitempty"` // A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	Username string `json:"username,omitempty"` // The name that uniquely identifies this user among all active users.
	Extra map[string]interface{} `json:"extra,omitempty"` // Any additional information provided by the authenticator.
	Groups []string `json:"groups,omitempty"` // The names of groups this user is a part of.
}

// Iok8sapiappsv1StatefulSetCondition represents the Iok8sapiappsv1StatefulSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetCondition struct {
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of statefulset condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
}

// Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetList represents the Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetList schema from the OpenAPI specification
type Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetList struct {
	Items []Iok8sapipolicyv1beta1PodDisruptionBudget `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapibatchv2alpha1CronJobList represents the Iok8sapibatchv2alpha1CronJobList schema from the OpenAPI specification
type Iok8sapibatchv2alpha1CronJobList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapibatchv2alpha1CronJob `json:"items"` // items is the list of CronJobs.
}

// Iok8sapiextensionsv1beta1IngressBackend represents the Iok8sapiextensionsv1beta1IngressBackend schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1IngressBackend struct {
	Servicename string `json:"serviceName"` // Specifies the name of the referenced service.
	Serviceport string `json:"servicePort"`
}

// Iok8sapipolicyv1beta1RunAsUserStrategyOptions represents the Iok8sapipolicyv1beta1RunAsUserStrategyOptions schema from the OpenAPI specification
type Iok8sapipolicyv1beta1RunAsUserStrategyOptions struct {
	Ranges []Iok8sapipolicyv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of uids that may be used.
	Rule string `json:"rule"` // Rule is the strategy that will dictate the allowable RunAsUser values that may be set.
}

// Iok8sapicorev1HostAlias represents the Iok8sapicorev1HostAlias schema from the OpenAPI specification
type Iok8sapicorev1HostAlias struct {
	Ip string `json:"ip,omitempty"` // IP address of the host file entry.
	Hostnames []string `json:"hostnames,omitempty"` // Hostnames for the above IP address.
}

// Iok8sapicorev1ResourceQuotaList represents the Iok8sapicorev1ResourceQuotaList schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuotaList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ResourceQuota `json:"items"` // Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
}

// Iok8skubernetespkgapiv1LimitRangeSpec represents the Iok8skubernetespkgapiv1LimitRangeSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LimitRangeSpec struct {
	Limits []Iok8sapicorev1LimitRangeItem `json:"limits"` // Limits is the list of LimitRangeItem objects that are enforced.
}

// Iok8sapiappsv1DaemonSetSpec represents the Iok8sapiappsv1DaemonSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetSpec struct {
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Updatestrategy Iok8sapiappsv1DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"` // DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscalerList represents the Iok8sapiautoscalingv1HorizontalPodAutoscalerList schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscalerList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiautoscalingv1HorizontalPodAutoscaler `json:"items"` // list of horizontal pod autoscaler objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1PodTemplateSpec represents the Iok8skubernetespkgapiv1PodTemplateSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodTemplateSpec struct {
	Spec Iok8sapicorev1PodSpec `json:"spec,omitempty"` // PodSpec is a description of a pod.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiextensionsv1beta1RollbackConfig represents the Iok8sapiextensionsv1beta1RollbackConfig schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1RollbackConfig struct {
	Revision int64 `json:"revision,omitempty"` // The revision to rollback to. If set to 0, rollback to the last revision.
}

// Iok8sapibatchv1JobStatus represents the Iok8sapibatchv1JobStatus schema from the OpenAPI specification
type Iok8sapibatchv1JobStatus struct {
	Succeeded int `json:"succeeded,omitempty"` // The number of pods which reached phase Succeeded.
	Active int `json:"active,omitempty"` // The number of actively running pods.
	Completiontime string `json:"completionTime,omitempty"`
	Conditions []Iok8sapibatchv1JobCondition `json:"conditions,omitempty"` // The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Failed int `json:"failed,omitempty"` // The number of pods which reached phase Failed.
	Starttime string `json:"startTime,omitempty"`
}

// Iok8sapiauthorizationv1SubjectAccessReview represents the Iok8sapiauthorizationv1SubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectAccessReview struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReview represents the Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReview struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapiv1ComponentCondition represents the Iok8skubernetespkgapiv1ComponentCondition schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ComponentCondition struct {
	ErrorField string `json:"error,omitempty"` // Condition error code for a component. For example, a health check error code.
	Message string `json:"message,omitempty"` // Message about the condition for a component. For example, information about a health check.
	Status string `json:"status"` // Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
	TypeField string `json:"type"` // Type of condition for a component. Valid value: "Healthy"
}

// Iok8skubernetespkgapisextensionsv1beta1DaemonSet represents the Iok8skubernetespkgapisextensionsv1beta1DaemonSet schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DaemonSet struct {
	Spec Iok8sapiextensionsv1beta1DaemonSetSpec `json:"spec,omitempty"` // DaemonSetSpec is the specification of a daemon set.
	Status Iok8sapiextensionsv1beta1DaemonSetStatus `json:"status,omitempty"` // DaemonSetStatus represents the current status of a daemon set.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ScaleIOVolumeSource represents the Iok8sapicorev1ScaleIOVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ScaleIOVolumeSource struct {
	Storagepool string `json:"storagePool,omitempty"` // The ScaleIO Storage Pool associated with the protection domain.
	Protectiondomain string `json:"protectionDomain,omitempty"` // The name of the ScaleIO Protection Domain for the configured storage.
	Sslenabled bool `json:"sslEnabled,omitempty"` // Flag to enable/disable SSL communication with Gateway, default false
	Volumename string `json:"volumeName,omitempty"` // The name of a volume already created in the ScaleIO system that is associated with this volume source.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Gateway string `json:"gateway"` // The host address of the ScaleIO API Gateway.
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	System string `json:"system"` // The name of the storage system as configured in ScaleIO.
	Storagemode string `json:"storageMode,omitempty"` // Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
}

// Iok8sapicorev1NodeSelectorTerm represents the Iok8sapicorev1NodeSelectorTerm schema from the OpenAPI specification
type Iok8sapicorev1NodeSelectorTerm struct {
	Matchexpressions []Iok8sapicorev1NodeSelectorRequirement `json:"matchExpressions"` // Required. A list of node selector requirements. The requirements are ANDed.
}

// Iok8sapicorev1LimitRange represents the Iok8sapicorev1LimitRange schema from the OpenAPI specification
type Iok8sapicorev1LimitRange struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1LimitRangeSpec `json:"spec,omitempty"` // LimitRangeSpec defines a min/max usage limit for resources that match on kind.
}

// Iok8sapiappsv1ControllerRevision represents the Iok8sapiappsv1ControllerRevision schema from the OpenAPI specification
type Iok8sapiappsv1ControllerRevision struct {
	Revision int64 `json:"revision"` // Revision indicates the revision of the state represented by Data.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.Object `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // External package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.RawExtension `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // On the wire, the JSON will look something like this: { 	"kind":"MyAPIObject", 	"apiVersion":"v1", 	"myPlugin": { 		"kind":"PluginA", 		"aOption":"foo", 	}, } So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiextensionsv1beta1SupplementalGroupsStrategyOptions represents the Iok8sapiextensionsv1beta1SupplementalGroupsStrategyOptions schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1SupplementalGroupsStrategyOptions struct {
	Ranges []Iok8sapiextensionsv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of supplemental groups. If you would like to force a single supplemental group then supply a single range with the same start and end.
	Rule string `json:"rule,omitempty"` // Rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
}

// Iok8skubernetespkgapisauthenticationv1beta1TokenReviewSpec represents the Iok8skubernetespkgapisauthenticationv1beta1TokenReviewSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1beta1TokenReviewSpec struct {
	Token string `json:"token,omitempty"` // Token is the opaque bearer token.
}

// Iok8sapiautoscalingv2beta1ResourceMetricStatus represents the Iok8sapiautoscalingv2beta1ResourceMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1ResourceMetricStatus struct {
	Currentaverageutilization int `json:"currentAverageUtilization,omitempty"` // currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. It will only be present if `targetAverageValue` was set in the corresponding metric specification.
	Currentaveragevalue string `json:"currentAverageValue"`
	Name string `json:"name"` // name is the name of the resource in question.
}

// Iok8skubernetespkgapiv1ResourceRequirements represents the Iok8skubernetespkgapiv1ResourceRequirements schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ResourceRequirements struct {
	Requests map[string]interface{} `json:"requests,omitempty"` // Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Limits map[string]interface{} `json:"limits,omitempty"` // Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
}

// Iok8skubernetespkgapisnetworkingv1NetworkPolicyPeer represents the Iok8skubernetespkgapisnetworkingv1NetworkPolicyPeer schema from the OpenAPI specification
type Iok8skubernetespkgapisnetworkingv1NetworkPolicyPeer struct {
	Ipblock Iok8sapinetworkingv1IPBlock `json:"ipBlock,omitempty"` // IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiappsv1beta1RollingUpdateDeployment represents the Iok8sapiappsv1beta1RollingUpdateDeployment schema from the OpenAPI specification
type Iok8sapiappsv1beta1RollingUpdateDeployment struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
	Maxsurge string `json:"maxSurge,omitempty"`
}

// Iok8skubernetespkgapiv1NodeSelector represents the Iok8skubernetespkgapiv1NodeSelector schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeSelector struct {
	Nodeselectorterms []Iok8sapicorev1NodeSelectorTerm `json:"nodeSelectorTerms"` // Required. A list of node selector terms. The terms are ORed.
}

// Iok8sapicorev1LimitRangeList represents the Iok8sapicorev1LimitRangeList schema from the OpenAPI specification
type Iok8sapicorev1LimitRangeList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1LimitRange `json:"items"` // Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1AzureDiskVolumeSource represents the Iok8sapicorev1AzureDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AzureDiskVolumeSource struct {
	Cachingmode string `json:"cachingMode,omitempty"` // Host Caching mode: None, Read Only, Read Write.
	Diskname string `json:"diskName"` // The Name of the data disk in the blob storage
	Diskuri string `json:"diskURI"` // The URI the data disk in the blob storage
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Kind string `json:"kind,omitempty"` // Expected values Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8skubernetespkgapiv1ReplicationControllerCondition represents the Iok8skubernetespkgapiv1ReplicationControllerCondition schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ReplicationControllerCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of replication controller condition.
}

// Iok8sapicorev1DownwardAPIVolumeSource represents the Iok8sapicorev1DownwardAPIVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1DownwardAPIVolumeSource struct {
	Items []Iok8sapicorev1DownwardAPIVolumeFile `json:"items,omitempty"` // Items is a list of downward API volume file
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8skubernetespkgapisrbacv1beta1ClusterRoleBindingList represents the Iok8skubernetespkgapisrbacv1beta1ClusterRoleBindingList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1ClusterRoleBindingList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1beta1ClusterRoleBinding `json:"items"` // Items is a list of ClusterRoleBindings
}

// Iok8skubernetespkgapisextensionsv1beta1ReplicaSetStatus represents the Iok8skubernetespkgapisextensionsv1beta1ReplicaSetStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ReplicaSetStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replica set.
	Replicas int `json:"replicas"` // Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replica set.
	Conditions []Iok8sapiextensionsv1beta1ReplicaSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replica set's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replicaset.
}

// Iok8sapirbacv1ClusterRoleBindingList represents the Iok8sapirbacv1ClusterRoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRoleBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1ClusterRoleBinding `json:"items"` // Items is a list of ClusterRoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudget represents the Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudget schema from the OpenAPI specification
type Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudget struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapipolicyv1beta1PodDisruptionBudgetSpec `json:"spec,omitempty"` // PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
	Status Iok8sapipolicyv1beta1PodDisruptionBudgetStatus `json:"status,omitempty"` // PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscaler represents the Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscaler schema from the OpenAPI specification
type Iok8skubernetespkgapisautoscalingv1HorizontalPodAutoscaler struct {
	Spec Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec `json:"spec,omitempty"` // specification of a horizontal pod autoscaler.
	Status Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus `json:"status,omitempty"` // current status of a horizontal pod autoscaler
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // Name is the name of the service
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the service
}

// Iok8sapicorev1EnvVar represents the Iok8sapicorev1EnvVar schema from the OpenAPI specification
type Iok8sapicorev1EnvVar struct {
	Value string `json:"value,omitempty"` // Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Valuefrom Iok8sapicorev1EnvVarSource `json:"valueFrom,omitempty"` // EnvVarSource represents a source for the value of an EnvVar.
	Name string `json:"name"` // Name of the environment variable. Must be a C_IDENTIFIER.
}

// Iok8skubernetespkgapisextensionsv1beta1DeploymentCondition represents the Iok8skubernetespkgapisextensionsv1beta1DeploymentCondition schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DeploymentCondition struct {
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
}

// Iok8sapiextensionsv1beta1NetworkPolicySpec represents the Iok8sapiextensionsv1beta1NetworkPolicySpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicySpec struct {
	Policytypes []string `json:"policyTypes,omitempty"` // List of rule types that the NetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	Egress []Iok8sapiextensionsv1beta1NetworkPolicyEgressRule `json:"egress,omitempty"` // List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
	Ingress []Iok8sapiextensionsv1beta1NetworkPolicyIngressRule `json:"ingress,omitempty"` // List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default).
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8skubernetespkgapiv1PersistentVolumeSpec represents the Iok8skubernetespkgapiv1PersistentVolumeSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeSpec struct {
	Local Iok8sapicorev1LocalVolumeSource `json:"local,omitempty"` // Local represents directly-attached storage with node affinity
	Mountoptions []string `json:"mountOptions,omitempty"` // A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	Gcepersistentdisk Iok8sapicorev1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"` // Represents a Persistent Disk resource in Google Compute Engine. A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	Hostpath Iok8sapicorev1HostPathVolumeSource `json:"hostPath,omitempty"` // Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	Csi Iok8sapicorev1CSIPersistentVolumeSource `json:"csi,omitempty"` // Represents storage that is managed by an external CSI volume driver (Beta feature)
	Scaleio Iok8sapicorev1ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty"` // ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
	Capacity map[string]interface{} `json:"capacity,omitempty"` // A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Nfs Iok8sapicorev1NFSVolumeSource `json:"nfs,omitempty"` // Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Glusterfs Iok8sapicorev1GlusterfsVolumeSource `json:"glusterfs,omitempty"` // Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Accessmodes []string `json:"accessModes,omitempty"` // AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	Nodeaffinity Iok8sapicorev1VolumeNodeAffinity `json:"nodeAffinity,omitempty"` // VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
	Claimref Iok8sapicorev1ObjectReference `json:"claimRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Azuredisk Iok8sapicorev1AzureDiskVolumeSource `json:"azureDisk,omitempty"` // AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	Flexvolume Iok8sapicorev1FlexPersistentVolumeSource `json:"flexVolume,omitempty"` // FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
	Fc Iok8sapicorev1FCVolumeSource `json:"fc,omitempty"` // Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Flocker Iok8sapicorev1FlockerVolumeSource `json:"flocker,omitempty"` // Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Cephfs Iok8sapicorev1CephFSPersistentVolumeSource `json:"cephfs,omitempty"` // Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Awselasticblockstore Iok8sapicorev1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"` // Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	Persistentvolumereclaimpolicy string `json:"persistentVolumeReclaimPolicy,omitempty"` // What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	Rbd Iok8sapicorev1RBDPersistentVolumeSource `json:"rbd,omitempty"` // Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Volumemode string `json:"volumeMode,omitempty"` // volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec. This is an alpha feature and may change in the future.
	Portworxvolume Iok8sapicorev1PortworxVolumeSource `json:"portworxVolume,omitempty"` // PortworxVolumeSource represents a Portworx volume resource.
	Storageclassname string `json:"storageClassName,omitempty"` // Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	Quobyte Iok8sapicorev1QuobyteVolumeSource `json:"quobyte,omitempty"` // Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Iscsi Iok8sapicorev1ISCSIPersistentVolumeSource `json:"iscsi,omitempty"` // ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Photonpersistentdisk Iok8sapicorev1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"` // Represents a Photon Controller persistent disk resource.
	Cinder Iok8sapicorev1CinderVolumeSource `json:"cinder,omitempty"` // Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Vspherevolume Iok8sapicorev1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"` // Represents a vSphere volume resource.
	Azurefile Iok8sapicorev1AzureFilePersistentVolumeSource `json:"azureFile,omitempty"` // AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	Storageos Iok8sapicorev1StorageOSPersistentVolumeSource `json:"storageos,omitempty"` // Represents a StorageOS persistent volume resource.
}

// Iok8sapiautoscalingv1ScaleStatus represents the Iok8sapiautoscalingv1ScaleStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv1ScaleStatus struct {
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
	Selector string `json:"selector,omitempty"` // label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors
}

// Iok8sapiauthorizationv1ResourceAttributes represents the Iok8sapiauthorizationv1ResourceAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1ResourceAttributes struct {
	Subresource string `json:"subresource,omitempty"` // Subresource is one of the existing resource types. "" means none.
	Verb string `json:"verb,omitempty"` // Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Version string `json:"version,omitempty"` // Version is the API Version of the Resource. "*" means all.
	Group string `json:"group,omitempty"` // Group is the API Group of the Resource. "*" means all.
	Name string `json:"name,omitempty"` // Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the action being requested. Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Resource string `json:"resource,omitempty"` // Resource is one of the existing resource types. "*" means all.
}

// Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicySpec represents the Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicySpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1PodSecurityPolicySpec struct {
	Runasuser Iok8sapiextensionsv1beta1RunAsUserStrategyOptions `json:"runAsUser"` // Run A sUser Strategy Options defines the strategy type and any options used to create the strategy.
	Requireddropcapabilities []string `json:"requiredDropCapabilities,omitempty"` // RequiredDropCapabilities are the capabilities that will be dropped from the container. These are required to be dropped and cannot be added.
	Allowedhostpaths []Iok8sapiextensionsv1beta1AllowedHostPath `json:"allowedHostPaths,omitempty"` // is a white list of allowed host paths. Empty indicates that all host paths may be used.
	Defaultaddcapabilities []string `json:"defaultAddCapabilities,omitempty"` // DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability. You may not list a capability in both DefaultAddCapabilities and RequiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the AllowedCapabilities list.
	Hostipc bool `json:"hostIPC,omitempty"` // hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	Supplementalgroups Iok8sapiextensionsv1beta1SupplementalGroupsStrategyOptions `json:"supplementalGroups"` // SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
	Defaultallowprivilegeescalation bool `json:"defaultAllowPrivilegeEscalation,omitempty"` // DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	Volumes []string `json:"volumes,omitempty"` // volumes is a white list of allowed volume plugins. Empty indicates that all plugins may be used.
	Hostports []Iok8sapiextensionsv1beta1HostPortRange `json:"hostPorts,omitempty"` // hostPorts determines which host port ranges are allowed to be exposed.
	Selinux Iok8sapiextensionsv1beta1SELinuxStrategyOptions `json:"seLinux"` // SELinux Strategy Options defines the strategy type and any options used to create the strategy.
	Allowprivilegeescalation bool `json:"allowPrivilegeEscalation,omitempty"` // AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	Fsgroup Iok8sapiextensionsv1beta1FSGroupStrategyOptions `json:"fsGroup"` // FSGroupStrategyOptions defines the strategy type and options used to create the strategy.
	Hostnetwork bool `json:"hostNetwork,omitempty"` // hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	Hostpid bool `json:"hostPID,omitempty"` // hostPID determines if the policy allows the use of HostPID in the pod spec.
	Allowedflexvolumes []Iok8sapiextensionsv1beta1AllowedFlexVolume `json:"allowedFlexVolumes,omitempty"` // AllowedFlexVolumes is a whitelist of allowed Flexvolumes. Empty or nil indicates that all Flexvolumes may be used. This parameter is effective only when the usage of the Flexvolumes is allowed in the "Volumes" field.
	Allowedcapabilities []string `json:"allowedCapabilities,omitempty"` // AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities.
	Privileged bool `json:"privileged,omitempty"` // privileged determines if a pod can request to be run as privileged.
	Readonlyrootfilesystem bool `json:"readOnlyRootFilesystem,omitempty"` // ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system. If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
}

// Iok8skubernetespkgapiv1Namespace represents the Iok8skubernetespkgapiv1Namespace schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Namespace struct {
	Spec Iok8sapicorev1NamespaceSpec `json:"spec,omitempty"` // NamespaceSpec describes the attributes on a Namespace.
	Status Iok8sapicorev1NamespaceStatus `json:"status,omitempty"` // NamespaceStatus is information about the current status of a Namespace.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscaler represents the Iok8sapiautoscalingv1HorizontalPodAutoscaler schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscaler struct {
	Status Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus `json:"status,omitempty"` // current status of a horizontal pod autoscaler
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec `json:"spec,omitempty"` // specification of a horizontal pod autoscaler.
}

// Iok8sapiextensionsv1beta1NetworkPolicyList represents the Iok8sapiextensionsv1beta1NetworkPolicyList schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1NetworkPolicyList struct {
	Items []Iok8sapiextensionsv1beta1NetworkPolicy `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapimachinerypkgapismetav1APIGroup represents the Iok8sapimachinerypkgapismetav1APIGroup schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIGroup struct {
	Versions []Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery `json:"versions"` // versions are the versions supported in this group.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name"` // name is the name of the group.
	Preferredversion Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery `json:"preferredVersion,omitempty"` // GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.
	Serveraddressbyclientcidrs []Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs"` // a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.
}

// Iok8sapicorev1NodeSelectorRequirement represents the Iok8sapicorev1NodeSelectorRequirement schema from the OpenAPI specification
type Iok8sapicorev1NodeSelectorRequirement struct {
	Key string `json:"key"` // The label key that the selector applies to.
	Operator string `json:"operator"` // Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Values []string `json:"values,omitempty"` // An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
}

// Iok8sapiappsv1beta1ControllerRevisionList represents the Iok8sapiappsv1beta1ControllerRevisionList schema from the OpenAPI specification
type Iok8sapiappsv1beta1ControllerRevisionList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1beta1ControllerRevision `json:"items"` // Items is the list of ControllerRevisions
}

// Iok8sapicorev1DownwardAPIVolumeFile represents the Iok8sapicorev1DownwardAPIVolumeFile schema from the OpenAPI specification
type Iok8sapicorev1DownwardAPIVolumeFile struct {
	Mode int `json:"mode,omitempty"` // Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Path string `json:"path"` // Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Resourcefieldref Iok8sapicorev1ResourceFieldSelector `json:"resourceFieldRef,omitempty"` // ResourceFieldSelector represents container resources (cpu, memory) and their output format
	Fieldref Iok8sapicorev1ObjectFieldSelector `json:"fieldRef,omitempty"` // ObjectFieldSelector selects an APIVersioned field of an object.
}

// Iok8skubernetespkgapisextensionsv1beta1ReplicaSet represents the Iok8skubernetespkgapisextensionsv1beta1ReplicaSet schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ReplicaSet struct {
	Status Iok8sapiextensionsv1beta1ReplicaSetStatus `json:"status,omitempty"` // ReplicaSetStatus represents the current status of a ReplicaSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiextensionsv1beta1ReplicaSetSpec `json:"spec,omitempty"` // ReplicaSetSpec is the specification of a ReplicaSet.
}

// Iok8sapiappsv1beta2DeploymentStatus represents the Iok8sapiappsv1beta2DeploymentStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta2DeploymentStatus struct {
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Conditions []Iok8sapiappsv1beta2DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of ready pods targeted by this deployment.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminated pods targeted by this deployment that have the desired template spec.
}

// Iok8skubernetespkgapisrbacv1beta1Role represents the Iok8skubernetespkgapisrbacv1beta1Role schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1Role struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1beta1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this Role
}

// Iok8sapiappsv1ReplicaSetList represents the Iok8sapiappsv1ReplicaSetList schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiappsv1ReplicaSet `json:"items"` // List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
}

// Iok8skubernetespkgapiv1NodeAffinity represents the Iok8skubernetespkgapiv1NodeAffinity schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution Iok8sapicorev1NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8sapiextensionsv1beta1RollingUpdateDeployment represents the Iok8sapiextensionsv1beta1RollingUpdateDeployment schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1RollingUpdateDeployment struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"`
	Maxsurge string `json:"maxSurge,omitempty"`
}

// Iok8sapicorev1PodAntiAffinity represents the Iok8sapicorev1PodAntiAffinity schema from the OpenAPI specification
type Iok8sapicorev1PodAntiAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution []Iok8sapicorev1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
}

// Iok8sapicorev1SecurityContext represents the Iok8sapicorev1SecurityContext schema from the OpenAPI specification
type Iok8sapicorev1SecurityContext struct {
	Runasgroup int64 `json:"runAsGroup,omitempty"` // The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Runasnonroot bool `json:"runAsNonRoot,omitempty"` // Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Runasuser int64 `json:"runAsUser,omitempty"` // The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
	Allowprivilegeescalation bool `json:"allowPrivilegeEscalation,omitempty"` // AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	Capabilities Iok8sapicorev1Capabilities `json:"capabilities,omitempty"` // Adds and removes POSIX capabilities from running containers.
	Privileged bool `json:"privileged,omitempty"` // Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
	Readonlyrootfilesystem bool `json:"readOnlyRootFilesystem,omitempty"` // Whether this container has a read-only root filesystem. Default is false.
}

// Iok8sapicorev1NodeSpec represents the Iok8sapicorev1NodeSpec schema from the OpenAPI specification
type Iok8sapicorev1NodeSpec struct {
	Externalid string `json:"externalID,omitempty"` // External ID of the node assigned by some machine database (e.g. a cloud provider). Deprecated.
	Podcidr string `json:"podCIDR,omitempty"` // PodCIDR represents the pod IP range assigned to the node.
	Providerid string `json:"providerID,omitempty"` // ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
	Taints []Iok8sapicorev1Taint `json:"taints,omitempty"` // If specified, the node's taints.
	Unschedulable bool `json:"unschedulable,omitempty"` // Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	Configsource Iok8sapicorev1NodeConfigSource `json:"configSource,omitempty"` // NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil.
}

// Iok8sapirbacv1RoleBindingList represents the Iok8sapirbacv1RoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1RoleBindingList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1RoleBinding `json:"items"` // Items is a list of RoleBindings
}

// Iok8sapicorev1KeyToPath represents the Iok8sapicorev1KeyToPath schema from the OpenAPI specification
type Iok8sapicorev1KeyToPath struct {
	Mode int `json:"mode,omitempty"` // Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Path string `json:"path"` // The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	Key string `json:"key"` // The key to project.
}

// Iok8skubernetespkgapisrbacv1alpha1RoleBindingList represents the Iok8skubernetespkgapisrbacv1alpha1RoleBindingList schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1RoleBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1RoleBinding `json:"items"` // Items is a list of RoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisrbacv1alpha1ClusterRole represents the Iok8skubernetespkgapisrbacv1alpha1ClusterRole schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1alpha1ClusterRole struct {
	Rules []Iok8sapirbacv1alpha1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this ClusterRole
	Aggregationrule Iok8sapirbacv1alpha1AggregationRule `json:"aggregationRule,omitempty"` // AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapiv1PodSpec represents the Iok8skubernetespkgapiv1PodSpec schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodSpec struct {
	Containers []Iok8sapicorev1Container `json:"containers"` // List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.
	Terminationgraceperiodseconds int64 `json:"terminationGracePeriodSeconds,omitempty"` // Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
	Volumes []Iok8sapicorev1Volume `json:"volumes,omitempty"` // List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes
	Schedulername string `json:"schedulerName,omitempty"` // If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
	Restartpolicy string `json:"restartPolicy,omitempty"` // Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	Serviceaccountname string `json:"serviceAccountName,omitempty"` // ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Subdomain string `json:"subdomain,omitempty"` // If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Tolerations []Iok8sapicorev1Toleration `json:"tolerations,omitempty"` // If specified, the pod's tolerations.
	Hostpid bool `json:"hostPID,omitempty"` // Use the host's pid namespace. Optional: Default to false.
	Hostname string `json:"hostname,omitempty"` // Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
	Imagepullsecrets []Iok8sapicorev1LocalObjectReference `json:"imagePullSecrets,omitempty"` // ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	Dnspolicy string `json:"dnsPolicy,omitempty"` // Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	Securitycontext Iok8sapicorev1PodSecurityContext `json:"securityContext,omitempty"` // PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext. Field values of container.securityContext take precedence over field values of PodSecurityContext.
	Priorityclassname string `json:"priorityClassName,omitempty"` // If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	Shareprocessnamespace bool `json:"shareProcessNamespace,omitempty"` // Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. This field is alpha-level and is honored only by servers that enable the PodShareProcessNamespace feature.
	Automountserviceaccounttoken bool `json:"automountServiceAccountToken,omitempty"` // AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	Initcontainers []Iok8sapicorev1Container `json:"initContainers,omitempty"` // List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	Nodeselector map[string]interface{} `json:"nodeSelector,omitempty"` // NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	Hostipc bool `json:"hostIPC,omitempty"` // Use the host's ipc namespace. Optional: Default to false.
	Hostnetwork bool `json:"hostNetwork,omitempty"` // Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.
	Nodename string `json:"nodeName,omitempty"` // NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	Priority int `json:"priority,omitempty"` // The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
	Activedeadlineseconds int64 `json:"activeDeadlineSeconds,omitempty"` // Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
	Hostaliases []Iok8sapicorev1HostAlias `json:"hostAliases,omitempty"` // HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.
	Serviceaccount string `json:"serviceAccount,omitempty"` // DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
	Dnsconfig Iok8sapicorev1PodDNSConfig `json:"dnsConfig,omitempty"` // PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
	Affinity Iok8sapicorev1Affinity `json:"affinity,omitempty"` // Affinity is a group of affinity scheduling rules.
}

// Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfigurationList represents the Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfigurationList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfigurationList struct {
	Items []Iok8sapiadmissionregistrationv1beta1ValidatingWebhookConfiguration `json:"items"` // List of ValidatingWebhookConfiguration.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiautoscalingv2beta1PodsMetricStatus represents the Iok8sapiautoscalingv2beta1PodsMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1PodsMetricStatus struct {
	Currentaveragevalue string `json:"currentAverageValue"`
	Metricname string `json:"metricName"` // metricName is the name of the metric in question
}

// Iok8skubernetespkgapisappsv1beta1DeploymentRollback represents the Iok8skubernetespkgapisappsv1beta1DeploymentRollback schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1DeploymentRollback struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name"` // Required: This must match the Name of a deployment.
	Rollbackto Iok8sapiappsv1beta1RollbackConfig `json:"rollbackTo"` // DEPRECATED.
	Updatedannotations map[string]interface{} `json:"updatedAnnotations,omitempty"` // The annotations to be updated to a deployment
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapicorev1ResourceQuotaStatus represents the Iok8sapicorev1ResourceQuotaStatus schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuotaStatus struct {
	Hard map[string]interface{} `json:"hard,omitempty"` // Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Used map[string]interface{} `json:"used,omitempty"` // Used is the current observed total usage of the resource in the namespace.
}

// Iok8sapicorev1SELinuxOptions represents the Iok8sapicorev1SELinuxOptions schema from the OpenAPI specification
type Iok8sapicorev1SELinuxOptions struct {
	Level string `json:"level,omitempty"` // Level is SELinux level label that applies to the container.
	Role string `json:"role,omitempty"` // Role is a SELinux role label that applies to the container.
	TypeField string `json:"type,omitempty"` // Type is a SELinux type label that applies to the container.
	User string `json:"user,omitempty"` // User is a SELinux user label that applies to the container.
}

// Iok8skubernetespkgapiv1PersistentVolumeStatus represents the Iok8skubernetespkgapiv1PersistentVolumeStatus schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeStatus struct {
	Message string `json:"message,omitempty"` // A human-readable message indicating details about why the volume is in this state.
	Phase string `json:"phase,omitempty"` // Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
	Reason string `json:"reason,omitempty"` // Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.
}

// Iok8skubernetespkgapisnetworkingv1NetworkPolicyPort represents the Iok8skubernetespkgapisnetworkingv1NetworkPolicyPort schema from the OpenAPI specification
type Iok8skubernetespkgapisnetworkingv1NetworkPolicyPort struct {
	Port string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"` // The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.
}

// Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec represents the Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec struct {
	Nonresourceattributes Iok8sapiauthorizationv1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
}

// Iok8skubernetespkgapiv1SecretVolumeSource represents the Iok8skubernetespkgapiv1SecretVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SecretVolumeSource struct {
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or it's keys must be defined
	Secretname string `json:"secretName,omitempty"` // Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8skubernetespkgapiv1Container represents the Iok8skubernetespkgapiv1Container schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Container struct {
	Volumemounts []Iok8sapicorev1VolumeMount `json:"volumeMounts,omitempty"` // Pod volumes to mount into the container's filesystem. Cannot be updated.
	Securitycontext Iok8sapicorev1SecurityContext `json:"securityContext,omitempty"` // SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext. When both are set, the values in SecurityContext take precedence.
	Image string `json:"image,omitempty"` // Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
	Stdin bool `json:"stdin,omitempty"` // Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Env []Iok8sapicorev1EnvVar `json:"env,omitempty"` // List of environment variables to set in the container. Cannot be updated.
	Command []string `json:"command,omitempty"` // Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Readinessprobe Iok8sapicorev1Probe `json:"readinessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Stdinonce bool `json:"stdinOnce,omitempty"` // Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	Envfrom []Iok8sapicorev1EnvFromSource `json:"envFrom,omitempty"` // List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	Workingdir string `json:"workingDir,omitempty"` // Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	Volumedevices []Iok8sapicorev1VolumeDevice `json:"volumeDevices,omitempty"` // volumeDevices is the list of block devices to be used by the container. This is an alpha feature and may change in the future.
	Lifecycle Iok8sapicorev1Lifecycle `json:"lifecycle,omitempty"` // Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Imagepullpolicy string `json:"imagePullPolicy,omitempty"` // Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	Name string `json:"name"` // Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	Terminationmessagepolicy string `json:"terminationMessagePolicy,omitempty"` // Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	Ports []Iok8sapicorev1ContainerPort `json:"ports,omitempty"` // List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.
	Args []string `json:"args,omitempty"` // Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Terminationmessagepath string `json:"terminationMessagePath,omitempty"` // Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	Tty bool `json:"tty,omitempty"` // Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Livenessprobe Iok8sapicorev1Probe `json:"livenessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
}

// Iok8sapicorev1ComponentCondition represents the Iok8sapicorev1ComponentCondition schema from the OpenAPI specification
type Iok8sapicorev1ComponentCondition struct {
	ErrorField string `json:"error,omitempty"` // Condition error code for a component. For example, a health check error code.
	Message string `json:"message,omitempty"` // Message about the condition for a component. For example, information about a health check.
	Status string `json:"status"` // Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
	TypeField string `json:"type"` // Type of condition for a component. Valid value: "Healthy"
}

// Iok8skubernetespkgapiv1HTTPGetAction represents the Iok8skubernetespkgapiv1HTTPGetAction schema from the OpenAPI specification
type Iok8skubernetespkgapiv1HTTPGetAction struct {
	Port string `json:"port"`
	Scheme string `json:"scheme,omitempty"` // Scheme to use for connecting to the host. Defaults to HTTP.
	Host string `json:"host,omitempty"` // Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	Httpheaders []Iok8sapicorev1HTTPHeader `json:"httpHeaders,omitempty"` // Custom headers to set in the request. HTTP allows repeated headers.
	Path string `json:"path,omitempty"` // Path to access on the HTTP server.
}

// Iok8sapicorev1DownwardAPIProjection represents the Iok8sapicorev1DownwardAPIProjection schema from the OpenAPI specification
type Iok8sapicorev1DownwardAPIProjection struct {
	Items []Iok8sapicorev1DownwardAPIVolumeFile `json:"items,omitempty"` // Items is a list of DownwardAPIVolume file
}

// Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec represents the Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec struct {
	Namespace string `json:"namespace,omitempty"` // Namespace to evaluate rules for. Required.
}

// Iok8sapicorev1CinderVolumeSource represents the Iok8sapicorev1CinderVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CinderVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	Volumeid string `json:"volumeID"` // volume id used to identify the volume in cinder More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
}

// Iok8skubernetespkgapiv1NodeDaemonEndpoints represents the Iok8skubernetespkgapiv1NodeDaemonEndpoints schema from the OpenAPI specification
type Iok8skubernetespkgapiv1NodeDaemonEndpoints struct {
	Kubeletendpoint Iok8sapicorev1DaemonEndpoint `json:"kubeletEndpoint,omitempty"` // DaemonEndpoint contains information about a single Daemon endpoint.
}

// Iok8sapiextensionsv1beta1PodSecurityPolicyList represents the Iok8sapiextensionsv1beta1PodSecurityPolicyList schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1PodSecurityPolicyList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1PodSecurityPolicy `json:"items"` // Items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiauthorizationv1LocalSubjectAccessReview represents the Iok8sapiauthorizationv1LocalSubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1LocalSubjectAccessReview struct {
	Spec Iok8sapiauthorizationv1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ContainerStateWaiting represents the Iok8sapicorev1ContainerStateWaiting schema from the OpenAPI specification
type Iok8sapicorev1ContainerStateWaiting struct {
	Message string `json:"message,omitempty"` // Message regarding why the container is not yet running.
	Reason string `json:"reason,omitempty"` // (brief) reason the container is not yet running.
}

// Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestList represents the Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestList schema from the OpenAPI specification
type Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicertificatesv1beta1CertificateSigningRequest `json:"items"`
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1ContainerPort represents the Iok8sapicorev1ContainerPort schema from the OpenAPI specification
type Iok8sapicorev1ContainerPort struct {
	Hostport int `json:"hostPort,omitempty"` // Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	Name string `json:"name,omitempty"` // If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
	Protocol string `json:"protocol,omitempty"` // Protocol for port. Must be UDP or TCP. Defaults to "TCP".
	Containerport int `json:"containerPort"` // Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
	Hostip string `json:"hostIP,omitempty"` // What host IP to bind the external port to.
}

// Iok8skubernetespkgapisappsv1beta1ControllerRevision represents the Iok8skubernetespkgapisappsv1beta1ControllerRevision schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1ControllerRevision struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Revision int64 `json:"revision"` // Revision indicates the revision of the state represented by Data.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.Object `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // External package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.RawExtension `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // On the wire, the JSON will look something like this: { 	"kind":"MyAPIObject", 	"apiVersion":"v1", 	"myPlugin": { 		"kind":"PluginA", 		"aOption":"foo", 	}, } So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
}

// Iok8sapicorev1ConfigMapList represents the Iok8sapicorev1ConfigMapList schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1ConfigMap `json:"items"` // Items is the list of ConfigMaps.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiextensionsv1beta1ScaleStatus represents the Iok8sapiextensionsv1beta1ScaleStatus schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1ScaleStatus struct {
	Selector map[string]interface{} `json:"selector,omitempty"` // label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Targetselector string `json:"targetSelector,omitempty"` // label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
}

// Iok8sapicorev1PodAffinity represents the Iok8sapicorev1PodAffinity schema from the OpenAPI specification
type Iok8sapicorev1PodAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution []Iok8sapicorev1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
}

// Iok8sapicorev1LocalObjectReference represents the Iok8sapicorev1LocalObjectReference schema from the OpenAPI specification
type Iok8sapicorev1LocalObjectReference struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8skubernetespkgapisextensionsv1beta1ReplicaSetCondition represents the Iok8skubernetespkgapisextensionsv1beta1ReplicaSetCondition schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ReplicaSetCondition struct {
	TypeField string `json:"type"` // Type of replica set condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
}

// Iok8skubernetespkgapiv1ProjectedVolumeSource represents the Iok8skubernetespkgapiv1ProjectedVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ProjectedVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Sources []Iok8sapicorev1VolumeProjection `json:"sources"` // list of volume projections
}

// Iok8sapiappsv1beta1ControllerRevision represents the Iok8sapiappsv1beta1ControllerRevision schema from the OpenAPI specification
type Iok8sapiappsv1beta1ControllerRevision struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Revision int64 `json:"revision"` // Revision indicates the revision of the state represented by Data.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.Object `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // External package: type MyAPIObject struct { 	runtime.TypeMeta `json:",inline"` 	MyPlugin runtime.RawExtension `json:"myPlugin"` } type PluginA struct { 	AOption string `json:"aOption"` } // On the wire, the JSON will look something like this: { 	"kind":"MyAPIObject", 	"apiVersion":"v1", 	"myPlugin": { 		"kind":"PluginA", 		"aOption":"foo", 	}, } So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfiguration represents the Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfiguration schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1beta1MutatingWebhookConfiguration struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Webhooks []Iok8sapiadmissionregistrationv1beta1Webhook `json:"webhooks,omitempty"` // Webhooks is a list of webhooks and the affected resources and operations.
}

// Iok8skubernetespkgapiv1Toleration represents the Iok8skubernetespkgapiv1Toleration schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Toleration struct {
	Effect string `json:"effect,omitempty"` // Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	Key string `json:"key,omitempty"` // Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	Operator string `json:"operator,omitempty"` // Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
	Tolerationseconds int64 `json:"tolerationSeconds,omitempty"` // TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	Value string `json:"value,omitempty"` // Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
}

// Iok8sapirbacv1ClusterRoleList represents the Iok8sapirbacv1ClusterRoleList schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRoleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1ClusterRole `json:"items"` // Items is a list of ClusterRoles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisauthorizationv1beta1NonResourceAttributes represents the Iok8skubernetespkgapisauthorizationv1beta1NonResourceAttributes schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1NonResourceAttributes struct {
	Verb string `json:"verb,omitempty"` // Verb is the standard HTTP verb
	Path string `json:"path,omitempty"` // Path is the URL path of the request
}

// Iok8sapiappsv1DaemonSetCondition represents the Iok8sapiappsv1DaemonSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetCondition struct {
	TypeField string `json:"type"` // Type of DaemonSet condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
}

// Iok8skubernetespkgapisappsv1beta1ScaleStatus represents the Iok8skubernetespkgapisappsv1beta1ScaleStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1ScaleStatus struct {
	Replicas int `json:"replicas"` // actual number of observed instances of the scaled object.
	Selector map[string]interface{} `json:"selector,omitempty"` // label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Targetselector string `json:"targetSelector,omitempty"` // label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version string `json:"version"` // Version is the version this resource belongs in
	Group string `json:"group"` // Group is the group this resource belongs in
	Names GeneratedType `json:"names"` // CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
	Scope string `json:"scope"` // Scope indicates whether this resource is cluster or namespace scoped. Default is namespaced
	Subresources GeneratedType `json:"subresources,omitempty"` // CustomResourceSubresources defines the status and scale subresources for CustomResources.
	Validation GeneratedType `json:"validation,omitempty"` // CustomResourceValidation is a list of validation methods for CustomResources.
}

// Iok8skubernetespkgapiv1AzureDiskVolumeSource represents the Iok8skubernetespkgapiv1AzureDiskVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1AzureDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Kind string `json:"kind,omitempty"` // Expected values Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Cachingmode string `json:"cachingMode,omitempty"` // Host Caching mode: None, Read Only, Read Write.
	Diskname string `json:"diskName"` // The Name of the data disk in the blob storage
	Diskuri string `json:"diskURI"` // The URI the data disk in the blob storage
}

// Iok8sapicorev1PersistentVolumeClaimStatus represents the Iok8sapicorev1PersistentVolumeClaimStatus schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimStatus struct {
	Accessmodes []string `json:"accessModes,omitempty"` // AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Represents the actual resources of the underlying volume.
	Conditions []Iok8sapicorev1PersistentVolumeClaimCondition `json:"conditions,omitempty"` // Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
	Phase string `json:"phase,omitempty"` // Phase represents the current phase of PersistentVolumeClaim.
}

// Iok8sapicorev1ExecAction represents the Iok8sapicorev1ExecAction schema from the OpenAPI specification
type Iok8sapicorev1ExecAction struct {
	Command []string `json:"command,omitempty"` // Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
}

// Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestCondition represents the Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestCondition schema from the OpenAPI specification
type Iok8skubernetespkgapiscertificatesv1beta1CertificateSigningRequestCondition struct {
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
	Message string `json:"message,omitempty"` // human readable message with details about the request state
	Reason string `json:"reason,omitempty"` // brief reason for the request state
	TypeField string `json:"type"` // request approval state, currently Approved or Denied.
}

// Iok8sapiappsv1ReplicaSet represents the Iok8sapiappsv1ReplicaSet schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSet struct {
	Status Iok8sapiappsv1ReplicaSetStatus `json:"status,omitempty"` // ReplicaSetStatus represents the current status of a ReplicaSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1ReplicaSetSpec `json:"spec,omitempty"` // ReplicaSetSpec is the specification of a ReplicaSet.
}

// Iok8sapiautoscalingv1ScaleSpec represents the Iok8sapiautoscalingv1ScaleSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // desired number of instances for the scaled object.
}

// Iok8sapirbacv1alpha1ClusterRoleBindingList represents the Iok8sapirbacv1alpha1ClusterRoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1alpha1ClusterRoleBindingList struct {
	Items []Iok8sapirbacv1alpha1ClusterRoleBinding `json:"items"` // Items is a list of ClusterRoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiappsv1RollingUpdateDeployment represents the Iok8sapiappsv1RollingUpdateDeployment schema from the OpenAPI specification
type Iok8sapiappsv1RollingUpdateDeployment struct {
	Maxsurge string `json:"maxSurge,omitempty"`
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8sapipolicyv1beta1PodDisruptionBudgetSpec represents the Iok8sapipolicyv1beta1PodDisruptionBudgetSpec schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodDisruptionBudgetSpec struct {
	Minavailable string `json:"minAvailable,omitempty"`
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Maxunavailable string `json:"maxUnavailable,omitempty"`
}

// Iok8skubernetespkgapiv1PodList represents the Iok8skubernetespkgapiv1PodList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PodList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Pod `json:"items"` // List of pods. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
}

// Iok8sapicorev1EndpointsList represents the Iok8sapicorev1EndpointsList schema from the OpenAPI specification
type Iok8sapicorev1EndpointsList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Endpoints `json:"items"` // List of endpoints.
}

// Iok8sapicorev1LoadBalancerIngress represents the Iok8sapicorev1LoadBalancerIngress schema from the OpenAPI specification
type Iok8sapicorev1LoadBalancerIngress struct {
	Hostname string `json:"hostname,omitempty"` // Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
	Ip string `json:"ip,omitempty"` // IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
}

// Iok8skubernetespkgapisrbacv1beta1RoleRef represents the Iok8skubernetespkgapisrbacv1beta1RoleRef schema from the OpenAPI specification
type Iok8skubernetespkgapisrbacv1beta1RoleRef struct {
	Apigroup string `json:"apiGroup"` // APIGroup is the group for the resource being referenced
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
}

// Iok8sapicorev1SecretKeySelector represents the Iok8sapicorev1SecretKeySelector schema from the OpenAPI specification
type Iok8sapicorev1SecretKeySelector struct {
	Key string `json:"key"` // The key of the secret to select from. Must be a valid secret key.
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or it's key must be defined
}

// Iok8skubernetespkgapisauthorizationv1beta1LocalSubjectAccessReview represents the Iok8skubernetespkgapisauthorizationv1beta1LocalSubjectAccessReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1LocalSubjectAccessReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1beta1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1beta1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
}

// Iok8skubernetespkgapisappsv1beta1DeploymentSpec represents the Iok8skubernetespkgapisappsv1beta1DeploymentSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1DeploymentSpec struct {
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused.
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Strategy Iok8sapiappsv1beta1DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 2.
	Rollbackto Iok8sapiappsv1beta1RollbackConfig `json:"rollbackTo,omitempty"` // DEPRECATED.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
}

// Iok8sapibatchv1beta1CronJobSpec represents the Iok8sapibatchv1beta1CronJobSpec schema from the OpenAPI specification
type Iok8sapibatchv1beta1CronJobSpec struct {
	Startingdeadlineseconds int64 `json:"startingDeadlineSeconds,omitempty"` // Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.
	Successfuljobshistorylimit int `json:"successfulJobsHistoryLimit,omitempty"` // The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 3.
	Suspend bool `json:"suspend,omitempty"` // This flag tells the controller to suspend subsequent executions, it does not apply to already started executions. Defaults to false.
	Concurrencypolicy string `json:"concurrencyPolicy,omitempty"` // Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	Failedjobshistorylimit int `json:"failedJobsHistoryLimit,omitempty"` // The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Jobtemplate Iok8sapibatchv1beta1JobTemplateSpec `json:"jobTemplate"` // JobTemplateSpec describes the data a Job should have when created from a template
	Schedule string `json:"schedule"` // The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
}

// Iok8skubernetespkgapisextensionsv1beta1IngressRule represents the Iok8skubernetespkgapisextensionsv1beta1IngressRule schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IngressRule struct {
	Host string `json:"host,omitempty"` // Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the 	 IP in the Spec of the parent Ingress. 2. The `:` delimiter is not respected because ports are not allowed. 	 Currently the port of an Ingress is implicitly :80 for http and 	 :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
	Http Iok8sapiextensionsv1beta1HTTPIngressRuleValue `json:"http,omitempty"` // HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
}

// Iok8sapimachinerypkgapismetav1ListMeta represents the Iok8sapimachinerypkgapismetav1ListMeta schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ListMeta struct {
	Resourceversion string `json:"resourceVersion,omitempty"` // String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	Selflink string `json:"selfLink,omitempty"` // selfLink is a URL representing this object. Populated by the system. Read-only.
	ContinueField string `json:"continue,omitempty"` // continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response.
}

// Iok8sapicorev1Capabilities represents the Iok8sapicorev1Capabilities schema from the OpenAPI specification
type Iok8sapicorev1Capabilities struct {
	Add []string `json:"add,omitempty"` // Added capabilities
	Drop []string `json:"drop,omitempty"` // Removed capabilities
}

// Iok8skubernetespkgapisauthorizationv1SubjectAccessReviewSpec represents the Iok8skubernetespkgapisauthorizationv1SubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1SubjectAccessReviewSpec struct {
	Nonresourceattributes Iok8sapiauthorizationv1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
	Uid string `json:"uid,omitempty"` // UID information about the requesting user.
	User string `json:"user,omitempty"` // User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra corresponds to the user.Info.GetExtra() method from the authenticator. Since that is input to the authorizer it needs a reflection here.
	Groups []string `json:"groups,omitempty"` // Groups is the groups you're testing for.
}

// Iok8skubernetespkgapiv1EndpointsList represents the Iok8skubernetespkgapiv1EndpointsList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EndpointsList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Endpoints `json:"items"` // List of endpoints.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1ObjectReference represents the Iok8skubernetespkgapiv1ObjectReference schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ObjectReference struct {
	Fieldpath string `json:"fieldPath,omitempty"` // If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
	Kind string `json:"kind,omitempty"` // Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Namespace string `json:"namespace,omitempty"` // Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Resourceversion string `json:"resourceVersion,omitempty"` // Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	Uid string `json:"uid,omitempty"` // UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent.
}

// Iok8skubernetespkgapisbatchv2alpha1CronJob represents the Iok8skubernetespkgapisbatchv2alpha1CronJob schema from the OpenAPI specification
type Iok8skubernetespkgapisbatchv2alpha1CronJob struct {
	Spec Iok8sapibatchv2alpha1CronJobSpec `json:"spec,omitempty"` // CronJobSpec describes how the job execution will look like and when it will actually run.
	Status Iok8sapibatchv2alpha1CronJobStatus `json:"status,omitempty"` // CronJobStatus represents the current state of a cron job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus represents the Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // most recent generation observed by this autoscaler.
	Currentcpuutilizationpercentage int `json:"currentCPUUtilizationPercentage,omitempty"` // current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
	Currentreplicas int `json:"currentReplicas"` // current number of replicas of pods managed by this autoscaler.
	Desiredreplicas int `json:"desiredReplicas"` // desired number of replicas of pods managed by this autoscaler.
	Lastscaletime string `json:"lastScaleTime,omitempty"`
}

// Iok8sapicorev1LimitRangeSpec represents the Iok8sapicorev1LimitRangeSpec schema from the OpenAPI specification
type Iok8sapicorev1LimitRangeSpec struct {
	Limits []Iok8sapicorev1LimitRangeItem `json:"limits"` // Limits is the list of LimitRangeItem objects that are enforced.
}

// Iok8skubernetespkgapiv1ReplicationControllerList represents the Iok8skubernetespkgapiv1ReplicationControllerList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ReplicationControllerList struct {
	Items []Iok8sapicorev1ReplicationController `json:"items"` // List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisextensionsv1beta1SupplementalGroupsStrategyOptions represents the Iok8skubernetespkgapisextensionsv1beta1SupplementalGroupsStrategyOptions schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1SupplementalGroupsStrategyOptions struct {
	Ranges []Iok8sapiextensionsv1beta1IDRange `json:"ranges,omitempty"` // Ranges are the allowed ranges of supplemental groups. If you would like to force a single supplemental group then supply a single range with the same start and end.
	Rule string `json:"rule,omitempty"` // Rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
}

// Iok8sapiappsv1beta1StatefulSet represents the Iok8sapiappsv1beta1StatefulSet schema from the OpenAPI specification
type Iok8sapiappsv1beta1StatefulSet struct {
	Spec Iok8sapiappsv1beta1StatefulSetSpec `json:"spec,omitempty"` // A StatefulSetSpec is the specification of a StatefulSet.
	Status Iok8sapiappsv1beta1StatefulSetStatus `json:"status,omitempty"` // StatefulSetStatus represents the current state of a StatefulSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8skubernetespkgapiv1PersistentVolumeClaim represents the Iok8skubernetespkgapiv1PersistentVolumeClaim schema from the OpenAPI specification
type Iok8skubernetespkgapiv1PersistentVolumeClaim struct {
	Spec Iok8sapicorev1PersistentVolumeClaimSpec `json:"spec,omitempty"` // PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
	Status Iok8sapicorev1PersistentVolumeClaimStatus `json:"status,omitempty"` // PersistentVolumeClaimStatus is the current status of a persistent volume claim.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1CSIPersistentVolumeSource represents the Iok8sapicorev1CSIPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CSIPersistentVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // Optional: The value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	Volumeattributes map[string]interface{} `json:"volumeAttributes,omitempty"` // Attributes of the volume to publish.
	Volumehandle string `json:"volumeHandle"` // VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
	Controllerpublishsecretref Iok8sapicorev1SecretReference `json:"controllerPublishSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Driver string `json:"driver"` // Driver is the name of the driver to use for this volume. Required.
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Nodepublishsecretref Iok8sapicorev1SecretReference `json:"nodePublishSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Nodestagesecretref Iok8sapicorev1SecretReference `json:"nodeStageSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
}

// Iok8skubernetespkgapiv1EmptyDirVolumeSource represents the Iok8skubernetespkgapiv1EmptyDirVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1EmptyDirVolumeSource struct {
	Sizelimit string `json:"sizeLimit,omitempty"`
	Medium string `json:"medium,omitempty"` // What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
}

// Iok8sapicorev1ServiceSpec represents the Iok8sapicorev1ServiceSpec schema from the OpenAPI specification
type Iok8sapicorev1ServiceSpec struct {
	Ports []Iok8sapicorev1ServicePort `json:"ports,omitempty"` // The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externaltrafficpolicy string `json:"externalTrafficPolicy,omitempty"` // externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. "Local" preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. "Cluster" obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.
	Loadbalancersourceranges []string `json:"loadBalancerSourceRanges,omitempty"` // If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature." More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	Selector map[string]interface{} `json:"selector,omitempty"` // Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
	Sessionaffinityconfig Iok8sapicorev1SessionAffinityConfig `json:"sessionAffinityConfig,omitempty"` // SessionAffinityConfig represents the configurations of session affinity.
	TypeField string `json:"type,omitempty"` // type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ExternalName" maps to the specified externalName. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a stable IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services---service-types
	Clusterip string `json:"clusterIP,omitempty"` // clusterIP is the IP address of the service and is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. This field can not be changed through updates. Valid values are "None", empty string (""), or a valid IP address. "None" can be specified for headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externalips []string `json:"externalIPs,omitempty"` // externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.
	Externalname string `json:"externalName,omitempty"` // externalName is the external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.
	Loadbalancerip string `json:"loadBalancerIP,omitempty"` // Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.
	Publishnotreadyaddresses bool `json:"publishNotReadyAddresses,omitempty"` // publishNotReadyAddresses, when set to true, indicates that DNS implementations must publish the notReadyAddresses of subsets for the Endpoints associated with the Service. The default value is false. The primary use case for setting this field is to use a StatefulSet's Headless Service to propagate SRV records for its Pods without respect to their readiness for purpose of peer discovery. This field will replace the service.alpha.kubernetes.io/tolerate-unready-endpoints when that annotation is deprecated and all clients have been converted to use this field.
	Sessionaffinity string `json:"sessionAffinity,omitempty"` // Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Healthchecknodeport int `json:"healthCheckNodePort,omitempty"` // healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified, HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use user-specified nodePort value if specified by the client. Only effects when Type is set to LoadBalancer and ExternalTrafficPolicy is set to Local.
}

// Iok8skubernetespkgapiv1LocalObjectReference represents the Iok8skubernetespkgapiv1LocalObjectReference schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LocalObjectReference struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8sapirbacv1beta1Role represents the Iok8sapirbacv1beta1Role schema from the OpenAPI specification
type Iok8sapirbacv1beta1Role struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1beta1PolicyRule `json:"rules"` // Rules holds all the PolicyRules for this Role
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []GeneratedType `json:"items"`
}

// Iok8sapicorev1PodCondition represents the Iok8sapicorev1PodCondition schema from the OpenAPI specification
type Iok8sapicorev1PodCondition struct {
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	TypeField string `json:"type"` // Type is the type of the condition. Currently only Ready. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Lastprobetime string `json:"lastProbeTime,omitempty"`
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
}

// Iok8skubernetespkgapiv1Service represents the Iok8skubernetespkgapiv1Service schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Service struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ServiceSpec `json:"spec,omitempty"` // ServiceSpec describes the attributes that a user creates on a service.
	Status Iok8sapicorev1ServiceStatus `json:"status,omitempty"` // ServiceStatus represents the current status of a service.
}

// Iok8sapicorev1Binding represents the Iok8sapicorev1Binding schema from the OpenAPI specification
type Iok8sapicorev1Binding struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Target Iok8sapicorev1ObjectReference `json:"target"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiextensionsv1beta1PodSecurityPolicySpec represents the Iok8sapiextensionsv1beta1PodSecurityPolicySpec schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1PodSecurityPolicySpec struct {
	Readonlyrootfilesystem bool `json:"readOnlyRootFilesystem,omitempty"` // ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system. If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	Runasuser Iok8sapiextensionsv1beta1RunAsUserStrategyOptions `json:"runAsUser"` // Run A sUser Strategy Options defines the strategy type and any options used to create the strategy.
	Requireddropcapabilities []string `json:"requiredDropCapabilities,omitempty"` // RequiredDropCapabilities are the capabilities that will be dropped from the container. These are required to be dropped and cannot be added.
	Allowedhostpaths []Iok8sapiextensionsv1beta1AllowedHostPath `json:"allowedHostPaths,omitempty"` // is a white list of allowed host paths. Empty indicates that all host paths may be used.
	Defaultaddcapabilities []string `json:"defaultAddCapabilities,omitempty"` // DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability. You may not list a capability in both DefaultAddCapabilities and RequiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the AllowedCapabilities list.
	Hostipc bool `json:"hostIPC,omitempty"` // hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	Supplementalgroups Iok8sapiextensionsv1beta1SupplementalGroupsStrategyOptions `json:"supplementalGroups"` // SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
	Defaultallowprivilegeescalation bool `json:"defaultAllowPrivilegeEscalation,omitempty"` // DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	Volumes []string `json:"volumes,omitempty"` // volumes is a white list of allowed volume plugins. Empty indicates that all plugins may be used.
	Hostports []Iok8sapiextensionsv1beta1HostPortRange `json:"hostPorts,omitempty"` // hostPorts determines which host port ranges are allowed to be exposed.
	Selinux Iok8sapiextensionsv1beta1SELinuxStrategyOptions `json:"seLinux"` // SELinux Strategy Options defines the strategy type and any options used to create the strategy.
	Allowprivilegeescalation bool `json:"allowPrivilegeEscalation,omitempty"` // AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	Fsgroup Iok8sapiextensionsv1beta1FSGroupStrategyOptions `json:"fsGroup"` // FSGroupStrategyOptions defines the strategy type and options used to create the strategy.
	Hostnetwork bool `json:"hostNetwork,omitempty"` // hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	Hostpid bool `json:"hostPID,omitempty"` // hostPID determines if the policy allows the use of HostPID in the pod spec.
	Allowedflexvolumes []Iok8sapiextensionsv1beta1AllowedFlexVolume `json:"allowedFlexVolumes,omitempty"` // AllowedFlexVolumes is a whitelist of allowed Flexvolumes. Empty or nil indicates that all Flexvolumes may be used. This parameter is effective only when the usage of the Flexvolumes is allowed in the "Volumes" field.
	Allowedcapabilities []string `json:"allowedCapabilities,omitempty"` // AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities.
	Privileged bool `json:"privileged,omitempty"` // privileged determines if a pod can request to be run as privileged.
}

// Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReviewSpec represents the Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1beta1SubjectAccessReviewSpec struct {
	Resourceattributes Iok8sapiauthorizationv1beta1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
	Uid string `json:"uid,omitempty"` // UID information about the requesting user.
	User string `json:"user,omitempty"` // User is the user you're testing for. If you specify "User" but not "Group", then is it interpreted as "What if User were not a member of any groups
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra corresponds to the user.Info.GetExtra() method from the authenticator. Since that is input to the authorizer it needs a reflection here.
	Group []string `json:"group,omitempty"` // Groups is the groups you're testing for.
	Nonresourceattributes Iok8sapiauthorizationv1beta1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
}

// Iok8sapimachinerypkgapismetav1APIGroupList represents the Iok8sapimachinerypkgapismetav1APIGroupList schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIGroupList struct {
	Groups []Iok8sapimachinerypkgapismetav1APIGroup `json:"groups"` // groups is a list of APIGroup.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiauthorizationv1ResourceRule represents the Iok8sapiauthorizationv1ResourceRule schema from the OpenAPI specification
type Iok8sapiauthorizationv1ResourceRule struct {
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed. "*" means all.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed. "*" means all.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. "*" means all in the specified apiGroups. "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
	Verbs []string `json:"verbs"` // Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy. "*" means all.
}

// Iok8sapicorev1RBDPersistentVolumeSource represents the Iok8sapicorev1RBDPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1RBDPersistentVolumeSource struct {
	Monitors []string `json:"monitors"` // A collection of Ceph monitors. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"` // The rados pool name. Default is rbd. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	User string `json:"user,omitempty"` // The rados user name. Default is admin. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Fstype string `json:"fsType,omitempty"` // Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	Image string `json:"image"` // The rados image name. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"` // Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
}

// Iok8skubernetespkgapisextensionsv1beta1IngressBackend represents the Iok8skubernetespkgapisextensionsv1beta1IngressBackend schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1IngressBackend struct {
	Servicename string `json:"serviceName"` // Specifies the name of the referenced service.
	Serviceport string `json:"servicePort"`
}

// Iok8sapistoragev1alpha1VolumeAttachmentSpec represents the Iok8sapistoragev1alpha1VolumeAttachmentSpec schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttachmentSpec struct {
	Attacher string `json:"attacher"` // Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Nodename string `json:"nodeName"` // The node that the volume should be attached to.
	Source Iok8sapistoragev1alpha1VolumeAttachmentSource `json:"source"` // VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
}

// Iok8sapipolicyv1beta1PodSecurityPolicySpec represents the Iok8sapipolicyv1beta1PodSecurityPolicySpec schema from the OpenAPI specification
type Iok8sapipolicyv1beta1PodSecurityPolicySpec struct {
	Allowedcapabilities []string `json:"allowedCapabilities,omitempty"` // AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities.
	Defaultallowprivilegeescalation bool `json:"defaultAllowPrivilegeEscalation,omitempty"` // DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	Requireddropcapabilities []string `json:"requiredDropCapabilities,omitempty"` // RequiredDropCapabilities are the capabilities that will be dropped from the container. These are required to be dropped and cannot be added.
	Hostnetwork bool `json:"hostNetwork,omitempty"` // hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	Hostipc bool `json:"hostIPC,omitempty"` // hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	Runasuser Iok8sapipolicyv1beta1RunAsUserStrategyOptions `json:"runAsUser"` // Run A sUser Strategy Options defines the strategy type and any options used to create the strategy.
	Selinux Iok8sapipolicyv1beta1SELinuxStrategyOptions `json:"seLinux"` // SELinux Strategy Options defines the strategy type and any options used to create the strategy.
	Readonlyrootfilesystem bool `json:"readOnlyRootFilesystem,omitempty"` // ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system. If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	Supplementalgroups Iok8sapipolicyv1beta1SupplementalGroupsStrategyOptions `json:"supplementalGroups"` // SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
	Allowprivilegeescalation bool `json:"allowPrivilegeEscalation,omitempty"` // AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	Fsgroup Iok8sapipolicyv1beta1FSGroupStrategyOptions `json:"fsGroup"` // FSGroupStrategyOptions defines the strategy type and options used to create the strategy.
	Defaultaddcapabilities []string `json:"defaultAddCapabilities,omitempty"` // DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability. You may not list a capability in both DefaultAddCapabilities and RequiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the AllowedCapabilities list.
	Volumes []string `json:"volumes,omitempty"` // volumes is a white list of allowed volume plugins. Empty indicates that all plugins may be used.
	Hostports []Iok8sapipolicyv1beta1HostPortRange `json:"hostPorts,omitempty"` // hostPorts determines which host port ranges are allowed to be exposed.
	Allowedhostpaths []Iok8sapipolicyv1beta1AllowedHostPath `json:"allowedHostPaths,omitempty"` // is a white list of allowed host paths. Empty indicates that all host paths may be used.
	Allowedflexvolumes []Iok8sapipolicyv1beta1AllowedFlexVolume `json:"allowedFlexVolumes,omitempty"` // AllowedFlexVolumes is a whitelist of allowed Flexvolumes. Empty or nil indicates that all Flexvolumes may be used. This parameter is effective only when the usage of the Flexvolumes is allowed in the "Volumes" field.
	Hostpid bool `json:"hostPID,omitempty"` // hostPID determines if the policy allows the use of HostPID in the pod spec.
	Privileged bool `json:"privileged,omitempty"` // privileged determines if a pod can request to be run as privileged.
}

// Iok8skubernetespkgapiv1ResourceFieldSelector represents the Iok8skubernetespkgapiv1ResourceFieldSelector schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ResourceFieldSelector struct {
	Containername string `json:"containerName,omitempty"` // Container name: required for volumes, optional for env vars
	Divisor string `json:"divisor,omitempty"`
	Resource string `json:"resource"` // Required: resource to select
}

// Iok8skubernetespkgapisextensionsv1beta1ReplicaSetSpec represents the Iok8skubernetespkgapisextensionsv1beta1ReplicaSetSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ReplicaSetSpec struct {
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
}

// Iok8sapicorev1NodeList represents the Iok8sapicorev1NodeList schema from the OpenAPI specification
type Iok8sapicorev1NodeList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1Node `json:"items"` // List of nodes
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8sapicorev1ServiceAccount represents the Iok8sapicorev1ServiceAccount schema from the OpenAPI specification
type Iok8sapicorev1ServiceAccount struct {
	Imagepullsecrets []Iok8sapicorev1LocalObjectReference `json:"imagePullSecrets,omitempty"` // ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Secrets []Iok8sapicorev1ObjectReference `json:"secrets,omitempty"` // Secrets is the list of secrets allowed to be used by pods running using this ServiceAccount. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Automountserviceaccounttoken bool `json:"automountServiceAccountToken,omitempty"` // AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
}

// Iok8sapirbacv1beta1PolicyRule represents the Iok8sapirbacv1beta1PolicyRule schema from the OpenAPI specification
type Iok8sapirbacv1beta1PolicyRule struct {
	Verbs []string `json:"verbs"` // Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. '*' represents all resources in the specified apiGroups. '*/foo' represents the subresource 'foo' for all resources in the specified apiGroups.
}

// Iok8sapiadmissionregistrationv1alpha1Initializer represents the Iok8sapiadmissionregistrationv1alpha1Initializer schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1Initializer struct {
	Name string `json:"name"` // Name is the identifier of the initializer. It will be added to the object that needs to be initialized. Name should be fully qualified, e.g., alwayspullimages.kubernetes.io, where "alwayspullimages" is the name of the webhook, and kubernetes.io is the name of the organization. Required
	Rules []Iok8sapiadmissionregistrationv1alpha1Rule `json:"rules,omitempty"` // Rules describes what resources/subresources the initializer cares about. The initializer cares about an operation if it matches _any_ Rule. Rule.Resources must not include subresources.
}

// Iok8sapicorev1CephFSPersistentVolumeSource represents the Iok8sapicorev1CephFSPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CephFSPersistentVolumeSource struct {
	Secretfile string `json:"secretFile,omitempty"` // Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	User string `json:"user,omitempty"` // Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors"` // Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Path string `json:"path,omitempty"` // Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Readonly bool `json:"readOnly,omitempty"` // Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
}

// Iok8sapiappsv1beta2ReplicaSet represents the Iok8sapiappsv1beta2ReplicaSet schema from the OpenAPI specification
type Iok8sapiappsv1beta2ReplicaSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta2ReplicaSetSpec `json:"spec,omitempty"` // ReplicaSetSpec is the specification of a ReplicaSet.
	Status Iok8sapiappsv1beta2ReplicaSetStatus `json:"status,omitempty"` // ReplicaSetStatus represents the current status of a ReplicaSet.
}

// Iok8skubernetespkgapisnetworkingv1NetworkPolicyIngressRule represents the Iok8skubernetespkgapisnetworkingv1NetworkPolicyIngressRule schema from the OpenAPI specification
type Iok8skubernetespkgapisnetworkingv1NetworkPolicyIngressRule struct {
	From []Iok8sapinetworkingv1NetworkPolicyPeer `json:"from,omitempty"` // List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.
	Ports []Iok8sapinetworkingv1NetworkPolicyPort `json:"ports,omitempty"` // List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
}

// Iok8skubernetespkgapiv1DownwardAPIVolumeFile represents the Iok8skubernetespkgapiv1DownwardAPIVolumeFile schema from the OpenAPI specification
type Iok8skubernetespkgapiv1DownwardAPIVolumeFile struct {
	Path string `json:"path"` // Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Resourcefieldref Iok8sapicorev1ResourceFieldSelector `json:"resourceFieldRef,omitempty"` // ResourceFieldSelector represents container resources (cpu, memory) and their output format
	Fieldref Iok8sapicorev1ObjectFieldSelector `json:"fieldRef,omitempty"` // ObjectFieldSelector selects an APIVersioned field of an object.
	Mode int `json:"mode,omitempty"` // Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8skubernetespkgapisextensionsv1beta1DaemonSetStatus represents the Iok8skubernetespkgapisextensionsv1beta1DaemonSetStatus schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DaemonSetStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The most recent generation observed by the daemon set controller.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Numberunavailable int `json:"numberUnavailable,omitempty"` // The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numbermisscheduled int `json:"numberMisscheduled"` // The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Updatednumberscheduled int `json:"updatedNumberScheduled,omitempty"` // The total number of nodes that are running updated daemon pod
	Conditions []Iok8sapiextensionsv1beta1DaemonSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a DaemonSet's current state.
	Currentnumberscheduled int `json:"currentNumberScheduled"` // The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Desirednumberscheduled int `json:"desiredNumberScheduled"` // The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Numberavailable int `json:"numberAvailable,omitempty"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numberready int `json:"numberReady"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
}

// Iok8skubernetespkgapiv1LimitRangeList represents the Iok8skubernetespkgapiv1LimitRangeList schema from the OpenAPI specification
type Iok8skubernetespkgapiv1LimitRangeList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapicorev1LimitRange `json:"items"` // Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1ConfigMapEnvSource represents the Iok8skubernetespkgapiv1ConfigMapEnvSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1ConfigMapEnvSource struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap must be defined
}

// Iok8skubernetespkgapiv1VsphereVirtualDiskVolumeSource represents the Iok8skubernetespkgapiv1VsphereVirtualDiskVolumeSource schema from the OpenAPI specification
type Iok8skubernetespkgapiv1VsphereVirtualDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Storagepolicyid string `json:"storagePolicyID,omitempty"` // Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	Storagepolicyname string `json:"storagePolicyName,omitempty"` // Storage Policy Based Management (SPBM) profile name.
	Volumepath string `json:"volumePath"` // Path that identifies vSphere volume vmdk
}

// Iok8sapiextensionsv1beta1DaemonSetList represents the Iok8sapiextensionsv1beta1DaemonSetList schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1DaemonSetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1DaemonSet `json:"items"` // A list of daemon sets.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1beta2Deployment represents the Iok8sapiappsv1beta2Deployment schema from the OpenAPI specification
type Iok8sapiappsv1beta2Deployment struct {
	Status Iok8sapiappsv1beta2DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1beta2DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
}

// Iok8skubernetespkgapisauthenticationv1TokenReviewSpec represents the Iok8skubernetespkgapisauthenticationv1TokenReviewSpec schema from the OpenAPI specification
type Iok8skubernetespkgapisauthenticationv1TokenReviewSpec struct {
	Token string `json:"token,omitempty"` // Token is the opaque bearer token.
}

// Iok8sapischedulingv1alpha1PriorityClassList represents the Iok8sapischedulingv1alpha1PriorityClassList schema from the OpenAPI specification
type Iok8sapischedulingv1alpha1PriorityClassList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapischedulingv1alpha1PriorityClass `json:"items"` // items is the list of PriorityClasses
}

// Iok8sapiautoscalingv2beta1MetricSpec represents the Iok8sapiautoscalingv2beta1MetricSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1MetricSpec struct {
	External Iok8sapiautoscalingv2beta1ExternalMetricSource `json:"external,omitempty"` // ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one "target" type should be set.
	Object Iok8sapiautoscalingv2beta1ObjectMetricSource `json:"object,omitempty"` // ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
	Pods Iok8sapiautoscalingv2beta1PodsMetricSource `json:"pods,omitempty"` // PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
	Resource Iok8sapiautoscalingv2beta1ResourceMetricSource `json:"resource,omitempty"` // ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory). The values will be averaged together before being compared to the target. Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source. Only one "target" type should be set.
	TypeField string `json:"type"` // type is the type of metric source. It should be one of "Object", "Pods" or "Resource", each mapping to a matching field in the object.
}

// Iok8skubernetespkgapiv1SELinuxOptions represents the Iok8skubernetespkgapiv1SELinuxOptions schema from the OpenAPI specification
type Iok8skubernetespkgapiv1SELinuxOptions struct {
	Level string `json:"level,omitempty"` // Level is SELinux level label that applies to the container.
	Role string `json:"role,omitempty"` // Role is a SELinux role label that applies to the container.
	TypeField string `json:"type,omitempty"` // Type is a SELinux type label that applies to the container.
	User string `json:"user,omitempty"` // User is a SELinux user label that applies to the container.
}

// Iok8sapicorev1AzureFileVolumeSource represents the Iok8sapicorev1AzureFileVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AzureFileVolumeSource struct {
	Secretname string `json:"secretName"` // the name of secret that contains Azure Storage Account Name and Key
	Sharename string `json:"shareName"` // Share Name
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapicorev1PortworxVolumeSource represents the Iok8sapicorev1PortworxVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1PortworxVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Volumeid string `json:"volumeID"` // VolumeID uniquely identifies a Portworx volume
	Fstype string `json:"fsType,omitempty"` // FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
}

// Iok8sapiappsv1DeploymentStrategy represents the Iok8sapiappsv1DeploymentStrategy schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentStrategy struct {
	Rollingupdate Iok8sapiappsv1RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
}

// Iok8skubernetespkgapisextensionsv1beta1ReplicaSetList represents the Iok8skubernetespkgapisextensionsv1beta1ReplicaSetList schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1ReplicaSetList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapiextensionsv1beta1ReplicaSet `json:"items"` // List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
}

// Iok8sapiautoscalingv2beta1ExternalMetricStatus represents the Iok8sapiautoscalingv2beta1ExternalMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2beta1ExternalMetricStatus struct {
	Currentaveragevalue string `json:"currentAverageValue,omitempty"`
	Currentvalue string `json:"currentValue"`
	Metricname string `json:"metricName"` // metricName is the name of a metric used for autoscaling in metric system.
	Metricselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"metricSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiappsv1beta1DeploymentStatus represents the Iok8sapiappsv1beta1DeploymentStatus schema from the OpenAPI specification
type Iok8sapiappsv1beta1DeploymentStatus struct {
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of ready pods targeted by this deployment.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Conditions []Iok8sapiappsv1beta1DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
}

// Iok8sapiauthenticationv1beta1TokenReview represents the Iok8sapiauthenticationv1beta1TokenReview schema from the OpenAPI specification
type Iok8sapiauthenticationv1beta1TokenReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthenticationv1beta1TokenReviewSpec `json:"spec"` // TokenReviewSpec is a description of the token authentication request.
	Status Iok8sapiauthenticationv1beta1TokenReviewStatus `json:"status,omitempty"` // TokenReviewStatus is the result of the token authentication request.
}

// Iok8skubernetespkgapisextensionsv1beta1DeploymentList represents the Iok8skubernetespkgapisextensionsv1beta1DeploymentList schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1DeploymentList struct {
	Items []Iok8sapiextensionsv1beta1Deployment `json:"items"` // Items is the list of Deployments.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8sapiappsv1beta2DaemonSetCondition represents the Iok8sapiappsv1beta2DaemonSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1beta2DaemonSetCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of DaemonSet condition.
}

// Iok8sapicorev1Handler represents the Iok8sapicorev1Handler schema from the OpenAPI specification
type Iok8sapicorev1Handler struct {
	Exec Iok8sapicorev1ExecAction `json:"exec,omitempty"` // ExecAction describes a "run in container" action.
	Httpget Iok8sapicorev1HTTPGetAction `json:"httpGet,omitempty"` // HTTPGetAction describes an action based on HTTP Get requests.
	Tcpsocket Iok8sapicorev1TCPSocketAction `json:"tcpSocket,omitempty"` // TCPSocketAction describes an action based on opening a socket
}

// Iok8sapicorev1ComponentStatus represents the Iok8sapicorev1ComponentStatus schema from the OpenAPI specification
type Iok8sapicorev1ComponentStatus struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Conditions []Iok8sapicorev1ComponentCondition `json:"conditions,omitempty"` // List of component conditions observed
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
}

// Iok8skubernetespkgapiv1Endpoints represents the Iok8skubernetespkgapiv1Endpoints schema from the OpenAPI specification
type Iok8skubernetespkgapiv1Endpoints struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Subsets []Iok8sapicorev1EndpointSubset `json:"subsets,omitempty"` // The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
}

// Iok8skubernetespkgapisextensionsv1beta1HostPortRange represents the Iok8skubernetespkgapisextensionsv1beta1HostPortRange schema from the OpenAPI specification
type Iok8skubernetespkgapisextensionsv1beta1HostPortRange struct {
	Max int `json:"max"` // max is the end of the range, inclusive.
	Min int `json:"min"` // min is the start of the range, inclusive.
}

// Iok8skubernetespkgapisauthorizationv1NonResourceAttributes represents the Iok8skubernetespkgapisauthorizationv1NonResourceAttributes schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1NonResourceAttributes struct {
	Verb string `json:"verb,omitempty"` // Verb is the standard HTTP verb
	Path string `json:"path,omitempty"` // Path is the URL path of the request
}

// Iok8sapiauthorizationv1beta1NonResourceAttributes represents the Iok8sapiauthorizationv1beta1NonResourceAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1beta1NonResourceAttributes struct {
	Verb string `json:"verb,omitempty"` // Verb is the standard HTTP verb
	Path string `json:"path,omitempty"` // Path is the URL path of the request
}

// Iok8skubernetespkgapisappsv1beta1DeploymentCondition represents the Iok8skubernetespkgapisappsv1beta1DeploymentCondition schema from the OpenAPI specification
type Iok8skubernetespkgapisappsv1beta1DeploymentCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"`
	Lastupdatetime string `json:"lastUpdateTime,omitempty"`
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
}

// Iok8sapiextensionsv1beta1Deployment represents the Iok8sapiextensionsv1beta1Deployment schema from the OpenAPI specification
type Iok8sapiextensionsv1beta1Deployment struct {
	Spec Iok8sapiextensionsv1beta1DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
	Status Iok8sapiextensionsv1beta1DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapirbacv1alpha1RoleBindingList represents the Iok8sapirbacv1alpha1RoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1alpha1RoleBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Items []Iok8sapirbacv1alpha1RoleBinding `json:"items"` // Items is a list of RoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8skubernetespkgapisauthorizationv1LocalSubjectAccessReview represents the Iok8skubernetespkgapisauthorizationv1LocalSubjectAccessReview schema from the OpenAPI specification
type Iok8skubernetespkgapisauthorizationv1LocalSubjectAccessReview struct {
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
}

// Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetStatus represents the Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetStatus schema from the OpenAPI specification
type Iok8skubernetespkgapispolicyv1beta1PodDisruptionBudgetStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // Most recent generation observed when updating this PDB status. PodDisruptionsAllowed and other status informatio is valid only if observedGeneration equals to PDB's object generation.
	Currenthealthy int `json:"currentHealthy"` // current number of healthy pods
	Desiredhealthy int `json:"desiredHealthy"` // minimum desired number of healthy pods
	Disruptedpods map[string]interface{} `json:"disruptedPods"` // DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
	Disruptionsallowed int `json:"disruptionsAllowed"` // Number of pod disruptions that are currently allowed.
	Expectedpods int `json:"expectedPods"` // total number of pods counted by this disruption budget
}

// Iok8sapicorev1PodDNSConfigOption represents the Iok8sapicorev1PodDNSConfigOption schema from the OpenAPI specification
type Iok8sapicorev1PodDNSConfigOption struct {
	Name string `json:"name,omitempty"` // Required.
	Value string `json:"value,omitempty"`
}
