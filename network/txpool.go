package network

import (
	"github.com/KumazakiRyoha/blockchain/core"
	"github.com/KumazakiRyoha/blockchain/types"
)

type TxPool struct {
	Transactions map[types.Hash]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		Transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	p.Transactions[hash] = tx
	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.Transactions[hash]
	return ok
}

func (p *TxPool) Len() int {
	return len(p.Transactions)
}

func (p *TxPool) Flush() {
	p.Transactions = make(map[types.Hash]*core.Transaction)
}
