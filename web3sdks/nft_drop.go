package web3sdks

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3sdks/go-sdk/internal/abi"
)

// You can access the NFT Drop interface from the SDK as follows:
//
// 	import (
// 		"github.com/web3sdks/go-sdk/web3sdks"
// 	)
//
// 	privateKey = "..."
//
// 	sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
//		PrivateKey: privateKey,
// 	})
//
//	contract, err := sdk.GetNFTDrop("{{contract_address}}")
type NFTDrop struct {
	abi    *abi.DropERC721
	helper *contractHelper
	*ERC721
	ClaimConditions *NFTDropClaimConditions
	Encoder         *ContractEncoder
}

func newNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTDrop, error) {
	if contractAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				claimConditions, err := newNFTDropClaimConditions(address, provider, helper, storage)
				if err != nil {
					return nil, err
				}

				encoder, err := newContractEncoder(abi.DropERC721ABI, helper)
				if err != nil {
					return nil, err
				}

				nftCollection := &NFTDrop{
					contractAbi,
					helper,
					erc721,
					claimConditions,
					encoder,
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
// 	owner := "{{wallet_address}}"
// 	nfts, err := contract.GetOwned(owner)
// 	name := nfts[0].Metadata.Name
func (nft *NFTDrop) GetOwned(address string) ([]*NFTMetadataOwner, error) {
	if address == "" {
		address = nft.helper.GetSignerAddress().String()
	}

	if tokenIds, err := nft.GetOwnedTokenIDs(address); err != nil {
		return nil, err
	} else {
		return nft.fetchNFTsByTokenId(tokenIds)
	}
}

// Get the tokenIds of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the tokenIds of all the NFTs owned by the address
func (nft *NFTDrop) GetOwnedTokenIDs(address string) ([]*big.Int, error) {
	if address == "" {
		address = nft.helper.GetSignerAddress().String()
	}

	if balance, err := nft.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address)); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}

		for i := 0; i < int(balance.Int64()); i++ {
			if tokenId, err := nft.abi.TokenOfOwnerByIndex(&bind.CallOpts{}, common.HexToAddress(address), big.NewInt(int64(i))); err == nil {
				tokenIds = append(tokenIds, tokenId)
			}
		}

		return tokenIds, nil
	}
}

// Get a list of all the NFTs that have been claimed from this contract.
//
// returns: a list of the metadatas of the claimed NFTs
//
// Example
//
// 	claimedNfts, err := contract.GetAllClaimed()
// 	firstOwner := claimedNfts[0].Owner
func (drop *NFTDrop) GetAllClaimed() ([]*NFTMetadataOwner, error) {
	if maxId, err := drop.abi.NextTokenIdToClaim(&bind.CallOpts{}); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for i := 0; i < int(maxId.Int64()); i++ {
			if nft, err := drop.Get(i); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

// Get a list of all the NFTs on this contract that have not yet been claimed.
//
// returns: a list of the metadatas of the unclaimed NFTs
//
// Example
//
// 	unclaimedNfts, err := contract.GetAllUnclaimed()
// 	firstNftName := unclaimedNfts[0].Name
func (drop *NFTDrop) GetAllUnclaimed() ([]*NFTMetadata, error) {
	maxId, err := drop.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	unmintedId, err := drop.abi.NextTokenIdToClaim(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	nfts := []*NFTMetadata{}
	for i := int(unmintedId.Int64()); i < int(maxId.Int64()); i++ {
		if nft, err := drop.getTokenMetadata(int(unmintedId.Int64()) + i); err == nil {
			nfts = append(nfts, nft)
		}
	}

	return nfts, nil
}

// Create a batch of NFTs on this contract.
//
// metadatas: a list of the metadatas of the NFTs to create
//
// returns: the transaction receipt of the batch creation
//
// Example
//
// 	image0, err := os.Open("path/to/image/0.jpg")
// 	defer image0.Close()
//
// 	image1, err := os.Open("path/to/image/1.jpg")
// 	defer image1.Close()
//
// 	metadatas := []*web3sdks.NFTMetadataInput{
// 		&web3sdks.NFTMetadataInput{
// 			Name: "Cool NFT",
// 			Description: "This is a cool NFT",
// 			Image: image1
// 		}
// 		&web3sdks.NFTMetadataInput{
// 			Name: "Cool NFT 2",
// 			Description: "This is also a cool NFT",
// 			Image: image2
// 		}
// 	}
//
// 	tx, err := contract.CreateBatch(metadatas)
func (drop *NFTDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	startNumber, err := drop.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	fileStartNumber := int(startNumber.Int64())

	contractAddress := drop.helper.getAddress().String()
	signerAddress := drop.helper.GetSignerAddress().String()

	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}
	batch, err := drop.storage.UploadBatch(
		data,
		fileStartNumber,
		contractAddress,
		signerAddress,
	)

	txOpts, err := drop.helper.getTxOptions()
	if err != nil {
		return nil, err
	}
	tx, err := drop.abi.LazyMint(
		txOpts,
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.helper.awaitTx(tx.Hash())
}

// Claim NFTs from this contract to the connect wallet.
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *NFTDrop) Claim(quantity int) (*types.Transaction, error) {
	address := drop.helper.GetSignerAddress().String()
	return drop.ClaimTo(address, quantity)
}

// Claim NFTs from this contract to the connect wallet.
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
//
// Example
//
// 	address := "{{wallet_address}}"
// 	quantity = 1
//
// 	tx, err := contract.ClaimTo(address, quantity)
func (drop *NFTDrop) ClaimTo(destinationAddress string, quantity int) (*types.Transaction, error) {
	claimVerification, err := drop.prepareClaim(quantity)
	if err != nil {
		return nil, err
	}

	txOpts, err := drop.helper.getTxOptions()
	if err != nil {
		return nil, err
	}
	tx, err := drop.abi.Claim(
		txOpts,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(claimVerification.currencyAddress),
		big.NewInt(int64(claimVerification.price)),
		claimVerification.proofs,
		big.NewInt(int64(claimVerification.maxQuantityPerTransaction)),
	)
	if err != nil {
		return nil, err
	}

	return drop.helper.awaitTx(tx.Hash())
}

func (drop *NFTDrop) prepareClaim(quantity int) (*ClaimVerification, error) {
	claimCondition, err := drop.ClaimConditions.GetActive()
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		quantity,
		claimCondition,
		drop.helper,
		drop.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}
