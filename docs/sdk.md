
## Web3sdksSDK

```go
type Web3sdksSDK struct {
    *ProviderHandler
    Storage  IpfsStorage
    Deployer ContractDeployer
    Auth     WalletAuthenticator
}
```

### func [NewWeb3sdksSDK](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L25>)

```go
func NewWeb3sdksSDK(rpcUrlOrChainName string, options *SDKOptions) (*Web3sdksSDK, error)
```

#### NewWeb3sdksSDK

Create a new instance of the Web3sdks SDK

rpcUrlOrName: the name of the chain to connection to \(e\.g\. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche"\) or the RPC URL to connect to

options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL

### func \(\*Web3sdksSDK\) [GetContract](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L183>)

```go
func (sdk *Web3sdksSDK) GetContract(address string) (*SmartContract, error)
```

#### GetContract

Get an instance of a custom contract deployed with web3sdks deploy

address: the address of the contract

### func \(\*Web3sdksSDK\) [GetContractFromAbi](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L199>)

```go
func (sdk *Web3sdksSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error)
```

#### GetContractFromABI

Get an instance of ant custom contract from its ABI

address: the address of the contract

abi: the ABI of the contract

### func \(\*Web3sdksSDK\) [GetEdition](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L97>)

```go
func (sdk *Web3sdksSDK) GetEdition(address string) (*Edition, error)
```

#### GetEdition

Get an Edition contract SDK instance

address: the address of the Edition contract

### func \(\*Web3sdksSDK\) [GetEditionDrop](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L141>)

```go
func (sdk *Web3sdksSDK) GetEditionDrop(address string) (*EditionDrop, error)
```

#### GetEditionDrop

Get an Edition Drop contract SDK instance

address: the address of the Edition Drop contract

### func \(\*Web3sdksSDK\) [GetMarketplace](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L169>)

```go
func (sdk *Web3sdksSDK) GetMarketplace(address string) (*Marketplace, error)
```

#### GetMarketplace

Get a Marketplace contract SDK instance

address: the address of the Marketplace contract

### func \(\*Web3sdksSDK\) [GetMultiwrap](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L155>)

```go
func (sdk *Web3sdksSDK) GetMultiwrap(address string) (*Multiwrap, error)
```

#### GetMultiwrap

Get a Multiwrap contract SDK instance

address: the address of the Multiwrap contract

### func \(\*Web3sdksSDK\) [GetNFTCollection](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L83>)

```go
func (sdk *Web3sdksSDK) GetNFTCollection(address string) (*NFTCollection, error)
```

#### GetNFTCollection

Get an NFT Collection contract SDK instance

address: the address of the NFT Collection contract

### func \(\*Web3sdksSDK\) [GetNFTDrop](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L127>)

```go
func (sdk *Web3sdksSDK) GetNFTDrop(address string) (*NFTDrop, error)
```

#### GetNFTDrop

Get an NFT Drop contract SDK instance

address: the address of the NFT Drop contract

### func \(\*Web3sdksSDK\) [GetToken](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/sdk.go#L113>)

```go
func (sdk *Web3sdksSDK) GetToken(address string) (*Token, error)
```

#### GetToken

Returns a Token contract SDK instance

address: address of the token contract

Returns a Token contract SDK instance
