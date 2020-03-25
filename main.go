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
package main

import (
	"flag"
	"math/rand"
	"strings"
	"time"

	log4 "github.com/alecthomas/log4go"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ontio/cross_chain_test/config"
	_ "github.com/ontio/cross_chain_test/testcase"
	"github.com/ontio/cross_chain_test/testframework"
	"github.com/ontio/cross_chain_test/utils"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
)

var (
	TestConfig string //Test config file
	LogConfig  string //Log config file
	TestCases  string //TestCase list in cmdline
)

func init() {
	flag.StringVar(&TestConfig, "cfg", "./config.json", "Config of cross_chain_test")
	flag.StringVar(&LogConfig, "lfg", "./log4go.xml", "Log config of cross_chain_test")
	flag.StringVar(&TestCases, "t", "", "Test case to run. use ',' to split test case")
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log4.LoadConfiguration(LogConfig)
	defer time.Sleep(time.Second)

	err := config.DefConfig.Init(TestConfig)
	if err != nil {
		log4.Error("DefConfig.Init error:%s", err)
		return
	}

	ontSdk := ontology_go_sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(config.DefConfig.OntJsonRpcAddress)

	ethClient, err := ethclient.Dial(config.DefConfig.EthURL)
	if err != nil {
		log4.Error("cannot dial sync node, err: %s", err)
		return
	}

	cli := utils.NewRestCli(config.DefConfig.BtcRestAddr, config.DefConfig.BtcRestUser, config.DefConfig.BtcRestPwd)

	testCases := make([]string, 0)
	if TestCases != "" {
		testCases = strings.Split(TestCases, ",")
	}
	testframework.TFramework.SetOntSdk(ontSdk)
	testframework.TFramework.SetEthClient(ethClient)
	testframework.TFramework.SetBtcCli(cli)
	//Start run test case
	testframework.TFramework.Start(testCases)
}
