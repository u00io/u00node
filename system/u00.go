package system

import (
	"fmt"
	"time"

	"github.com/u00io/u00client/u00client"
	"github.com/u00io/u00node/localstorage"
)

type U00 struct {
	client *u00client.U00Client
}

func NewU00() *U00 {
	var u U00
	return &u
}

func (c *U00) Run() {
	var privateKey []byte
	bs, err := localstorage.Read("private.key")
	if err != nil || len(bs) < 64 {
		privateKey, _ = u00client.GenerateKeyPair()
		localstorage.Write("private.key", privateKey)
	} else {
		privateKey = bs
	}
	c.client = u00client.NewClientWithKey(privateKey)
	fmt.Println("Client Address:", "https://u00.io/native/"+c.client.Address())
}

func (c *U00) WriteValue(value string) {
	c.client.WriteValue("MyItem", time.Now(), value)
}
