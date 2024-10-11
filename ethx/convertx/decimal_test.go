package convertx

import (
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func TestToDecimal(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name    string
		args    args
		want    *decimal.Decimal
		wantErr bool
	}{
		{
			name:    "",
			args:    args{x: 0},
			want:    &zeroD,
			wantErr: false,
		},

		{
			name:    "",
			args:    args{x: 1},
			want:    &oneD,
			wantErr: false,
		},

		{
			name:    "",
			args:    args{x: 1.23},
			want:    &f1,
			wantErr: false,
		},

		{
			name:    "",
			args:    args{x: "1.23"},
			want:    &f1,
			wantErr: false,
		},
		{
			name:    "",
			args:    args{x: false},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToDecimal(tt.args.x)
			t.Logf("got=%+v, want=%+v args=%+v, wantErr=%+v err=%v\n", got, tt.want, tt.args.x, tt.wantErr, err)
			if err != nil { 
				if !tt.wantErr {
					t.Fatalf("ToDecimal() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ToDecimal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulByDecimal(t *testing.T) {
	type args struct {
		value  any
		factor any
	}
	tests := []struct {
		name    string
		args    args
		want    *decimal.Decimal
		wantErr bool
	}{
		{
			args: args{
				value:  1,
				factor: 9,
			},
			want:    &gwei,
			wantErr: false,
		},
		{
			args: args{
				value:  1,
				factor: 18,
			},
			want:    &ether,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MulByDecimal(tt.args.value, tt.args.factor)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToSmallUnit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ToSmallUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivByDecimal(t *testing.T) {
	type args struct {
		value  any
		factor any
	}
	tests := []struct {
		name    string
		args    args
		want    *decimal.Decimal
		wantErr bool
	}{
		{
			args: args{
				value:  1,
				factor: 9,
			},
			want:    &oneGWeiInEther,
			wantErr: false,
		},
		{
			args: args{
				value:  1,
				factor: 18,
			},
			want:    &oneWeiInEther,
			wantErr: false,
		},
		{
			args: args{
				value:  "1",
				factor: "9",
			},
			want:    &oneGWeiInEther,
			wantErr: false,
		},
		{
			args: args{
				value:  "1",
				factor: "18",
			},
			want:    &oneWeiInEther,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DivByDecimal(tt.args.value, tt.args.factor)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToLargeUnit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Fatalf("ToLargeUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
