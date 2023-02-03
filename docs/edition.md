
## Edition

You can access the Edition interface from the SDK as follows:

```
import (
	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

privateKey = "..."

sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetEdition("{{contract_address}}")
```

```go
type Edition struct {
    *ERC1155
    Signature *ERC1155SignatureMinting
    Encoder   *ContractEncoder
    Events    *ContractEvents
}
```

### func \(\*Edition\) [Mint](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/edition.go#L80>)

```go
func (edition *Edition) Mint(ctx context.Context, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error)
```

Mint an NFT to the connected wallet\.

metadataWithSupply: nft metadata with supply of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintAdditionalSupply](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/edition.go#L141>)

```go
func (edition *Edition) MintAdditionalSupply(ctx context.Context, tokenId int, additionalSupply int) (*types.Transaction, error)
```

Mint additionaly supply of a token to the connected wallet\.

tokenId: token ID to mint additional supply of

additionalSupply: additional supply to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintAdditionalSupplyTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/edition.go#L155>)

```go
func (edition *Edition) MintAdditionalSupplyTo(ctx context.Context, to string, tokenId int, additionalSupply int) (*types.Transaction, error)
```

Mint additional supply of a token to the specified wallet\.

to: address of the wallet to mint NFTs to

tokenId: token Id to mint additional supply of

additionalySupply: additional supply to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintBatch](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/edition.go#L184>)

```go
func (edition *Edition) MintBatch(ctx context.Context, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error)
```

Mint a batch of NFTs to the connected wallet\.

metadatasWithSupply: list of NFT metadatas with supplies to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintBatchTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/edition.go#L216>)

```go
func (edition *Edition) MintBatchTo(ctx context.Context, to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error)
```

Mint a batch of NFTs to a specific wallet\.

to: address of the wallet to mint NFTs to

metadatasWithSupply: list of NFT metadatas with supplies to mint

returns: the transaction receipt of the mint

#### Example

```
metadatasWithSupply := []*web3sdks.EditionMetadataInput{
	&web3sdks.EditionMetadataInput{
		Metadata: &web3sdks.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
		},
		Supply: 100,
	},
	&web3sdks.EditionMetadataInput{
		Metadata: &web3sdks.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
		},
		Supply: 100,
	},
}

tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatasWithSupply)
```

### func \(\*Edition\) [MintTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/edition.go#L109>)

```go
func (edition *Edition) MintTo(ctx context.Context, address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the specified wallet\.

address: the wallet address to mint the NFT to

metadataWithSupply: nft metadata with supply of the NFT to mint

returns: the transaction receipt of the mint

#### Example

```
image, err := os.Open("path/to/image.jpg")
	defer image.Close()

	metadataWithSupply := &web3sdks.EditionMetadataInput{
        context.Background(),
		Metadata: &web3sdks.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
			Image: image,
		},
		Supply: 100,
	}

	tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", metadataWithSupply)
```
