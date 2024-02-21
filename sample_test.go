package sample

import (
	"reflect"
	"testing"
)

func TestNewExample1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Example
	}{
		{
			name: "",
			args: args{
				name: "Test",
			},
			want: &Example{
				Name: "Test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExample(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
