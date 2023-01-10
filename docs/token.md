
## Token

You can access the Token interface from the SDK as follows:

```
import (
	"github.com/web3sdks/go-sdk/web3sdks"
)

privateKey = "..."

sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetToken("{{contract_address}}")
```

```go
type Token struct {
    *ERC20
    Encoder *ContractEncoder
}
```

### func \(\*Token\) [DelegateTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L192>)

```go
func (token *Token) DelegateTo(delegatreeAddress string) (*types.Transaction, error)
```

Delegate the connected wallets tokens to a specified wallet\.

delegateeAddress: wallet address to delegate tokens to

returns: transaction receipt of the delegation

### func \(\*Token\) [GetDelegation](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L80>)

```go
func (token *Token) GetDelegation() (string, error)
```

Get the connected wallets delegatee address for this token\.

returns: delegation address of the connected wallet

### func \(\*Token\) [GetDelegationOf](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L87>)

```go
func (token *Token) GetDelegationOf(address string) (string, error)
```

Get a specified wallets delegatee for this token\.

returns: delegation address of the connected wallet

### func \(\*Token\) [GetVoteBalance](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L59>)

```go
func (token *Token) GetVoteBalance() (*CurrencyValue, error)
```

Get the connected wallets voting power in this token\.

returns: vote balance of the connected wallet

### func \(\*Token\) [GetVoteBalanceOf](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L68>)

```go
func (token *Token) GetVoteBalanceOf(address string) (*CurrencyValue, error)
```

Get the voting power of the specified wallet in this token\.

address: wallet address to check the vote balance of

returns: vote balance of the specified wallet

### func \(\*Token\) [Mint](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L101>)

```go
func (token *Token) Mint(amount float64) (*types.Transaction, error)
```

Mint tokens to the connected wallet\.

amount: amount of tokens to mint

returns: transaction receipt of the mint

### func \(\*Token\) [MintBatchTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L154>)

```go
func (token *Token) MintBatchTo(args []*TokenAmount) (*types.Transaction, error)
```

Mint tokens to a list of wallets\.

args: list of wallet addresses and amounts to mint

returns: transaction receipt of the mint

#### Example

```
args = []*web3sdks.TokenAmount{
	&web3sdks.TokenAmount{
		ToAddress: "{{wallet_address}}",
		Amount:    1
	}
	&web3sdks.TokenAmount{
		ToAddress: "{{wallet_address}}",
		Amount:    2
	}
}

tx, err := contract.MintBatchTo(args)
```

### func \(\*Token\) [MintTo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/token.go#L116>)

```go
func (token *Token) MintTo(to string, amount float64) (*types.Transaction, error)
```

Mint tokens to a specified wallet\.

to: wallet address to mint tokens to

amount: amount of tokens to mint

returns: transaction receipt of the mint

#### Example

```
tx, err := contract.MintTo("{{wallet_address}}", 1)
```

## type [TokenAmount](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/types.go#L94-L97>)

```go
type TokenAmount struct {
    ToAddress string
    Amount    float64
}
```
