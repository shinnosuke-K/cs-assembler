package corres_table

import "testing"

func TestGetCompBinary(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{word: "0"},
			want: "0101010",
		},
		{
			name: "D",
			args: args{word: "D"},
			want: "0001100",
		},
		{
			name: "Fail",
			args: args{word: "K"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompBinary(tt.args.word); got != tt.want {
				t.Errorf("GetCompBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
