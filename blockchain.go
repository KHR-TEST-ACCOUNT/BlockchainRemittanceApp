package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transaction  []string
}

type Blockchaine struct {
	transctionPool []string
	chain          []*Block
}

// BlockChaineを生成して返す。
func NewBlockChain() *Blockchaine {
	// create pointer
	b := &Block{}
	bc := new(Blockchaine)
	bc.CreateBlock(0, b.Hash())
	return bc
}

/*
BlockをBlockchainの配列に格納し、
return: Blockポインタ
*/
func (bc *Blockchaine) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

/*
最後のブロックを取り出す。
*/
func (bc *Blockchaine) LastBlcok() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Blockを生成して返す。
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Blockの情報を出力
func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previousHash    %x\n", b.previousHash)
	fmt.Printf("transaction     %s\n", b.transaction)
}

/*
JsonマーシャルにBlcokを格納。
*/
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

/*
Json マーシャルへの格納をオーバーライドして値が取得できるように変更
*/
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json: "nonce"`
		PreviousHash [32]byte `json: "previous_hash"`
		Timestamp    int64    `json: "timestamp"`
		Transaction  []string `json: "transaction"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transaction:  b.transaction,
	})
}

/*
ブロックチェーンの配列にあるブロックの一覧を表示する。
*/
func (bc *Blockchaine) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain  %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchane: ")
}

func main() {
	//1 - block
	blockchaine := NewBlockChain()
	blockchaine.Print()
	// 2- block
	previousHash := blockchaine.LastBlcok().Hash()
	blockchaine.CreateBlock(1, previousHash)
	blockchaine.Print()
	// 3 - block
	previousHash = blockchaine.LastBlcok().Hash()
	blockchaine.CreateBlock(2, previousHash)
	blockchaine.Print()
}
