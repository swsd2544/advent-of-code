package day2

import (
	"testing"
)

func TestSolveTotalWrappingArea(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test_0",
			args: args{
				puzzle: "2x3x4",
			},
			want: 58,
		},
		{
			name: "test_1",
			args: args{
				puzzle: "1x1x10",
			},
			want: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveWrappingArea(tt.args.puzzle)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveTotalWrappingArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveTotalWrappingArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveRibbonLength(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test_0",
			args: args{
				puzzle: "2x3x4",
			},
			want: 34,
		},
		{
			name: "test_1",
			args: args{
				puzzle: "1x1x10",
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveRibbonLength(tt.args.puzzle)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveRibbonLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveRibbonLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
