package processor

import (
	"sort"
	"strconv"
	"time"

	web3types "github.com/alethio/web3-go/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
)

func (p *Processor) preprocess() error {
	err := p.parseBlockData()
	if err != nil {
		return errors.Wrap(err, "could not parse block data")
	}

	err = p.parseReceipts()
	if err != nil {
		return errors.Wrap(err, "could not parse receipts")
	}

	return nil
}

func (p *Processor) parseBlockData() error {
	raw := p.Raw.Block

	if raw.Miner == "" {
		raw.Miner = raw.Author
	}

	var b types.Block

	b.BlockHash = utils.Trim0x(raw.Hash)
	b.ParentBlockHash = utils.Trim0x(raw.ParentHash)

	// -- ints
	number, err := strconv.ParseInt(raw.Number, 0, 64)
	if err != nil {
		return errors.Wrap(err, "could not decode block number")
	}
	b.Number = number

	// --timestamp
	timestamp, err := strconv.ParseInt(raw.Timestamp, 0, 64)
	if err != nil {
		return errors.Wrap(err, "could not decode block timestamp")
	}
	b.BlockCreationTime = types.DatetimeToJSONUnix(time.Unix(timestamp, 0))

	p.Block = &b

	return nil
}

func (p *Processor) parseReceipts() error {
	for index, tx := range p.Raw.Block.Transactions {
		receipt := p.Raw.Receipts[index]

		storableTx, err := p.parseTx(tx, receipt)
		if err != nil {
			return err
		}

		for i, log := range receipt.Logs {
			logEntry, err := p.parseLog(*storableTx, log, i)
			if err != nil {
				return errors.Wrap(err, "could not parse log")
			}

			storableTx.LogEntries = append(storableTx.LogEntries, *logEntry)
		}

		sort.Sort(storableTx.LogEntries)

		p.Block.Txs = append(p.Block.Txs, *storableTx)
	}

	sort.Sort(p.Block.Txs)

	return nil
}

func (p *Processor) parseTx(tx web3types.Transaction, receipt web3types.Receipt) (*types.Tx, error) {
	sTx := &types.Tx{}
	sTx.IncludedInBlock = p.Block.Number
	sTx.BlockCreationTime = p.Block.BlockCreationTime

	sTx.TxHash = utils.NormalizeAddress(tx.Hash)
	sTx.From = utils.NormalizeAddress(tx.From)
	sTx.To = utils.NormalizeAddress(tx.To)
	if tx.To == "" {
		if contractAddress, ok := receipt.ContractAddress.(string); ok && contractAddress != "" {
			sTx.To = utils.NormalizeAddress(contractAddress)
			sTx.Creates = utils.NormalizeAddress(contractAddress)
		}
	}

	if tx.Creates != "" {
		sTx.Creates = utils.NormalizeAddress(tx.Creates)
	}

	sTx.MsgPayload = types.ByteArray(utils.Trim0x(tx.Input))
	sTx.TxLogsBloom = types.ByteArray(utils.Trim0x(receipt.LogsBloom))
	sTx.MsgStatus = receipt.Status

	// -- int
	txIndex, err := strconv.ParseInt(tx.TransactionIndex, 0, 32)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode transaction index")
	}
	sTx.TxIndex = txIndex

	txNonce, err := strconv.ParseInt(tx.Nonce, 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode tx nonce")
	}
	sTx.TxNonce = txNonce

	// -- bigint
	gasLimit, err := utils.HexStrToBigIntStr(tx.Gas)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode gas")
	}
	sTx.MsgGasLimit = gasLimit

	value, err := utils.HexStrToBigIntStr(tx.Value)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode tx value")
	}
	sTx.Value = value

	gasUsed, err := utils.HexStrToBigIntStr(receipt.GasUsed)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode tx gas used")
	}
	sTx.TxGasUsed = gasUsed

	gasPrice, err := utils.HexStrToBigIntStr(tx.GasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode tx gas price")
	}
	sTx.TxGasPrice = gasPrice

	cumulativeGasUsed, err := utils.HexStrToBigIntStr(receipt.CumulativeGasUsed)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode tx cumulative gas used")
	}
	sTx.CumulativeGasUsed = cumulativeGasUsed

	return sTx, nil
}

func (p *Processor) parseLog(tx types.Tx, log web3types.Log, index int) (*types.LogEntry, error) {
	l := &types.LogEntry{}
	l.IncludedInBlock = tx.IncludedInBlock
	l.TxHash = utils.Trim0x(tx.TxHash)
	l.LogIndex = int32(index) // = transaction log index

	l.LogData = types.ByteArray(utils.Trim0x(log.Data))
	l.LoggedBy = utils.NormalizeAddress(log.Address)

	if len(log.Topics) > 0 {
		l.Topic0 = utils.Trim0x(log.Topics[0])
	}

	if len(log.Topics) > 1 {
		l.Topic1 = utils.Trim0x(log.Topics[1])
	}

	if len(log.Topics) > 2 {
		l.Topic2 = utils.Trim0x(log.Topics[2])
	}

	if len(log.Topics) > 3 {
		l.Topic3 = utils.Trim0x(log.Topics[3])
	}

	return l, nil
}
