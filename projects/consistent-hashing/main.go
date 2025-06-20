package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"sync/atomic"
)

const (
	MACHINES_TO_START_WITH = 20
	NUM_OF_KEYS            = 10000
)

var machineIdCounter int64 = 1

type Machine struct {
	MachineId int64
	IP        string
	Hash      string
}

func CalculateHashInStr(str string) string {
	hash := sha256.Sum256([]byte(str))
	//Convert the hash byte array to a slice arr[:] short for arr[0:len(arr)]
	hashStr := hex.EncodeToString(hash[:])

	return hashStr
}

func NewMachine(ip string) Machine {
	newMachine := Machine{
		MachineId: atomic.LoadInt64(&machineIdCounter),
		IP:        ip,
		Hash:      CalculateHashInStr(ip),
	}

	atomic.AddInt64(&machineIdCounter, 1)

	return newMachine
}

func FindMachineToAssign(machines []Machine, key string) *Machine {
	var keyHash string = CalculateHashInStr(key)
	sort.Slice(machines, func(i, j int) bool {
		return machines[i].Hash < machines[j].Hash
	})

	for _, m := range machines {
		if m.Hash > keyHash {
			return &m
		}
	}

	if len(machines) > 0 {
		return &machines[0]
	}

	return nil
}

func generateRandomIP() string {
	ip := make([]byte, 4)
	_, err := rand.Read(ip)
	if err != nil {
		return ""
	}

	// Avoid special ranges like 0.x.x.x, 127.x.x.x, 255.x.x.x
	if ip[0] == 0 || ip[0] == 127 || ip[0] == 255 {
		ip[0] = 1 + ip[0]%223 // ensure it's in a usable range
	}

	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func generateRandomIPList(n int) []string {
	ipSet := make(map[string]bool)
	result := make([]string, 0)

	for len(result) < n {
		ip := generateRandomIP()
		_, exists := ipSet[ip]

		if !exists {
			ipSet[ip] = true
			result = append(result, ip)
		}
	}

	return result
}

func generateRandomMachineList(n int) []Machine {
	ipList := generateRandomIPList(n)
	machines := make([]Machine, 0)

	for _, ip := range ipList {
		machines = append(machines, NewMachine(ip))
	}

	return machines
}

func generateKeyList(n int) []string {
	keys := make([]string, 0)

	for i := range n {
		keys = append(keys, fmt.Sprintf("Key_%d", i))
	}

	return keys
}

func assignKeysForMachines(machines []Machine, keys []string) map[int64][]string {
	machineToKeyCountMap := make(map[int64][]string)

	for _, k := range keys {
		machineForK := FindMachineToAssign(machines, k)

		if machineToKeyCountMap[machineForK.MachineId] == nil {
			machineToKeyCountMap[machineForK.MachineId] = make([]string, 0)
		}

		keyArr := machineToKeyCountMap[machineForK.MachineId]
		keyArr = append(keyArr, k)
		machineToKeyCountMap[machineForK.MachineId] = keyArr
	}

	return machineToKeyCountMap
}

func main() {
	machines := generateRandomMachineList(MACHINES_TO_START_WITH)
	keys := generateKeyList(NUM_OF_KEYS)

	machineToKeyCountMap := assignKeysForMachines(machines, keys)

	printMap(machineToKeyCountMap)

	newMachines := generateRandomMachineList(2)

	fmt.Println("==================================")

	machineToKeyCountMap = assignKeysForMachines(append(machines, newMachines...), keys)
	printMap(machineToKeyCountMap)
}

func printMap(machineToKeyMap map[int64][]string) {
	for machineId, keyArr := range machineToKeyMap {
		fmt.Println("Machine: ", machineId)
		if len(keyArr) == 0 {
			fmt.Println("No keys assigned")
		} else {
			fmt.Println("Keys assigned: ", len(keyArr))
			//fmt.Println("Keys: ", keyArr)
		}
	}
}
