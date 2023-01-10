package main

import (
	"log"

	"github.com/web3sdks/go-sdk/web3sdks"
)

var (
	web3sdksSDK *web3sdks.Web3sdksSDK
)

func initSdk() {
	if sdk, err := web3sdks.NewWeb3sdksSDK(
		chainRpcUrl,
		&web3sdks.SDKOptions{
			PrivateKey: privateKey,
		},
	); err != nil {
		panic(err)
	} else {
		web3sdksSDK = sdk
	}
}

func getNftCollection() (*web3sdks.NFTCollection, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Collection on chain %v, contract %v\n", chainRpcUrl, nftContractAddress)

	if contract, err := web3sdksSDK.GetNFTCollection(nftContractAddress); err != nil {
		log.Println("Failed to create an NFT Collection object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEdition() (*web3sdks.Edition, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition on chain %v, contract %v\n", chainRpcUrl, editionAddress)

	if contract, err := web3sdksSDK.GetEdition(editionAddress); err != nil {
		log.Println("Failed to create an Edition object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getToken() (*web3sdks.Token, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Token on chain %v, contract %v\n", chainRpcUrl, tokenAddress)

	if contract, err := web3sdksSDK.GetToken(tokenAddress); err != nil {
		log.Println("Failed to create an Token object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getNftDrop() (*web3sdks.NFTDrop, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Drop on chain %v, contract %v\n", chainRpcUrl, nftDropContractAddress)

	if contract, err := web3sdksSDK.GetNFTDrop(nftDropContractAddress); err != nil {
		log.Println("Failed to create an NFT Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEditionDrop() (*web3sdks.EditionDrop, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition Drop on chain %v, contract %v\n", chainRpcUrl, editionDropContractAddress)

	if contract, err := web3sdksSDK.GetEditionDrop(editionDropContractAddress); err != nil {
		log.Println("Failed to create an Edition Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMultiwrap() (*web3sdks.Multiwrap, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Multiwrap on chain %v, contract %v\n", chainRpcUrl, multiwrapContractAddress)

	if contract, err := web3sdksSDK.GetMultiwrap(multiwrapContractAddress); err != nil {
		log.Println("Failed to create a Multiwrap object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMarketplace() (*web3sdks.Marketplace, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Marketplace on chain %v, contract %v\n", chainRpcUrl, marketplaceAddress)

	if contract, err := web3sdksSDK.GetMarketplace(marketplaceAddress); err != nil {
		log.Println("Failed to create a Marketplace object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getCustom() (*web3sdks.SmartContract, error) {
	if web3sdksSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Custom on chain %v, contract %v\n", chainRpcUrl, customContractAddress)

	if contract, err := web3sdksSDK.GetContract(customContractAddress); err != nil {
		log.Println("Failed to create an Custom object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getStorage() web3sdks.IpfsStorage {
	if web3sdksSDK == nil {
		initSdk()
	}

	return web3sdksSDK.Storage
}
