package core

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/storage"
	"github.com/medibloc/go-medibloc/util"
)

type states struct {
	accState   *AccountStateBatch
	txsState   *TrieBatch
	usageState *TrieBatch

	storage storage.Storage
}

func newStates(stor storage.Storage) (*states, error) {
	accState, err := NewAccountStateBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	txsState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	usageState, err := NewTrieBatch(nil, stor)
	if err != nil {
		return nil, err
	}

	return &states{
		accState:   accState,
		txsState:   txsState,
		usageState: usageState,
		storage:    stor,
	}, nil
}

func (st *states) Clone() (*states, error) {
	accState, err := NewAccountStateBatch(st.accState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	txsState, err := NewTrieBatch(st.txsState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	usageState, err := NewTrieBatch(st.usageState.RootHash(), st.storage)
	if err != nil {
		return nil, err
	}

	return &states{
		accState:   accState,
		txsState:   txsState,
		usageState: usageState,
		storage:    st.storage,
	}, nil
}

func (st *states) BeginBatch() error {
	if err := st.accState.BeginBatch(); err != nil {
		return err
	}
	if err := st.txsState.BeginBatch(); err != nil {
		return err
	}
	if err := st.usageState.BeginBatch(); err != nil {
		return err
	}
	return nil
}

func (st *states) RollBack() error {
	if err := st.accState.RollBack(); err != nil {
		return err
	}
	if err := st.txsState.RollBack(); err != nil {
		return err
	}
	if err := st.usageState.RollBack(); err != nil {
		return err
	}
	return nil
}

func (st *states) Commit() error {
	if err := st.accState.Commit(); err != nil {
		return err
	}
	if err := st.txsState.Commit(); err != nil {
		return err
	}
	if err := st.usageState.Commit(); err != nil {
		return err
	}
	return nil
}

func (st *states) AccountsRoot() common.Hash {
	return common.BytesToHash(st.accState.RootHash())
}

func (st *states) TransactionsRoot() common.Hash {
	return common.BytesToHash(st.txsState.RootHash())
}

func (st *states) UsageRoot() common.Hash {
	return common.BytesToHash(st.usageState.RootHash())
}

func (st *states) LoadAccountsRoot(rootHash common.Hash) error {
	accState, err := NewAccountStateBatch(rootHash.Bytes(), st.storage)
	if err != nil {
		return err
	}
	st.accState = accState
	return nil
}

func (st *states) LoadTransactionsRoot(rootHash common.Hash) error {
	txsState, err := NewTrieBatch(rootHash.Bytes(), st.storage)
	if err != nil {
		return err
	}
	st.txsState = txsState
	return nil
}

func (st *states) LoadUsageRoot(rootHash common.Hash) error {
	usageState, err := NewTrieBatch(rootHash.Bytes(), st.storage)
	if err != nil {
		return err
	}
	st.usageState = usageState
	return nil
}

func (st *states) GetAccount(address common.Address) (Account, error) {
	return st.accState.GetAccount(address.Bytes())
}

func (st *states) AddBalance(address common.Address, amount *util.Uint128) error {
	return st.accState.AddBalance(address.Bytes(), amount)
}

func (st *states) AddWriter(address common.Address, writer common.Address) error {
	return st.accState.AddWriter(address.Bytes(), writer.Bytes())
}

func (st *states) RemoveWriter(address common.Address, writer common.Address) error {
	return st.accState.RemoveWriter(address.Bytes(), writer.Bytes())
}

func (st *states) SubBalance(address common.Address, amount *util.Uint128) error {
	return st.accState.SubBalance(address.Bytes(), amount)
}

func (st *states) incrementNonce(address common.Address) error {
	return st.accState.IncrementNonce(address.Bytes())
}

func (st *states) GetTx(txHash common.Hash) ([]byte, error) {
	return st.txsState.Get(txHash.Bytes())
}

func (st *states) PutTx(txHash common.Hash, txBytes []byte) error {
	return st.txsState.Put(txHash.Bytes(), txBytes)
}

func (st *states) updateUsage(tx *Transaction, blockTime int64) error {
	weekSec := int64(604800)

	if tx.Timestamp() < blockTime-weekSec {
		return ErrTooOldTransaction
	}

	usageBytes, err := st.usageState.Get(tx.from.Bytes())
	switch err {
	case nil:
	case ErrNotFound:
		usage := &corepb.Usage{
			Timestamps: []*corepb.TxTimestamp{
				{
					Hash:      tx.Hash().Bytes(),
					Timestamp: tx.Timestamp(),
				},
			},
		}
		usageBytes, err = proto.Marshal(usage)
		return st.usageState.Put(tx.from.Bytes(), usageBytes)
	default:
		return err
	}

	pbUsage := new(corepb.Usage)
	if err := proto.Unmarshal(usageBytes, pbUsage); err != nil {
		return err
	}

	var idx int
	for idx = range pbUsage.Timestamps {
		if blockTime-weekSec < tx.Timestamp() {
			break
		}
	}
	pbUsage.Timestamps = append(pbUsage.Timestamps[idx:], &corepb.TxTimestamp{Hash: tx.Hash().Bytes(), Timestamp: tx.Timestamp()})
	sort.Slice(pbUsage.Timestamps, func(i, j int) bool { return pbUsage.Timestamps[i].Timestamp < pbUsage.Timestamps[j].Timestamp })

	pbBytes, err := proto.Marshal(pbUsage)
	if err != nil {
		return err
	}

	return st.usageState.Put(tx.from.Bytes(), pbBytes)
}

func (st *states) GetUsage(addr common.Address) ([]*corepb.TxTimestamp, error) {
	usageBytes, err := st.usageState.Get(addr.Bytes())
	switch err {
	case nil:
	case ErrNotFound:
		return []*corepb.TxTimestamp{}, nil
	default:
		return nil, err
	}

	pbUsage := new(corepb.Usage)
	if err := proto.Unmarshal(usageBytes, pbUsage); err != nil {
		return nil, err
	}
	return pbUsage.Timestamps, nil
}

// BlockState possesses every states a block should have
type BlockState struct {
	*states
	snapshot *states
}

// NewBlockState creates a new block state
func NewBlockState(stor storage.Storage) (*BlockState, error) {
	states, err := newStates(stor)
	if err != nil {
		return nil, err
	}
	return &BlockState{
		states:   states,
		snapshot: nil,
	}, nil
}

// Clone clones block state
func (bs *BlockState) Clone() (*BlockState, error) {
	states, err := bs.states.Clone()
	if err != nil {
		return nil, err
	}
	return &BlockState{
		states:   states,
		snapshot: nil,
	}, nil
}

// BeginBatch begins batch
func (bs *BlockState) BeginBatch() error {
	snapshot, err := bs.states.Clone()
	if err != nil {
		return err
	}
	if err := bs.states.BeginBatch(); err != nil {
		return err
	}
	bs.snapshot = snapshot
	return nil
}

// RollBack rolls back batch
func (bs *BlockState) RollBack() error {
	bs.states = bs.snapshot
	bs.snapshot = nil
	return nil
}

// Commit saves batch updates
func (bs *BlockState) Commit() error {
	if err := bs.states.Commit(); err != nil {
		return err
	}
	bs.snapshot = nil
	return nil
}

// ExecuteTx and update internal states
func (bs *BlockState) ExecuteTx(tx *Transaction) error {
	return tx.ExecuteOnState(bs)
}

// AcceptTransaction and update internal txsStates
func (bs *BlockState) AcceptTransaction(tx *Transaction, blockTime int64) error {
	pbTx, err := tx.ToProto()
	if err != nil {
		return err
	}

	txBytes, err := proto.Marshal(pbTx)
	if err != nil {
		return err
	}

	if err := bs.PutTx(tx.hash, txBytes); err != nil {
		return err
	}

	if err := bs.updateUsage(tx, blockTime); err != nil {
		return err
	}

	if err = bs.incrementNonce(tx.from); err != nil {
		return err
	}
	return nil
}

func (bs *BlockState) checkNonce(tx *Transaction) error {
	fromAcc, err := bs.GetAccount(tx.from)
	if err != nil {
		return err
	}

	expectedNonce := fromAcc.Nonce() + 1
	if tx.nonce > expectedNonce {
		return ErrLargeTransactionNonce
	} else if tx.nonce < expectedNonce {
		return ErrSmallTransactionNonce
	}
	return nil
}