/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package bitcoincash

import (
	"github.com/blocktree/bitcoin-adapter/bitcoin"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

type WalletManager struct {
	*bitcoin.WalletManager
	Blockscanner openwallet.BlockScanner
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = bitcoin.NewWalletManager()
	wm.Config = bitcoin.NewConfig(Symbol, CurveType, Decimals)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.TxDecoder = NewTransactionDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.Blockscanner = NewBCHBlockScanner(&wm)

	return &wm
}

func (wm *WalletManager) ListAddresses() ([]string, error) {
	var (
		addresses = make([]string, 0)
	)

	request := []interface{}{
		"",
	}

	result, err := wm.WalletClient.Call("getaddressesbyaccount", request)
	if err != nil {
		return nil, err
	}

	array := result.Array()
	for _, a := range array {
		addresses = append(addresses, a.String())
	}

	return addresses, nil
}

//ListUnspent 获取未花记录
func (wm *WalletManager) ListUnspent(min uint64, addresses ...string) ([]*bitcoin.Unspent, error) {

	//:分页限制

	var (
		limit       = 100
		searchAddrs = make([]string, 0)
		max         = len(addresses)
		step        = max / limit
		utxo        = make([]*bitcoin.Unspent, 0)
		pice        []*bitcoin.Unspent
		err         error
	)

	for i := 0; i <= step; i++ {
		begin := i * limit
		end := (i + 1) * limit
		if end > max {
			end = max
		}

		searchAddrs = addresses[begin:end]

		if len(searchAddrs) == 0 {
			continue
		}

		pice, err = wm.getListUnspentByCore(min, searchAddrs...)
		if err != nil {
			return nil, err
		}
		utxo = append(utxo, pice...)
	}
	return utxo, nil
}

//getTransactionByCore 获取交易单
func (wm *WalletManager) getListUnspentByCore(min uint64, addresses ...string) ([]*bitcoin.Unspent, error) {

	var (
		utxos = make([]*bitcoin.Unspent, 0)
	)

	request := []interface{}{
		min,
		9999999,
	}

	if len(addresses) > 0 {
		request = append(request, addresses)
	} else {
		request = append(request, []string{})
	}

	//request = append(request, 3)

	result, err := wm.WalletClient.Call("listunspent", request)
	if err != nil {
		return nil, err
	}

	array := result.Array()
	for _, a := range array {
		utxos = append(utxos, bitcoin.NewUnspent(&a))
	}

	return utxos, nil
}
