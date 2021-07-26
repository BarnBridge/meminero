package syncer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path"
	"reflect"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/db"
)

var log = logrus.WithField("module", "syncer")

type Data struct {
	Labels                Labels                `json:"labels"`
	MonitoredAccounts     MonitoredAccounts     `json:"monitored-accounts"`
	MonitoredERC20        MonitoredERC20        `json:"monitored-erc20"`
	SmartExposurePools    SmartExposurePools    `json:"smart-exposure-pools"`
	SmartYieldPools       SmartYieldPools       `json:"smart-yield-pools"`
	SmartYieldRewardPools SmartYieldRewardPools `json:"smart-yield-reward-pools"`
	Tokens                Tokens                `json:"tokens"`
}

type SyncAble interface {
	Sync(tx pgx.Tx) error
}

func readFileInto(file string, v interface{}) error {
	path := path.Join(config.Store.Syncer.Path, config.Store.Syncer.Network, file)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func Run() error {
	var data Data

	for _, set := range config.Store.Syncer.Datasets {
		file := set + ".json"
		err := readFileInto(file, &data)
		if err != nil {
			return errors.Wrapf(err, "could not read file %s", file)
		}
	}

	return syncDb(data)
}

func syncDb(data Data) error {
	d, err := db.New()
	if err != nil {
		return err
	}

	tx, err := d.Connection().Begin(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not start transaction")
	}

	x := reflect.ValueOf(data)

	for i := 0; i < x.NumField(); i++ {
		err := x.Field(i).Interface().(SyncAble).Sync(tx)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
