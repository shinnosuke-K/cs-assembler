package corres_table

import "testing"

func TestGetJumpBinary(t *testing.T) {
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
			name: "null",
			args: args{word: "null"},
			want: "000",
		},
		{
			name: "JGT",
			args: args{word: "JGT"},
			want: "001",
		},
		{
			name: "JMP",
			args: args{word: "JMP"},
			want: "111",
		},
		{
			name: "FAIL",
			args: args{word: "JFT"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJumpBinary(tt.args.word); got != tt.want {
				t.Errorf("GetJumpBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
