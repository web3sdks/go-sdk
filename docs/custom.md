
## Custom Contracts

### Custom Contracts

With the web3sdks SDK, you can get a contract instance for any contract\. Additionally, if you deployed your contract using web3sdks deploy, you can get a more explicit and intuitive interface to interact with your contracts\.

\# Getting a Custom Contract Instance

Let's take a look at how you can get a custom contract instance for one of your contracts deployed using the web3sdks deploy flow:

```
import (
	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

privateKey = "..."

sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
	PrivateKey: privateKey,
})

// You can replace your own contract address here
contractAddress := "{{contract_address}}"

// Now you have a contract instance ready to go
contract, err := sdk.GetContract(contractAddress)
```

Alternatively, if you didn't deploy your contract with web3sdks deploy, you can still get a contract instance for any contract using your contracts ABI:

```
import (
	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

privateKey = "..."

sdk, err := web3sdks.NewWeb3sdksSDK("mumbai", &web3sdks.SDKOptions{
	PrivateKey: privateKey,
})

// You can replace your own contract address here
contractAddress := "{{contract_address}}"

// Add your contract ABI here
abi := "[...]"

// Now you have a contract instance ready to go
contract, err := sdk.GetContractFromAbi(contractAddress, abi)
```

\# Calling Contract Functions

Now that you have an SDK instance for your contract, you can easily call any function on your contract with the contract "call" method as follows:

```
// The first parameter to the call function is the method name
// All other parameters to the call function get passed as arguments to your contract
balance, err := contract.Call("balanceOf", "{{wallet_address}}")

// You can also make a transaction to your contract with the call method
tx, err := contract.Call("mintTo", "{{wallet_address}}", "ipfs://...")
```

```go
type SmartContract struct {
    Encoder *ContractEncoder
    Events  *ContractEvents
}
```

### func \(\*SmartContract\) [Call](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/smart_contract.go#L134>)

```go
func (c *SmartContract) Call(ctx context.Context, method string, args ...interface{}) (interface{}, error)
```

Call any function on your contract\.

method: the name of the method on your contract you want to call

args: the arguments to pass to the method

#### Example

```
// The first parameter to the call function is the method name
// All other parameters to the call function get passed as arguments to your contract
balance, err := contract.Call("balanceOf", "{{wallet_address}}")

// You can also make a transaction to your contract with the call method
tx, err := contract.Call(context.Background(), "mintTo", "{{wallet_address}}", "ipfs://...")
```

## type [SnapshotClaim](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/snapshots.go#L17-L21>)

```go
type SnapshotClaim struct {
    Address      string   `json:"address"`
    MaxClaimable int      `json:"maxClaimable"`
    Proof        []string `json:"proof"`
}
```

## type [SnapshotEntry](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/types.go#L573-L578>)

```go
type SnapshotEntry struct {
    Address         string `json:"address"`
    MaxClaimable    string `json:"maxClaimable"`
    Price           string `json:"price"`
    CurrencyAddress string `json:"currencyAddress"`
}
```

## type [SnapshotEntryWithProof](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/types.go#L551-L557>)

```go
type SnapshotEntryWithProof struct {
    Address         string
    MaxClaimable    string
    Price           string
    CurrencyAddress string
    Proof           [][32]byte
}
```

## type [SnapshotInfo](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/snapshots.go#L23-L26>)

```go
type SnapshotInfo struct {
    MerkleRoot string          `json:"merkleRoot"`
    Claims     []SnapshotClaim `json:"claims"`
}
```

## type [SnapshotInfos](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/snapshots.go#L28-L32>)

```go
type SnapshotInfos struct {
    Snapshot    SnapshotInfo
    MerkleRoot  string
    SnapshotUri string
}
```

## type [SnapshotInput](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/snapshots.go#L12-L15>)

```go
type SnapshotInput struct {
    Address      string
    MaxClaimable int
}
```
