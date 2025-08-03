package utils

import (
	"crypto/ed25519"
	"encoding/hex"
)

type Key struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewKeyFromPrivate(privateKey []byte) *Key {
	var c Key
	c.PrivateKey = privateKey
	return &c
}

func NewKey() *Key {
	var k Key
	pubKey, privKey, _ := ed25519.GenerateKey(nil)
	k.PrivateKey = privKey
	k.PublicKey = pubKey
	return &k
}

func (c *Key) GetPrivateKey() []byte {
	return c.PrivateKey
}

func (c *Key) GetPublicKey() []byte {
	return c.PublicKey
}

func (c *Key) SetPrivateKey(privateKey []byte) {
	c.PrivateKey = privateKey
}

func (c *Key) SetPublicKey(publicKey []byte) {
	c.PublicKey = publicKey
}

func (c *Key) String() string {
	return hex.EncodeToString(c.PublicKey)
}
