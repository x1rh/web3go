package convertx

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestToLowerUnit(t *testing.T) {
}

func TestWeiToGwei(t *testing.T) {

}

func TestWeiToEther(t *testing.T) {

}

func TestGweiToEther(t *testing.T) {

}

func TestGweiToWei(t *testing.T) {

}

func TestEtherToWei(t *testing.T) {
	x_wei, err := decimal.NewFromString("114514000000000000000000000")
	if err != nil {
		t.Fatal("invalid", err)
	}

	type args struct {
		x string
	}
	tests := []struct {
		name    string
		args    args
		want    decimal.Decimal
		wantErr bool
	}{
		{name: "test1", args: args{
			x: "114514000",
		}, want: x_wei,
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EtherToWei(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("EtherToWei() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if x_wei.Cmp(*got) != 0 {
				// if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EtherToWei() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEtherToGwei(t *testing.T) {

}
