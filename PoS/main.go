package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/gp-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
	Validator string
}

var Blockchain []Block
var tempBlocks []Block

var candidateBlocks = make(chan Block)

var announcements = make(chan string)

var mutex = &sync.Mutex{}4

var validators = make(map[string]int)

func calculateHash (s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func calculateBlockHash(block Block) string{
	s := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash + block.Validator
	return calculateHash(s)
}

func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index+1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}

func isBlockValid(newblock, oldBlock Block) bool {
	if oldBlock.Index + 1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if newBlock.Hash != calculateBlockHash(newBlock){
		return false
	}
	return true
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
			for {
				mas := <-announcements
				io.WriteString(conn, msg)
			}
	}()

	var address string

	io.WriteString(conn, "Enter the balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan(){
		balance, err = strconv.Atoi(scanBalance.Text())
		
	}
}