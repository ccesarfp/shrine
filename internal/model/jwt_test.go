package model

import (
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"testing"
)

func TestNewJwt(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		{
			name: "creating_a_token_with_a_valid_jwt",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiUG9zdG1hbiIsImV4cCI6MTcwNTQzNjc2MSwiaWQiOjEsIm5hbWUiOiJDYWlvIn0.GP0aS6gBGV3Wnb3s5pp5_1-7dvjDa-cEyL_YOYUPAGA",
			},
			want: &Token{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiUG9zdG1hbiIsImV4cCI6MTcwNTQzNjc2MSwiaWQiOjEsIm5hbWUiOiJDYWlvIn0.GP0aS6gBGV3Wnb3s5pp5_1-7dvjDa-cEyL_YOYUPAGA",
			},
			wantErr: false,
		},
		{
			name: "creating_a_token_with_a_invalid_jwt",
			args: args{
				token: "aaaa",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJwtWithId1(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		{
			name: "creating_token_using_valid_id",
			args: args{
				id: "1-Shrine",
			},
			want: &Token{
				Id: "1-Shrine",
			},
			wantErr: false,
		},
		{
			name: "creating_token_using_number_only",
			args: args{
				id: "1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creating_token_using_name_only",
			args: args{
				id: "Shrine",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creating_token_using_dash_only",
			args: args{
				id: "-",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTokenWithId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenWithId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenWithId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtValidity(t1 *testing.T) {
	t1.Setenv("jwtSecretKey", "ItsASecret")
	type fields struct {
		Id    string
		Token string
	}
	type args struct {
		jwtSecretKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "checking_valid_token",
			fields: fields{
				Id:    "1-Postman",
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiUG9zdG1hbiIsImV4cCI6MTcwNTQ2NzQ2MSwiaWQiOjEsIm5hbWUiOiJDYWlvIn0.v7HtG4G-V647nmZ7hJzxPXAKSnYx1-7k4YIZj_0gT3M",
			},
			args: args{
				jwtSecretKey: "ItsASecret",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				Id:    tt.fields.Id,
				Token: tt.fields.Token,
			}
			got, err := t.CheckValidity(tt.args.jwtSecretKey)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CheckValidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("CheckValidity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtCreateToken(t1 *testing.T) {
	type fields struct {
		Id    string
		Token string
	}
	type args struct {
		claims       jwt.MapClaims
		jwtSecretKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "creating_token_with_all_args",
			fields: fields{},
			args: args{claims: jwt.MapClaims{
				"id":          1,
				"name":        "Caio",
				"appOrigin":   "Test",
				"accessLevel": 1,
				"exp":         99999999,
			}},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6MSwiYXBwT3JpZ2luIjoiVGVzdCIsImV4cCI6OTk5OTk5OTksImlkIjoxLCJuYW1lIjoiQ2FpbyJ9.BOwI0zuG5zlKQ5FUWUtECbE6uECH4Gkthp7yxBiYEa0",
			wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				Id:    tt.fields.Id,
				Token: tt.fields.Token,
			}
			got, err := t.CreateToken(tt.args.claims, tt.args.jwtSecretKey)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
