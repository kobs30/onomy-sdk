package client

import (
	"github.com/onomyprotocol/onomy-sdk/x/distribution/client/cli"
	"github.com/onomyprotocol/onomy-sdk/x/distribution/client/rest"
	govclient "github.com/onomyprotocol/onomy-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
