package convertx

import (
	"reflect"
	"testing"

	"github.com/x1rh/ethx/types"
)

func TestPrettyGasInfo(t *testing.T) {
	type args struct {
		gasLimit any
		gasPrice any
	}
	tests := []struct {
		name    string
		args    args
		want    *types.GasInfo
		wantErr bool
	}{
		{
			args: args{
				gasLimit: nil,
				gasPrice: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrettyGasInfo(tt.args.gasLimit, tt.args.gasPrice)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrettyGasInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrettyGasInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
