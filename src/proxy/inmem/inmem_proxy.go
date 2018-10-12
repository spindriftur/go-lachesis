package inmem

import (
	"github.com/andrecronje/lachesis/src/poset"
	"github.com/sirupsen/logrus"
)

//InmemProxy implements the AppProxy interface natively
type InmemProxy struct {
	commitHandler   CommitHandler
	snapshotHandler SnapshotHandler
	restoreHandler  RestoreHandler
	submitCh        chan []byte
	logger          *logrus.Logger
}

// NewInmemProxy instantiates an InmemProxy from a set of handlers.
// If no logger, a new one is created
func NewInmemProxy(commitHandler CommitHandler,
	snapshotHandler SnapshotHandler,
	restoreHandler RestoreHandler,
	logger *logrus.Logger) *InmemProxy {

	if logger == nil {
		logger = logrus.New()
		logger.Level = logrus.DebugLevel
	}

	return &InmemProxy{
		commitHandler:   commitHandler,
		snapshotHandler: snapshotHandler,
		restoreHandler:  restoreHandler,
		submitCh:        make(chan []byte),
		logger:          logger,
	}
}

/*******************************************************************************
* SubmitTx                                                                     *
*******************************************************************************/

//SubmitTx is called by the App to submit a transaction to Lachesis
func (p *InmemProxy) SubmitTx(tx []byte) {
	//have to make a copy, or the tx will be garbage collected and weird stuff
	//happens in transaction pool
	t := make([]byte, len(tx), len(tx))
	copy(t, tx)
	p.submitCh <- t
}

/*******************************************************************************
* Implement AppProxy Interface                                                 *
*******************************************************************************/

//SubmitCh returns the channel of raw transactions
func (p *InmemProxy) SubmitCh() chan []byte {
	return p.submitCh
}

//CommitBlock calls the commitHandler
func (p *InmemProxy) CommitBlock(block poset.Block) ([]byte, error) {

	stateHash, err := p.commitHandler(block)

	p.logger.WithFields(logrus.Fields{
		"round_received": block.RoundReceived(),
		"txs":            len(block.Transactions()),
		"state_hash":     stateHash,
		"err":            err,
	}).Debug("InmemProxy.CommitBlock")

	return stateHash, err
}

//GetSnapshot calls the snapshotHandler
func (p *InmemProxy) GetSnapshot(blockIndex int) ([]byte, error) {

	snapshot, err := p.snapshotHandler(blockIndex)

	p.logger.WithFields(logrus.Fields{
		"block":    blockIndex,
		"snapshot": snapshot,
		"err":      err,
	}).Debug("InmemProxy.GetSnapshot")

	return snapshot, err
}

//Restore calls the restoreHandler
func (p *InmemProxy) Restore(snapshot []byte) error {

	stateHash, err := p.restoreHandler(snapshot)

	p.logger.WithFields(logrus.Fields{
		"state_hash": stateHash,
		"err":        err,
	}).Debug("InmemProxy.Restore")

	return err
}