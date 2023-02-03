.PHONY: abi docs publish local-test

SHELL := /bin/bash

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out abi/token_erc1155.go --type TokenERC1155
	# If you want to generate drop contracts, you'll have to delete a struct
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721_V3.json --out abi/drop_erc721.go --type DropERC721
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC1155_V2.json --out abi/drop_erc1155.go --type DropERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Multiwrap.json --out abi/multiwrap.go --type Multiwrap
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Marketplace.json --out abi/marketplace.go --type Marketplace

	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TWFactory.json --out abi/twfactory.go --type TWFactory
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out abi/ierc20.go --type IERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC721.json --out abi/ierc721.go --type IERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC1155.json --out abi/ierc1155.go --type IERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC165.json --out abi/ierc165.go --type IERC165

docs:
	rm -rf docs
	mkdir docs
	gomarkdoc --output docs/doc.md --repository.default-branch main ./web3sdks
	node ./scripts/generate-docs.mjs
	rm ./docs/doc.md ./docs/start.md ./docs/delete.md
	node ./scripts/generate-snippets.mjs

cmd: FORCE
	cd cmd/web3sdks && go build -o ../../bin/web3sdks && cd -

test-nft-read:
	./bin/web3sdks nft getAll -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks nft getOwned -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-write:
	./bin/web3sdks nft mint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks nft mintLink -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-sigmint:
	./bin/web3sdks nft sigmint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-read:
	./bin/web3sdks edition getAll -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks edition getOwned -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-write:
	./bin/web3sdks edition mint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-sigmint:
	./bin/web3sdks edition sigmint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks edition sigmint-tokenid -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-read:
	./bin/web3sdks token get -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-write:
	./bin/web3sdks token mint -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks token mintBatch -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-read:
	./bin/web3sdks multiwrap getAll -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	# ./bin/web3sdks multiwrap getContents -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-write:
	./bin/web3sdks multiwrap wrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC} -n ${GO_NFT_COLLECTION} -e ${GO_EDITION} -t ${GO_TOKEN}
	./bin/web3sdks multiwrap unwrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-read:
	./bin/web3sdks nftdrop getAll -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks nftdrop getActive -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-write:
	./bin/web3sdks nftdrop createBatch -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks nftdrop claim -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-read:
	./bin/web3sdks editiondrop getAll -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-write:
	./bin/web3sdks editiondrop createBatch -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks editiondrop claim -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-storage:
	./bin/web3sdks storage upload
	./bin/web3sdks storage uploadBatch
	./bin/web3sdks storage uploadImage
	./bin/web3sdks storage uploadImageLink

test-custom:
	./bin/web3sdks custom set -a ${GO_CUSTOM} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-deploy:
	./bin/web3sdks deploy nft -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks deploy edition -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks deploy token -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks deploy nftdrop -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks deploy editiondrop -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks deploy multiwrap -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/web3sdks deploy marketplace -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-encoder:
	./bin/web3sdks marketplace encode -a ${GO_MARKETPLACE} -k ${GO_PRIVATE_KEY} -u ${GO_AVAX_RPC}

test-cmd:
	make cmd
	make test-nft-read
	make test-nft-write
	make test-edition-read
	make test-edition-write
	make test-token-read
	make test-token-write
	make test-drop-read
	make test-drop-write
	make test-edition-drop-read
	make test-edition-drop-write
	make test-multiwrap-read
	make test-multiwrap-write
	make test-storage

stop-docker:
	docker stop hardhat-node
	docker rm hardhat-node

test: FORCE
	docker build . -t hardhat-mainnet-fork
	docker start hardhat-node || docker run --name hardhat-node -d -p 8545:8545 -e SDK_ALCHEMY_KEY=${SDK_ALCHEMY_KEY} hardhat-mainnet-fork
	sudo bash ./scripts/test/await-hardhat.sh
	go clean -testcache
	go test -v ./web3sdks
	docker stop hardhat-node
	docker rm hardhat-node

local-test:
  # Needs to be run along with npx hardhat node from this repo, and needs to be a mainnet fork hardhat
	go clean -testcache
	go test -v ./web3sdks

publish:
	# Make sure to pass the TAG variable to this command ex: `make publish TAG=v2.0.0`
	# Publish following https://go.dev/doc/modules/publishing
	go mod tidy
	git tag $(TAG)
	git push origin $(TAG)
	GOPROXY=proxy.golang.org go list -m github.com/web3sdks/go-sdk/v2@$(TAG)

FORCE: ;