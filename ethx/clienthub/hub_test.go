package clienthub

import (
	"reflect"
	"testing"

	"github.com/x1rh/web3go/ethx/chain"
)

func TestNewHub(t *testing.T) {
	type args struct {
		chains map[int]chain.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *ClientHub
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
