// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package gwpool_test

import (
	"context"
	"testing"
	"time"

	"github.com/TheThingsNetwork/ttn/pkg/band"
	"github.com/TheThingsNetwork/ttn/pkg/gatewayserver/gwpool"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
	"github.com/TheThingsNetwork/ttn/pkg/util/test"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestPoolUplinks(t *testing.T) {
	a := assertions.New(t)

	p := gwpool.NewPool(test.GetLogger(t), time.Millisecond)

	gatewayID := "gateway"
	gatewayIdentifier := ttnpb.GatewayIdentifier{GatewayID: gatewayID}
	link := &dummyLink{
		AcceptSendingUplinks: true,

		NextUplink: make(chan *ttnpb.GatewayUp),
	}
	emptyUplink := &ttnpb.GatewayUp{}
	upstream, err := p.Subscribe(gatewayIdentifier, link, ttnpb.FrequencyPlan{BandID: band.EU_863_870})
	a.So(err, should.BeNil)

	obs, err := p.GetGatewayObservations(&gatewayIdentifier)
	a.So(err, should.BeNil)
	a.So(obs.UplinkCount, should.Equal, 0)
	a.So(obs.LastUplinkReceived, should.BeNil)

	go func() { link.NextUplink <- emptyUplink }()
	newUplink := <-upstream
	a.So(newUplink, should.Equal, emptyUplink)

	obs, err = p.GetGatewayObservations(&gatewayIdentifier)
	a.So(err, should.BeNil)
	a.So(obs.UplinkCount, should.Equal, 0)
	a.So(obs.LastUplinkReceived, should.BeNil)

	go func() {
		link.NextUplink <- &ttnpb.GatewayUp{
			UplinkMessages: []*ttnpb.UplinkMessage{
				{},
			},
		}
	}()
	select {
	case <-upstream:
	case <-time.After(time.Second * 10):
		t.Log("Timeout: uplink was not received")
		t.Fail()
	}

	obs, err = p.GetGatewayObservations(&gatewayIdentifier)
	a.So(err, should.BeNil)
	a.So(obs.UplinkCount, should.Equal, 1)
	a.So(obs.LastUplinkReceived.Unix(), should.AlmostEqual, time.Now().Unix(), 1)

	link.AcceptSendingUplinks = false
	go func() { link.NextUplink <- emptyUplink }()
	newUplink = <-upstream
	a.So(newUplink, should.BeNil)
}

func TestDoneContextUplinks(t *testing.T) {
	a := assertions.New(t)

	p := gwpool.NewPool(test.GetLogger(t), time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())

	gatewayID := "gateway"
	link := &dummyLink{
		AcceptSendingUplinks: true,

		context:       ctx,
		cancelContext: cancel,

		NextUplink: make(chan *ttnpb.GatewayUp),
	}
	cancel()

	emptyUplink := &ttnpb.GatewayUp{}
	upstream, err := p.Subscribe(ttnpb.GatewayIdentifier{GatewayID: gatewayID}, link, ttnpb.FrequencyPlan{BandID: band.EU_863_870})
	a.So(err, should.BeNil)

	go func() { link.NextUplink <- emptyUplink }()
	time.Sleep(time.Millisecond)
	select {
	case _, ok := <-upstream:
		if ok {
			t.Error("Stream not closed, message received")
		}
	default:
		t.Error("Stream not closed, no message")
	}
}

func TestSubscribeTwice(t *testing.T) {
	a := assertions.New(t)

	p := gwpool.NewPool(test.GetLogger(t), time.Millisecond)

	gateway := ttnpb.GatewayIdentifier{GatewayID: "gateway"}

	link := &dummyLink{}
	newLink := &dummyLink{}

	_, err := p.Subscribe(gateway, link, ttnpb.FrequencyPlan{BandID: band.EU_863_870})
	a.So(err, should.BeNil)
	_, err = p.Subscribe(gateway, newLink, ttnpb.FrequencyPlan{BandID: band.EU_863_870})
	a.So(err, should.BeNil)
}
