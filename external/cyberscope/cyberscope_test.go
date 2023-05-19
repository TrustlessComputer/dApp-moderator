package cyberscope

import "testing"

func TestCheckBytescode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    []BytescodeInfo
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				code: "0x15f570dc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckBytescode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckBytescode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
