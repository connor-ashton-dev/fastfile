package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	litenAddr := ":4000"
	tr := NewTCPTransport(litenAddr)
	assert.Equal(t, tr.listenAddress, litenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
