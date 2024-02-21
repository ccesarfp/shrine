package model

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ipAddress     string
		hoursToExpire int32
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "creating_user_with_all_args",
			args: args{
				ipAddress:     "127.0.0.1",
				hoursToExpire: 4,
			},
			want: &User{
				IpAddress:     "127.0.0.1",
				HoursToExpire: 4,
			},
			wantErr: false,
		},
		{
			name:    "creating_user_with_none_args",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.ipAddress, tt.args.hoursToExpire)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
