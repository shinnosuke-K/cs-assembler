package main

import "testing"

func Test_encodeBinary(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{num: 1},
			want: "0000000000000001",
		},
		{
			name: "2",
			args: args{num: 2},
			want: "0000000000000010",
		},
		{
			name: "3",
			args: args{num: 3},
			want: "0000000000000011",
		},
		{
			name: "4",
			args: args{num: 4},
			want: "0000000000000100",
		},
		{
			name: "5",
			args: args{num: 5},
			want: "0000000000000101",
		},
		{
			name: "6",
			args: args{num: 6},
			want: "0000000000000110",
		},
		{
			name: "7",
			args: args{num: 7},
			want: "0000000000000111",
		},
		{
			name: "8",
			args: args{num: 8},
			want: "0000000000001000",
		},
		{
			name: "9",
			args: args{num: 9},
			want: "0000000000001001",
		},
		{
			name: "10",
			args: args{num: 10},
			want: "0000000000001010",
		},
		{
			name: "11",
			args: args{num: 11},
			want: "0000000000001011",
		},
		{
			name: "12",
			args: args{num: 12},
			want: "0000000000001100",
		},
		{
			name: "13",
			args: args{num: 13},
			want: "0000000000001101",
		},
		{
			name: "14",
			args: args{num: 14},
			want: "0000000000001110",
		},
		{
			name: "15",
			args: args{num: 15},
			want: "0000000000001111",
		},
		{
			name: "16",
			args: args{num: 16},
			want: "0000000000010000",
		},
		{
			name: "16033",
			args: args{num: 16033},
			want: "0011111010100001",
		},
		{
			name: "-1",
			args: args{num: -1},
			want: "1111111111111111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBinary(tt.args.num); got != tt.want {
				t.Errorf("encodeBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
