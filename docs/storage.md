
## IPFS Storage

```go
type IpfsStorage struct {
    Url string
}
```

### func \(\*IpfsStorage\) [Get](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/ipfs_storage.go#L55>)

```go
func (ipfs *IpfsStorage) Get(uri string) ([]byte, error)
```

#### Get

Get IPFS data at a given hash and return it as byte data

uri: the IPFS URI to fetch data from

returns: byte data of the IPFS data at the URI

### func \(\*IpfsStorage\) [Upload](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/ipfs_storage.go#L84>)

```go
func (ipfs *IpfsStorage) Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
```

#### Upload

Upload method can be used to upload a generic payload to IPFS\.

data: the individual data to upload to IPFS

contractAddress: the optional contractAddress upload is being called from

signerAddress: the optional signerAddress upload is being called from

returns: the URI of the IPFS upload

### func \(\*IpfsStorage\) [UploadBatch](<https://github.com/web3sdks/go-sdk/blob/main/web3sdks/ipfs_storage.go#L105>)

```go
func (ipfs *IpfsStorage) UploadBatch(data []interface{}, fileStartNumber int, contractAddress string, signerAddress string) (*baseUriWithUris, error)
```

#### UploadBatch

UploadBatch method can be used to upload a batch of generic payloads to IPFS\.

data: the array of data to upload to IPFS

contractAddress: the optional contractAddress upload is being called from

signerAddress: the optional signerAddress upload is being called from

returns: the base URI of the IPFS upload folder with the URIs of each subfile
