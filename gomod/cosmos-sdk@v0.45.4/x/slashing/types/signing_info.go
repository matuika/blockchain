package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



func NewValidatorSigningInfo(
	condAddr sdk.ConsAddress, startHeight, indexOffset int64,
	jailedUntil time.Time, tombstoned bool, missedBlocksCounter int64,
) ValidatorSigningInfo {

	return ValidatorSigningInfo{
		Address:             condAddr.String(),
		StartHeight:         startHeight,
		IndexOffset:         indexOffset,
		JailedUntil:         jailedUntil,
		Tombstoned:          tombstoned,
		MissedBlocksCounter: missedBlocksCounter,
	}
}


func (i ValidatorSigningInfo) String() string {
	return fmt.Sprintf(`Validator Signing Info:
  Address:               %s
  Start Height:          %d
  Index Offset:          %d
  Jailed Until:          %v
  Tombstoned:            %t
  Missed Blocks Counter: %d`,
		i.Address, i.StartHeight, i.IndexOffset, i.JailedUntil,
		i.Tombstoned, i.MissedBlocksCounter)
}


func UnmarshalValSigningInfo(cdc codec.Codec, value []byte) (signingInfo ValidatorSigningInfo, err error) {
	err = cdc.Unmarshal(value, &signingInfo)
	return signingInfo, err
}
