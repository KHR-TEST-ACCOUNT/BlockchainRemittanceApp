package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
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
	bc := new(Blockchaine)
	bc.CreateBlock(0, "Init Hash!")
	return bc
}

/*
BlockをBlockchainの配列に格納し、
return: Blockポインタ
*/
func (bc *Blockchaine) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// Blockを生成して返す。
func NewBlock(nonce int, previousHash string) *Block {
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
	fmt.Printf("previousHash    %s\n", b.previousHash)
	fmt.Printf("transaction     %s\n", b.transaction)
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
	blockchaine := NewBlockChain()
	blockchaine.Print()
	blockchaine.CreateBlock(1, "hash 1")
	blockchaine.Print()
	blockchaine.CreateBlock(2, "hash 2")
	blockchaine.Print()
}
