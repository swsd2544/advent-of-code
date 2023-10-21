package day3

import "testing"

func TestSolveNumberOfHousesWithPresents(t *testing.T) {
	type args struct {
		puzzle  string
		workers int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_0",
			args: args{
				puzzle:  ">",
				workers: 1,
			},
			want: 2,
		},
		{
			name: "test_1",
			args: args{
				puzzle:  "^>v<",
				workers: 1,
			},
			want: 4,
		},
		{
			name: "test_2",
			args: args{
				puzzle:  "^v^v^v^v^v",
				workers: 1,
			},
			want: 2,
		},
		{
			name: "test_3",
			args: args{
				puzzle:  "^v",
				workers: 2,
			},
			want: 3,
		},
		{
			name: "test_4",
			args: args{
				puzzle:  "^>v<",
				workers: 2,
			},
			want: 3,
		},
		{
			name: "test_5",
			args: args{
				puzzle:  "^v^v^v^v^v",
				workers: 2,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveNumberOfHousesWithPresents(tt.args.puzzle, tt.args.workers); got != tt.want {
				t.Errorf("SolveNumberOfHousesWithPresents() = %v, want %v", got, tt.want)
			}
		})
	}
}
