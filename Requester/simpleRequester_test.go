package Requester

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	type args struct {
		method string
		url    string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "simpleExample",
			args: args{
				`GET`,
				`https://www.baidu.com`,
			},
			want:    []byte{0012},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fetch(tt.args.method, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(got))
		})
	}
}
