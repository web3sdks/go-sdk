package web3sdks

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/web3sdks/go-sdk/v2/abi"
)

// You can access the Token interface from the SDK as follows:
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
//	contract, err := sdk.GetToken("{{contract_address}}")
type Token struct {
	*ERC20
	abi     *abi.TokenERC20
	Helper  *contractHelper
	Encoder *ContractEncoder
	Events  *ContractEvents
}

func newToken(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Token, error) {
	if contractAbi, err := abi.NewTokenERC20(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		if erc20, err := newERC20(provider, address, privateKey, storage); err != nil {
			return nil, err
		} else {
			encoder, err := newContractEncoder(abi.TokenERC20ABI, helper)
			if err != nil {
				return nil, err
			}

			events, err := newContractEvents(abi.TokenERC20ABI, helper)
			if err != nil {
				return nil, err
			}

			token := &Token{
				erc20,
				contractAbi,
				helper,
				encoder,
				events,
			}
			return token, nil
		}
	}
}

// Get the connected wallets voting power in this token.
//
// returns: vote balance of the connected wallet
func (token *Token) GetVoteBalance(ctx context.Context) (*CurrencyValue, error) {
	return token.GetVoteBalanceOf(ctx, token.Helper.GetSignerAddress().String())
}

// Get the voting power of the specified wallet in this token.
//
// address: wallet address to check the vote balance of
//
// returns: vote balance of the specified wallet
func (token *Token) GetVoteBalanceOf(ctx context.Context, address string) (*CurrencyValue, error) {
	votes, err := token.abi.GetVotes(&bind.CallOpts{Context: ctx}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return token.getValue(ctx, votes)
}

// Get the connected wallets delegatee address for this token.
//
// returns: delegation address of the connected wallet
func (token *Token) GetDelegation(ctx context.Context) (string, error) {
	return token.GetDelegationOf(ctx, token.Helper.GetSignerAddress().String())
}

// Get a specified wallets delegatee for this token.
//
// returns: delegation address of the connected wallet
func (token *Token) GetDelegationOf(ctx context.Context, address string) (string, error) {
	delegation, err := token.abi.Delegates(&bind.CallOpts{Context: ctx}, common.HexToAddress(address))
	if err != nil {
		return "", err
	}

	return delegation.String(), nil
}

// Mint tokens to the connected wallet.
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
func (token *Token) Mint(ctx context.Context, amount float64) (*types.Transaction, error) {
	return token.MintTo(ctx, token.Helper.GetSignerAddress().String(), amount)
}

// Mint tokens to a specified wallet.
//
// to: wallet address to mint tokens to
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
//
// Example
//
//	tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", 1)
func (token *Token) MintTo(ctx context.Context, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := token.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := token.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := token.abi.MintTo(txOpts, common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return token.Helper.AwaitTx(ctx, tx.Hash())
}

// Mint tokens to a list of wallets.
//
// args: list of wallet addresses and amounts to mint
//
// returns: transaction receipt of the mint
//
// Example
//
//	args = []*web3sdks.TokenAmount{
//		&web3sdks.TokenAmount{
//			ToAddress: "{{wallet_address}}",
//			Amount:    1
//		}
//		&web3sdks.TokenAmount{
//			ToAddress: "{{wallet_address}}",
//			Amount:    2
//		}
//	}
//
//	tx, err := contract.MintBatchTo(context.Background(), args)
func (token *Token) MintBatchTo(ctx context.Context, args []*TokenAmount) (*types.Transaction, error) {
	encoded := [][]byte{}

	for _, arg := range args {
		amountWithDecimals, err := token.normalizeAmount(ctx, arg.Amount)
		if err != nil {
			return nil, err
		}

		txOpts, err := token.Helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := token.abi.MintTo(txOpts, common.HexToAddress(arg.ToAddress), amountWithDecimals)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := token.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := token.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return token.Helper.AwaitTx(ctx, tx.Hash())
}

// Delegate the connected wallets tokens to a specified wallet.
//
// delegateeAddress: wallet address to delegate tokens to
//
// returns: transaction receipt of the delegation
func (token *Token) DelegateTo(ctx context.Context, delegatreeAddress string) (*types.Transaction, error) {
	txOpts, err := token.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := token.abi.Delegate(txOpts, common.HexToAddress(delegatreeAddress))
	if err != nil {
		return nil, err
	}

	return token.Helper.AwaitTx(ctx, tx.Hash())
}
