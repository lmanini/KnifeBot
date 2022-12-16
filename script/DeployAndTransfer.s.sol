// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import {SpyNFT, KnifeNFT} from "../src/NFT.sol";

contract DeployAndTransferScript is Script {
    address deployerAddr = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
    uint256 deployerPrivKey =
        0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

    address peer1 = address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8);
    address peer2 = address(0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC);
    address peer3 = address(0x90F79bf6EB2c4f870365E785982E1f101E93b906);

    function setUp() public {}

    function run() public {
        vm.startBroadcast(deployerPrivKey);

        SpyNFT spy = new SpyNFT();
        uint256 id1 = spy.mint(deployerAddr, block.timestamp * 10);
        uint256 id2 = spy.mint(deployerAddr, block.timestamp * 10);
        uint256 id3 = spy.mint(deployerAddr, block.timestamp * 10);

        spy.sudoTransferFrom(deployerAddr, peer1, id1);
        spy.sudoTransferFrom(deployerAddr, peer2, id2);
        spy.sudoTransferFrom(deployerAddr, peer3, id3);

        /**
        
        deployerAddr -- id1 --> peer1
        deployerAddr -- id2 --> peer2
        deployerAddr -- id3 --> peer3
    
        deployerAddr : {}
        peer1: {id1}
        peer2: {id2}
        peer3: {id3}
         
         */
        vm.stopBroadcast();
    }
}
