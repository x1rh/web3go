package convertx

import (
	"math/big"
	"reflect"
	"testing"
)

func TestToLowerUnit(t *testing.T) {
	res, ok := big.NewInt(0).SetString("114514000000000000000000000", 10)
	if !ok {
		t.Fatal("invalid")
	}

	type args struct {
		x        string
		accuracy string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{name: "test1", args: args{
			x:        "114514000",
			accuracy: "18",
		}, want: res,
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DivByDecimal(tt.args.x, tt.args.accuracy)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLowerUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLowerUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
