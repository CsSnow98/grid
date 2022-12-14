#!/bin/bash


echo "Linux"
export PATH=${PWD}/bin:$PATH

echo "一、清理环境"
./stop.sh

echo "二、生成证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
cryptogen generate --config=crypto-config.yaml

echo "三、创建排序通道创世区块"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID fabric-channel

echo "四、生成通道配置事务'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel

echo "五、为 Org1 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP

echo "六、为 Org2 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP

#echo "---please execute clean_docker---"
#docker-compose up -d
#echo "正在等待节点的启动完成，等待10秒"
#sleep 10
#
echo "---please execute clean_docker---"
##
#TaobaoPeer0Cli="CORE_PEER_ADDRESS=peer0.org1.example.com:7051 CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/org1.example.com/users/Admin@org1.example.com/msp"
##TaobaoPeer1Cli="CORE_PEER_ADDRESS=peer1.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp"
#JDPeer0Cli="CORE_PEER_ADDRESS=peer0.org2.example.com:7051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/org2.example.com/users/Admin@org2.example.com/msp"
##JDPeer1Cli="CORE_PEER_ADDRESS=peer1.jd.com:7051 CORE_PEER_LOCALMSPID=JDMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/jd.com/users/Admin@jd.com/msp"
##
#echo "七、创建通道"
#docker exec cli bash -c "$TaobaoPeer0Cli peer channel create -o orderer.example.com:7050 -c mychannel -f /etc/hyperledger/config/channel.tx"
##
#echo "八、将所有节点加入通道"
#docker exec cli bash -c "$TaobaoPeer0Cli peer channel join -b channel.block"
##docker exec cli bash -c "$TaobaoPeer1Cli peer channel join -b appchannel.block"
#docker exec cli bash -c "$JDPeer0Cli peer channel join -b channel.block"
##docker exec cli bash -c "$JDPeer1Cli peer channel join -b appchannel.block"
##
#echo "九、更新锚节点"
#docker exec cli bash -c "$TaobaoPeer0Cli peer channel update -o orderer.example.com:7050 -c mychannel -f /etc/hyperledger/config/Org1MSPanchors.tx"
#docker exec cli bash -c "$JDPeer0Cli peer channel update -o orderer.example.com:7050 -c mychannel -f /etc/hyperledger/config/Org2MSPanchors.tx"
##
### -n 链码名，可以自己随便设置
### -v 版本号
### -p 链码目录，在 /opt/gopath/src/ 目录下
#echo "十、安装链码"
#docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"
#docker exec cli bash -c "$JDPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"
#
## 只需要其中一个节点实例化
## -n 对应上一步安装链码的名字
## -v 版本号
## -C 是通道，在fabric的世界，一个通道就是一条不同的链
## -c 为传参，传入init参数
#echo "十一、实例化链码"
#docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode instantiate -o orderer.qq.com:7050 -C appchannel -n fabric-realty -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('TaobaoMSP.member','JDMSP.member')\""
#
#echo "正在等待链码实例化完成，等待5秒"
#sleep 5
#
## 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
#echo "十二、验证链码"
#docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"
#docker exec cli bash -c "$JDPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"