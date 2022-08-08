package main

import (
	"fmt"
	"grid/sdkInit"
	"os"
)

const (
	cc_name    = "simplecc"
	cc_version = "1.0.0"
)

func main() {
	// init orgs information
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org2",
			OrgMspId:      "Org2MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/grid/fixtures/channel-artifacts/Org2MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/grid/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
	}

	// init sdk env info 初始化sdk相关信息
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/grid/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH") + "/src/grid/chaincode/",
		ChaincodeVersion: cc_version,
	}

	// sdk setup 调用setup方法将sdk初始化
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}
	//fmt.Printf("%+v\n%+v\n", *orgs[0], *orgs[0])
	// create channel and join 调用CreateAndJoinChannel方法，创建并加入通道
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle 调用CreateCCLifecycle方法实现链码生命周期
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	//edu := service.Education{
	//	Name:           "张三",
	//	Gender:         "男",
	//	Nation:         "汉",
	//	EntityID:       "101",
	//	Place:          "北京",
	//	BirthDay:       "1991年01月01日",
	//	EnrollDate:     "2009年9月",
	//	GraduationDate: "2013年7月",
	//	SchoolName:     "中国政法大学",
	//	Major:          "社会学",
	//	QuaType:        "普通",
	//	Length:         "四年",
	//	Mode:           "普通全日制",
	//	Level:          "本科",
	//	Graduation:     "毕业",
	//	CertNo:         "111",
	//	Photo:          "/static/photo/11.png",
	//}
	//
	//serviceSetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	//if err != nil {
	//	fmt.Println()
	//	os.Exit(-1)
	//}
	//msg, err := serviceSetup.SaveEdu(edu)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("信息发布成功, 交易编号为: " + msg)
	//}
	//
	//result, err := serviceSetup.FindEduInfoByEntityID("101")
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	var edu service.Education
	//	json.Unmarshal(result, &edu)
	//	fmt.Println("根据身份证号码查询信息成功：")
	//	fmt.Println(edu)
	//}
	//
	//app := controller.Application{
	//	Setup: serviceSetup,
	//}
	//web.WebStart(app)
}
