/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

//common use fot ontology-tool
package config

import (
	"encoding/json"
	"fmt"
	log4 "github.com/alecthomas/log4go"
	"io/ioutil"
	"os"
)

//Default config instance
var DefConfig = NewTestConfig()

const (
	ONT_CHAIN_ID = 3
	ETH_CHAIN_ID = 2
	BTC_CHAIN_ID = 1
)

//Config object used by ontology-instance
type TestConfig struct {
	BtcRestAddr         string
	BtcRestUser         string
	BtcRestPwd          string
	BtceContractAddress string
	BtcoContractAddress string
	BtcFee              int64
	BtcRedeem           string

	EthURL                       string
	EthProxyContract             string
	EthCrossChainContractAddress string

	//JsonRpcAddress of ontology
	OntJsonRpcAddress string
	//password of ont wallet
	OntWalletPassword string
	//ontology proxy contract address
	OntProxyContractAddress string
	GasPrice                uint64
	GasLimit                uint64

	RchainJsonRpcAddress string

	ReportInterval uint64
}

//NewTestConfig retuen a TestConfig instance
func NewTestConfig() *TestConfig {
	return &TestConfig{}
}

//Init TestConfig with a config file
func (this *TestConfig) Init(fileName string) error {
	err := this.loadConfig(fileName)
	if err != nil {
		return fmt.Errorf("loadConfig error:%s", err)
	}
	return nil
}

func (this *TestConfig) loadConfig(fileName string) error {
	data, err := this.readFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}

func (this *TestConfig) readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log4.Error("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}
