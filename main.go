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
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	log4 "github.com/alecthomas/log4go"
	"github.com/ontio/cross_chain_test/config"
	"github.com/ontio/cross_chain_test/testcase"
	_ "github.com/ontio/cross_chain_test/testcase"
	"github.com/ontio/cross_chain_test/testframework"
	"github.com/ontio/cross_chain_test/utils"
	multi_chain_go_sdk "github.com/ontio/multi-chain-go-sdk"
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

	rcSdk := multi_chain_go_sdk.NewMultiChainSdk()
	rcSdk.NewRpcClient().SetAddress(config.DefConfig.RchainJsonRpcAddress)

	ethTools := utils.NewEthTools(config.DefConfig.EthURL)

	noncemanager := utils.NewNonceManager(ethTools.GetEthClient())
	cli := utils.NewRestCli(config.DefConfig.BtcRestAddr, config.DefConfig.BtcRestUser, config.DefConfig.BtcRestPwd)

	testCases := make([]string, 0)
	if TestCases != "" {
		testCases = strings.Split(TestCases, ",")
	}
	testframework.TFramework.SetOntSdk(ontSdk)
	testframework.TFramework.SetRcSdk(rcSdk)
	testframework.TFramework.SetEthTools(ethTools)
	testframework.TFramework.SetNonceManager(noncemanager)
	testframework.TFramework.SetBtcCli(cli)

	testcase.InitAccount()
	//Start run test case
	testframework.TFramework.Start(testCases)

	waitToExit()
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sc {
			fmt.Println("cross chain test received exit signal: ", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}
