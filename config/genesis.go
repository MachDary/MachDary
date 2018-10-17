package config

import (
	"encoding/hex"

	log "github.com/sirupsen/logrus"

	"github.com/MachDary/MachDary/consensus"
	"github.com/MachDary/MachDary/protocol/bc"
	"github.com/MachDary/MachDary/protocol/bc/types"
)

func genesisTx() *types.Tx {
	contract, err := hex.DecodeString("0014678f9a43d1de0809ff2bbf9b00312a166dfacce8")
	if err != nil {
		log.Panicf("fail on decode genesis tx output control program")
	}

	txData := types.TxData{
		Version: 1,
		Inputs: []*types.TxInput{
			types.NewCoinbaseInput([]byte("Knowledge is power. Learning to get rid of poverty. -- Sep/01/2018.")),
		},
		Outputs: []*types.TxOutput{
			types.NewTxOutput(*consensus.NativeAssetID, consensus.InitialBlockSubsidy, contract),
		},
	}
	return types.NewTx(txData)
}

func mainNetGenesisBlock() *types.Block {
	tx := genesisTx()
	txStatus := bc.NewTransactionStatus()
	txStatus.SetStatus(0, false)
	txStatusHash, err := types.TxStatusMerkleRoot(txStatus.VerifyStatus)
	if err != nil {
		log.Panicf("fail on calc genesis tx status merkle root")
	}

	merkleRoot, err := types.TxMerkleRoot([]*bc.Tx{tx.Tx})
	if err != nil {
		log.Panicf("fail on calc genesis tx merkel root")
	}

	block := &types.Block{
		BlockHeader: types.BlockHeader{
			Version:   1,
			Height:    0,
			Nonce:     1530935882,
			Timestamp: 1539673382,
			Bits:      2305843009214892324,
			BlockCommitment: types.BlockCommitment{
				TransactionsMerkleRoot: merkleRoot,
				TransactionStatusHash:  txStatusHash,
				StateRoot:              bc.Hash{},
			},
		},
		Transactions: []*types.Tx{tx},
	}
	if SupportBalanceInStateDB {
		block.Nonce = 1530935912
	}
	return block
}

func testNetGenesisBlock() *types.Block {
	tx := genesisTx()
	txStatus := bc.NewTransactionStatus()
	txStatus.SetStatus(0, false)
	txStatusHash, err := types.TxStatusMerkleRoot(txStatus.VerifyStatus)
	if err != nil {
		log.Panicf("fail on calc genesis tx status merkle root")
	}

	merkleRoot, err := types.TxMerkleRoot([]*bc.Tx{tx.Tx})
	if err != nil {
		log.Panicf("fail on calc genesis tx merkel root")
	}

	block := &types.Block{
		BlockHeader: types.BlockHeader{
			Version:   1,
			Height:    0,
			Nonce:     1530936095,
			Timestamp: 1539673382,
			Bits:      2305843009214892324,
			BlockCommitment: types.BlockCommitment{
				TransactionsMerkleRoot: merkleRoot,
				TransactionStatusHash:  txStatusHash,
				StateRoot:              bc.Hash{},
			},
		},
		Transactions: []*types.Tx{tx},
	}
	if SupportBalanceInStateDB {
		block.Nonce = 1530936107
	}
	return block
}

func soloNetGenesisBlock() *types.Block {
	tx := genesisTx()
	txStatus := bc.NewTransactionStatus()
	txStatus.SetStatus(0, false)
	txStatusHash, err := types.TxStatusMerkleRoot(txStatus.VerifyStatus)
	if err != nil {
		log.Panicf("fail on calc genesis tx status merkle root")
	}

	merkleRoot, err := types.TxMerkleRoot([]*bc.Tx{tx.Tx})
	if err != nil {
		log.Panicf("fail on calc genesis tx merkel root")
	}

	block := &types.Block{
		BlockHeader: types.BlockHeader{
			Version:   1,
			Height:    0,
			Nonce:     68,
			Timestamp: 1539673382,
			Bits:      2305843009214892324,
			BlockCommitment: types.BlockCommitment{
				TransactionsMerkleRoot: merkleRoot,
				TransactionStatusHash:  txStatusHash,
				StateRoot:              bc.Hash{},
			},
		},
		Transactions: []*types.Tx{tx},
	}
	if SupportBalanceInStateDB {
		block.Nonce = 85
	}
	return block
}

// GenesisBlock will return genesis block
func GenesisBlock() *types.Block {
	return map[string]func() *types.Block{
		"main": mainNetGenesisBlock,
		"test": testNetGenesisBlock,
		"solo": soloNetGenesisBlock,
	}[consensus.ActiveNetParams.Name]()
}

var SupportBalanceInStateDB = false

func GenesisBlockHash() *bc.Hash {
	if !SupportBalanceInStateDB {
		return map[string]*bc.Hash{
			"main": {
				V0: uint64(5296793296687494021),
				V1: uint64(259185367918554011),
				V2: uint64(6617776325189918720),
				V3: uint64(6385282310144889357),
			},
			"test": {
				V0: uint64(8902951982539050259),
				V1: uint64(5181020502328952258),
				V2: uint64(6447955034809385332),
				V3: uint64(6167668015356892148),
			},
			"solo": {
				V0: uint64(15265895030219202748),
				V1: uint64(15126226954191477309),
				V2: uint64(6984997729432543837),
				V3: uint64(14619828310239418552),
			},
		}[consensus.ActiveNetParams.Name]
	} else {
		return map[string]*bc.Hash{
			"main": {
				V0: uint64(8325532997334157898),
				V1: uint64(4189984660549501270),
				V2: uint64(5027510122468721539),
				V3: uint64(18379455307324088015),
			},
			"test": {
				V0: uint64(16672815734135948520),
				V1: uint64(14559210567994881926),
				V2: uint64(426186010121443363),
				V3: uint64(2053255217068846892),
			},
			"solo": {
				V0: uint64(195960844143915746),
				V1: uint64(17468542655660531027),
				V2: uint64(17784038276451449838),
				V3: uint64(661814422175617024),
			},
		}[consensus.ActiveNetParams.Name]
	}
}
