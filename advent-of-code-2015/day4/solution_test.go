package day4

import "testing"

func TestSolveLowestPositiveNumber(t *testing.T) {
	type args struct {
		puzzle string
		prefix string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test_0",
			args: args{
				puzzle: "abcdef",
				prefix: "00000",
			},
			want: 609043,
		},
		{
			name: "test_1",
			args: args{
				puzzle: "pqrstuv",
				prefix: "00000",
			},
			want: 1048970,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveLowestPositiveNumber(tt.args.puzzle, tt.args.prefix); got != tt.want {
				t.Errorf("SolveLowestPositiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
