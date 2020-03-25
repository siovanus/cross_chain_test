package testcase

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/cross_chain_test/config"
	"github.com/ontio/cross_chain_test/testframework"
	"github.com/ontio/cross_chain_test/utils"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	ontcommon "github.com/ontio/ontology/common"
	nutils "github.com/ontio/ontology/smartcontract/service/native/utils"
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
	txHash, err := ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		signer,
		signer,
		assetContractAddress,
		[]interface{}{"approve", []interface{}{signer.Address[:], proxyContractAddress[:], amount}})
	if err != nil {
		return fmt.Errorf("ApproveOEP4, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	ctx.LogInfo("ApproveOEP4, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendOntCrossEth(ctx *testframework.TestFrameworkContext, signer *ontology_go_sdk.Account,
	ethAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	txHash, err := ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		signer,
		signer,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OntContractAddress[:], signer.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	ctx.LogInfo("SendOntCrossEth, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendOngCrossEth(ctx *testframework.TestFrameworkContext, signer *ontology_go_sdk.Account,
	ethAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntProxyContractAddress)
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ethcommon.HexToAddress(ethAddress).Bytes()
	txHash, err := ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		signer,
		signer,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OngContractAddress[:], signer.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	ctx.LogInfo("SendOngCrossEth, tx success, txHash is: %s", txHash.ToHexString())
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
	txHash, err := ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		signer,
		signer,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetContractAddress[:], signer.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	ctx.LogInfo("SendOEP4CrossEth, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendBtcxCrossBtc(ctx *testframework.TestFrameworkContext, signer *ontology_go_sdk.Account,
	btcAddress string, amount uint64) error {
	btcxContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.BtcxContractAddress)
	if err != nil {
		return fmt.Errorf("SendBtcxCrossBtc, ontcommon.AddressFromHexString error: %s", err)
	}
	to := hex.EncodeToString([]byte(btcAddress))
	txHash, err := ctx.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		signer,
		signer,
		btcxContractAddress,
		[]interface{}{"lock", []interface{}{config.ETH_CHAIN_ID, signer.Address[:], to, amount}})
	if err != nil {
		return fmt.Errorf("SendBtcxCrossBtc, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	ctx.LogInfo("SendBtcxCrossBtc, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendBtcCrossOnt(ctx *testframework.TestFrameworkContext, btcPriKey, ontAddress string, amount int64) error {
	err := sendBtcCross(ctx, config.ONT_CHAIN_ID, btcPriKey, ontAddress, amount)
	return fmt.Errorf("SendBtcCrossOnt, sendBtcCross error: %s", err)
}

func SendBtcCrossEth(ctx *testframework.TestFrameworkContext, btcPriKey, ontAddress string, amount int64) error {
	err := sendBtcCross(ctx, config.ETH_CHAIN_ID, btcPriKey, ontAddress, amount)
	return fmt.Errorf("SendBtcCrossEth, sendBtcCross error: %s", err)
}

func sendBtcCross(ctx *testframework.TestFrameworkContext, chainID uint64, btcPriKey, ontAddress string, amount int64) error {
	value := float64(amount) / btcutil.SatoshiPerBitcoin
	fee := float64(config.DefConfig.BtcFee) / btcutil.SatoshiPerBitcoin
	wif, err := btcutil.DecodeWIF(btcPriKey)
	if err != nil {
		return fmt.Errorf("sendBtcCross, failed to decode wif: %v", err)
	}

	addrPubk, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		return fmt.Errorf("sendBtcCross, Failed to new an address pubkey: %v", err)
	}
	pubkScript, err := txscript.PayToAddrScript(addrPubk.AddressPubKeyHash())
	if err != nil {
		return fmt.Errorf("sendBtcCross, Failed to build pubk script: %v", err)
	}
	data, err := utils.BuildData(chainID, 0, ontAddress)
	if err != nil {
		return fmt.Errorf("sendBtcCross, Failed to ge data: %v", err)
	}

	addr := addrPubk.EncodeAddress()
	//err = ra.Cli.ImportAddress(addr)
	//if err != nil {
	//	log.Errorf("rpc failed: %v", err)
	//	return nil
	//}
HERE:
	cnt, err := ctx.BtcCli.GetBlockCount()
	if err != nil {
		return fmt.Errorf("sendBtcCross, rpc failed: %v", err)
	}
	utxos, err := ctx.BtcCli.ListUnspent(1, cnt, addr)
	if err != nil {
		return fmt.Errorf("sendBtcCross, rpc failed: %v", err)
	}
	total, err := btcutil.NewAmount(value + fee)
	if err != nil {
		return fmt.Errorf("sendBtcCross, failed to new amount: %v", err)
	}
	selected, sumVal, err := utils.SelectUtxos(utxos, int64(total))
	if err != nil {
		return fmt.Errorf("sendBtcCross, failed to select utxo when build btc tx: %v", err)
	}

	//var prevPkScripts [][]byte
	var ipts []btcjson.TransactionInput
	for _, v := range selected {
		ipts = append(ipts, btcjson.TransactionInput{
			Txid: v.Txid,
			Vout: v.Vout,
		})
		//sb, err := hex.DecodeString(v.ScriptPubKey)
		//if err != nil {
		//	log.Errorf("failed to decode hex string pubk %s: %v", err)
		//	return nil
		//}
		//prevPkScripts = append(prevPkScripts, sb)
	}

	b, err := utils.NewBuilder(&utils.BuildCrossChainTxParam{
		Redeem:       config.DefConfig.BtcRedeem,
		Data:         data,
		Inputs:       ipts,
		NetParam:     &chaincfg.TestNet3Params,
		PrevPkScript: pubkScript,
		Privk:        wif.PrivKey,
		Locktime:     nil,
		ToMultiValue: value,
		Changes: func() map[string]float64 {
			if changeVal := float64(sumVal)/btcutil.SatoshiPerBitcoin - value - fee; changeVal > 0 {
				return map[string]float64{addrPubk.EncodeAddress(): changeVal}
			} else {
				return map[string]float64{}
			}
		}(),
	})
	if err != nil {
		return fmt.Errorf("sendBtcCross, Failed to new an instance of Builder: %v", err)
	}

	var buf bytes.Buffer
	//if ra.Privkb58 == "" {
	//	err = b.Tx.BtcEncode(&buf, wire.ProtocolVersion, wire.LatestEncoding)
	//	if err != nil {
	//		log.Errorf("Failed to encode transaction: %v", err)
	//		return nil
	//	}
	//	log.Infof("------------------------Your unsigned cross chain transaction------------------------\n%x\n", buf.Bytes())
	//	return
	//}
	err = b.BuildSignedTx()
	if err != nil || !b.IsSigned {
		return fmt.Errorf("sendBtcCross, Failed to build signed transaction: %v", err)
	}
	err = b.Tx.BtcEncode(&buf, wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		return fmt.Errorf("sendBtcCross, Failed to encode transaction: %v", err)
	}

	txid, err := ctx.BtcCli.SendRawTx(hex.EncodeToString(buf.Bytes()))
	if err != nil {
		if strings.Contains(err.Error(), "min relay fee not met") {
			arr := strings.Split(err.Error(), " ")
			newFee, _ := strconv.ParseUint(arr[len(arr)-1], 10, 64)
			fee = float64(newFee) / btcutil.SatoshiPerBitcoin
			goto HERE
		}
		return fmt.Errorf("sendBtcCross, failed to send tx: %v", err)
	}
	ctx.LogInfo("sendBtcCross, send tx %s(%f btc) to regression net(%s)", txid, value, ctx.BtcCli.Addr)
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
