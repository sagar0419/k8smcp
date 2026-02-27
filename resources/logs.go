package resources

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	structvalues "github.com/sagar0419/k8smcp/structValues"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) PodLogs(ctx context.Context, req *mcp.CallToolRequest, in structvalues.LogsPodinput) (*mcp.CallToolResult, structvalues.PodLogsOutput, error) {
	if in.Namespace == "" {
		return nil, structvalues.PodLogsOutput{}, fmt.Errorf("namespace is required")
	}
	if in.PodName == "" {
		return nil, structvalues.PodLogsOutput{}, fmt.Errorf("PodName is required to fetch logs")
	}

	log.Printf("k8s_pod_logs ns=%q pod=%q container=%q follow=%v tail=%v", in.Namespace, in.PodName, in.Container, in.Follow, in.TailLines)

	opts := &corev1.PodLogOptions{
		Container:  in.Container,
		Follow:     in.Follow,
		Previous:   in.Previous,
		Timestamps: in.Timestamps,
	}

	if opts.Follow {
		if in.TailLines == nil {
			var tail int64 = 1000
			opts.TailLines = &tail
		} else {
			opts.TailLines = in.TailLines
		}
	} else {
		if in.LimitBytes != nil {
			opts.LimitBytes = in.LimitBytes
		} else {
			var limit int64 = 1_000_000
			opts.LimitBytes = &limit
		}
	}

	if in.SinceSeconds != nil {
		opts.SinceSeconds = in.SinceSeconds
	}
	if in.SinceTime != nil {

		mt := metav1.NewTime(*in.SinceTime)
		opts.SinceTime = &mt
	}
	reqObj := s.k8s.CoreV1().Pods(in.Namespace).GetLogs(in.PodName, opts)

	stream, err := reqObj.Stream(ctx)
	if err != nil {
		return nil, structvalues.PodLogsOutput{}, fmt.Errorf("stream logs: %w", err)
	}
	defer stream.Close()

	buf, err := io.ReadAll(stream)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, structvalues.PodLogsOutput{}, ctx.Err()
		default:
			return nil, structvalues.PodLogsOutput{}, fmt.Errorf("read logs: %w", err)
		}
	}

	out := structvalues.PodLogsOutput{
		Pods: []structvalues.LogsPod{
			{
				Namespace:  in.Namespace,
				PodName:    in.PodName,
				Container:  in.Container,
				Logs:       string(buf),
				Timestamps: in.Timestamps,
				CapturedAt: time.Now().UTC(),
			},
		},
	}
	return nil, out, nil
}
