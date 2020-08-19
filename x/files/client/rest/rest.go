package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers files-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/files/claim", listClaimHandler(cliCtx, "files")).Methods("GET")
	r.HandleFunc("/files/claim", createClaimHandler(cliCtx)).Methods("POST")
}
