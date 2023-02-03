package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"

	"github.com/web3sdks/go-sdk/v2/web3sdks"
)

var storageCmd = &cobra.Command{
	Use:   "storage [command]",
	Short: "Interact with the IPFS storage interface",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var storageUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload data with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		assetToUpload := map[string]interface{}{
			"psilocybin": map[string]interface{}{
				"strength": 5,
				"dosage":   "1mg",
				"potency":  "strong",
				"side_effects": []interface{}{
					"headache",
					"nausea",
				},
			},
		}

		uri, err := storage.Upload(context.Background(), assetToUpload, "", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uri)
	},
}

var storageUploadBatchCmd = &cobra.Command{
	Use:   "uploadBatch",
	Short: "Upload data with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		asset := []interface{}{
			&web3sdks.NFTMetadataInput{
				Name:        "Test NFT 2",
				Description: "Description 2",
				Image:       imageFile,
				Properties: map[string]interface{}{
					"health": "100",
					"image":  imageFile,
				},
			},
			&web3sdks.NFTMetadataInput{Name: "Test NFT 3", Description: "Description 3"},
		}
		assetToUpload := []map[string]interface{}{}
		if err := mapstructure.Decode(asset, &assetToUpload); err != nil {
			panic(err)
		}

		uriWithBaseUris, err := storage.UploadBatch(
			context.Background(),
			assetToUpload,
			0,
			"",
			"",
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uriWithBaseUris)
	},
}

var storageUploadImageCmd = &cobra.Command{
	Use:   "uploadImage",
	Short: "Upload image with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		asset := &web3sdks.NFTMetadataInput{
			Name:  "Test NFT 1",
			Image: imageFile,
		}
		assetToUpload := map[string]interface{}{}
		err = mapstructure.Decode(asset, &assetToUpload)
		if err != nil {
			panic(err)
		}

		uri, err := storage.Upload(context.Background(), assetToUpload, "", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uri)
	},
}

var storageUploadImageLinkCmd = &cobra.Command{
	Use:   "uploadImageLink",
	Short: "Upload image with link with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		asset := &web3sdks.NFTMetadataInput{
			Name:  "Test NFT 1",
			Image: "ipfs://QmcCJC4T37rykDjR6oorM8hpB9GQWHKWbAi2YR1uTabUZu/0",
		}
		assetToUpload := map[string]interface{}{}
		err = mapstructure.Decode(asset, &assetToUpload)
		if err != nil {
			panic(err)
		}

		uri, err := storage.Upload(context.Background(), assetToUpload, "", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uri)
	},
}

func init() {
	storageCmd.AddCommand(storageUploadCmd)
	storageCmd.AddCommand(storageUploadBatchCmd)
	storageCmd.AddCommand(storageUploadImageCmd)
	storageCmd.AddCommand(storageUploadImageLinkCmd)
}
