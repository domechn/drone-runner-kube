package engine

import (
	"bytes"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func TestKubernetes_exec(t *testing.T) {
	cfg, _ := clientcmd.BuildConfigFromFlags("", "/Users/dmc/.kube/ali.config")
	k := kubernetes.NewForConfigOrDie(cfg)

	type fields struct {
		client *kubernetes.Clientset
		config *rest.Config
	}
	type args struct {
		podNamespace string
		podName      string
		container    string
		commands     string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		{
			name: "case1",
			fields: fields{
				client: k,
				config: cfg,
			},
			args: args{
				podNamespace: "default",
				podName:      "test-pod",
				container:    "image-1",
				commands:     "home/test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kubernetes{
				client: tt.fields.client,
				config: tt.fields.config,
			}
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			err := k.exec(tt.args.podNamespace, tt.args.podName, tt.args.container, tt.args.commands, stdout, stderr)
			if (err != nil) != tt.wantErr {
				t.Errorf("exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("exec() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
			if gotStderr := stderr.String(); gotStderr != tt.wantStderr {
				t.Errorf("exec() gotStderr = %v, want %v", gotStderr, tt.wantStderr)
			}
		})
	}
}
