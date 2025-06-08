package mainform

import (
	"fmt"
	"time"

	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00client/u00client"
	"github.com/u00io/u00node/localstorage"
)

func Run() {
	var privateKey []byte

	bs, err := localstorage.Read("private.key")
	if err != nil || len(bs) < 64 {
		privateKey, _ = u00client.GenerateKeyPair()
		localstorage.Write("private.key", privateKey)
	} else {
		privateKey = bs
	}

	client := u00client.NewClientWithKey(privateKey)
	fmt.Println("Client Address:", "https://u00.io/native/"+client.Address())

	c := ui.NewForm()
	c.SetTitle("U00 Node")
	c.SetSize(800, 600)

	btnSend := ui.NewButton()
	btnSend.SetPosition(10, 10)
	btnSend.SetSize(100, 30)
	btnSend.SetProp("text", "Send")
	btnSend.SetProp("onClick", func() {
		client.WriteValue("MyItem", time.Now(), time.Now().Format("2006-01-02 15:04:05.000000"))
	})
	c.Panel().AddWidget(btnSend)
	c.Exec()
}
