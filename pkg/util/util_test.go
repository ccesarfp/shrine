package util

import "testing"

func TestPrepareKey(t *testing.T) {
	type args struct {
		id      int64
		appName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "creating_key", args: args{id: 1, appName: "Test"}, want: "1-test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrepareKey(tt.args.id, tt.args.appName); got != tt.want {
				t.Errorf("PrepareKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateUsingRegex(t *testing.T) {
	type args struct {
		pattern string
		value   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "validating_value", args: args{pattern: "^\\d+-[A-Za-z]+$", value: "1-test"}, want: true, wantErr: false},
		{name: "validating_invalid_value", args: args{pattern: "^\\d+-[A-Za-z]+$", value: "1"}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateUsingRegex(tt.args.pattern, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUsingRegex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateUsingRegex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
