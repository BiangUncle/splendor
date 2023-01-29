package main

import (
	"net/http"
	"testing"
)

func TestClient_ConstructURL(t *testing.T) {
	type fields struct {
		client  *http.Client
		Address string
	}
	type args struct {
		uri  string
		args map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"test",
			fields{
				nil,
				"127.0.0.1:8765",
			},
			args{
				"join",
				make(map[string]any, 0),
			},
			"http://127.0.0.1:8765/join",
		},
		{
			"test",
			fields{
				nil,
				"127.0.0.1:8765",
			},
			args{
				"join",
				map[string]any{
					"username": "biang",
				},
			},
			"http://127.0.0.1:8765/join?username=biang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				Address: tt.fields.Address,
			}
			if got := c.ConstructURL(tt.args.uri, tt.args.args); got != tt.want {
				t.Errorf("ConstructURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
