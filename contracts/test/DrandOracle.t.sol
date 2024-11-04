// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {stdJson} from "forge-std/StdJson.sol";
import {Test, console} from "forge-std/Test.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {IDrandOracle} from "../src/IDrandOracle.sol";
import {DrandOracle} from "../src/DrandOracle.sol";

contract DrandOracleTest is Test {
    using stdJson for string;

    uint256 internal ownerPrivateKey;
    uint256 internal signerPrivateKey;
    address internal owner;
    address internal signer;

    string constant oracleName = "DrandOracle";
    string constant oracleVersion = "1.0.0";

    bytes32 constant chainHash = bytes32(0x8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce);

    DrandOracle public oracle;

    bytes32 internal constant _TYPE_HASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
    bytes32 _HASHED_NAME;
    bytes32 _HASHED_VERSION;

    function setUp() public {
        ownerPrivateKey = 0xA11CE;
        signerPrivateKey = 0xB0B;
        owner = vm.addr(ownerPrivateKey);
        signer = vm.addr(signerPrivateKey);

        oracle = new DrandOracle{salt: "DrandOracle"}(owner, signer, chainHash);

        bytes32 hashedName = keccak256(bytes(oracleName));
        bytes32 hashedVersion = keccak256(bytes(oracleVersion));
        _HASHED_NAME = hashedName;
        _HASHED_VERSION = hashedVersion;
    }

    function test_setRandomness_firstTime_success() public {
        uint64 round = 4493690;
        bytes32 randomnessBytes = bytes32(hex"a502d9e94fd02472fbd292e054893fb26a37490610a4e6ec29734a20f359b9c5");
        bytes memory signatureBytes =
            hex"b699e3214449474050112c80e01515b3250084e57f7fae9138e0a509aee2a641098450f9a5be3c6485c0fb8e8b3c8581029d1192ec6f9eb8d3a5e307f54ffd512a184562dc7f842d44b09e8481bc4ed38cfbd9c2d5b4e92f97e1baf6e97767c5";
        uint64 timestamp = 1724995200;

        IDrandOracle.Random memory randomData = IDrandOracle.Random({
            round: round,
            timestamp: timestamp,
            randomness: randomnessBytes,
            signature: signatureBytes
        });

        bytes32 hash = _hashSetRandomness(randomData);
        bytes memory signature = _signMessage(hash, signerPrivateKey);

        oracle.setRandomness(randomData, signature);

        IDrandOracle.Random memory retrievedData = oracle.getRandomnessFromRound(round);
        assertEq(retrievedData.round, round);
        assertEq(retrievedData.randomness, randomnessBytes);
        assertEq(retrievedData.signature, signatureBytes);
        assertEq(retrievedData.timestamp, timestamp);
        assertEq(oracle.earliestRound(), round);
        assertEq(oracle.latestRound(), round);
    }

    function test_getRandomnessFromTimestamp_success() public {
        uint64 round = 4493690;
        bytes32 randomnessBytes = bytes32(hex"a502d9e94fd02472fbd292e054893fb26a37490610a4e6ec29734a20f359b9c5");
        bytes memory signatureBytes =
            hex"b699e3214449474050112c80e01515b3250084e57f7fae9138e0a509aee2a641098450f9a5be3c6485c0fb8e8b3c8581029d1192ec6f9eb8d3a5e307f54ffd512a184562dc7f842d44b09e8481bc4ed38cfbd9c2d5b4e92f97e1baf6e97767c5";
        uint64 timestamp = 1724995200;

        IDrandOracle.Random memory randomData = IDrandOracle.Random({
            round: round,
            timestamp: timestamp,
            randomness: randomnessBytes,
            signature: signatureBytes
        });

        bytes32 hash = _hashSetRandomness(randomData);
        bytes memory signature = _signMessage(hash, signerPrivateKey);

        oracle.setRandomness(randomData, signature);

        // Test getting the randomness from a timestamp that is close to the set randomness timestamp
        uint64 newTimestamp = timestamp + 20;

        IDrandOracle.Random memory retrievedData = oracle.getRandomnessFromTimestamp(newTimestamp);
        assertEq(retrievedData.round, round);
        assertEq(retrievedData.randomness, randomnessBytes);
        assertEq(retrievedData.signature, signatureBytes);
        assertEq(retrievedData.timestamp, timestamp);

        // Test getting the randomness from a timestamp that is far from the set randomness timestamp
        // In cases like this, our oracle is likely haven't been updated for a while, this tests that the binary search works
        newTimestamp = timestamp + 2000000;

        retrievedData = oracle.getRandomnessFromTimestamp(newTimestamp);
        assertEq(retrievedData.round, round);
        assertEq(retrievedData.randomness, randomnessBytes);
        assertEq(retrievedData.signature, signatureBytes);
        assertEq(retrievedData.timestamp, timestamp);
    }

    function test_getRandomnessFromTimestamp_multipleRounds_success() public {
        IDrandOracle.Random memory randomData1 = IDrandOracle.Random({
            round: 4508563,
            timestamp: 1730687910,
            randomness: bytes32(hex"152fbcf71e680a5bc43e0374c0fd5e8bd5c87e21884fbc8c8396bb372a49c088"),
            signature: hex"a0d0eb129fb178dc76f749edd34b3fa333c4fabea9cfa4e6e93c3ddd65d2a75d431eb62cd20a3425331bcd1000a5897314e851377c314a707cc2abbcca2acc7d6bc858171940d45ed0bc834bc67286c016b33454c8b8fd45a74d025b2a7ba923"
        });
        bytes memory signature1 = _signMessage(_hashSetRandomness(randomData1), signerPrivateKey);
        oracle.setRandomness(randomData1, signature1);

        IDrandOracle.Random memory randomData2 = IDrandOracle.Random({
            round: 4508564,
            timestamp: 1730687940,
            randomness: bytes32(hex"32fd64310176d074234a3cef76dce4ba63c9c0dfb7941d4dbdca46ebf7a7afad"),
            signature: hex"ac6c0641c64317951e640a5601a47c1dcc5d0a24d99b2456fd2d870efb3abd43f656d90116aa89fd932ba601e3ef997615372fd2c4c97d35f3e0eca948c0d054276c698537649ab373acda4bb2a454e7fcee31f04a5541325f407a0f2d602fa5"
        });
        bytes memory signature2 = _signMessage(_hashSetRandomness(randomData2), signerPrivateKey);

        oracle.setRandomness(randomData2, signature2);

        // Test getting the randomness from a timestamp that is right after the first round timestamp
        uint64 newTimestamp1 = randomData1.timestamp + 1;
        IDrandOracle.Random memory retrievedData1 = oracle.getRandomnessFromTimestamp(newTimestamp1);
        assertEq(retrievedData1.round, randomData1.round);
        assertEq(retrievedData1.randomness, randomData1.randomness);
        assertEq(retrievedData1.signature, randomData1.signature);
        assertEq(retrievedData1.timestamp, randomData1.timestamp);

        // Test getting the randomness from a timestamp that is right after the second round timestamp
        uint64 newTimestamp2 = randomData2.timestamp + 1;
        IDrandOracle.Random memory retrievedData2 = oracle.getRandomnessFromTimestamp(newTimestamp2);
        assertEq(retrievedData2.round, randomData2.round);
        assertEq(retrievedData2.randomness, randomData2.randomness);
        assertEq(retrievedData2.signature, randomData2.signature);
        assertEq(retrievedData2.timestamp, randomData2.timestamp);

        // Test getting the randomness from a timestamp that is far from the second round timestamp
        // In cases like this, our oracle is likely haven't been updated for a while, this tests that the binary search works
        uint64 newTimestamp3 = randomData2.timestamp + 2000000;
        IDrandOracle.Random memory retrievedData3 = oracle.getRandomnessFromTimestamp(newTimestamp3);
        assertEq(retrievedData3.round, randomData2.round);
        assertEq(retrievedData3.randomness, randomData2.randomness);
        assertEq(retrievedData3.signature, randomData2.signature);
        assertEq(retrievedData3.timestamp, randomData2.timestamp);
    }

    function test_getRandomnessFromTimestamp_invalidTimestamp() public {
        uint64 round = 4493690;
        bytes32 randomnessBytes = bytes32(hex"a502d9e94fd02472fbd292e054893fb26a37490610a4e6ec29734a20f359b9c5");
        bytes memory signatureBytes =
            hex"b699e3214449474050112c80e01515b3250084e57f7fae9138e0a509aee2a641098450f9a5be3c6485c0fb8e8b3c8581029d1192ec6f9eb8d3a5e307f54ffd512a184562dc7f842d44b09e8481bc4ed38cfbd9c2d5b4e92f97e1baf6e97767c5";
        uint64 timestamp = 1724995200;

        IDrandOracle.Random memory randomData = IDrandOracle.Random({
            round: round,
            timestamp: timestamp,
            randomness: randomnessBytes,
            signature: signatureBytes
        });

        bytes32 hash = _hashSetRandomness(randomData);
        bytes memory signature = _signMessage(hash, signerPrivateKey);

        oracle.setRandomness(randomData, signature);

        uint64 newTimestamp = timestamp - 1;

        vm.expectRevert(abi.encodeWithSelector(IDrandOracle.InvalidRoundTimestamp.selector));
        oracle.getRandomnessFromTimestamp(newTimestamp);
    }

    function _signMessage(bytes32 hash, uint256 privateKey) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, hash);
        return abi.encodePacked(r, s, v);
    }

    // --------------- hash functions ------------- //
    function _hashSetRandomness(IDrandOracle.Random memory _random) private view returns (bytes32) {
        return _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256("SetRandomness(uint64 round,uint64 timestamp,bytes32 randomness,bytes signature)"),
                    _random.round,
                    _random.timestamp,
                    _random.randomness,
                    keccak256(_random.signature)
                )
            )
        );
    }

    // --------------- EIP712 signature tools ------------- //
    function _hashTypedDataV4(bytes32 structHash) internal view virtual returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _buildDomainSeparator(), structHash));
    }

    function _buildDomainSeparator() private view returns (bytes32) {
        return keccak256(abi.encode(_TYPE_HASH, _HASHED_NAME, _HASHED_VERSION, _getChainId(), address(oracle)));
    }

    function _getChainId() private view returns (uint256 chainId) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        // solhint-disable-next-line no-inline-assembly
        assembly {
            chainId := chainid()
        }
    }
}
