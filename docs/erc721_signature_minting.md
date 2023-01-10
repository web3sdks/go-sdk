
## ERC721 Signature Minting

You can access this interface from the NFT Collection contract under the signature interface\.

```go
type ERC721SignatureMinting struct {}
```

### func \(\*ERC721SignatureMinting\) [Generate](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/erc721_signature_minting.go#L175>)

```go
func (signature *ERC721SignatureMinting) Generate(payloadToSign *Signature721PayloadInput) (*SignedPayload721, error)
```

#### Generate a new payload from the given data

payloadToSign: the payload containing the data for the signature mint

returns: the payload signed by the minter's private key

#### Example

```
payload := &web3sdks.Signature721PayloadInput{
	To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6", // address to mint to
	Price:                0,                                            // cost of minting
	CurrencyAddress:      "0x0000000000000000000000000000000000000000", // currency to pay in order to mint
	MintStartTime:        0,                                            // time where minting is allowed to start (epoch seconds)
	MintEndTime:          100000000000000,                              // time when this signature expires (epoch seconds)
	PrimarySaleRecipient: "0x0000000000000000000000000000000000000000", // address to receive the primary sales of this mint
	Metadata: &web3sdks.NFTMetadataInput{																// metadata of the NFT to mint
 		Name:  "ERC721 Sigmint!",
	},
	RoyaltyRecipient: "0x0000000000000000000000000000000000000000",     // address to receive royalties of this mint
	RoyaltyBps:       0,                                                // royalty cut of this mint in basis points
}

signedPayload, err := contract.Signature.Generate(payload)
```

### func \(\*ERC721SignatureMinting\) [GenerateBatch](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/erc721_signature_minting.go#L224>)

```go
func (signature *ERC721SignatureMinting) GenerateBatch(payloadsToSign []*Signature721PayloadInput) ([]*SignedPayload721, error)
```

#### Generate a batch of new payload from the given data

payloadToSign: the payloads containing the data for the signature mint

returns: the payloads signed by the minter's private key

#### Example

```
payload := []*web3sdks.Signature721PayloadInput{
	&web3sdks.Signature721PayloadInput{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &web3sdks.NFTMetadataInput{
 			Name:  "ERC721 Sigmint!",
 			Image: imageFile,
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
	},
	&web3sdks.Signature721PayloadInput{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &web3sdks.NFTMetadataInput{
 			Name:  "ERC721 Sigmint!",
 			Image: imageFile,
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
	},
}

signedPayload, err := contract.Signature.GenerateBatch(payload)
```

### func \(\*ERC721SignatureMinting\) [Mint](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/erc721_signature_minting.go#L51>)

```go
func (signature *ERC721SignatureMinting) Mint(signedPayload *SignedPayload721) (*types.Transaction, error)
```

Mint a token with the data in given payload\.

signedPayload: the payload signed by the minters private key being used to mint

returns: the transaction receipt of the mint

#### Example

```
// Learn more about how to craft a payload in the Generate() function
signedPayload, err := contract.Signature.Generate(payload)
tx, err := contract.Signature.Mint(signedPayload)
```

### func \(\*ERC721SignatureMinting\) [MintBatch](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/erc721_signature_minting.go#L87>)

```go
func (signature *ERC721SignatureMinting) MintBatch(signedPayloads []*SignedPayload721) (*types.Transaction, error)
```

Mint a batch of token with the data in given payload\.

signedPayload: the list of payloads signed by the minters private key being used to mint

returns: the transaction receipt of the batch mint

#### Example

```
// Learn more about how to craft multiple payloads in the GenerateBatch() function
signedPayloads, err := contract.Signature.GenerateBatch(payloads)
tx, err := contract.Signature.MintBatch(signedPayloads)
```

### func \(\*ERC721SignatureMinting\) [Verify](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/erc721_signature_minting.go#L139>)

```go
func (signature *ERC721SignatureMinting) Verify(signedPayload *SignedPayload721) (bool, error)
```

#### Verify that a signed payload is valid

signedPayload: the payload to verify

returns: true if the payload is valid, otherwise false\.

#### Example

```
// Learn more about how to craft a payload in the Generate() function
signedPayload, err := contract.Signature.Generate(payload)
isValid, err := contract.Signature.Verify(signedPayload)
```
