package theater

import (
	"github.com/OSHeroes/bfheroesFesl/inter/network"
	"github.com/OSHeroes/bfheroesFesl/inter/network/codec"
)

type reqPING struct {
	TID string `fesl:"TID"`
}

type ansPING struct {
	TID string `fesl:"TID"`
}

func (tm *Theater) PING(event *network.EventNewClient) {
	event.Client.Answer(&codec.Packet{		
		Message: thtrPING,
		Content: ansPING{"0"},
	})
}
