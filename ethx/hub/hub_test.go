package hub

import (
	"reflect"
	"testing"
	"github.com/x1rh/ethx/config"
)

func TestNewHub(t *testing.T) {
	type args struct {
		chains map[int]config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Hub
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.chains)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProxy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxy() got = %v, want %v", got, tt.want)
			}
		})
	}
}
