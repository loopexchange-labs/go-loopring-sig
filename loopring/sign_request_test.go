package loopring

import (
	"testing"
)

func TestSignRequest(t *testing.T) {
	type args struct {
		privateKey string
		method     string
		baseUrl    string
		path       string
		data       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getUserApiKey",
			args: args{
				privateKey: "0x1234",
				method:     "GET",
				baseUrl:    "https://api3.loopring.io",
				path:       "/api/v3/apiKey",
				data:       "accountId%3D12345",
			},
			want:    "0x0b29b2bacc6718d86e55822e689735b11c78785846045644dfc9875fd1eff63b1497bfb8ffb3dc24023d93877ea003602c4d45664c14322047e88404a3f5770009694914fb154906f3819fdb469f02920d6f83b58d1348b5cdbb1ae98f211562",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pk := NewPrivateKeyFromString(tt.args.privateKey)
			got, err := SignRequest(pk, tt.args.method, tt.args.baseUrl, tt.args.path, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
