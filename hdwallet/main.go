package main

import (
	"fmt"

	"github.com/bcext/cashutil/hdkeychain"
	"github.com/bcext/gcash/chaincfg"
)

func main() {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		panic(err)
	}

	key, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}

	keyStr := key.String()
	// output: tprv8ZgxMBicQKsPdFH9KwqSxs4mhxxKQ1fhAXvTHsTrztRzn4MF42ySDtS1AkTciqWY6FRfQy8pnBrrmJ1CbcQXZ7iaiWiXiM8W24K3KabEYtY
	fmt.Println(keyStr)

	generatedKey, err := hdkeychain.NewKeyFromString(keyStr)
	if err != nil {
		panic(err)
	}

	child0, err := generatedKey.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		panic(err)
	}

	child0Address, err := child0.Address(&chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}

	fmt.Println("child0 address:", child0Address.EncodeAddress(true))

	child00Extend, err := child0.Child(0)
	if err != nil {
		panic(err)
	}
	child00ExtendAddress, err := child00Extend.Address(&chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}
	fmt.Println("child0's extended address:", child00ExtendAddress.EncodeAddress(true))

	child01Internal, err := child0.Child(1)
	if err != nil {
		panic(err)
	}
	// 可以使用主网编码地址
	child01InternalAddress, err := child01Internal.Address(&chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("child0's internal address:", child01InternalAddress.EncodeAddress(true))

	fmt.Println("child00Extended is privkey?", child00Extend.IsPrivate())
	fmt.Println("child01Internal is privkey?", child01Internal.IsPrivate())

	extenedKey, err := child01Internal.Neuter()
	if err != nil {
		panic(err)
	}
	fmt.Println("extendedKey string:", extenedKey.String())
}
