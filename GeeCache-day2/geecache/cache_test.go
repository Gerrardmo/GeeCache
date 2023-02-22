package geecache

import (
	"reflect"
	"testing"
)

func TestNewGroup(t *testing.T) {
	type args struct {
		name       string
		cacheBytes int64
		getter     Getter
	}
	tests := []struct {
		name string
		args args
		want *Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroup(tt.args.name, tt.args.cacheBytes, tt.args.getter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroup(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGroup(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
