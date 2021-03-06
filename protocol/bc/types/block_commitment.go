package types

import (
	"io"

	"github.com/MachDary/MachDary/basis/encoding/blockchain"
	"github.com/MachDary/MachDary/protocol/bc"
)

// BlockCommitment store the TransactionsMerkleRoot && TransactionStatusHash
type BlockCommitment struct {
	// TransactionsMerkleRoot is the root hash of the Merkle binary hash tree
	// formed by the hashes of all transactions included in the block.
	TransactionsMerkleRoot bc.Hash `json:"transaction_merkle_root"`

	// TransactionStatusHash is the root hash of the Merkle binary hash tree
	// formed by the hashes of all transaction verify results
	TransactionStatusHash bc.Hash `json:"transaction_status_hash"`

	StateRoot bc.Hash `json:"state_root"`
}

func (bc *BlockCommitment) readFrom(r *blockchain.Reader) error {
	if _, err := bc.TransactionsMerkleRoot.ReadFrom(r); err != nil {
		return err
	}

	if _, err := bc.TransactionStatusHash.ReadFrom(r); err != nil {
		return err
	}

	_, err := bc.StateRoot.ReadFrom(r)
	return err
}

func (bc *BlockCommitment) writeTo(w io.Writer) error {
	if _, err := bc.TransactionsMerkleRoot.WriteTo(w); err != nil {
		return err
	}

	if _, err := bc.TransactionStatusHash.WriteTo(w); err != nil {
		return err
	}

	_, err := bc.StateRoot.WriteTo(w)
	return err
}
