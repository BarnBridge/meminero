package api

import (
	apiTypes "github.com/barnbridge/smartbackend/api/types"
	"github.com/barnbridge/smartbackend/types"
)

func (a *API) getBlockTxs(number int64) ([]apiTypes.Tx, error) {
	var txs = make([]apiTypes.Tx, 0)

	rows, err := a.db.Query(`select tx_index, tx_hash, value, "from", "to", msg_gas_limit, tx_gas_used, tx_gas_price from txs where included_in_block = $1 order by tx_index`, number)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			txIndex  int32
			txHash   string
			value    string
			from     types.ByteArray
			to       types.ByteArray
			gasLimit string
			gasUsed  string
			gasPrice string
		)

		err = rows.Scan(&txIndex, &txHash, &value, &from, &to, &gasLimit, &gasUsed, &gasPrice)
		if err != nil {
			a.logger.Error(err)
			return nil, err
		}

		txs = append(txs, apiTypes.Tx{
			TxIndex:     &txIndex,
			TxHash:      &txHash,
			Value:       &value,
			From:        &from,
			To:          &to,
			MsgGasLimit: &gasLimit,
			TxGasUsed:   &gasUsed,
			TxGasPrice:  &gasPrice,
		})
	}

	return txs, nil
}
