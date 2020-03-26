package testframework

import (
	"encoding/hex"
	"math/big"
	"os"
	"time"

	"github.com/ontio/cross_chain_test/config"
)

func MonitorOnt(ctx *TestFrameworkContext) {
	currentHeight, err := ctx.OntSdk.GetCurrentBlockHeight()
	if err != nil {
		ctx.LogError("ctx.OntSdk.GetCurrentBlockHeight error: %s", err)
		os.Exit(1)
	}
	ontHeight := currentHeight

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := ctx.OntSdk.GetCurrentBlockHeight()
			if err != nil {
				ctx.LogError("self.ontsdk.GetCurrentBlockHeight error: %s", err)
				continue
			}

			ctx.LogInfo("ont chain current height: %d", currentHeight)

			if currentHeight <= ontHeight {
				continue
			}

			for currentHeight > ontHeight {
				ontHeight++
				err = parseOntologyChainBlock(ctx, ontHeight)
				if err != nil {
					ctx.LogError("parseOntologyChainBlock error: %s", err)
					ontHeight--
					break
				}
			}
		}
	}
}

func parseOntologyChainBlock(ctx *TestFrameworkContext, height uint32) error {
	events, err := ctx.OntSdk.GetSmartContractEventByBlock(height)
	if err != nil {
		return err
	}

	ctx.LogInfo("parseOntologyChainBlock, ont chain block height: %d, events num: %d", height, len(events))
	for _, event := range events {
		for _, notify := range event.Notify {
			if notify.ContractAddress != "0900000000000000000000000000000000000000" {
				continue
			}

			states := notify.States.([]interface{})
			contractMethod, _ := states[0].(string)
			if contractMethod != "verifyToOntProof" {
				continue
			}
			// try to get all data
			//
			if len(states) < 5 {
				continue
			}
			allianceTxHash := states[1].(string)
			rawTxHash := states[2].(string)
			ctx.LogInfo("receive cross chain tx on ontology, tx hash: %s, alliance tx hash: %s, raw tx hash: %s", event.TxHash, allianceTxHash, rawTxHash)
			ctx.Lock.Lock()
			delete(ctx.TxMap, rawTxHash)
			ctx.Lock.Unlock()
		}
	}
	return nil
}

func MonitorEthChain(ctx *TestFrameworkContext) {
	currentHeight, err := ctx.EthTools.GetNodeHeight()
	if err != nil {
		ctx.LogError("ctx.EthTools.GetNodeHeight error: %s", err)
		os.Exit(1)
	}
	ethHeight := uint32(currentHeight)
	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := ctx.EthTools.GetNodeHeight()
			if err != nil {
				ctx.LogError("ctx.EthTools.GetNodeHeight error: %s", err)
				continue
			}

			ctx.LogInfo("eth chain current height: %d", currentHeight)

			if uint32(currentHeight) <= ethHeight {
				continue
			}

			for uint32(currentHeight) > ethHeight {
				ethHeight++
				err = parseEthChainBlock(ctx, ethHeight)
				if err != nil {
					ctx.LogError("parseEthChainBlock error: %s", err)
					ethHeight--
					break
				}
			}
		}
	}
}

func parseEthChainBlock(ctx *TestFrameworkContext, height uint32) error {
	// contract is different
	lockevents, unlockevents, err := ctx.EthTools.GetSmartContractEventByBlock(config.DefConfig.EthCrossChainContractAddress, uint64(height))
	if err != nil {
		return err
	}
	ctx.LogInfo("eth chain, block height: %d, unlock events num: %d, lock events num: %d", height, len(unlockevents), len(lockevents))
	for _, event := range lockevents {
		// try to get all data
		//
		ethTxHash := event.TxHash[2:]
		var ethTxId_byte []byte
		index_int := big.NewInt(0)
		index_int.SetBytes(event.Txid)
		for i := len(index_int.Bytes()); i < 32; i++ {
			ethTxId_byte = append(ethTxId_byte, 0)
		}
		ethTxId_byte = append(ethTxId_byte, index_int.Bytes()...)
		ethTxId_string := hex.EncodeToString(ethTxId_byte)
		ctx.LogInfo("send cross chain tx on eth, tx hash: %s, tx id: %s", ethTxHash, ethTxId_string)
		ctx.Lock.Lock()
		ctx.TxMap[ethTxId_string] = ethTxHash
		delete(ctx.TxMap, ethTxHash)
		ctx.Lock.Unlock()
	}

	for _, event := range unlockevents {
		// try to get all data
		//
		allianceTxHash := event.RTxid
		rawTxHash := event.FromTxId

		ctx.LogInfo("receive cross chain tx on eth, txhash: %s, alliance tx hash: %s, raw tx hash: %s", event.Txid, allianceTxHash, rawTxHash)
		ctx.Lock.Lock()
		delete(ctx.TxMap, rawTxHash)
		ctx.Lock.Unlock()
	}
	return nil
}

func MonitorRChain(ctx *TestFrameworkContext) {
	currentHeight, err := ctx.RcSdk.GetCurrentBlockHeight()
	if err != nil {
		ctx.LogError("ctx.RcSdk.GetCurrentBlockHeight error: %s", err)
		os.Exit(1)
	}
	rcHeight := currentHeight

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := ctx.RcSdk.GetCurrentBlockHeight()
			if err != nil {
				ctx.LogError("self.RcSdk.GetCurrentBlockHeight error: %s", err)
				continue
			}

			ctx.LogInfo("rc chain current height: %d", currentHeight)

			if currentHeight <= rcHeight {
				continue
			}

			for currentHeight > rcHeight {
				rcHeight++
				err = parseRelayChainBlock(ctx, rcHeight)
				if err != nil {
					ctx.LogError("parseRelayChainBlock error: %s", err)
					rcHeight--
					break
				}
			}
		}
	}
}

func parseRelayChainBlock(ctx *TestFrameworkContext, height uint32) error {
	events, err := ctx.RcSdk.GetSmartContractEventByBlock(height)
	if err != nil {
		return err
	}

	ctx.LogInfo("parseRelayChainBlock, relay chain block height: %d, events num: %d", height, len(events))
	for _, event := range events {
		for _, notify := range event.Notify {
			states, ok := notify.States.([]interface{})
			if !ok {
				continue
			}
			name, ok := states[0].(string)
			if ok && name == "btcTxToRelay" {
				txHash, _ := states[4].(string)
				ctx.LogInfo("receive cross chain tx on relay chain, tx hash: %s, raw tx hash: %s", event.TxHash, txHash)
				ctx.Lock.Lock()
				delete(ctx.TxMap, txHash)
				ctx.Lock.Unlock()
			}
		}
	}
	return nil
}

func ReportPending(ctx *TestFrameworkContext) {
	reportTicker := time.NewTicker(time.Second * time.Duration(config.DefConfig.ReportInterval))
	for {
		select {
		case <-reportTicker.C:
			for txHash, info := range ctx.TxMap {
				ctx.LogInfo("ReportPending, txhash is %s, info is %s", txHash, info)
			}
		}
	}
}
