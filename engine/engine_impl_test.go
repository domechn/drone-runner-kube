package engine

import (
	"bytes"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"reflect"
	"testing"
)

func TestKubernetes_start(t *testing.T) {
	kk , _ := NewFromConfig("/Users/dmc/.kube/ali.config")
	type fields struct {
		client *kubernetes.Clientset
		config *rest.Config
	}
	type args struct {
		spec *Spec
		step *Step
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		want       *State
		wantErr    bool
	}{
		{
			name:       "test",
			fields:     fields{
				client: kk.client,
				config: kk.config,
			},
			args:       args{
				spec: &Spec{
					PodSpec:    PodSpec{
						Name:"test",
						Namespace:"beta",
					},
				},
				step: nil,
			},
			wantOutput: "",
			want:       nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kubernetes{
				client: tt.fields.client,
				config: tt.fields.config,
			}
			output := &bytes.Buffer{}
			got, err := k.start(tt.args.spec, tt.args.step, output)
			if (err != nil) != tt.wantErr {
				t.Errorf("start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("start() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("start() got = %v, want %v", got, tt.want)
			}
		})
	}
}
