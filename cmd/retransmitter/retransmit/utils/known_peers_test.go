package utils_test

import (
	"github.com/wavesplatform/gowaves/cmd/retransmitter/retransmit/utils"
	"net"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func TestKnownPeers_SaveDisk(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := utils.NewFileBasedStorage(fs, "/known_peers.json")
	require.NoError(t, err)

	knownPeers, err := utils.NewKnownPeersInterval(s, time.NewTicker(1*time.Second))
	require.NoError(t, err)
	defer knownPeers.Stop()
	knownPeers.Add(proto.PeerInfo{Addr: net.IPv4(10, 10, 10, 10), Port: 90}, proto.Version{})

	assert.Equal(t, []string{"10.10.10.10:90"}, knownPeers.GetAll())
	assert.Len(t, knownPeers.Addresses(), 1)
}