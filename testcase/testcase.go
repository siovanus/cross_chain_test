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
package testcase

import (
	"fmt"
	"github.com/ontio/cross_chain_test/testframework"
	"github.com/ontio/cross_chain_test/utils"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"os"
)

var ontSigner1 *ontology_go_sdk.Account
var ethSigner1 *utils.EthSigner

func InitAccount() {
	var err error
	ontSigner1, err = GetAccountByPath("ont_wallets/wallet1.dat")
	if err != nil {
		fmt.Printf("init, GetAccountByPath error: %s", err)
		os.Exit(1)
	}

	ethSigner1, err = utils.NewEthSigner("56B446A2DE5EDFCCEE1581FBBA79E8BB5C269E28AB4C0487860AFB7E2C2D2B6E")
	if err != nil {
		fmt.Printf("init, utils.NewEthSigner error: %s", err)
		os.Exit(1)
	}
}

func SendOntToEthChain(ctx *testframework.TestFrameworkContext) bool {
	err := SendOntCrossEth(ctx, ontSigner1, ethSigner1, 100)
	if err != nil {
		ctx.LogError("SendOntToEthChain, SendOntCrossEth error: %s", err)
		return false
	}
	return true
}
