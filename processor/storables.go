package processor

import (
	"github.com/barnbridge/smartbackend/processor/storable"
)

// registerStorables instantiates all the storables defined via code with the requested raw data
// Only the storables that are registered will be executed when the Store function is called
func (p *Processor) registerStorables() {
	p.storables = append(p.storables, storable.NewStorableTxs(p.Block))
	p.storables = append(p.storables, storable.NewStorableUncles(p.Block, p.Raw))
	p.storables = append(p.storables, storable.NewStorableLogEntries(p.Block))
	p.storables = append(p.storables, storable.NewStorableAccountTxs(p.Block))
}
