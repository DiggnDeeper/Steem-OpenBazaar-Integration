package litecoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/chaincfg"

	"github.com/OpenBazaar/spvwallet"
	wi "github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	btc "github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/coinset"
	hd "github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcutil/txsort"
	"github.com/btcsuite/btcwallet/wallet/txauthor"
	"github.com/ltcsuite/ltcutil"
	"github.com/ltcsuite/ltcwallet/wallet/txrules"

	laddr "github.com/OpenBazaar/multiwallet/litecoin/address"
	"github.com/OpenBazaar/multiwallet/util"
)

func (w *SteemWallet) buildTx(amount int64, addr btc.Address, feeLevel wi.FeeLevel, optionalOutput *wire.TxOut) (*wire.MsgTx, error) {
}

func (w *SteemWallet) buildSpendAllTx(addr btc.Address, feeLevel wi.FeeLevel) (*wire.MsgTx, error) {
}

func newUnsignedTransaction(outputs []*wire.TxOut, feePerKb btc.Amount, fetchInputs txauthor.InputSource, fetchChange txauthor.ChangeSource) (*txauthor.AuthoredTx, error) {
}

func (w *SteemWallet) bumpFee(txid chainhash.Hash) (*chainhash.Hash, error) {
}

func (w *SteemWallet) sweepAddress(ins []wi.TransactionInput, address *btc.Address, key *hd.ExtendedKey, redeemScript *[]byte, feeLevel wi.FeeLevel) (*chainhash.Hash, error) {
}

func (w *SteemWallet) createMultisigSignature(ins []wi.TransactionInput, outs []wi.TransactionOutput, key *hd.ExtendedKey, redeemScript []byte, feePerByte uint64) ([]wi.Signature, error) {
}

func (w *SteemWallet) multisign(ins []wi.TransactionInput, outs []wi.TransactionOutput, sigs1 []wi.Signature, sigs2 []wi.Signature, redeemScript []byte, feePerByte uint64, broadcast bool) ([]byte, error) {
}

func (w *SteemWallet) generateMultisigScript(keys []hd.ExtendedKey, threshold int, timeout time.Duration, timeoutKey *hd.ExtendedKey) (addr btc.Address, redeemScript []byte, err error) {
}

func (w *SteemWallet) estimateSpendFee(amount int64, feeLevel wi.FeeLevel) (uint64, error) {
}
