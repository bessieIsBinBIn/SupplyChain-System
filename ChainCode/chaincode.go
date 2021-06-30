package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"time"
)

type BlockChainRealEstate struct {
}

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "StoreOrder":
		return routers.StoreOrder(stub,args)
	case "QueryOrder":
		return routers.QueryOrder(stub,nil)
	case "QueryOrderByID":
		return routers.QueryOrderByID(stub, args)
	case "QueryOrderByBuyer":
		return routers.QueryOrderByBuyer(stub, args)
	case "QueryOrderBySeller":
		return routers.QueryOrderBySeller(stub, args)
	case "UpdateStatus":
		return routers.UpdateStatus(stub, args)
	case "UpdateTransferStatus":
		return routers.UpdateTransferStatus(stub, args)
	case "DeleteOrder":
		return routers.DeleteOrder(stub, args)
	case "StoreProduct":
		return routers.StoreProduct(stub, args)
	case "UpdateProductByAmount":
		return routers.UpdateProductByAmount(stub, args)
	case "QueryProduct":
		return routers.QueryProduct(stub, args)
	case "QueryProductByID":
		return routers.QueryProductByID(stub, args)
	case "QueryProductByGoodType":
		return routers.QueryProductByGoodType(stub, args)
	case "QueryProductByMaterial":
		return routers.QueryProductByMaterial(stub, args)
	case "QueryProductByWorkmanship":
		return routers.QueryProductByWorkmanship(stub, args)
	case "QueryProductBySupplier":
		return routers.QueryProductBySupplier(stub, args)
	case "DeleteProduct":
		return routers.DeleteProduct(stub, args)
	case "StoreFinance":
		return routers.StoreFinance(stub, args)
	case "UpdateFinanceStatus":
		return routers.UpdateFinanceStatus(stub, args)
	case "QueryFinance":
		return routers.QueryFinance(stub, args)
	case "QueryFinanceByStatus":
		return routers.QueryFinanceByStatus(stub, args)
	case "QueryFinanceByOrderID":
		return routers.QueryFinanceByOrderID(stub, args)
	case "StoreUser":
		return routers.StoreUser(stub, args)
	case "QueryUser":
		return routers.QueryUser(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}