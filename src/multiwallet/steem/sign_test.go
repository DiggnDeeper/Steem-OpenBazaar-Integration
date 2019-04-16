package steem

import (
	"bytes"
	"encoding/hex"
	"testing"
	"time"

	"github.com/OpenBazaar/multiwallet/cache"
	"github.com/OpenBazaar/multiwallet/datastore"
	"github.com/OpenBazaar/multiwallet/keys"
	laddr "github.com/OpenBazaar/multiwallet/litecoin/address"
	"github.com/OpenBazaar/multiwallet/model/mock"
	"github.com/OpenBazaar/multiwallet/service"
	"github.com/OpenBazaar/multiwallet/util"
	"github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

type FeeResponse struct {
	Priority int `json:"priority"`
	Normal   int `json:"normal"`
	Economic int `json:"economic"`
}

func newMockWallet() (*SteemWallet, error) {
}

func TestSteemcoinWallet_buildTx(t *testing.T) {
}

func TestSteemWallet_buildSpendAllTx(t *testing.T) {
}

func TestSteemWallet_GenerateMultisigScript(t *testing.T) {
}

func TestSteemWallet_newUnsignedTransaction(t *testing.T) {
}

func TestSteemWallet_CreateMultisigSignature(t *testing.T) {
}

func TestSteemWallet_Multisign(t *testing.T) {
}

func TestSteemWallet_bumpFee(t *testing.T) {
}

func TestSteemWallet_sweepAddress(t *testing.T) {
}

func TestSteemWallet_estimateSpendFee(t *testing.T) {
}
