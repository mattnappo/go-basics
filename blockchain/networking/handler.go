package networking

import (
	"io"
	"net/http"

	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

// HandleGetBlockchain - Handle a get request (to return the blockchain)
func HandleGetBlockchain(chain *blockchain.Blockchain, w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, string(chain.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
