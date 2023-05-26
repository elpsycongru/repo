#!/bin/bash

version="1.4.12" # 2.5

if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    echo "Version : 1.4.12"
    export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH
fi

if [[ `uname` == 'Darwin' ]]; then
    echo "Darwin"
    echo "Version : 1.4.12"
    export PATH=${PWD}/hyperledger-fabric-darwin-amd64-1.4.12/bin:$PATH
fi

#if [ "$version" == "2.5" ]; then
#    echo "version : 2.5" 
#    export PATH=${PWD}/hyperledger-fabric-linux-amd64-2.5/bin:$PATH
#fi

#if [ "$version" == "1.4.12" ]; then 
#    echo "Version : 1.4.12"
#    export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH
#fi

echo "一、清理环境"
./stop.sh

echo "二、生成证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
cryptogen generate --config=./crypto-config.yaml

echo "三、创建排序通道创世区块"
configtxgen  -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID firstchannel
#configtxgen -configPath ./network/configtx -profile SampleSingleMSPSolo -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

echo "四、生成通道配置事务'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo "五、为 ORG1 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/ORG1Anchor.tx -channelID appchannel -asOrg ORG1

echo "六、为 ORG2 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/ORG2Anchor.tx -channelID appchannel -asOrg ORG2
#configtxgen -profile SampleSingleMSPChannel -outputAnchorPeersUpdate ./config/TESLAAnchor.tx -channelID appchannel -asOrg TESLA

echo "区块链 ： 启动"
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10

#MYORGPeer0Cli="CORE_PEER_ADDRESS=peer0.myorg.com:7051 CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_MSPCONFIGPATH=/root/go/src/github.com/elpsycongru/communication-system-blockchain/network/crypto-config/peerOrganizations/myorg.com/users/Admin@myorg.com/msp"
#MYORGPeer1Cli="CORE_PEER_ADDRESS=peer1.myorg.com:7051 CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_MSPCONFIGPATH=/root/go/src/github.com/elpsycongru/communication-system-blockchain/network/crypto-config/peerOrganizations/myorg.com/users/Admin@myorg.com/msp"
#MYORGPeer2Cli="CORE_PEER_ADDRESS=peer2.myorg.com:7051 CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_MSPCONFIGPATH=/root/go/src/github.com/elpsycongru/communication-system-blockchain/network/crypto-config/peerOrganizations/myorg.com/users/Admin@myorg.com/msp"
ORG1Peer0Cli="CORE_PEER_ADDRESS=peer0.org1.com:7051 CORE_PEER_LOCALMSPID=ORG1MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/org1.com/users/Admin@org1.com/msp"
ORG1Peer1Cli="CORE_PEER_ADDRESS=peer1.org1.com:7051 CORE_PEER_LOCALMSPID=ORG1MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/org1.com/users/Admin@org1.com/msp"
ORG2Peer0Cli="CORE_PEER_ADDRESS=peer0.org2.com:7051 CORE_PEER_LOCALMSPID=ORG2MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/org2.com/users/Admin@org2.com/msp"
ORG2Peer1Cli="CORE_PEER_ADDRESS=peer1.org2.com:7051 CORE_PEER_LOCALMSPID=ORG2MSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/org2.com/users/Admin@org2.com/msp"



#MYORGPeer1Cli="CORE_PEER_ADDRESS=peer1.myorg.com:7051 CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/myorg.com/users/Admin@myorg.com/msp"
#MYORGPeer2Cli="CORE_PEER_ADDRESS=peer2.myorg.com:7051 CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/myorg.com/users/Admin@myorg.com/msp"

#MYORGAdminCli="CORE_PEER_ADDRESS=cli:7051 CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/myorg.com/users/Admin@myorg.com/msp"

#TESLAPeer0Cli="CORE_PEER_ADDRESS=peer0.tesla.com:7051 CORE_PEER_LOCALMSPID=TESLAMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/tesla.com/users/Admin@tesla.com/msp"
#TESLAPeer1Cli="CORE_PEER_ADDRESS=peer1.tesla.com:7051 CORE_PEER_LOCALMSPID=TESLAMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/tesla.com/users/Admin@tesla.com/msp"

echo "七、创建通道"

#MYORGCli="CORE_PEER_ADDRESS=peer0.myorg.com:7051 CORE_PEER_TLS_ENABLED=true CORE_PEER_LOCALMSPID=MYORGMSP CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto-config/peerOrganizations/myorg.com//peer0.myorg.com/tls/ca.crt " #CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/myorg.com/users/Admin@myorg.com/msp"
#export CORE_PEER_TLS_ENABLED=true
#export CORE_PEER_LOCALMSPID="MYORGMSP"
#export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto-config/peerOrganizations/myorg.com/peers/peer0.myorg.com/tls/ca.crt
#export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/myorg.com/users/Admin@myorg.com/msp
#export CORE_PEER_ADDRESS=localhost:7051 

#
#docker exec cli bash -c "$ORG1Peer0Cli peer channel list"


docker exec cli bash -c "$ORG1Peer0Cli peer channel create -o orderer.carunion.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx"
#docker exec cli bash -c "$MYORGPeer0Cli peer channel create -o orderer.carunion.com:7050 -c appchannel -f /root/go/src/github.com/elpsycongru/communication-system-blockchain/network/config/appchannel.tx"


echo "八、将所有节点加入通道"
docker exec cli bash -c "$ORG1Peer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$ORG1Peer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$ORG2Peer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$ORG2Peer1Cli peer channel join -b appchannel.block"

echo "九、更新锚节点"
docker exec cli bash -c "$ORG1Peer0Cli peer channel update -o orderer.carunion.com:7050 -c appchannel -f /etc/hyperledger/config/ORG1Anchor.tx"
docker exec cli bash -c "$ORG2Peer0Cli peer channel update -o orderer.carunion.com:7050 -c appchannel -f /etc/hyperledger/config/ORG2Anchor.tx"

# -n 链码名，可以自己随便设置
# -v 版本号
# -p 链码目录，在 /opt/gopath/src/ 目录下
echo "十、安装链码"
docker exec cli bash -c "$ORG1Peer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"
docker exec cli bash -c "$ORG2Peer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"
#docker exec cli bash -c "$MYORGPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"

# 只需要其中一个节点实例化
# -n 对应上一步安装链码的名字
# -v 版本号
# -C 是通道，在fabric的世界，一个通道就是一条不同的链
# -c 为传参，传入init参数
echo "十一、实例化链码"
docker exec cli bash -c "$ORG1Peer0Cli peer chaincode instantiate -o orderer.carunion.com:7050 -C appchannel -n fabric-realty -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('ORG1MSP.member','ORG2MSP.member')\""

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作

echo "十二、验证链码"
docker exec cli bash -c "$ORG1Peer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"
#docker exec cli bash -c "$TESLAPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"