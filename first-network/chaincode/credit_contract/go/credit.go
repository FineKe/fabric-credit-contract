package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

const (
	FunctionRegister = "register"
	FunctionTransfer = "transfer"
	FunctionQuery    = "query"
	FunctionConsume  = "consume"
	FunctionCharge   = "charge"
)

type CreditContract struct {
}

// init creditContract
func (t *CreditContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("CreditContract")
	return shim.Success([]byte("init successful"))
}

func (t *CreditContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	switch function {
	case FunctionRegister:
		return t.register(stub, args)
	case FunctionTransfer:
		return t.transfer(stub, args)
	case FunctionQuery:
		return t.query(stub, args)
	case FunctionConsume:
		return t.consume(stub, args)
	case FunctionCharge:
		return t.charge(stub, args)
	default:
		return shim.Error("Invalid function")
	}
}

// register
func (t *CreditContract) register(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	account := args[0]
	// check account exists
	bytes, err := stub.GetState(account)
	if err != nil && bytes != nil {
		return shim.Error("can't register account,the account exists")
	}
	// register account and init credit as zero
	if err := stub.PutState(account, []byte(strconv.Itoa(0))); err != nil {
		return shim.Error("register account failed: " + err.Error())
	}
	return shim.Success([]byte("register account success"))
}

// transfer credit
func (t *CreditContract) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments")
	}
	// source account
	from := args[0]
	// dest account
	to := args[1]
	// convert int value
	creditValue, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid param :" + err.Error())
	}
	// check from account exists
	if bytes, err := stub.GetState(from); err != nil || bytes == nil {
		return shim.Error("The from account not exists")
	}
	// check to account exists
	if bytes, err := stub.GetState(to); err != nil || bytes == nil {
		return shim.Error("The to account not exists")
	}
	// check from account balance
	bytesf, _ := stub.GetState(from)
	fromCredit, _ := strconv.Atoi(string(bytesf))
	if fromCredit < creditValue {
		return shim.Error("Sorry, from account haven't enough credit")
	}
	// get to account credit
	bytest, _ := stub.GetState(to)
	toCredit, _ := strconv.Atoi(string(bytest))
	// transfer credit
	fromCredit = fromCredit - creditValue
	toCredit = toCredit + creditValue
	stub.PutState(from, []byte(strconv.Itoa(fromCredit)))
	stub.PutState(to, []byte(strconv.Itoa(toCredit)))
	return shim.Success([]byte("transfer success"))
}

// query credit
func (t *CreditContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	bytes, err := stub.GetState(args[0])
	if err != nil || bytes == nil {
		return shim.Error("query failed")
	}
	return shim.Success(bytes)
}

// consume credit
func (t *CreditContract) consume(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments")
	}
	account:=args[0]
	credit,err:=strconv.Atoi(args[1])
	if err!=nil {
		return shim.Error("convert credit failed : "+err.Error())
	}
	// check account exists
	bytes,err:=stub.GetState(account)
	if err!=nil||bytes==nil {
		return shim.Error("account not exists")
	}
	balance,err:=strconv.Atoi(string(bytes))
	if err!=nil {
		return shim.Error("get balance failed: "+err.Error())
	}
	if balance < credit{
		return shim.Error("balance not enough")
	}
	balance -= credit
	if err:=stub.PutState(account,[]byte(strconv.Itoa(balance)));err!=nil{
		return shim.Error("consume failed: "+err.Error())
	}
	return shim.Success([]byte("consume "+args[1] + "credit"))
}


// charge credit
func (t *CreditContract) charge(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args)!=2 {
		return shim.Error("Incorrect number of arguments")
	}
	account:=args[0]
	chargeValue,err := strconv.Atoi(args[1])
	if err!=nil {
		return shim.Error("convert charge failed: "+err.Error())
	}
	if chargeValue<=0 {
		return shim.Error("charge credit should be more than 0")
	}
	bytes,err:=stub.GetState(account)
	if err!=nil||bytes==nil{
		return shim.Error("account not exists")
	}
	balance,err:=strconv.Atoi(string(bytes))
	if err!=nil{
		return shim.Error("convert balance failed: "+err.Error())
	}
	balance += chargeValue
	if err:=stub.PutState(account,[]byte(strconv.Itoa(balance)));err!=nil{
		return shim.Error("charge failed :"+err.Error())
	}
	return shim.Success([]byte("charge success"))
}

func main()  {
	err:=shim.Start(new(CreditContract))
	if err!=nil{
		fmt.Printf("Error starting CreditContract: %s",err)
	}
}
