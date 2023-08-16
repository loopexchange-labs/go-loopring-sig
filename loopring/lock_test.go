package loopring

import "testing"

func TestGetLockHashAndEddsaSignature(t *testing.T) {
	type args struct {
		privateKey      string
		exchangeAddress string
		accountId       string
		tokenId         string
		volume          string
		timestamp       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			// Happy path
			name: "Happy path",
			args: args{
				privateKey:      "0x1234",
				exchangeAddress: "0x2345",
				accountId:       "1",
				tokenId:         "2",
				volume:          "3",
				timestamp:       "4",
			},
			want:    "0x021711906668a2a79c71b30fe72bff29bbd55c8a974bb6acf1610437a2f06310205a66cc59027eb27491cf1b481a7d4955e1b69bf43ba93e95e52bee67b43f3a1069ebe2b2e49723c9e0344fae8e13b19e5e26f0bc2ec81286fe0ad791898a77",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLockHashAndEddsaSignature(tt.args.privateKey, tt.args.exchangeAddress, tt.args.accountId, tt.args.tokenId, tt.args.volume, tt.args.timestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLockHashAndEddsaSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLockHashAndEddsaSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
