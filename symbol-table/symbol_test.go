package symbol_table

import "testing"

func TestGetSymbolValue(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name:  "R0",
			args:  args{"R0"},
			want:  0,
			want1: true,
		},
		{
			name:  "R16",
			args:  args{symbol: "R16"},
			want:  0,
			want1: false,
		},
		{
			name:  "SCREEN",
			args:  args{symbol: "SCREEN"},
			want:  16384,
			want1: true,
		},
		{
			name:  "unknown",
			args:  args{symbol: "dsafsa"},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetSymbolValue(tt.args.symbol)
			if got != tt.want {
				t.Errorf("GetSymbolValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetSymbolValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
