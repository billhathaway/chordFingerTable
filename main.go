package main

// generates CHORD finger tables

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"code.google.com/p/intmath/intgr"
)

var m int
var peers []int

func findNextPeer(val int) int {
	for i := 0; i < len(peers); i++ {
		if peers[i] >= val {
			return peers[i]
		}
	}
	return peers[len(peers)-1]
}

func genFingerTable(nodeID int) {
	fmt.Printf("node %d\n", nodeID)
	for i := 0; i < m; i++ {
		n := (nodeID + intgr.Pow(2, i)) % (intgr.Pow(2, m))
		fmt.Printf("%d: [%d] %d\n", i, n, findNextPeer(n))
	}
	fmt.Println()
}

func main() {
	flag.IntVar(&m, "m", 8, "number of bits to use")
	peersString := flag.String("peers", "", "csv list of peer IDs ex: 1,19,33,41,46,99,120")
	var val int
	flag.Parse()

	var err error
	for _, p := range strings.Split(*peersString, ",") {
		if val, err = strconv.Atoi(p); err != nil {
			log.Fatalln("invalid peer value ", p)
		}
		peers = append(peers, val)
	}
	if len(peers) == 0 {
		log.Printf("ERROR: at least one value needs to be set for peers")
		flag.Usage()
		os.Exit(1)
	}
	sort.Ints(peers)
	for _, peer := range peers {
		genFingerTable(peer)
	}
}
