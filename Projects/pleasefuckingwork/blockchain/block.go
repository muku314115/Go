package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(prevHash []byte, data string) *Block {
	hash := sha256.Sum256(bytes.Join([][]byte{[]byte(data), prevHash}, []byte{}))
	return &Block{Hash: hash[:], Data: []byte(data), PrevHash: prevHash}
}

func Genesis() *Block {
	return CreateBlock([]byte{}, "Genesis")
}

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) *BlockChain {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(prevBlock.Hash, data)
	chain.Blocks = append(chain.Blocks, newBlock)

	return chain
}
