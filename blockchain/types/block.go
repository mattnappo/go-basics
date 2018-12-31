package types

// Block - A block in the chain
type Block struct {
	Index        int            `json:"Index"`
	Transactions []*Transaction `json:"Transactions"`
	PrevHash     []byte         `json:"PrevHash"`
	Timestamp    string         `json:"Timestamp"`
	Hash         []byte         `json:"Hash"`
}

// NewBlock - Create a new block
func NewBlock(index int, transactions, prevHash string) {

}
