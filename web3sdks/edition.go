package web3sdks

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/web3sdks/go-sdk/v2/abi"
)

// You can access the Edition interface from the SDK as follows:
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
//	contract, err := sdk.GetEdition("{{contract_address}}")
type Edition struct {
	*ERC1155
	abi       *abi.TokenERC1155
	Helper    *contractHelper
	Signature *ERC1155SignatureMinting
	Encoder   *ContractEncoder
	Events    *ContractEvents
}

func newEdition(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Edition, error) {
	if contractAbi, err := abi.NewTokenERC1155(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			erc1155, err := newERC1155(provider, address, privateKey, storage)
			if err != nil {
				return nil, err
			}

			signature, err := newERC1155SignatureMinting(provider, address, privateKey, storage)
			if err != nil {
				return nil, err
			}

			encoder, err := newContractEncoder(abi.TokenERC1155ABI, helper)
			if err != nil {
				return nil, err
			}

			events, err := newContractEvents(abi.TokenERC1155ABI, helper)
			if err != nil {
				return nil, err
			}

			edition := &Edition{
				erc1155,
				contractAbi,
				helper,
				signature,
				encoder,
				events,
			}
			return edition, nil
		}
	}
}

// Mint an NFT to the connected wallet.
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) Mint(ctx context.Context, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	address := edition.Helper.GetSignerAddress().String()
	return edition.MintTo(ctx, address, metadataWithSupply)
}

// Mint a new NFT to the specified wallet.
//
// address: the wallet address to mint the NFT to
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//		image, err := os.Open("path/to/image.jpg")
//		defer image.Close()
//
//		metadataWithSupply := &web3sdks.EditionMetadataInput{
//	        context.Background(),
//			Metadata: &web3sdks.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//				Image: image,
//			},
//			Supply: 100,
//		}
//
//		tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", metadataWithSupply)
func (edition *Edition) MintTo(ctx context.Context, address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(ctx, metadataWithSupply.Metadata, edition.storage)
	if err != nil {
		return nil, err
	}

	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	txOpts, err := edition.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := edition.abi.MintTo(
		txOpts,
		common.HexToAddress(address),
		MaxUint256,
		uri,
		big.NewInt(int64(metadataWithSupply.Supply)),
	)
	if err != nil {
		return nil, err
	}

	return edition.Helper.AwaitTx(ctx, tx.Hash())
}

// Mint additionaly supply of a token to the connected wallet.
//
// tokenId: token ID to mint additional supply of
//
// additionalSupply: additional supply to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintAdditionalSupply(ctx context.Context, tokenId int, additionalSupply int) (*types.Transaction, error) {
	address := edition.Helper.GetSignerAddress().String()
	return edition.MintAdditionalSupplyTo(ctx, address, tokenId, additionalSupply)
}

// Mint additional supply of a token to the specified wallet.
//
// to: address of the wallet to mint NFTs to
//
// tokenId: token Id to mint additional supply of
//
// additionalySupply: additional supply to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintAdditionalSupplyTo(ctx context.Context, to string, tokenId int, additionalSupply int) (*types.Transaction, error) {
	metadata, err := edition.getTokenMetadata(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	txOpts, err := edition.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := edition.abi.MintTo(
		txOpts,
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
		metadata.Uri,
		big.NewInt(int64(additionalSupply)),
	)
	if err != nil {
		return nil, err
	}

	return edition.Helper.AwaitTx(ctx, tx.Hash())
}

// Mint a batch of NFTs to the connected wallet.
//
// metadatasWithSupply: list of NFT metadatas with supplies to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintBatch(ctx context.Context, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	return edition.MintBatchTo(ctx, edition.Helper.GetSignerAddress().String(), metadatasWithSupply)
}

// Mint a batch of NFTs to a specific wallet.
//
// to: address of the wallet to mint NFTs to
//
// metadatasWithSupply: list of NFT metadatas with supplies to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	metadatasWithSupply := []*web3sdks.EditionMetadataInput{
//		&web3sdks.EditionMetadataInput{
//			Metadata: &web3sdks.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//			},
//			Supply: 100,
//		},
//		&web3sdks.EditionMetadataInput{
//			Metadata: &web3sdks.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//			},
//			Supply: 100,
//		},
//	}
//
//	tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatasWithSupply)
func (edition *Edition) MintBatchTo(ctx context.Context, to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	metadatas := []*NFTMetadataInput{}
	for _, metadataWithSupply := range metadatasWithSupply {
		metadatas = append(metadatas, metadataWithSupply.Metadata)
	}

	supplies := []int{}
	for _, metadataWithSupply := range metadatasWithSupply {
		supplies = append(supplies, metadataWithSupply.Supply)
	}

	uris, err := uploadOrExtractUris(ctx, metadatas, edition.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	for index, uri := range uris {
		txOpts, err := edition.Helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := edition.abi.MintTo(
			txOpts,
			common.HexToAddress(to),
			MaxUint256,
			uri,
			big.NewInt(int64(supplies[index])),
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := edition.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := edition.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return edition.Helper.AwaitTx(ctx, tx.Hash())
}
