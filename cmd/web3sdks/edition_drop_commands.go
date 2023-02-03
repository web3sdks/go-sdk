package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

var (
	editionDropContractAddress string
)

var editionDropCmd = &cobra.Command{
	Use:   "editiondrop [command]",
	Short: "Interact with an Edition Drop contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var editionDropGetAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		editionDrop, err := getEditionDrop()
		if err != nil {
			panic(err)
		}

		allNfts, err := editionDrop.GetAll(context.Background())
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got drop nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var editionDropGetActiveCmd = &cobra.Command{
	Use:   "getActive",
	Short: "Get active claim in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		editionDrop, err := getEditionDrop()
		if err != nil {
			panic(err)
		}
		all, err := editionDrop.ClaimConditions.GetAll(context.Background(), 0)
		if err != nil {
			panic(err)
		}

		for i, c := range all {
			fmt.Printf(fmt.Sprintf("\n\nClaim Condition %d\n================\n", i))
			fmt.Println("Start Time:", c.StartTime)
			fmt.Println("Available:", c.AvailableSupply)
			fmt.Println("Quantity:", c.MaxClaimableSupply)
			fmt.Println("Quantity Limit:", c.MaxClaimablePerWallet)
			fmt.Println("Price:", c.Price)
			fmt.Println("Wait In Seconds", c.WaitInSeconds)
		}

		all, err = editionDrop.ClaimConditions.GetAll(context.Background(), 1)
		if err != nil {
			panic(err)
		}

		for i, c := range all {
			fmt.Printf(fmt.Sprintf("\n\nClaim Condition %d\n================\n", i))
			fmt.Println("Start Time:", c.StartTime)
			fmt.Println("Available:", c.AvailableSupply)
			fmt.Println("Quantity:", c.MaxClaimableSupply)
			fmt.Println("Quantity Limit:", c.MaxClaimablePerWallet)
			fmt.Println("Price:", c.Price)
			fmt.Println("Wait In Seconds", c.WaitInSeconds)
		}
	},
}

var editionDropClaimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Claim an nft",
	Run: func(cmd *cobra.Command, args []string) {
		editionDrop, err := getEditionDrop()
		if err != nil {
			panic(err)
		}

		if tx, err := editionDrop.Claim(context.Background(), 0, 1); err != nil {
			panic(err)
		} else {
			log.Printf("Claimed nft successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

var editionDropCreateBatchCmd = &cobra.Command{
	Use:   "createBatch",
	Short: "Create a batch of nfts",
	Run: func(cmd *cobra.Command, args []string) {
		editionDrop, err := getEditionDrop()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		if tx, err := editionDrop.CreateBatch(
			context.Background(),
			[]*web3sdks.NFTMetadataInput{
				{
					Name:  "Drop NFT 1",
					Image: imageFile,
				},
				{
					Name:  "Drop NFT 2",
					Image: imageFile,
				},
			},
		); err != nil {
			panic(err)
		} else {
			log.Printf("Created batch of nfts successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

func init() {
	editionDropCmd.PersistentFlags().StringVarP(&editionDropContractAddress, "address", "a", "", "edition drop contract address")
	editionDropCmd.AddCommand(editionDropGetAllCmd)
	editionDropCmd.AddCommand(editionDropGetActiveCmd)
	editionDropCmd.AddCommand(editionDropClaimCmd)
	editionDropCmd.AddCommand(editionDropCreateBatchCmd)
}
