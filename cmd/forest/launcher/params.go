package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://3a2df61c1236394c5ab2405ba94a6e194696a059e3b311a76b80651cab453d5cd5e1fece50fc89a8913fa02af74978533bdcde9cd47186986887c80bc75f8489@38.242.208.184:5050",
		"enode://27fede1a17e40e0787e099c516646102a5a5887dbaee561bedaedb6764ba589c16d0724832e235480153df8fa908645d8f02492b8921246a25ae3b8f74d8b41f@38.242.208.183:5050",
		"enode://5c4c0a63c445cdd2756594d99c918ae828cb5cef0da14a64b8386279a710c5947de1ff34a5683a71560c9d37c63126d17e8332938748d6d034e83712229d3a08@38.242.202.244:5050",
		"enode://d9370c2d4e36213b74710f9b7b405bfa54d6b3634beab4bed514d6719d2d8e2dc75fb368f8bdd362855013c01892b3a640b5334249b279873da8efd4e336839b@185.169.252.41:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
