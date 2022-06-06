// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/spruce-solutions/go-quai/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Quai Network.
var MainnetBootnodes = []string{
	"enode://ac169b82df3350085dbdb2cd236008fac52cc0cb81481c0b0a957e01da1a9552201c8c6cfb54ebb28b7c23fbd065cf5e9cfa1f30d6802ecf4831ccfb6d36dc92@34.135.197.187", // ds-team-node
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://a99f5dfcd7c642521b01873befef566829f732ef1a05d664f0737cc5f7877da00a086c3599619272eb220e1e5c7dcf3d527290c6dfa289a71f77db8dd2d27a2d@216.128.131.59", // vultr-full-node-4-ropsten
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetPrimeGenesisHash:
		net = "mainnet-prime"
	case MainnetRegionGenesisHash:
		net = "mainnet-region"
	case MainnetZoneGenesisHash:
		net = "mainnet-zone"
	case RopstenPrimeGenesisHash:
		net = "ropsten-prime"
	case RopstenRegionGenesisHash:
		net = "ropsten-region"
	case RopstenZoneGenesisHash:
		net = "ropsten-zone"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
