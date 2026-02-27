package structvalues

import "time"

// -------- Tool input/output types --------
// Pod Structs
type ListPodsInput struct {
	Namespace     string `json:"namespace" jsonschema:"required,Namespace to list pods from"`
	LabelSelector string `json:"labelSelector,omitempty" jsonschema:"Optional Kubernetes label selector, e.g. app=nginx"`
	PodName       string `json:"podname" jsonschema:"required,podname to get pod logs"`
}

type PodInfo struct {
	Name     string `json:"name"`
	Phase    string `json:"phase"`
	Ready    string `json:"ready"`
	Restarts int32  `json:"restarts"`
	Age      string `json:"age"`
}

type ListPodsOutput struct {
	Pods []PodInfo `json:"pods"`
}

type LogsPodinput struct {
	Namespace     string     `json:"namespace" jsonschema:"required,Namespace to list pods from"`
	LabelSelector string     `json:"labelSelector,omitempty" jsonschema:"Optional Kubernetes label selector, e.g. app=nginx"`
	PodName       string     `json:"podname" jsonschema:"required,podname to get pod logs"`
	Container     string     `json:"container,omitempty"`
	TailLines     *int64     `json:"taillines,omitempty"`
	SinceSeconds  *int64     `json:"sinceseconds,omitempty"`
	SinceTime     *time.Time `json:"sincetime,omitempty"`
	LimitBytes    *int64     `json:"limitbytes,omitempty"`
	Previous      bool       `json:"previous,omitempty"`
	Follow        bool       `json:"follow,omitempty"`
	Timestamps    bool       `json:"timestamps,omitempty"`
}

type LogsPod struct {
	Namespace  string    `json:"namespace"`
	PodName    string    `json:"podName"`
	Container  string    `json:"container,omitempty"`
	Logs       string    `json:"logs"`
	Timestamps bool      `json:"timestamps"`
	CapturedAt time.Time `json:"capturedAt"`
}

type PodLogsOutput struct {
	Pods []LogsPod `json:"logspods"`
}

// Namespace Struct
type ListNamespacesOutput struct {
	Namespaces []string `json:"namespaces"`
}

// Volume Struct
type VolumeInfo struct {
	Name                  string `json:"name"`
	Capacity              int    `json:"capacity"`
	AccessModes           string `json:"accessmodes"`
	ReclaimPolicy         string `json:"reclaimpolicy"`
	Status                string `json:"status"`
	Claim                 string `json:"claim"`
	StorageClass          string `json:"storageclass"`
	VolumeAttributesClass string `json:"volumeattributesclass,omitempty"`
	Reason                string `json:"reason,omitempty"`
	Age                   string `json:"age"`
}

type ListPvOutput struct {
	PersistentVolumes []VolumeInfo `json:"persistentvolumes"`
}
