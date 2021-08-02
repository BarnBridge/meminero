package smartalpha

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/barnbridge/meminero/processor/storables/smartalpha"
	"github.com/barnbridge/meminero/utils"
)

type SmartAlpha struct {
	db *pgxpool.Pool

	Pools []smartalpha.Pool
}

func New(db *pgxpool.Pool) *SmartAlpha {
	return &SmartAlpha{
		db: db,
	}
}

func (sa *SmartAlpha) PoolByAddress(addr string) *smartalpha.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sa.Pools {
		if p.PoolAddress == addr {
			return &p
		}
	}

	return nil
}

func (sa *SmartAlpha) IsERC20OfInterest(addr string) bool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sa.Pools {
		if p.JuniorTokenAddress == addr || p.SeniorTokenAddress == addr {
			return true
		}
	}

	return false
}

func (sa *SmartAlpha) PoolByJuniorTokenAddress(addr string) *smartalpha.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sa.Pools {
		if p.JuniorTokenAddress == addr {
			return &p
		}
	}

	return nil
}

func (sa *SmartAlpha) PoolBySeniorTokenAddress(addr string) *smartalpha.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sa.Pools {
		if p.SeniorTokenAddress == addr {
			return &p
		}
	}

	return nil
}
