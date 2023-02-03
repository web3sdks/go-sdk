package web3sdks

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventListener(t *testing.T) {
	nft := getNft()

	events := []ContractEvent{}
	subscription := nft.Events.AddEventListener(
		context.Background(), 
		"TokensMinted", 
		func (event ContractEvent) {
			events = append(events, event)
		},
	)

	_, err := nft.Mint(context.Background(), &NFTMetadataInput{})
	assert.Nil(t, err)

	_, err = nft.Mint(context.Background(), &NFTMetadataInput{})
	assert.Nil(t, err)

	<- time.After(2 * time.Second)

	subscription.Unsubscribe()

	_, err = nft.Mint(context.Background(), &NFTMetadataInput{})
	assert.Nil(t, err)

	<- time.After(2 * time.Second)

	assert.Equal(t, 2, len(events))
	assert.Equal(t, events[0].EventName, "TokensMinted")
}

func TestGetEvents(t *testing.T) {
	nft := getNft()

	_, err := nft.Mint(context.Background(), &NFTMetadataInput{})
	assert.Nil(t, err)

	events, err := nft.Events.GetEvents(context.Background(), "TokensMinted", EventQueryOptions{})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(events))
}