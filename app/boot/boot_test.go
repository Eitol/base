package boot

import "testing"

func TestBoot(t *testing.T) {
	type args struct {
		isTest bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				isTest: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Boot(tt.args.isTest)
		})
	}
}
