package resources

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	structvalues "github.com/sagar0419/k8smcp/structValues"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Server struct {
	k8s *kubernetes.Clientset
}

// NewServer constructs a Server backed by the provided Kubernetes clientset.
func NewServer(clientset *kubernetes.Clientset) *Server {
	return &Server{k8s: clientset}
}

// Tool: k8s_list_namespaces
func (s *Server) ListNamespaces(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, structvalues.ListNamespacesOutput, error) {
	nsList, err := s.k8s.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, structvalues.ListNamespacesOutput{}, err
	}
	out := structvalues.ListNamespacesOutput{Namespaces: make([]string, 0, len(nsList.Items))}
	for _, ns := range nsList.Items {
		out.Namespaces = append(out.Namespaces, ns.Name)
	}
	return nil, out, nil
}

// Tool: k8s_list_pv
func (s *Server) ListPersistentVolumes(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, structvalues.ListPvOutput, error) {
	pvlist, err := s.k8s.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, structvalues.ListPvOutput{}, err
	}
	out := structvalues.ListPvOutput{PersistentVolumes: make([]structvalues.VolumeInfo, 0, len(pvlist.Items))}

	for _, pvs := range pvlist.Items {
		claimName := ""
		if pvs.Spec.ClaimRef != nil {
			claimName = pvs.Spec.ClaimRef.Name
		}
		vac := ""
		if pvs.Spec.VolumeAttributesClassName != nil {
			vac = *pvs.Spec.VolumeAttributesClassName
		}

		capacity := pvs.Spec.Capacity[v1.ResourceStorage]
		age := time.Since(pvs.CreationTimestamp.Time).Round(time.Second).String()
		out.PersistentVolumes = append(out.PersistentVolumes, structvalues.VolumeInfo{
			Name:                  pvs.Name,
			Capacity:              int(capacity.Value()),
			AccessModes:           fmt.Sprintf("%v", pvs.Spec.AccessModes),
			ReclaimPolicy:         string(pvs.Spec.PersistentVolumeReclaimPolicy),
			Status:                string(pvs.Status.Phase),
			Claim:                 claimName,
			StorageClass:          pvs.Spec.StorageClassName,
			VolumeAttributesClass: vac,
			Reason:                pvs.Status.Reason,
			Age:                   age,
		})
	}
	return nil, out, nil
}

// Tool: k8s_list_pods
func (s *Server) ListPods(ctx context.Context, req *mcp.CallToolRequest, in structvalues.ListPodsInput) (*mcp.CallToolResult, structvalues.ListPodsOutput, error) {
	if in.Namespace == "" {
		return nil, structvalues.ListPodsOutput{}, fmt.Errorf("namespace is required")
	}

	log.Printf("k8s_list_pods ns=%q selector=%q", in.Namespace, in.LabelSelector)

	opts := metav1.ListOptions{}
	if in.LabelSelector != "" {
		opts.LabelSelector = in.LabelSelector
	}

	pods, err := s.k8s.CoreV1().Pods(in.Namespace).List(ctx, opts)
	if err != nil {
		return nil, structvalues.ListPodsOutput{}, err
	}

	out := structvalues.ListPodsOutput{Pods: make([]structvalues.PodInfo, 0, len(pods.Items))}
	now := time.Now()
	for _, p := range pods.Items {
		var ready, total int
		var restarts int32
		for _, cs := range p.Status.ContainerStatuses {
			total++
			if cs.Ready {
				ready++
			}
			restarts += cs.RestartCount
		}
		age := now.Sub(p.CreationTimestamp.Time).Round(time.Second)
		out.Pods = append(out.Pods, structvalues.PodInfo{
			Name:     p.Name,
			Phase:    string(p.Status.Phase),
			Ready:    fmt.Sprintf("%d/%d", ready, total),
			Restarts: restarts,
			Age:      age.String(),
		})
	}

	return nil, out, nil
}
