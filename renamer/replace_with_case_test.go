package renamer

import "testing"

func TestReplaceWithCase(t *testing.T) {
	type args struct {
		text string
		old  string
		new  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple replace",
			args: args{
				text: "abc",
				old:  "a",
				new:  "b",
			},
			want: "bbc",
		},
		{
			name: "replace with same length",
			args: args{
				text: "abc",
				old:  "a",
				new:  "def",
			},
			want: "defbc",
		},
		{
			name: "replace with longer length",
			args: args{
				text: "abc",
				old:  "a",
				new:  "defg",
			},
			want: "defgbc",
		},
		{
			name: "replace with shorter length",
			args: args{
				text: "abc",
				old:  "ab",
				new:  "d",
			},
			want: "dc",
		},
		{
			name: "replace with case",
			args: args{
				text: "aBc",
				old:  "a",
				new:  "d",
			},
			want: "dBc",
		},
		{
			name: "replace with case",
			args: args{
				text: "aBc",
				old:  "aB",
				new:  "d",
			},
			want: "dc",
		},
		{
			name: "replace with case",
			args: args{
				text: "ABC",
				old:  "a",
				new:  "d",
			},
			want: "DBC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceWithCase(tt.args.text, tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("ReplaceWithCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
