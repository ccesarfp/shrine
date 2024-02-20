package model

import (
	"reflect"
	"testing"
)

func TestNewOpaqueToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *OpaqueToken
		wantErr bool
	}{
		{
			name: "creating_a_opaque_token_with_a_invalid_token",
			args: args{
				token: "aaaaa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creating_a_opaque_token_with_a_valid_token",
			args: args{
				token: "e87b2ef4-1a75-5dfd-814e-ba5868b8fa4d",
			},
			want: &OpaqueToken{
				Token: "e87b2ef4-1a75-5dfd-814e-ba5868b8fa4d",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOpaqueToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOpaqueToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOpaqueToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOpaqueTokenWithJwt(t *testing.T) {
	type args struct {
		token string
		jwt   string
	}
	tests := []struct {
		name    string
		args    args
		want    *OpaqueToken
		wantErr bool
	}{
		{
			name: "creating_a_opaque_token_with_a_invalid_token_and_valid_jwt",
			args: args{
				token: "aaaaa",
				jwt:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiUG9zdG1hbiIsImV4cCI6MTcwODM5NTAzMywiaWQiOjEsIm5hbWUiOiJDYWlvIn0.zNN7G5zwltkKCy01I79bMvmppB8I3IsyW5_ZUvB2ric",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creating_a_opaque_token_with_a_invalid_token_and_invalid_jwt",
			args: args{
				token: "aaaaa",
				jwt:   "aaaaa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creating_a_opaque_token_with_a_valid_token_and_invalid_jwt",
			args: args{
				token: "e87b2ef4-1a75-5dfd-814e-ba5868b8fa4d",
				jwt:   "aaaaa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creating_a_opaque_token_with_a_valid_token_and_valid_jwt",
			args: args{
				token: "e87b2ef4-1a75-5dfd-814e-ba5868b8fa4d",
				jwt:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiUG9zdG1hbiIsImV4cCI6MTcwODM5NTAzMywiaWQiOjEsIm5hbWUiOiJDYWlvIn0.zNN7G5zwltkKCy01I79bMvmppB8I3IsyW5_ZUvB2ric",
			},
			want: &OpaqueToken{
				Token: "e87b2ef4-1a75-5dfd-814e-ba5868b8fa4d",
				Jwt:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiUG9zdG1hbiIsImV4cCI6MTcwODM5NTAzMywiaWQiOjEsIm5hbWUiOiJDYWlvIn0.zNN7G5zwltkKCy01I79bMvmppB8I3IsyW5_ZUvB2ric",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOpaqueTokenWithJwt(tt.args.token, tt.args.jwt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOpaqueTokenWithJwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOpaqueTokenWithJwt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
