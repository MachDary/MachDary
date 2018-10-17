package config

import (
	"testing"
	"github.com/MachDary/MachDary/consensus/difficulty"
	"github.com/MachDary/MachDary/consensus"
)

func TestGenerateGenesisBlock(t *testing.T) {
	consensus.ActiveNetParams = consensus.MainNetParams
	block := GenesisBlock()
	nonce := block.Nonce
	for {
		hash := block.Hash()
		if difficulty.CheckProofOfWork(&hash, consensus.InitialSeed, block.Bits) {
			break
		}
		block.Nonce++
	}
	if block.Nonce != nonce {
		t.Errorf("correct nonce is %d, but get %d", block.Nonce, nonce)
	}
}
