package legacytx

import (
	"encoding/json"
	"fmt"

	"github.com/onomyprotocol/onomy-sdk/codec"
	"github.com/onomyprotocol/onomy-sdk/codec/legacy"
	codectypes "github.com/onomyprotocol/onomy-sdk/codec/types"
	cryptotypes "github.com/onomyprotocol/onomy-sdk/crypto/types"
	"github.com/onomyprotocol/onomy-sdk/crypto/types/multisig"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	sdkerrors "github.com/onomyprotocol/onomy-sdk/types/errors"
	"github.com/onomyprotocol/onomy-sdk/types/tx/signing"
	"gopkg.in/yaml.v2"
)

// StdSignDoc is replay-prevention structure.
// It includes the result of msg.GetSignBytes(),
// as well as the ChainID (prevent cross chain replay)
// and the Sequence numbers for each signature (prevent
// inchain replay and enforce tx ordering per account).
type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number" yaml:"account_number"`
	Sequence      uint64            `json:"sequence" yaml:"sequence"`
	TimeoutHeight uint64            `json:"timeout_height,omitempty" yaml:"timeout_height"`
	ChainID       string            `json:"chain_id" yaml:"chain_id"`
	Memo          string            `json:"memo" yaml:"memo"`
	Fee           json.RawMessage   `json:"fee" yaml:"fee"`
	Msgs          []json.RawMessage `json:"msgs" yaml:"msgs"`
}

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytes(chainID string, accnum, sequence, timeout uint64, fee StdFee, msgs []sdk.Msg, memo string) []byte {
	msgsBytes := make([]json.RawMessage, 0, len(msgs))
	for _, msg := range msgs {
		// If msg is a legacy Msg, then GetSignBytes is implemented.
		// If msg is a ServiceMsg, then GetSignBytes has graceful support of
		// calling GetSignBytes from its underlying Msg.
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}

	bz, err := legacy.Cdc.MarshalJSON(StdSignDoc{
		AccountNumber: accnum,
		ChainID:       chainID,
		Fee:           json.RawMessage(fee.Bytes()),
		Memo:          memo,
		Msgs:          msgsBytes,
		Sequence:      sequence,
		TimeoutHeight: timeout,
	})
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(bz)
}

// Deprecated: StdSignature represents a sig
type StdSignature struct {
	cryptotypes.PubKey `json:"pub_key" yaml:"pub_key"` // optional
	Signature          []byte                          `json:"signature" yaml:"signature"`
}

// Deprecated
func NewStdSignature(pk cryptotypes.PubKey, sig []byte) StdSignature {
	return StdSignature{PubKey: pk, Signature: sig}
}

// GetSignature returns the raw signature bytes.
func (ss StdSignature) GetSignature() []byte {
	return ss.Signature
}

// GetPubKey returns the public key of a signature as a cryptotypes.PubKey using the
// Amino codec.
func (ss StdSignature) GetPubKey() cryptotypes.PubKey {
	return ss.PubKey
}

// MarshalYAML returns the YAML representation of the signature.
func (ss StdSignature) MarshalYAML() (interface{}, error) {
	pk := ""
	if ss.PubKey != nil {
		pk = ss.PubKey.String()
	}

	bz, err := yaml.Marshal(struct {
		PubKey    string
		Signature string
	}{
		pk,
		fmt.Sprintf("%X", ss.Signature),
	})
	if err != nil {
		return nil, err
	}

	return string(bz), err
}

func (ss StdSignature) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return codectypes.UnpackInterfaces(ss.PubKey, unpacker)
}

// StdSignatureToSignatureV2 converts a StdSignature to a SignatureV2
func StdSignatureToSignatureV2(cdc *codec.LegacyAmino, sig StdSignature) (signing.SignatureV2, error) {
	pk := sig.GetPubKey()
	data, err := pubKeySigToSigData(cdc, pk, sig.Signature)
	if err != nil {
		return signing.SignatureV2{}, err
	}

	return signing.SignatureV2{
		PubKey: pk,
		Data:   data,
	}, nil
}

func pubKeySigToSigData(cdc *codec.LegacyAmino, key cryptotypes.PubKey, sig []byte) (signing.SignatureData, error) {
	multiPK, ok := key.(multisig.PubKey)
	if !ok {
		return &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
			Signature: sig,
		}, nil
	}
	var multiSig multisig.AminoMultisignature
	err := cdc.UnmarshalBinaryBare(sig, &multiSig)
	if err != nil {
		return nil, err
	}

	sigs := multiSig.Sigs
	sigDatas := make([]signing.SignatureData, len(sigs))
	pubKeys := multiPK.GetPubKeys()
	bitArray := multiSig.BitArray
	n := multiSig.BitArray.Count()
	signatures := multisig.NewMultisig(n)
	sigIdx := 0
	for i := 0; i < n; i++ {
		if bitArray.GetIndex(i) {
			data, err := pubKeySigToSigData(cdc, pubKeys[i], multiSig.Sigs[sigIdx])
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "Unable to convert Signature to SigData %d", sigIdx)
			}

			sigDatas[sigIdx] = data
			multisig.AddSignature(signatures, data, sigIdx)
			sigIdx++
		}
	}

	return signatures, nil
}
