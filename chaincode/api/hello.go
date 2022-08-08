package api

import (
	"grid/chaincode/pkg/utils"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// Hello 测试
func Hello(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	err := utils.WriteLedger(map[string]interface{}{"msg": "hello"}, stub, "hello", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("hello world"))
}
