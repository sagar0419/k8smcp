package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sagar0419/k8smcp/resources"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	clientset, err := resources.KubeClient()
	if err != nil {
		log.Fatalf("kube client: %v", err)
	}

	impl := &mcp.Implementation{Name: "k8s-mcp-go", Version: "0.1.1"}
	server := mcp.NewServer(impl, nil)
	s := resources.NewServer(clientset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "k8s_list_namespaces",
		Description: "List all Kubernetes namespaces in the current cluster context",
	}, s.ListNamespaces)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "k8s_list_pods",
		Description: "List pods in a namespace (optionally filtered by label selector)",
	}, s.ListPods)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "k8s_list_pvs",
		Description: "List persistent volumes in the current cluster context",
	}, s.ListPersistentVolumes)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "k8s_pod_logs",
		Description: "Get the logs of the defined pod in the given namespace",
	}, s.PodLogs)

	// Logs
	log.SetOutput(os.Stderr)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
