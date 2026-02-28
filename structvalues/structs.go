package structValues

import "time"

// -------- Tool input/output types --------

// List Pod Structs
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

// Logs Pod Structs
type LogsPodinput struct {
	Namespace     string     `json:"namespace" jsonschema:"required,Namespace to list pods from"`
	LabelSelector string     `json:"labelSelector,omitempty" jsonschema:"Optional Kubernetes label selector, e.g. app=nginx"`
	PodName       string     `json:"podname,omitempty" jsonschema:"podname to get pod logs"`
	Container     string     `json:"container,omitempty" jsonschema:"Name of container to get logs"`
	TailLines     *int64     `json:"taillines,omitempty" jsonschema:"only bring the last N lines (avoid carrying the whole history)"`
	SinceSeconds  *int64     `json:"sinceseconds,omitempty" jsonschema:"bring logs newer than this duration"`
	SinceTime     *time.Time `json:"sincetime,omitempty" jsonschema:"bring logs newer than this time"`
	LimitBytes    *int64     `json:"limitbytes,omitempty" jsonschema:"do not bring more than N bytes (size cap)"`
	Previous      bool       `json:"previous,omitempty" jsonschema:"bring logs from the previous instance of the container (useful after a crash/restart)"`
	Follow        bool       `json:"follow,omitempty" jsonschema:"keep bringing new pages as they are being written (like kubectl logs -f)"`
	Timestamps    bool       `json:"timestamps,omitempty" jsonschema:"include a timestamp at the start of each line"`
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
