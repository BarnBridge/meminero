package storable

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	web3types "github.com/alethio/web3-go/types"

	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
)

type UnclesStorable struct {
	Raw   *types.RawData
	Block *types.Block

	logger *logrus.Entry

	uncles []*Uncle
}

type Uncle struct {
	BlockHash         string
	IncludedInBlock   int64
	Number            int64
	BlockCreationTime types.DatetimeToJSONUnix
	UncleIndex        int32
	BlockGasLimit     string
	BlockGasUsed      string
	HasBeneficiary    types.ByteArray
	BlockDifficulty   string
	BlockExtraData    types.ByteArray
	BlockMixHash      types.ByteArray
	BlockNonce        types.ByteArray
	Sha3Uncles        types.ByteArray
}

func NewStorableUncles(block *types.Block, raw *types.RawData) *UnclesStorable {
	return &UnclesStorable{
		Block:  block,
		Raw:    raw,
		logger: logrus.WithField("module", "storable(uncles)"),
	}
}

func (s *UnclesStorable) Result() interface{} {
	return s.uncles
}

func (s *UnclesStorable) Rollback(tx *sql.Tx) error {
	_, err := tx.Exec(`delete from uncles where included_in_block = $1`, s.Block.Number)
	return err
}

func (s *UnclesStorable) Execute() error {
	if len(s.Raw.Uncles) == 0 {
		return nil
	}

	s.logger.Trace("processing uncles")
	start := time.Now()
	defer func() {
		s.logger.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"count":    len(s.uncles),
		}).Debug("done processing uncles")
	}()

	for index, uncle := range s.Raw.Uncles {
		storableUncle, err := s.buildStorableUncle(uncle, int32(index))
		if err != nil {
			return err
		}

		s.uncles = append(s.uncles, storableUncle)
	}

	return nil
}

func (s *UnclesStorable) SaveToDatabase(tx *sql.Tx) error {
	s.logger.Trace("storing uncles")
	start := time.Now()
	defer func() {
		s.logger.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"count":    len(s.uncles),
		}).Debug("done storing uncles")
	}()

	stmt, err := tx.Prepare(pq.CopyIn("uncles", "block_hash", "included_in_block", "number", "block_creation_time", "uncle_index", "block_gas_limit", "block_gas_used", "has_beneficiary", "block_difficulty", "block_extra_data", "block_mix_hash", "block_nonce", "sha3_uncles"))
	if err != nil {
		return err
	}

	for _, uncle := range s.uncles {
		_, err = stmt.Exec(uncle.BlockHash, uncle.IncludedInBlock, uncle.Number, uncle.BlockCreationTime, uncle.UncleIndex, uncle.BlockGasLimit, uncle.BlockGasUsed, uncle.HasBeneficiary, uncle.BlockDifficulty, uncle.BlockExtraData, uncle.BlockMixHash, uncle.BlockNonce, uncle.Sha3Uncles)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *UnclesStorable) buildStorableUncle(uncle web3types.Block, index int32) (*Uncle, error) {
	u := &Uncle{}
	u.IncludedInBlock = s.Block.Number
	u.UncleIndex = index

	if uncle.Miner == "" {
		uncle.Miner = uncle.Author
	}

	// -- raw
	u.BlockHash = utils.Trim0x(uncle.Hash)
	u.HasBeneficiary = types.ByteArray(utils.Trim0x(uncle.Miner))
	u.BlockExtraData = types.ByteArray(utils.Trim0x(uncle.ExtraData))
	u.BlockMixHash = types.ByteArray(utils.Trim0x(uncle.MixHash))
	u.BlockNonce = types.ByteArray(utils.Trim0x(uncle.Nonce))
	u.Sha3Uncles = types.ByteArray(utils.Trim0x(uncle.Sha3Uncles))

	// -- int64
	number, err := strconv.ParseInt(uncle.Number, 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode uncle number")
	}
	u.Number = number

	// -- hexes
	gasLimit, err := utils.HexStrToBigIntStr(uncle.GasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode gas limit")
	}
	u.BlockGasLimit = gasLimit

	gasUsed, err := utils.HexStrToBigIntStr(uncle.GasUsed)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode gas used")
	}
	u.BlockGasUsed = gasUsed

	difficulty, err := utils.HexStrToBigIntStr(uncle.Difficulty)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode difficulty")
	}
	u.BlockDifficulty = difficulty

	// -- timestamp
	timestamp, err := strconv.ParseInt(uncle.Timestamp, 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode timestamp")
	}
	u.BlockCreationTime = types.DatetimeToJSONUnix(time.Unix(timestamp, 0))

	return u, nil
}
