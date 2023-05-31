package encrypt_test

import (
	"testing"

	"github.com/atomgunlk/prime-vote/internal/encrypt"
)

func TestUnitCheckPassword(t *testing.T) {
	type args struct {
		password       string
		hashedPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				password:       "test0",
				hashedPassword: func(p string) string { h, _ := encrypt.HashPassword(p); return h }("test0"),
			},
			wantErr: false,
		},
		{
			name: "fail_mismatch",
			args: args{
				password:       "test0",
				hashedPassword: func(p string) string { h, _ := encrypt.HashPassword(p); return h }("test1"),
			},
			wantErr: true,
		},
		{
			name: "fail_empty#1",
			args: args{
				password:       "",
				hashedPassword: func(p string) string { h, _ := encrypt.HashPassword(p); return h }("test0"),
			},
			wantErr: true,
		},
		{
			name: "fail_empty#2",
			args: args{
				password:       "test0",
				hashedPassword: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encrypt.CheckPassword(tt.args.password, tt.args.hashedPassword); (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
