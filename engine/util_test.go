package engine

import "testing"

func Test_filterEmoji(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				content: "test🐂",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterEmoji(tt.args.content); got != tt.want {
				t.Errorf("filterEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}
