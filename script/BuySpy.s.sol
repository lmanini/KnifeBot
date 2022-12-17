// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import {SpyNFT, KnifeNFT} from "../src/NFT.sol";
import "../src/KnifeGame.sol";

contract BuySpyScript is Script {
    SpyNFT spies;
    KnifeNFT knives;
    KnifeGame game;

    address[] addresses = [
        0xbFc58f666741e3d6794cFeC9153f6CFBda90dabD, 
        0xC5B0a210A702d59774EFb08db8025954001D0345, 
        0x053723A8D91D56edC1d35f571c655ef83fBaabB9, 
        0x4e80FE3c57BB8C0D0EE85D956c4215108947E071
    ];

    function setUp() public {
        spies = SpyNFT(address(0x265f4EbCB310bB98e76436569D014D3D55087Aa8));
        knives = KnifeNFT(address(0x0434Ba7f2F795C083bF1f8692acd36721cD34799));
        game = KnifeGame(address(0x6488b54F9502e9A4B1D7BB90863012F682fe60c6));
    }

    function run() public {
        uint256 deployerPk1 = vm.envUint("PK1");
        vm.startBroadcast(deployerPk1);
        game.purchaseSpy{value: 0.1 ether}(addresses[0]);
        vm.stopBroadcast();
    }
}
