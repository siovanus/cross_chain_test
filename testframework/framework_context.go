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

package testframework

import (
	log4 "github.com/alecthomas/log4go"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ontio/cross_chain_test/utils"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
)

//TestFrameworkContext is the context for test case
type TestFrameworkContext struct {
	OntSdk    *ontology_go_sdk.OntologySdk //sdk to ontology
	EthClient *ethclient.Client
	BtcCli    *utils.RestCli
}

//NewTestFrameworkContext return a TestFrameworkContext instance
func NewTestFrameworkContext(ontSdk *ontology_go_sdk.OntologySdk, ethClient *ethclient.Client,
	btcCli *utils.RestCli) *TestFrameworkContext {
	return &TestFrameworkContext{
		OntSdk:    ontSdk,
		EthClient: ethClient,
		BtcCli:    btcCli,
	}
}

//LogInfo log info in test case
func (this *TestFrameworkContext) LogInfo(arg0 interface{}, args ...interface{}) {
	log4.Info(arg0, args...)
}

//LogError log error info  when error occur in test case
func (this *TestFrameworkContext) LogError(arg0 interface{}, args ...interface{}) {
	log4.Error(arg0, args...)
}

//LogWarn log warning info in test case
func (this *TestFrameworkContext) LogWarn(arg0 interface{}, args ...interface{}) {
	log4.Warn(arg0, args...)
}
