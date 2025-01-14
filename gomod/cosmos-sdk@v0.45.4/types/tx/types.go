package tx

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


const MaxGasWanted = uint64((1 << 63) - 1)


var _, _, _, _ codectypes.UnpackInterfacesMessage = &Tx{}, &TxBody{}, &AuthInfo{}, &SignerInfo{}
var _ sdk.Tx = &Tx{}


func (t *Tx) GetMsgs() []sdk.Msg {
	if t == nil || t.Body == nil {
		return nil
	}

	anys := t.Body.Messages
	res := make([]sdk.Msg, len(anys))
	for i, any := range anys {
		cached := any.GetCachedValue()
		if cached == nil {
			panic("Any cached value is nil. Transaction messages must be correctly packed Any values.")
		}
		res[i] = cached.(sdk.Msg)
	}
	return res
}


func (t *Tx) ValidateBasic() error {
	if t == nil {
		return fmt.Errorf("bad Tx")
	}

	body := t.Body
	if body == nil {
		return fmt.Errorf("missing TxBody")
	}

	authInfo := t.AuthInfo
	if authInfo == nil {
		return fmt.Errorf("missing AuthInfo")
	}

	fee := authInfo.Fee
	if fee == nil {
		return fmt.Errorf("missing fee")
	}

	if fee.GasLimit > MaxGasWanted {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid gas supplied; %d > %d", fee.GasLimit, MaxGasWanted,
		)
	}

	if fee.Amount.IsAnyNil() {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: null",
		)
	}

	if fee.Amount.IsAnyNegative() {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: %s", fee.Amount,
		)
	}

	if fee.Payer != "" {
		_, err := sdk.AccAddressFromBech32(fee.Payer)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid fee payer address (%s)", err)
		}
	}

	sigs := t.Signatures

	if len(sigs) == 0 {
		return sdkerrors.ErrNoSignatures
	}

	if len(sigs) != len(t.GetSigners()) {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			"wrong number of signers; expected %d, got %d", len(t.GetSigners()), len(sigs),
		)
	}

	return nil
}




func (t *Tx) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress
	seen := map[string]bool{}

	for _, msg := range t.GetMsgs() {
		for _, addr := range msg.GetSigners() {
			if !seen[addr.String()] {
				signers = append(signers, addr)
				seen[addr.String()] = true
			}
		}
	}

	
	feePayer := t.AuthInfo.Fee.Payer
	if feePayer != "" && !seen[feePayer] {
		payerAddr, err := sdk.AccAddressFromBech32(feePayer)
		if err != nil {
			panic(err)
		}
		signers = append(signers, payerAddr)
		seen[feePayer] = true
	}

	return signers
}

func (t *Tx) GetGas() uint64 {
	return t.AuthInfo.Fee.GasLimit
}
func (t *Tx) GetFee() sdk.Coins {
	return t.AuthInfo.Fee.Amount
}
func (t *Tx) FeePayer() sdk.AccAddress {
	feePayer := t.AuthInfo.Fee.Payer
	if feePayer != "" {
		payerAddr, err := sdk.AccAddressFromBech32(feePayer)
		if err != nil {
			panic(err)
		}
		return payerAddr
	}
	
	return t.GetSigners()[0]
}

func (t *Tx) FeeGranter() sdk.AccAddress {
	feePayer := t.AuthInfo.Fee.Granter
	if feePayer != "" {
		granterAddr, err := sdk.AccAddressFromBech32(feePayer)
		if err != nil {
			panic(err)
		}
		return granterAddr
	}
	return nil
}


func (t *Tx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if t.Body != nil {
		if err := t.Body.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}

	if t.AuthInfo != nil {
		return t.AuthInfo.UnpackInterfaces(unpacker)
	}

	return nil
}


func (m *TxBody) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, any := range m.Messages {
		var msg sdk.Msg
		err := unpacker.UnpackAny(any, &msg)
		if err != nil {
			return err
		}
	}

	return nil
}


func (m *AuthInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, signerInfo := range m.SignerInfos {
		err := signerInfo.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}


func (m *SignerInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return unpacker.UnpackAny(m.PublicKey, new(cryptotypes.PubKey))
}


func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.tx.v1beta1.Tx", (*sdk.Tx)(nil))
	registry.RegisterImplementations((*sdk.Tx)(nil), &Tx{})
}
