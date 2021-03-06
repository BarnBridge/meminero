package processor

import (
	"sort"
	"strconv"

	web3types "github.com/alethio/web3-go/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/lacasian/ethwheels/ethgen"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
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
	b.BlockCreationTime = timestamp

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

		for _, log := range receipt.Logs {
			logEntry, err := p.parseLog(log)
			if err != nil {
				return errors.Wrap(err, "could not parse log")
			}

			storableTx.LogEntries = append(storableTx.LogEntries, logEntry)
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

func (p *Processor) parseLog(log web3types.Log) (gethtypes.Log, error) {
	return ethgen.W3LogToLog(log)
}
