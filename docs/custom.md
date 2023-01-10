
## Custom Contracts

### Custom Contracts

With the web3sdks SDK, you can get a contract instance for any contract\. Additionally, if you deployed your contract using web3sdks deploy, you can get a more explicit and intuitive interface to interact with your contracts\.

### Getting a Custom Contract Instance

Let's take a look at how you can get a custom contract instance for one of your contracts deployed using the web3sdks deploy flow:

```
import (
	"github.com/web3sdks/go-sdk/web3sdks"
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
	"github.com/web3sdks/go-sdk/web3sdks"
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

### Calling Contract Functions

Now that you have an SDK instance for your contract, you can easily call any function on your contract with the contract "call" method as follows:

```
// The first parameter to the call function is the method name
// All other parameters to the call function get passed as arguments to your contract
balance, err := contract.Call("balanceOf", "{{wallet_address}}")

// You can also make a transaction to your contract with the call method
tx, err := contract.Call("mintTo", "{{wallet_address}}", "ipfs://...")
```

```go
type SmartContract struct {}
```

### func \(\*SmartContract\) [Call](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/smart_contract.go#L117>)

```go
func (c *SmartContract) Call(method string, args ...interface{}) (interface{}, error)
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
tx, err := contract.Call("mintTo", "{{wallet_address}}", "ipfs://...")
```
