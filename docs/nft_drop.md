
## NFT Drop

You can access the NFT Drop interface from the SDK as follows:

```
import (
	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

privateKey = "..."

sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetNFTDrop("{{contract_address}}")
```

```go
type NFTDrop struct {
    Abi *abi.DropERC721

    *ERC721
    ClaimConditions *NFTDropClaimConditions
    Encoder         *NFTDropEncoder
    Events          *ContractEvents
}
```

### func \(\*NFTDrop\) [Claim](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L285>)

```go
func (drop *NFTDrop) Claim(ctx context.Context, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*NFTDrop\) [ClaimTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L304>)

```go
func (drop *NFTDrop) ClaimTo(ctx context.Context, destinationAddress string, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

#### Example

```
address := "{{wallet_address}}"
quantity = 1

tx, err := contract.ClaimTo(context.Background(), address, quantity)
```

### func \(\*NFTDrop\) [CreateBatch](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L238>)

```go
func (drop *NFTDrop) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Create a batch of NFTs on this contract\.

metadatas: a list of the metadatas of the NFTs to create

returns: the transaction receipt of the batch creation

#### Example

```
image0, err := os.Open("path/to/image/0.jpg")
defer image0.Close()

image1, err := os.Open("path/to/image/1.jpg")
defer image1.Close()

metadatas := []*web3sdks.NFTMetadataInput{
	&web3sdks.NFTMetadataInput{
		Name: "Cool NFT",
		Description: "This is a cool NFT",
		Image: image1
	}
	&web3sdks.NFTMetadataInput{
		Name: "Cool NFT 2",
		Description: "This is also a cool NFT",
		Image: image2
	}
}

tx, err := contract.CreateBatch(context.Background(), metadatas)
```

### func \(\*NFTDrop\) [GetAllClaimed](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L140>)

```go
func (drop *NFTDrop) GetAllClaimed(ctx context.Context) ([]*NFTMetadataOwner, error)
```

Get a list of all the NFTs that have been claimed from this contract\.

returns: a list of the metadatas of the claimed NFTs

#### Example

```
claimedNfts, err := contract.GetAllClaimed(context.Background())
firstOwner := claimedNfts[0].Owner
```

### func \(\*NFTDrop\) [GetAllUnclaimed](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L164>)

```go
func (drop *NFTDrop) GetAllUnclaimed(ctx context.Context) ([]*NFTMetadata, error)
```

Get a list of all the NFTs on this contract that have not yet been claimed\.

returns: a list of the metadatas of the unclaimed NFTs

#### Example

```
unclaimedNfts, err := contract.GetAllUnclaimed(context.Background())
firstNftName := unclaimedNfts[0].Name
```

### func \(\*NFTDrop\) [GetClaimArguments](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L345-L352>)

```go
func (drop *NFTDrop) GetClaimArguments(ctx context.Context, destinationAddress string, quantity int) (*ClaimArguments, error)
```

### func \(\*NFTDrop\) [GetOwned](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L89>)

```go
func (nft *NFTDrop) GetOwned(ctx context.Context, address string) ([]*NFTMetadataOwner, error)
```

Get the metadatas of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the metadata of all the NFTs owned by the address

#### Example

```
owner := "{{wallet_address}}"
nfts, err := contract.GetOwned(context.Background(), owner)
name := nfts[0].Metadata.Name
```

### func \(\*NFTDrop\) [GetOwnedTokenIDs](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L106>)

```go
func (nft *NFTDrop) GetOwnedTokenIDs(ctx context.Context, address string) ([]*big.Int, error)
```

Get the tokenIds of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the tokenIds of all the NFTs owned by the address

### func \(\*NFTDrop\) [TotalClaimedSupply](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L185>)

```go
func (drop *NFTDrop) TotalClaimedSupply() (int, error)
```

Get the total number of NFTs that have been claimed\.

### func \(\*NFTDrop\) [TotalUnclaimedSupply](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/nft_drop.go#L195>)

```go
func (drop *NFTDrop) TotalUnclaimedSupply() (int, error)
```

Get the total number of NFTs that have not yet been claimed\.
