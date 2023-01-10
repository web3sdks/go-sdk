
## NFT Collection

You can access the NFT Collection interface from the SDK as follows:

```
import (
	"github.com/web3sdks/go-sdk/web3sdks"
)

privateKey = "..."

sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetNFTCollection("{{contract_address}}")
```

```go
type NFTCollection struct {
    *ERC721
    Signature *ERC721SignatureMinting
    Encoder   *ContractEncoder
}
```

### func \(\*NFTCollection\) [GetOwned](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_collection.go#L78>)

```go
func (nft *NFTCollection) GetOwned(address string) ([]*NFTMetadataOwner, error)
```

Get the metadatas of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the metadata of all the NFTs owned by the address

#### Example

```
owner := "{{wallet_address}}"
nfts, err := contract.GetOwned(owner)
name := nfts[0].Metadata.Name
```

### func \(\*NFTCollection\) [GetOwnedTokenIDs](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_collection.go#L95>)

```go
func (nft *NFTCollection) GetOwnedTokenIDs(address string) ([]*big.Int, error)
```

Get the tokenIds of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the tokenIds of all the NFTs owned by the address

### func \(\*NFTCollection\) [Mint](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_collection.go#L120>)

```go
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the connected wallet\.

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatch](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_collection.go#L172>)

```go
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Mint a batch of new NFTs to the connected wallet\.

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatchTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_collection.go#L199>)

```go
func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Mint a batch of new NFTs to the specified wallet\.

to: the wallet address to mint to

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

#### Example

```
metadatas := []*web3sdks.NFTMetadataInput{
	&web3sdks.NFTMetadataInput{
		Name: "Cool NFT",
		Description: "This is a cool NFT",
	}
	&web3sdks.NFTMetadataInput{
		Name: "Cool NFT 2",
		Description: "This is also a cool NFT",
	}
}

tx, err := contract.MintBatchTo("{{wallet_address}}", metadatas)
```

### func \(\*NFTCollection\) [MintTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_collection.go#L145>)

```go
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the specified wallet\.

address: the wallet address to mint to

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

#### Example

```
image, err := os.Open("path/to/image.jpg")
defer image.Close()

metadata := &web3sdks.NFTMetadataInput{
	Name: "Cool NFT",
	Description: "This is a cool NFT",
	Image: image,
}

tx, err := contract.MintTo("{{wallet_address}}", metadata)
```
