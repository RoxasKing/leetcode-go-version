package main

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			6,
		},
		{
			"2",
			args{[][]byte{}},
			0,
		},
		{
			"3",
			args{[][]byte{{'0'}}},
			0,
		},
		{
			"4",
			args{[][]byte{{'1'}}},
			1,
		},
		{
			"5",
			args{[][]byte{{'0', '0'}}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
