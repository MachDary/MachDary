package validation

import (
	"github.com/MachDary/MachDary/protocol/bc"
	"github.com/MachDary/MachDary/protocol/vm"
	evm_common "github.com/ethereum/go-ethereum/common"
	evm_state "github.com/ethereum/go-ethereum/core/state"
	"math/big"
	"math"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/MachDary/MachDary/consensus/segwit"
)

func EstimateContractGas(e bc.Entry, tx *bc.Tx, block *bc.Block, chain vm.ChainContext, stateDB *evm_state.StateDB) (gasStatus *GasState, err error) {

	stateDB.Prepare(tx.ID.Byte32(), evm_common.Hash{}, 0)

	gasStatus = &GasState{GasValid: true, GasLeft: math.MaxInt64}

	vs := &ValidationState{
		chain:     chain,
		stateDB:   stateDB,
		block:     block,
		tx:        tx,
		entryID:   tx.ID,
		gasStatus: gasStatus,
		cache:     make(map[bc.Hash]error),
	}

	gasLeft := int64(0)
	var args [][]byte

	switch e := e.(type) {
	case *bc.Creation:
		if vm.IsOpCreate(e.Input.Code) {
			from, err := segwit.GetHashFromStandardProg(e.From.Code)
			if err != nil {
				return nil, err
			}
			args = append(args, from)
			args = append(args, new(big.Int).SetUint64(e.Nonce).Bytes())
			_, gasLeft, err = vm.Verify(NewTxVMContext(vs, e, e.Input, args), vs.gasStatus.GasLeft)
		}
	case *bc.Call:
		if vm.IsOpCall(e.Input.Code) {
			from, err := segwit.GetHashFromStandardProg(e.From.Code)
			if err != nil {
				return nil, err
			}
			args = append(args, from)
			args = append(args, new(big.Int).SetUint64(e.Nonce).Bytes())
			args = append(args, e.To.Code)
			_, gasLeft, err = vm.Verify(NewTxVMContext(vs, e, e.Input, args), vs.gasStatus.GasLeft)
		}
	case *bc.Contract:
		if vm.IsOpContract(e.Input.Code) {
			from, err := segwit.GetHashFromStandardProg(e.From.Code)
			if err != nil {
				return nil, err
			}
			args = append(args, from)
			args = append(args, new(big.Int).SetUint64(e.Nonce).Bytes())
			args = append(args, e.To)
			_, gasLeft, err = vm.Verify(NewTxVMContext(vs, e, e.Input, args), vs.gasStatus.GasLeft)
		}
	default:
		return nil, errors.New("unknown program")
	}

	if err != nil {
		return gasStatus, err
	}

	log.WithField("gasUsed", gasStatus.GasLeft-gasLeft).Println("EstimateContractGas")
	err = gasStatus.updateUsage(gasLeft)
	if err != nil {
		return gasStatus, err
	}

	log.WithField("gasUsed", gasStatus.GasUsed).Println("EstimateContractGas")
	return gasStatus, nil

}
