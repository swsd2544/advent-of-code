package day1

import (
	"testing"
)

func TestSolveDeliveryFloor(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_0",
			args: args{
				puzzle: "(())",
			},
			want: 0,
		},
		{
			name: "test_1",
			args: args{
				puzzle: "()()",
			},
			want: 0,
		},
		{
			name: "test_2",
			args: args{
				puzzle: "(((",
			},
			want: 3,
		},
		{
			name: "test_3",
			args: args{
				puzzle: "(()(()(",
			},
			want: 3,
		},
		{
			name: "test_4",
			args: args{
				puzzle: "())",
			},
			want: -1,
		},
		{
			name: "test_5",
			args: args{
				puzzle: "))(",
			},
			want: -1,
		},
		{
			name: "test_6",
			args: args{
				puzzle: ")))",
			},
			want: -3,
		},
		{
			name: "test_7",
			args: args{
				puzzle: ")())())",
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveDeliveryFloor(tt.args.puzzle); got != tt.want {
				t.Errorf("SolveDeliveryFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveFirstCharacterIdxToBasement(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_0",
			args: args{
				puzzle: ")",
			},
			want: 1,
		},
		{
			name: "test_1",
			args: args{
				puzzle: "()())",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveFirstCharacterIdxToBasement(tt.args.puzzle); got != tt.want {
				t.Errorf("SolveFirstCharacterIdxToBasement() = %v, want %v", got, tt.want)
			}
		})
	}
}
