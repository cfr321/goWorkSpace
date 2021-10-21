package main

import (
	"fmt"
	"github.com/chenfar/npaillier"
	"math/big"
)

func main() {
	privateKey := npaillier.GenkeyPair(1024, 3)

	pubkey := privateKey.GetPk()
	//pubkey2 := privateKey.sks[0].Pk

	c1 := pubkey.Encrypt(new(big.Int).SetInt64(123))
	c2 := pubkey.Encrypt(new(big.Int).SetInt64(321))
	c := pubkey.Add(c2, c1)
	fmt.Println(privateKey.DecryptByx(c))
	fmt.Println(privateKey.Decrypt(c.T1))
	fmt.Println(privateKey.DecryptBy2(c.T1))
}
