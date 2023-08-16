package loopring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPair(t *testing.T) {
	type args struct {
		signature string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Happy path",
			args: args{
				signature: "0x1234",
			},
			want: map[string]interface{}{
				"keyPair": map[string]interface{}{
					"publicKeyX": "12485317073945829795163901166283443631921547304790744867554182714098609949000",
					"publicKeyY": "12527734297243956538968023182262726317403413118463934434378426571924900188005",
					"secretKey":  "2280752172643637060311335108033819750862509663263550628348196500509381754787",
				},
				"formatedPx": "0x1b9a7064d6f1b27979170f9a5ce2fa32cefb77e2ac81583404e544943d620948",
				"formatedPy": "0x1bb27243185bc780f46f964ad2011d55d0ff3ed0177b62bbe400d88ce9bc8f65",
				"sk":         "0x50adc27dea06292802e0696cecf816f4d141ad9fa80a877948b0251be51ffa3",
			},
		},
		{
			name: "Happy path",
			args: args{
				signature: "0x12345",
			},
			want: map[string]interface{}{
				"keyPair": map[string]interface{}{
					"publicKeyX": "17150263012821040753302789854468952052284392848527137254990915027830228822755",
					"publicKeyY": "4967225526344074395340056881251792399893955194356541630701296482692946284746",
					"secretKey":  "1211277797918594114398771323365461492872742227233686142033366613988626417152",
				},
				"formatedPx": "0x25eab47287d73b063cefe42d3fc585f4006132133635604e6f87599ba4daaae3",
				"formatedPy": "0x0afb599abdde1699590e16c4097fbff168b8e82e5db4fb3e39e07d55d2b150ca",
				"sk":         "0x2ad8f1695121d2e424c19cd7da6281b35d5933fe4f800b29447dccc58684600",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateKeyPair(tt.args.signature)

			assert.Equal(t, tt.want, got, "GenerateKeyPair() = %v, want %v", got, tt.want)
		})
	}
}
