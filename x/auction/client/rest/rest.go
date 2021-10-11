package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Rest variable names
// nolint
const (
	RestParamAuctionId = "auction-id"
	RestParamOwner     = "owner"
	RestParamBidder    = "bidder"
)

// RegisterHandlers registers token-related REST handlers to a router
func RegisterHandlers(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

type openAuctionReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	AuctionType int32        `json:"auction_type"`
	NftDenomId  string       `json:"nft_denom_id"`
	NftTokenId  string       `json:"nft_token_id"`
	Duration    uint64       `json:"duration"`
	MinAmount   string       `json:"min_amount"`
	Owner       string       `json:"owner"`
	Limit       uint32       `json:"limit"`
}

type editAuctionReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Duration uint64       `json:"duration"`
	Owner    string       `json:"owner"`
}

type cancelAuctionReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"owner"`
}

type openBidReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Bidder    string       `json:"bidder"`
	BidAmount string       `json:"bid_amount"`
}

type cancelBidReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Bidder  string       `json:"bidder"`
}

type withdrawBidReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Recipient string       `json:"recipient"`
}