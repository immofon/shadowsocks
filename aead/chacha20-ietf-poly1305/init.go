package chacha20_ietf_poly1305

import (
	"github.com/immofon/shadowsocks/aead"
	"golang.org/x/crypto/chacha20poly1305"
)

func init() {
	aead.RegisterCipher("chacha20-ietf-poly1305", 32, chacha20poly1305.New)
}
