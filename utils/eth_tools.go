package utils

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ontio/ont-eth-ft-tool/contract_cross_chain"
	"github.com/ontio/ont-eth-ft-tool/log"
	"github.com/ontio/ont-eth-ft-tool/restclient"
	"math/big"
	"strconv"
	"time"
)

type ETHTools struct {
	restclient *restclient.RestClient
	ethclient  *ethclient.Client
}

type LockEvent struct {
	Method   string
	TxHash   string
	Txid     []byte
	Saddress string
	Tchain   uint32
	Taddress string
	Height   uint64
	Value    []byte
	Token    string
}
type UnlockEvent struct {
	Method   string
	Txid     string
	RTxid    string
	FromTxId string
	Height   uint64
	Token    string
}

type heightReq struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      uint     `json:"id"`
}

type heightRep struct {
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      uint   `json:"id"`
}

type BlockReq struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      uint          `json:"id"`
}

type BlockRep struct {
	JsonRPC string        `json:"jsonrpc"`
	Result  *types.Header `json:"result"`
	Id      uint          `json:"id"`
}

func NewEthTools(url string) *ETHTools {
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		log.Error("NewEthTools: cannot dial sync node, err: %s", err)
		return nil
	}
	restclient := restclient.NewRestClient()
	restclient.SetAddr(url)
	tool := &ETHTools{
		restclient: restclient,
		ethclient:  ethclient,
	}
	return tool
}

func (self *ETHTools) GetEthClient() *ethclient.Client {
	return self.ethclient
}

func (self *ETHTools) GetNodeHeight() (uint64, error) {
	req := &heightReq{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  make([]string, 0),
		Id:      1,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	resp, err := self.restclient.SendRestRequest(data)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rep := &heightRep{}
	err = json.Unmarshal(resp, rep)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	height, err := strconv.ParseUint(rep.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, parse resp height %s failed", rep.Result)
	} else {
		return height, nil
	}
}

func (self *ETHTools) GetBlockHeader(height uint64) (*types.Header, error) {
	params := []interface{}{fmt.Sprintf("0x%x", height), true}
	req := &BlockReq{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  params,
		Id:      1,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	resp, err := self.restclient.SendRestRequest(data)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &BlockRep{}
	err = json.Unmarshal(resp, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}

	return rsp.Result, nil
}

func (self *ETHTools) GetSmartContractEventByBlock(contractAddr string, height uint64) ([]*LockEvent, []*UnlockEvent, error) {
	lockAddress := common.HexToAddress(contractAddr)
	instance, err := cross_chain_abi.NewEthCrossChain(lockAddress, self.ethclient)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}

	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	ethlockevents := make([]*LockEvent, 0)
	{
		events, err := instance.FilterCrossChainEvent(opt, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethlockevents = append(ethlockevents, &LockEvent{
				Method:   "lock",
				TxHash:   evt.Raw.TxHash.String(),
				Txid:     evt.TxId,
				Saddress: evt.Sender.String(),
				Tchain:   uint32(evt.ToChainId),
				Value:    evt.Rawdata,
				Height:   height,
				Token:    evt.Token.String(),
			})
		}
	}

	ethunlockevents := make([]*UnlockEvent, 0)
	{
		events, err := instance.FilterVerifyAndExecuteTxEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethunlockevents = append(ethunlockevents, &UnlockEvent{
				Method:   "unlock",
				Txid:     evt.Raw.TxHash.String(),
				RTxid:    hex.EncodeToString(evt.CrossChainTxHash),
				FromTxId: hex.EncodeToString(evt.FromChainTxHash),
				Token:    hex.EncodeToString(evt.ToContract),
				Height:   height,
			})
		}
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethunlockevents = append(ethunlockevents, &UnlockEvent{
				Method: "unlock",
				Txid:   evt.Raw.TxHash.String(),
				RTxid:  hex.EncodeToString(evt.CrossChainTxHash),
				Token:  hex.EncodeToString(evt.ToContract),
				Height: height,
			})
		}
	}
	return ethlockevents, ethunlockevents, nil
}

func EncodeBigInt(b *big.Int) string {
	if b.Uint64() == 0 {
		return "00"
	}
	return hex.EncodeToString(b.Bytes())
}

type getMoneyReq struct {
	Address string `json:"address"`
	Amount  uint   `json:"amount"`
}

type getMoneyRsp struct {
	Hash   string `json:"tx_hash"`
	Result string `json:"result"`
}

func (self *ETHTools) GetMoney(accounts []accounts.Account) []common.Hash {
	restclient := restclient.NewRestClient()
	restclient.SetAddr("https://api.bitaps.com/eth/testnet/v1/faucet/send/payment")
	hash := make([]common.Hash, 0)
	for _, account := range accounts {
		req := &getMoneyReq{
			Address: account.Address.String(),
			Amount:  1000000000000000000,
		}
		data, err := json.Marshal(req)
		if err != nil {
			log.Errorf("getMoney: marshal req err: %s", err)
			continue
		}
		resp, err := restclient.SendRestRequest([]byte(data))
		if err != nil {
			log.Errorf("getMoney err: %s", err)
			continue
		}
		rep := &getMoneyRsp{}
		err = json.Unmarshal(resp, rep)
		if err != nil {
			log.Errorf("getMoney, unmarshal resp err: %s", err)
			continue
		}
		hashBytes := common.HexToHash(rep.Hash)
		hash = append(hash, hashBytes)
	}
	return hash
}

func (self *ETHTools) WaitTransactionsConfirm(hashs []common.Hash) {
	hasPending := true
	for hasPending {
		time.Sleep(time.Second * 1)
		hasPending = false
		for _, hash := range hashs {
			_, ispending, err := self.ethclient.TransactionByHash(context.Background(), hash)
			log.Infof("transaction %s is pending: %d", hash.String(), ispending)
			if err != nil {
				hasPending = true
				continue
			}
			if ispending == true {
				hasPending = true
			} else {
			}
		}
	}
}

func (self *ETHTools) WaitTransactionConfirm(hash common.Hash) {
	errNum := 0
	for errNum < 100 {
		time.Sleep(time.Second * 1)
		_, ispending, err := self.ethclient.TransactionByHash(context.Background(), hash)
		log.Infof("transaction %s is pending: %d", hash.String(), ispending)
		if err != nil {
			errNum++
			continue
		}
		if ispending == true {
			continue
		} else {
			break
		}
	}
}
