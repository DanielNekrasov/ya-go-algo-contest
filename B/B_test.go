package main

import "testing"

func TestCalcMax(t *testing.T) {
	type args struct {
		k    int
		list []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{k: 3, list: []int{5, 8, 2, 1, 3, 4, 11}},
			want: 24,
		},
		{
			name: "Test #2",
			args: args{k: 5, list: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "Test #3",
			args: args{k: 4, list: []int{1, 1, 9, 2, 2, 2, 6}},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcMax(tt.args.list, tt.args.k); got != tt.want {
				t.Errorf("game() = %v, want %v", got, tt.want)
			}
		})
	}
}
