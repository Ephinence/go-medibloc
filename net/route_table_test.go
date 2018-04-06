package net

import (
	"testing"

	"time"

	"sync"

	"github.com/libp2p/go-libp2p-peer"
	"github.com/medibloc/go-medibloc/util/logging"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	RouteTableCheckIntervalSec = 10
)

func waitRouteTableSyncLoop(wg *sync.WaitGroup, node *Node, nodeIDs []peer.ID, sec time.Duration) {
	defer wg.Done()
	ids := make([]peer.ID, len(nodeIDs))
	copy(ids, nodeIDs)
	remainIteration := len(ids)

	for len(ids) > 0 {
		i := 0
		for _, v := range ids {
			if node.routeTable.peerStore.Addrs(v) == nil {
				ids[i] = v
				i++
			}
		}
		ids = ids[:i]

		time.Sleep(sec * time.Second)

		remainIteration--
		// fail if (sec * len(nodeIDs)) seconds left
		if remainIteration < 1 && len(ids) > 0 {
			logging.Console().WithFields(logrus.Fields{
				"node ID":                          node.ID(),
				"routeTable not synced node count": len(ids),
			}).Warn("route table not synced in time")
			return
		}
	}
}

func TestRouteTable_SyncWithPeer(t *testing.T) {
	logging.Console().Info("TestRouteTable_SyncWithPeer Start...")
	var wg sync.WaitGroup

	nodeNum := 5
	seedNum := 2
	nodeArr := make([]*Node, nodeNum)
	var err error
	var seedNodes []ma.Multiaddr
	var seedNodeIDs, allNodeIDs []peer.ID

	// make all test nodes
	for i := 0; i < nodeNum; i++ {
		nodeArr[i], err = makeNewTestNode("")
		assert.Nil(t, err)
	}

	// set value of seedNodes, seedNodeIDs
	for i := 0; i < seedNum; i++ {
		seedMultiaddrs, err := convertListenAddrToMultiAddr(nodeArr[i].config.Listen)
		assert.Nil(t, err)
		newSeedNodes, err := convertMultiAddrToIPFSMultiAddr(seedMultiaddrs, nodeArr[i].ID())
		assert.Nil(t, err)
		for _, v := range newSeedNodes {
			seedNodes = append(seedNodes, v)
		}
		seedNodeIDs = append(seedNodeIDs, nodeArr[i].id)
	}

	// set value of allNodeIDs
	for i := 0; i < nodeNum; i++ {
		allNodeIDs = append(allNodeIDs, nodeArr[i].id)
	}

	// setup seedNodes to every nodes
	for i := 0; i < nodeNum; i++ {
		nodeArr[i].routeTable.seedNodes = seedNodes
	}

	// start seed nodes and wait for sync route table
	for i := 0; i < seedNum; i++ {
		nodeArr[i].Start()
		wg.Add(1)
		go waitRouteTableSyncLoop(&wg, nodeArr[i], seedNodeIDs, RouteTableCheckIntervalSec)
	}

	// wait until seed nodes are synced
	logging.Console().Info("Waiting waitGroup Start...")
	wg.Wait()
	logging.Console().Info("Waiting waitGroup Finished")

	// start normal nodes
	for i := nodeNum - 1; i >= seedNum; i-- {
		nodeArr[i].Start()
	}

	// wait
	for i := 0; i < nodeNum; i++ {
		wg.Add(1)
		go waitRouteTableSyncLoop(&wg, nodeArr[i], allNodeIDs, RouteTableCheckIntervalSec)
	}

	// wait until all nodes are synced
	logging.Console().Info("Waiting waitGroup Start...")
	wg.Wait()
	logging.Console().Info("Waiting waitGroup Finished")

	// test whether route table peer list is correct
	for i := 0; i < nodeNum; i++ {
		got := nodeArr[i].routeTable.peerStore.Peers()
		want := allNodeIDs
		assert.Subset(t, got, want)
		assert.Subset(t, want, got)
	}

	// for debug
	for i := 0; i < nodeNum; i++ {
		nodeArr[i].routeTable.PrintPeers()
	}

	// stop all nodes
	for i := 0; i < nodeNum; i++ {
		nodeArr[i].Stop()
	}

	logging.Console().Info("TestRouteTable_SyncWithPeer Finished")
}