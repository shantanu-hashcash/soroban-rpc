package test

import (
	"context"
	"testing"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/stretchr/testify/assert"

	"github.com/shantanu-hashcash/soroban-rpc/cmd/soroban-rpc/internal/methods"
)

func TestGetNetworkSucceeds(t *testing.T) {
	test := NewTest(t, nil)

	ch := jhttp.NewChannel(test.sorobanRPCURL(), nil)
	client := jrpc2.NewClient(ch, nil)

	request := methods.GetNetworkRequest{}

	var result methods.GetNetworkResponse
	err := client.CallResult(context.Background(), "getNetwork", request, &result)
	assert.NoError(t, err)
	assert.Equal(t, friendbotURL, result.FriendbotURL)
	assert.Equal(t, StandaloneNetworkPassphrase, result.Passphrase)
	assert.GreaterOrEqual(t, result.ProtocolVersion, 20)
}
