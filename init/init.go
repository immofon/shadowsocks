package init

import (
	_ "github.com/immofon/shadowsocks/aead/aes-x-gcm"
	_ "github.com/immofon/shadowsocks/aead/chacha20-ietf-poly1305"
	_ "github.com/immofon/shadowsocks/dummy"
	_ "github.com/immofon/shadowsocks/stream/aes-x-cfb"
	_ "github.com/immofon/shadowsocks/stream/aes-x-ctr"
	_ "github.com/immofon/shadowsocks/stream/bf-cfb"
	_ "github.com/immofon/shadowsocks/stream/cast5-cfb"
	_ "github.com/immofon/shadowsocks/stream/chacha20"
	_ "github.com/immofon/shadowsocks/stream/des-cfb"
	_ "github.com/immofon/shadowsocks/stream/rc4-md5-x"
	_ "github.com/immofon/shadowsocks/stream/salsa20"
)
