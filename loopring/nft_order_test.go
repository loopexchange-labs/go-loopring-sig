package loopring

import "testing"

func TestGetEddsaSigNftOrder(t *testing.T) {
	type args struct {
		privateKey      string
		exchangeAddress string
		storageId       string
		accountId       string
		sellTokenId     string
		buyTokenId      string
		sellTokenAmount string
		buyTokenAmount  string
		validUntil      string
		maxFeeBips      string
		fillAmountBOrS  string
		takerAddress    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// Happy path
		{
			name: "Happy path",
			args: args{
				privateKey:      "0x1234",
				exchangeAddress: "0x2345",
				storageId:       "1",
				accountId:       "2",
				sellTokenId:     "3",
				buyTokenId:      "4",
				sellTokenAmount: "5",
				buyTokenAmount:  "6",
				validUntil:      "7",
				maxFeeBips:      "8",
				fillAmountBOrS:  "9",
				takerAddress:    "10",
			},
			want:    "0x006fd1c67c3824dbb82820610feec07b03f24a4546aa2487304f8791e4449d152160cfbce988ccdf274b3f307d2b65b55d699eaa9de7cd8e62192e55d7ddf8a02477e88cad04799c1f6af6c1635710417e3c18a5571f6b495bfd9fc13722ca63",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEddsaSigNftOrder(tt.args.privateKey, tt.args.exchangeAddress, tt.args.storageId, tt.args.accountId, tt.args.sellTokenId, tt.args.buyTokenId, tt.args.sellTokenAmount, tt.args.buyTokenAmount, tt.args.validUntil, tt.args.maxFeeBips, tt.args.fillAmountBOrS, tt.args.takerAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEddsaSigNftOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEddsaSigNftOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
