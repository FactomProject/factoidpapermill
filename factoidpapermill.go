package main

import (
	"fmt"
	"github.com/FactomProject/factoid"
	"github.com/FactomProject/factoid/wallet"
	"crypto/rand"
)

func main() {

	privatekey := make([]byte, 32)
	_, err := rand.Read(privatekey)
	if err == nil {
		privAddr := factoid.NewAddress(privatekey)
		
		privHuman := factoid.ConvertFctPrivateToUserStr(privAddr)
		fmt.Printf("New Factoid Private Key: %v\n", privHuman)

		pub, priv, err := wallet.GenerateKeyFromPrivateKey(privatekey)
		if err != nil {
			panic(err)
		}

		we := new(wallet.WalletEntry)
		we.AddKey(pub, priv)
		we.SetName([]byte("test"))
		we.SetRCD(factoid.NewRCD_1(pub))
		we.SetType("fct")

		address, _ := we.GetAddress()

		adr := factoid.ConvertFctAddressToUserStr(address)

		fmt.Printf("New Factoid Address:     %v\n", adr)
	} else {
		fmt.Printf("error\n")
	}
	

}
