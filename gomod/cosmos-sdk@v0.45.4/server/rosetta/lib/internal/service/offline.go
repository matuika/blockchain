package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"
)




func NewOffline(network *types.NetworkIdentifier, client crgtypes.Client) (crgtypes.API, error) {
	return OfflineNetwork{
		OnlineNetwork{
			client:         client,
			network:        network,
			networkOptions: networkOptionsFromClient(client, nil),
		},
	}, nil
}




type OfflineNetwork struct {
	OnlineNetwork
}


func (o OfflineNetwork) AccountBalance(_ context.Context, _ *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) Block(_ context.Context, _ *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) BlockTransaction(_ context.Context, _ *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) Mempool(_ context.Context, _ *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) MempoolTransaction(_ context.Context, _ *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) NetworkStatus(_ context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) ConstructionSubmit(_ context.Context, _ *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

func (o OfflineNetwork) ConstructionMetadata(_ context.Context, _ *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}