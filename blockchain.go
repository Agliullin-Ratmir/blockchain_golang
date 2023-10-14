package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func isValidToSmartContract(amount float64) bool {
	if amount > 2 {
		return true
	}
	log.Println("Can't add amount to node, because amount: ", amount,
		" is not valid!")
	return false
}

func GetAmountOfNodes() int {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Please, type amount of nodes:")
	amount, _ := reader.ReadString('\n')
	amountStr := strings.TrimSuffix(amount, "\r\n")
	amountStr = strings.TrimSuffix(amountStr, "\n")
	amountInt, err := strconv.Atoi(amountStr)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return amountInt
}

func CreateNodes(amount int) map[string]Blockchain {
	result := map[string]Blockchain{}
	result = make(map[string]Blockchain)

	for i := 1; i <= amount; i++ {
		key := strconv.FormatInt(int64(i), 10)
		newBlock := CreateBlockchain(2)
		newBlock.addBlock("One", "Two", 10)
		newBlock.addBlock("Two", "Three", 10)
		result[key] = newBlock
		log.Println("Created node: ", key)
	}

	return result
}

func ShowAllBlocks(mapNodes map[string]Blockchain) {
	for k, v := range mapNodes {
		lengthOfBlock := len(v.chain)
		log.Println("For key: ", k, " value is: ", v.chain[lengthOfBlock-1].hash,
			"amount of blocks is: ", lengthOfBlock)
	}
}

func Broadcast(mapNodes map[string]Blockchain, from, to string, amount float64) {
	if isValidToSmartContract(amount) {
		for _, v := range mapNodes {
			v.addBlock(from, to, amount)
		}
		log.Println("New block has been added to all nodes")
	} else {
		log.Println("The amount: ", strconv.FormatFloat(amount, 'f', 3, 64),
			" is not valid")
	}
}

func ShowBlockBYIndex(index string, mapNodes map[string]Blockchain) {
	v := mapNodes[index]
	lengthOfBlock := len(v.chain)
	log.Println("For node: ", index, " value is: ", v.chain[lengthOfBlock-1].hash,
		" amount of blocks: ", lengthOfBlock)
}
