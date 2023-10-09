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
		name  string
		args  args
		want  interface{}
		want2 KeyPairFormatted
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
			want2: KeyPairFormatted{
				PublicKeyX: "0x1b9a7064d6f1b27979170f9a5ce2fa32cefb77e2ac81583404e544943d620948",
				PublicKeyY: "0x1bb27243185bc780f46f964ad2011d55d0ff3ed0177b62bbe400d88ce9bc8f65",
				SecretKey:  "0x50adc27dea06292802e0696cecf816f4d141ad9fa80a877948b0251be51ffa3",
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
			want2: KeyPairFormatted{
				PublicKeyX: "0x25eab47287d73b063cefe42d3fc585f4006132133635604e6f87599ba4daaae3",
				PublicKeyY: "0x0afb599abdde1699590e16c4097fbff168b8e82e5db4fb3e39e07d55d2b150ca",
				SecretKey:  "0x2ad8f1695121d2e424c19cd7da6281b35d5933fe4f800b29447dccc58684600",
			},
		},
		{
			name: "Happy path",
			args: args{
				signature: "0xa6bd206701d1aaa1d45d6ab870e5e17801a6d98fb4440a02bab58e7ab762a55a5495da27508c529b7bab192baf819c600ed6fc474bac0fb87a7733b6257af92f1b",
			},
			want: map[string]interface{}{
				"keyPair": map[string]interface{}{
					"publicKeyX": "11466079104134713165495087776047160054094966511533634450353147444674915986288",
					"publicKeyY": "19764600549227660523215132843133188682086068475762266060254469321575810955040",
					"secretKey":  "2649207223406745302329198464486813355387508004074877954507584610740757597633",
				},
				"formatedPx": "0x1959921f5f8d4f486b4c57afd6459444be34c75eede50ff1bd00b0333887eb70",
				"formatedPy": "0x2bb25e133b1294626e1154e662b38d9332b66927451813c1615912fd705b7720",
				"sk":         "0x5db65ed466a3b154dcf83e2e4b06b66c0c305d7d2088f9f60031567cf080dc1",
			},
			want2: KeyPairFormatted{
				PublicKeyX: "0x1959921f5f8d4f486b4c57afd6459444be34c75eede50ff1bd00b0333887eb70",
				PublicKeyY: "0x2bb25e133b1294626e1154e662b38d9332b66927451813c1615912fd705b7720",
				SecretKey:  "0x5db65ed466a3b154dcf83e2e4b06b66c0c305d7d2088f9f60031567cf080dc1",
			},
		},
		{
			name: "00 prefix",
			args: args{
				signature: "0x00c4fafca785d8b4da5e15265e4e58767a3f2502cdea2fcc4408cee5149e2a697fb7aaba14dceb096abd78cd4e9576859dcef262a4b51f2bb4554ebfd4fb72f5cf02",
			},
			want: map[string]interface{}{
				"keyPair": map[string]interface{}{
					"publicKeyX": "13177128602008476316648007186770414467529839960903001891999762356783519376339",
					"publicKeyY": "12345804151819755343032722751269003874303422211927604135345729794881772258791",
					"secretKey":  "110363126876775936939184191137732252584845853294910358210190810231166593026",
				},
				"formatedPx": "0x1d21fd9096f5e99a270d0692b2e88127669a6980029513b1ca1459bea0e423d3",
				"formatedPy": "0x1b4b7a3ef37cb12bdd53c1dd25e71c032c6a06ab1490dbafc7c304bbbf1305e7",
				"sk":         "0x3e769be3e46e6ca29d3b493f3ddfd7b66b18f404f4bd6e9f3d51a791f9e802",
			},
			want2: KeyPairFormatted{
				PublicKeyX: "0x1d21fd9096f5e99a270d0692b2e88127669a6980029513b1ca1459bea0e423d3",
				PublicKeyY: "0x1b4b7a3ef37cb12bdd53c1dd25e71c032c6a06ab1490dbafc7c304bbbf1305e7",
				SecretKey:  "0x3e769be3e46e6ca29d3b493f3ddfd7b66b18f404f4bd6e9f3d51a791f9e802",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateKeyPair(tt.args.signature)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, map[string]interface{}{
				"keyPair": map[string]interface{}{
					"publicKeyX": got.PublicKeyX.String(),
					"publicKeyY": got.PublicKeyY.String(),
					"secretKey":  got.SecretKey.String(),
				},
				"formatedPx": got.FormatPublicKeyX(),
				"formatedPy": got.FormatPublicKeyY(),
				"sk":         got.FormatSecretKey(),
			}, "GenerateKeyPair() = %v, want %v", got, tt.want)

			assert.Equal(t, tt.want2, got.ToFormatted(), "GenerateKeyPair() = %v, want %v", got.ToFormatted(), tt.want2)
		})
	}
}
