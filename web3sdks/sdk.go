package web3sdks

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Web3sdksSDK struct {
	*ProviderHandler
	Storage  IpfsStorage
	Deployer ContractDeployer
	Auth     WalletAuthenticator
}

// NewWeb3sdksSDK
//
// Create a new instance of the Web3sdks SDK
//
// rpcUrlOrName: the name of the chain to connection to (e.g. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche") or the RPC URL to connect to
//
// options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL
func NewWeb3sdksSDK(rpcUrlOrChainName string, options *SDKOptions) (*Web3sdksSDK, error) {
	rpc, err := getDefaultRpcUrl(rpcUrlOrChainName)
	if err != nil {
		return nil, err
	}

	provider, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	// Define defaults for all the options
	privateKey := ""
	gatewayUrl := defaultIpfsGatewayUrl

	// Override defaults with the options that are defined
	if options != nil {
		if options.PrivateKey != "" {
			privateKey = options.PrivateKey
		}

		if options.GatewayUrl != "" {
			gatewayUrl = options.GatewayUrl
		}
	}

	storage := newIpfsStorage(gatewayUrl)

	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	deployer, err := newContractDeployer(provider, privateKey, storage)
	if err != nil {
		return nil, err
	}

	auth, err := newWalletAuthenticator(provider, privateKey)
	if err != nil {
		return nil, err
	}

	sdk := &Web3sdksSDK{
		ProviderHandler: handler,
		Storage:         *storage,
		Deployer:        *deployer,
		Auth:            *auth,
	}

	return sdk, nil
}

// GetNFTCollection
//
// Get an NFT Collection contract SDK instance
//
// address: the address of the NFT Collection contract
func (sdk *Web3sdksSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	return newNFTCollection(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEdition
//
// Get an Edition contract SDK instance
//
// address: the address of the Edition contract
func (sdk *Web3sdksSDK) GetEdition(address string) (*Edition, error) {
	return newEdition(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetToken
//
// Returns a Token contract SDK instance
//
// address: address of the token contract
//
// Returns a Token contract SDK instance
func (sdk *Web3sdksSDK) GetToken(address string) (*Token, error) {
	return newToken(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetNFTDrop
//
// Get an NFT Drop contract SDK instance
//
// address: the address of the NFT Drop contract
func (sdk *Web3sdksSDK) GetNFTDrop(address string) (*NFTDrop, error) {
	return newNFTDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEditionDrop
//
// Get an Edition Drop contract SDK instance
//
// address: the address of the Edition Drop contract
func (sdk *Web3sdksSDK) GetEditionDrop(address string) (*EditionDrop, error) {
	return newEditionDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMultiwrap
//
// Get a Multiwrap contract SDK instance
//
// address: the address of the Multiwrap contract
func (sdk *Web3sdksSDK) GetMultiwrap(address string) (*Multiwrap, error) {
	return newMultiwrap(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMarketplace
//
// Get a Marketplace contract SDK instance
//
// address: the address of the Marketplace contract
func (sdk *Web3sdksSDK) GetMarketplace(address string) (*Marketplace, error) {
	return newMarketplace(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetContract
//
// Get an instance of a custom contract deployed with web3sdks deploy
//
// address: the address of the contract
func (sdk *Web3sdksSDK) GetContract(address string) (*SmartContract, error) {
	abi, err := fetchContractMetadataFromAddress(address, sdk.GetProvider(), &sdk.Storage)
	if err != nil {
		return nil, err
	}

	return sdk.GetContractFromAbi(address, abi)
}

// GetContractFromABI
//
// Get an instance of ant custom contract from its ABI
//
// address: the address of the contract
//
// abi: the ABI of the contract
func (sdk *Web3sdksSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error) {
	return newSmartContract(
		sdk.GetProvider(),
		common.HexToAddress(address),
		abi,
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

func getDefaultRpcUrl(rpcUrlorName string) (string, error) {
	switch rpcUrlorName {
	case "mumbai":
		return "https://polygon-mumbai.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "rinkeby":
		return "https://eth-rinkeby.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "goerli":
		return "https://eth-goerli.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "polygon":
		return "https://polygon-mainnet.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "mainnet", "ethereum":
		return "https://eth-mainnet.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "fantom":
		return "https://rpc.ftm.tools", nil
	case "avalanche":
		return "https://rpc.ankr.com/avalanche", nil
	default:
		if strings.HasPrefix(rpcUrlorName, "http") {
			return rpcUrlorName, nil
		} else {
			return "", fmt.Errorf("invalid rpc url or chain name: %s", rpcUrlorName)
		}
	}
}
