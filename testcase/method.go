package testcase

import (
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/cross_chain_test/common"
	"github.com/ontio/cross_chain_test/testframework"
	"github.com/ontio/ont-eth-ft-tool/config"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	ontcommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func GetAccountByPath(path string) (*ontology_go_sdk.Account, error) {
	wallet, err := ontology_go_sdk.OpenWallet(path)
	if err != nil {
		return nil, fmt.Errorf("ontology_go_sdk.OpenWallet %s error:%s", path, err)
	}
	account, err := wallet.GetDefaultAccount([]byte(common.DefConfig.OntWalletPassword))
	if err != nil {
		return nil, fmt.Errorf("wallet.GetDefaultAccount error:%s", err)
	}
	return account, nil
}

func ApproveOEP4(ctx *testframework.TestFrameworkContext, contractAddress string, signer *ontology_go_sdk.Account,
	amount uint64) error {
	assetContractAddress, err := ontcommon.AddressFromHexString(contractAddress)
	if err != nil {
		return fmt.Errorf("ApproveOEP4, ontcommon.AddressFromHexString error: %s", err)
	}
	proxyContractAddress, err := ontcommon.AddressFromHexString(common.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("ApproveOEP4, ontcommon.AddressFromHexString error: %s", err)
	}
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(common.DefConfig.GasPrice, common.DefConfig.GasLimit,
		signer,
		signer,
		assetContractAddress,
		[]interface{}{"approve", []interface{}{signer.Address[:], proxyContractAddress[:], amount}})
	if err != nil {
		return fmt.Errorf("ApproveOEP4, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	return nil
}

func SendOntCrossEth(ctx *testframework.TestFrameworkContext, signer *ontology_go_sdk.Account,
	ethAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(common.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(common.DefConfig.GasPrice, common.DefConfig.GasLimit,
		signer,
		signer,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{utils.OntContractAddress[:], signer.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	return nil
}

func SendOngCrossEth(ctx *testframework.TestFrameworkContext, signer *ontology_go_sdk.Account,
	ethAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(common.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(common.DefConfig.GasPrice, common.DefConfig.GasLimit,
		signer,
		signer,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{utils.OngContractAddress[:], signer.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	return nil
}

func SendOEP4CrossEth(ctx *testframework.TestFrameworkContext, contractAddress string, signer *ontology_go_sdk.Account,
	ethAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(common.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	assetContractAddress, err := ontcommon.AddressFromHexString(contractAddress)
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(common.DefConfig.GasPrice, common.DefConfig.GasLimit,
		signer,
		signer,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetContractAddress[:], signer.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	return nil
}
