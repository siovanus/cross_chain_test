package testcase

import (
	"encoding/hex"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/cross_chain_test/config"
	"github.com/ontio/cross_chain_test/testframework"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	ontcommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func GetAccountByPath(path string) (*ontology_go_sdk.Account, error) {
	wallet, err := ontology_go_sdk.OpenWallet(path)
	if err != nil {
		return nil, fmt.Errorf("ontology_go_sdk.OpenWallet %s error:%s", path, err)
	}
	account, err := wallet.GetDefaultAccount([]byte(config.DefConfig.OntWalletPassword))
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
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("ApproveOEP4, ontcommon.AddressFromHexString error: %s", err)
	}
	_, err = ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
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
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	_, err = ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
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
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	_, err = ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
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
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	assetContractAddress, err := ontcommon.AddressFromHexString(contractAddress)
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	_, err = ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
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

func SendBtcxCrossBtc(ctx *testframework.TestFrameworkContext, signer *ontology_go_sdk.Account,
	btcAddress string, amount uint64) error {
	btcxContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.BtcxContractAddress)
	if err != nil {
		return fmt.Errorf("SendBtcxCrossBtc, ontcommon.AddressFromHexString error: %s", err)
	}
	to := hex.EncodeToString([]byte(btcAddress))
	_, err = ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		signer,
		signer,
		btcxContractAddress,
		[]interface{}{"lock", []interface{}{config.ETH_CHAIN_ID, signer.Address[:], to, amount}})
	if err != nil {
		return fmt.Errorf("SendBtcxCrossBtc, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	return nil
}

//func SendEthCrossOnt(ctx *testframework.TestFrameworkContext, ontAddress string, amount uint64) error {
//	gasPrice, err := ctx.EthClient.SuggestGasPrice(context.Background())
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, get suggest gas price failed error: %s", err.Error())
//	}
//	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, abi.JSON error:" + err.Error())
//	}
//	cross2OntAddress, err := ontcommon.AddressFromBase58(ontAddress)
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, ontcommon.AddressFromBase58 error:" + err.Error())
//	}
//	assetaddress := ethcommon.HexToAddress("0000000000000000000000000000000000000000")
//	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.ONT_CHAIN_ID), cross2OntAddress[:],
//		big.NewInt(int64(amount)))
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, contractabi.Pack error:" + err.Error())
//	}
//
//	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthProxyContract)
//	signerAccount := self.ethSigner.GetAccount(index)
//	callMsg := ethereum.CallMsg{
//		From: signerAccount.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
//		Value: big.NewInt(int64(self.cfg.Cross2OntAmount)), Data: txData,
//	}
//	gasLimit, err := ctx.EthClient.EstimateGas(context.Background(), callMsg)
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, estimate gas limit error: %s", err.Error())
//	}
//
//	nonceManager := tools.NewNonceManager(ctx.EthClient)
//	nonce := nonceManager.GetAddressNonce(signerAccount.Address)
//	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(self.cfg.Cross2OntAmount)), gasLimit, gasPrice, txData)
//	bf := new(bytes.Buffer)
//	rlp.Encode(bf, tx)
//
//	rawtx := hexutil.Encode(bf.Bytes())
//	signedtx, err := self.ethSigner.SignRawTx(rawtx, index)
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, sign raw tx error: %s", err.Error())
//	}
//
//	err = ctx.EthClient.SendTransaction(context.Background(), signedtx)
//	if err != nil {
//		return fmt.Errorf("SendEthCrossOnt, send transaction error:%s", err.Error())
//	}
//	WaitTransactionConfirm(ctx.EthClient, signedtx.Hash())
//
//	return nil
//}
//
//func WaitTransactionConfirm(ethclient *ethclient.Client, hash common.Hash) {
//	//
//	errNum := 0
//	for errNum < 100 {
//		time.Sleep(time.Second * 1)
//		_, ispending, err := ethclient.TransactionByHash(context.Background(), hash)
//		if err != nil {
//			log.Infof("query transaction %s error: %s", hash.String(), err.Error())
//			errNum++
//			continue
//		}
//		log.Infof("transaction %s is pending: %d", hash.String(), ispending)
//		if ispending == true {
//			continue
//		} else {
//			break
//		}
//	}
//}
