package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())

	amount := GetAmountOfNodes()
	nodes := CreateNodes(amount)
	reader := bufio.NewReader(os.Stdin)
	log.Println("Please, type a command:")
	commandStr := readString(*reader)
	for commandStr != "exit" {
		getCommand(nodes, commandStr, *reader)
		log.Println("Please, type a command:")
		commandStr = readString(*reader)
	}

	// 1) add console command to add amount of nodes
	// + 1.1 method for taking amount of nodes
	// + 1.2 method for initialize a map of nodes and logging all messages about nodes (starts from 1, it is important)
	// + check 1.1, 1.2
	// +2) add command to show all nodes and node by index
	// +2.1 put number of node and show all records
	// + check 2 and 2.1
	// + 2.2 or you can type add (zero) to add a new record to block
	// + 3) add simple smart contract (amount > 2) isValidToSmartContract
	// + 3.1 add method "broadcast" - add new record to all blocks.
	// + 3.2 add method about smart-contract
	// + add infinite cycle, stop word - exit
}

func getCommand(nodes map[string]Blockchain, commandStr string, reader bufio.Reader) {
	if commandStr == "all" {
		ShowAllBlocks(nodes)
	} else if commandStr == "get" {
		index := readString(reader)
		ShowBlockBYIndex(index, nodes)
	} else if commandStr == "add" {
		log.Println("Please, type from:")
		from := readString(reader)

		log.Println("Please, type to:")
		to := readString(reader)

		log.Println("Please, type amount:")
		amount := readString(reader)
		amountFloat, _ := strconv.ParseFloat(amount, 64)
		Broadcast(nodes, from, to, amountFloat)
		ShowAllBlocks(nodes)
	}
}

func readString(reader bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	str = strings.TrimSuffix(str, "\n")
	return str
}
