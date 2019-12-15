package corres_table

import "testing"

func TestGetDestBinary(t *testing.T) {
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
			name: "M",
			args: args{word: "M"},
			want: "001",
		},
		{
			name: "D",
			args: args{word: "D"},
			want: "010",
		},
		{
			name: "FAIL",
			args: args{word: "F"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDestBinary(tt.args.word); got != tt.want {
				t.Errorf("GetDestBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
