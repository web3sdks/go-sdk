package main

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [command]",
	Short: "Deploy a contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var deployNftCmd = &cobra.Command{
	Use:   "nft",
	Short: "Deploy an nft collection",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployNFTCollection(context.Background(), &web3sdks.DeployNFTCollectionMetadata{
			Name: "Goku NFT",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionCmd = &cobra.Command{
	Use:   "edition",
	Short: "Deploy an edition",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployEdition(context.Background(), &web3sdks.DeployEditionMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Deploy an token",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployToken(context.Background(), &web3sdks.DeployTokenMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployNFTDropCmd = &cobra.Command{
	Use:   "nftdrop",
	Short: "Deploy an nft drop",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployNFTDrop(context.Background(), &web3sdks.DeployNFTDropMetadata{
			Name: "Go Script Drop",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionDropCmd = &cobra.Command{
	Use:   "editiondrop",
	Short: "Deploy an edition drop",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployEditionDrop(context.Background(), &web3sdks.DeployEditionDropMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMultiwrapCmd = &cobra.Command{
	Use:   "multiwrap",
	Short: "Deploy a multiwrap",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployMultiwrap(context.Background(), &web3sdks.DeployMultiwrapMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMarketplaceCmd = &cobra.Command{
	Use:   "marketplace",
	Short: "Deploy a marketplace",
	Run: func(cmd *cobra.Command, args []string) {
		if web3sdksSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := web3sdksSDK.Deployer.DeployMarketplace(context.Background(), &web3sdks.DeployMarketplaceMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

func init() {
	deployCmd.AddCommand(deployNftCmd)
	deployCmd.AddCommand(deployEditionCmd)
	deployCmd.AddCommand(deployTokenCmd)
	deployCmd.AddCommand(deployNFTDropCmd)
	deployCmd.AddCommand(deployEditionDropCmd)
	deployCmd.AddCommand(deployMultiwrapCmd)
	deployCmd.AddCommand(deployMarketplaceCmd)
}
