package web3sdks

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/web3sdks/go-sdk/v2/abi"
)

// You can access the NFT Collection interface from the SDK as follows:
//
//	import (
//		"github.com/web3sdks/go-sdk/v2/web3sdks"
//	)
//
//	privateKey = "..."
//
//	sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	contract, err := sdk.GetNFTCollection("{{contract_address}}")
type NFTCollection struct {
	*ERC721
	abi       *abi.TokenERC721
	Helper    *contractHelper
	Signature *ERC721SignatureMinting
	Encoder   *ContractEncoder
	Events    *ContractEvents
}

func newNFTCollection(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTCollection, error) {
	if contractAbi, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				signature, err := newERC721SignatureMinting(provider, address, privateKey, storage)
				if err != nil {
					return nil, err
				}

				encoder, err := newContractEncoder(abi.TokenERC721ABI, helper)
				if err != nil {
					return nil, err
				}

				events, err := newContractEvents(abi.TokenERC721ABI, helper)
				if err != nil {
					return nil, err
				}

				nftCollection := &NFTCollection{
					erc721,
					contractAbi,
					helper,
					signature,
					encoder,
					events,
				}
				return nftCollection, nil
			}
		}
	}
}

// Get the metadatas of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the metadata of all the NFTs owned by the address
//
// Example
//
//	owner := "{{wallet_address}}"
//	nfts, err := contract.GetOwned(context.Background(), owner)
//	name := nfts[0].Metadata.Name
func (nft *NFTCollection) GetOwned(ctx context.Context, address string) ([]*NFTMetadataOwner, error) {
	if address == "" {
		address = nft.Helper.GetSignerAddress().String()
	}

	if tokenIds, err := nft.GetOwnedTokenIDs(ctx, address); err != nil {
		return nil, err
	} else {
		return nft.fetchNFTsByTokenId(ctx, tokenIds)
	}
}

// Get the tokenIds of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the tokenIds of all the NFTs owned by the address
func (nft *NFTCollection) GetOwnedTokenIDs(ctx context.Context, address string) ([]*big.Int, error) {
	if address == "" {
		address = nft.Helper.GetSignerAddress().String()
	}

	if balance, err := nft.abi.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(address)); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}

		for i := 0; i < int(balance.Int64()); i++ {
			if tokenId, err := nft.abi.TokenOfOwnerByIndex(&bind.CallOpts{Context: ctx}, common.HexToAddress(address), big.NewInt(int64(i))); err == nil {
				tokenIds = append(tokenIds, tokenId)
			}
		}

		return tokenIds, nil
	}
}

// Mint a new NFT to the connected wallet.
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) Mint(ctx context.Context, metadata *NFTMetadataInput) (*types.Transaction, error) {
	address := nft.Helper.GetSignerAddress().String()
	return nft.MintTo(ctx, address, metadata)
}

// Mint a new NFT to the specified wallet.
//
// address: the wallet address to mint to
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	image, err := os.Open("path/to/image.jpg")
//	defer image.Close()
//
//	metadata := &web3sdks.NFTMetadataInput{
//		Name: "Cool NFT",
//		Description: "This is a cool NFT",
//		Image: image,
//	}
//
//	tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", metadata)
func (nft *NFTCollection) MintTo(ctx context.Context, address string, metadata *NFTMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(ctx, metadata, nft.storage)
	if err != nil {
		return nil, err
	}

	txOpts, err := nft.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := nft.abi.MintTo(
		txOpts,
		common.HexToAddress(address),
		uri,
	)
	if err != nil {
		return nil, err
	}

	return nft.Helper.AwaitTx(ctx, tx.Hash())
}

// Mint a batch of new NFTs to the connected wallet.
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) MintBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	address := nft.Helper.GetSignerAddress().String()
	return nft.MintBatchTo(ctx, address, metadatas)
}

// Mint a batch of new NFTs to the specified wallet.
//
// to: the wallet address to mint to
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	metadatas := []*web3sdks.NFTMetadataInput{
//		&web3sdks.NFTMetadataInput{
//			Name: "Cool NFT",
//			Description: "This is a cool NFT",
//		}
//		&web3sdks.NFTMetadataInput{
//			Name: "Cool NFT 2",
//			Description: "This is also a cool NFT",
//		}
//	}
//
//	tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatas)
func (nft *NFTCollection) MintBatchTo(ctx context.Context, address string, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	uris, err := uploadOrExtractUris(ctx, metadatas, nft.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	for _, uri := range uris {
		txOpts, err := nft.Helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := nft.abi.MintTo(
			txOpts,
			common.HexToAddress(address),
			uri,
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := nft.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := nft.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return nft.Helper.AwaitTx(ctx, tx.Hash())
}
