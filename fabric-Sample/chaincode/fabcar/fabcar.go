/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

const REVOKED = "Revoked"

// Define the Smart Contract structure
type SmartContract struct {
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately

	switch function {
	case "initLedger":
		return s.initLedger(APIstub)
	case "addFile":
		return s.addFile(APIstub, args)
	case "verifyFile":
		return s.verifyFile(APIstub, args)
	case "revoke":
		return s.revoke(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

//@Parma
// Args 0 = Key
// Args 1 = Data subject ID
// Args 2 = Data Cust ID
// Args 3 = file Hash
func (s *SmartContract) addFile(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	const NUMBER_OF_ARGS = 4

	if len(args) != NUMBER_OF_ARGS {
		return shim.Error("Incorrect number of arguments. Expecting " + strconv.Itoa(NUMBER_OF_ARGS))
	}

	APIstub.PutState(args[1]+args[2]+args[0], []byte(args[1]+args[2]+args[3]))

	return shim.Success([]byte("ADDED: KEY:" + args[0] + " VALUE:" + args[1] + args[2] + args[3]))
}

//@Parma
// Args 0 = Key
// Args 1 = Data subject ID
// Args 2 = Data Cust ID
func (s *SmartContract) revoke(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	const NUMBER_OF_ARGS = 3

	if len(args) != NUMBER_OF_ARGS {
		return shim.Error("Incorrect number of arguments. Expecting " + strconv.Itoa(NUMBER_OF_ARGS))
	}

	APIstub.PutState(args[1]+args[2]+args[0], []byte(REVOKED))

	return shim.Success([]byte("REVOKED " + args[0] + " Cert"))
}

//@Parma
//ARGS 0 = file ID
//Args 1 = Data subject ID
//Args 2 = Data Cust ID
//ARGS 3 = hash to check
func (s *SmartContract) verifyFile(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	const NUMBER_OF_ARGS = 4

	if len(args) != NUMBER_OF_ARGS {
		return shim.Error("Incorrect number of arguments. Expecting " + strconv.Itoa(NUMBER_OF_ARGS))
	}

	hashAsBytes, err := APIstub.GetState(args[1] + args[2] + args[0])

	if err != nil {
		return shim.Error("error")
	}

	if len(hashAsBytes) == 0 {
		return shim.Error("Key not found")
	}

	hash := string(hashAsBytes[:len(hashAsBytes)])

	if hash == args[3] {

		if hash == REVOKED {
			return shim.Success([]byte(REVOKED))
		}

		return shim.Success([]byte("TRUE"))
	}

	return shim.Success([]byte("FALSE " + hash + " != " + args[1]))
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	i := 0
	for i < 10 {
		fmt.Println("i is ", i)
		//@TODO no idea why im getting addFile undfined
		//addFile(APIstub, []string{"FILE" + strconv.Itoa(i), "0", "0", "abc"})
		APIstub.PutState("FILE"+strconv.Itoa(i), []byte(strconv.Itoa(i)+"0"+"abc"))
		i = i + 1
	}

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
