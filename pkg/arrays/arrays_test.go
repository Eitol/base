package arrays

import "testing"

func TestIn(t *testing.T) {
	type args struct {
		array []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same type",
			args: args{
				array: []string{"A", "B", "C"},
				item:  "A",
			},
			want: true,
		},
		{
			name: "not in",
			args: args{
				array: []string{"A", "B", "C"},
				item:  "Z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InStrArray(tt.args.array, tt.args.item); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}
