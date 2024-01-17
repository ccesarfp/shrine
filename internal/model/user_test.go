package model

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id            int64
		name          string
		appOrigin     string
		accessLevel   int32
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
				id:            1,
				name:          "Caio",
				appOrigin:     "Test",
				accessLevel:   1,
				hoursToExpire: 4,
			},
			want: &User{
				Id:            1,
				Name:          "Caio",
				AppOrigin:     "Test",
				AccessLevel:   1,
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
			got, err := NewUser(tt.args.id, tt.args.name, tt.args.appOrigin, tt.args.accessLevel, tt.args.hoursToExpire)
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
