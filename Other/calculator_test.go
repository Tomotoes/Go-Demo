package calculator

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1, 2}, 3},
		{"test1", args{1, 2}, 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Sum(test.args.x, test.args.y); got != test.want {
				t.Errorf("Sum() = %v, want %v", got, test.want)
			}
		})
	}
}
