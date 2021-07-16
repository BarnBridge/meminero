package abi

import (
	"bytes"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"

	"github.com/barnbridge/smartbackend/config"
)

type store struct {
	abis map[string]abi.ABI
	mu   *sync.Mutex
}

var instance *store

func Store() *store {
	if instance == nil {
		instance = &store{
			abis: make(map[string]abi.ABI),
			mu:   new(sync.Mutex),
		}
	}

	return instance
}

func (s *store) Get(name string) (*abi.ABI, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	name = strings.ToLower(name)

	a, exists := s.abis[name]
	if exists {
		return &a, nil
	}

	// not found in memory, look for it in files
	files, err := ioutil.ReadDir(config.Store.EthTypes.AbiFolder)
	if err != nil {
		return nil, errors.Wrap(err, "could not read abis directory")
	}

	for _, file := range files {
		if strings.ToLower(file.Name()) == name+".json" {
			byteValue, err := ioutil.ReadFile(config.Store.EthTypes.AbiFolder + "/" + file.Name())
			if err != nil {
				return nil, errors.Wrap(err, "could not read abi file")
			}

			a, err := abi.JSON(bytes.NewReader(byteValue))
			if err != nil {
				return nil, errors.Wrap(err, "could not decode abi")
			}

			s.abis[name] = a

			return &a, nil
		}
	}

	return nil, nil
}
