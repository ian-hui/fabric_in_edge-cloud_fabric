#!/bin/bash

rm -rf crypto-config channel-artifacts && mkdir crypto-config channel-artifacts
cryptogen generate --config=crypto-config.yaml
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID fabric-channel
#generate channel config-channel.tx
configtxgen -profile userinfoChannel -outputCreateChannelTx ./channel-artifacts/userinfoChannel.tx -channelID myuserinfochannel
configtxgen -profile accessChannel -outputCreateChannelTx ./channel-artifacts/accessChannel.tx -channelID myaccesschannel
configtxgen -profile nodeinfoChannel -outputCreateChannelTx ./channel-artifacts/nodeinfoChannel.tx -channelID mynodeinfochannel
#generate anchorpeerfile for every channel
#org1
configtxgen -profile userinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors_userinfoChannel.tx -channelID myuserinfochannel -asOrg Org1MSP
configtxgen -profile accessChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors_accessChannel.tx -channelID myaccesschannel -asOrg Org1MSP
configtxgen -profile nodeinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors_nodeinfoChannel.tx -channelID mynodeinfochannel -asOrg Org1MSP
#org2
configtxgen -profile userinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors_userinfoChannel.tx -channelID myuserinfochannel -asOrg Org2MSP
configtxgen -profile accessChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors_accessChannel.tx -channelID myaccesschannel -asOrg Org2MSP
configtxgen -profile nodeinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors_nodeinfoChannel.tx -channelID mynodeinfochannel -asOrg Org2MSP
#org2
configtxgen -profile userinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors_userinfoChannel.tx -channelID myuserinfochannel -asOrg Org3MSP
configtxgen -profile accessChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors_accessChannel.tx -channelID myaccesschannel -asOrg Org3MSP
configtxgen -profile nodeinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors_nodeinfoChannel.tx -channelID mynodeinfochannel -asOrg Org3MSP
#org2
configtxgen -profile userinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org4MSPanchors_userinfoChannel.tx -channelID myuserinfochannel -asOrg Org4MSP
configtxgen -profile accessChannel -outputAnchorPeersUpdate ./channel-artifacts/Org4MSPanchors_accessChannel.tx -channelID myaccesschannel -asOrg Org4MSP
configtxgen -profile nodeinfoChannel -outputAnchorPeersUpdate ./channel-artifacts/Org4MSPanchors_nodeinfoChannel.tx -channelID mynodeinfochannel -asOrg Org4MSP
