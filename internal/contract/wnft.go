// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WnftMetaData contains all meta data concerning the Wnft contract.
var WnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600e81526020017f4d792057726170706564204e46540000000000000000000000000000000000008152506040518060400160405280600581526020017f4d574e4654000000000000000000000000000000000000000000000000000000815250815f908161008a91906102df565b50806001908161009a91906102df565b5050506103ae565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061011d57607f821691505b6020821081036101305761012f6100d9565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610157565b61019c8683610157565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101e06101db6101d6846101b4565b6101bd565b6101b4565b9050919050565b5f819050919050565b6101f9836101c6565b61020d610205826101e7565b848454610163565b825550505050565b5f5f905090565b610224610215565b61022f8184846101f0565b505050565b5b81811015610252576102475f8261021c565b600181019050610235565b5050565b601f8211156102975761026881610136565b61027184610148565b81016020851015610280578190505b61029461028c85610148565b830182610234565b50505b505050565b5f82821c905092915050565b5f6102b75f198460080261029c565b1980831691505092915050565b5f6102cf83836102a8565b9150826002028217905092915050565b6102e8826100a2565b67ffffffffffffffff811115610301576103006100ac565b5b61030b8254610106565b610316828285610256565b5f60209050601f831160018114610347575f8415610335578287015190505b61033f85826102c4565b8655506103a6565b601f19841661035586610136565b5f5b8281101561037c57848901518255600182019150602085019450602081019050610357565b868310156103995784890151610395601f8916826102a8565b8355505b6001600288020188555050505b505050505050565b61293b806103bb5f395ff3fe608060405234801561000f575f5ffd5b5060043610610109575f3560e01c806342842e0e116100a057806395d89b411161006f57806395d89b41146102ed578063a22cb4651461030b578063b88d4fde14610327578063c87b56dd14610343578063e985e9c51461037357610109565b806342842e0e146102415780634f6ccce71461025d5780636352211e1461028d57806370a08231146102bd57610109565b8063095ea7b3116100dc578063095ea7b3146101bb57806318160ddd146101d757806323b872dd146101f55780632f745c591461021157610109565b8063013828071461010d57806301ffc9a71461013d57806306fdde031461016d578063081812fc1461018b575b5f5ffd5b61012760048036038101906101229190611f06565b6103a3565b6040516101349190611f78565b60405180910390f35b61015760048036038101906101529190611fe6565b6103d9565b604051610164919061202b565b60405180910390f35b6101756103fa565b60405161018291906120a4565b60405180910390f35b6101a560048036038101906101a091906120ee565b610489565b6040516101b29190612128565b60405180910390f35b6101d560048036038101906101d09190612141565b6104a4565b005b6101df6104ba565b6040516101ec9190611f78565b60405180910390f35b61020f600480360381019061020a919061217f565b6104c6565b005b61022b60048036038101906102269190612141565b6105c5565b6040516102389190611f78565b60405180910390f35b61025b6004803603810190610256919061217f565b610669565b005b610277600480360381019061027291906120ee565b610688565b6040516102849190611f78565b60405180910390f35b6102a760048036038101906102a291906120ee565b6106fa565b6040516102b49190612128565b60405180910390f35b6102d760048036038101906102d291906121cf565b61070b565b6040516102e49190611f78565b60405180910390f35b6102f56107c1565b60405161030291906120a4565b60405180910390f35b61032560048036038101906103209190612224565b610851565b005b610341600480360381019061033c9190612300565b610867565b005b61035d600480360381019061035891906120ee565b61088c565b60405161036a91906120a4565b60405180910390f35b61038d60048036038101906103889190612380565b61089e565b60405161039a919061202b565b60405180910390f35b5f6103ae600b61092c565b5f6103b9600b610940565b90506103c5848261094c565b6103cf8184610a3f565b8091505092915050565b5f6103e382610a99565b806103f357506103f282610b12565b5b9050919050565b60605f8054610408906123eb565b80601f0160208091040260200160405190810160405280929190818152602001828054610434906123eb565b801561047f5780601f106104565761010080835404028352916020019161047f565b820191905f5260205f20905b81548152906001019060200180831161046257829003601f168201915b5050505050905090565b5f61049382610b72565b5061049d82610bf8565b9050919050565b6104b682826104b1610c31565b610c38565b5050565b5f600980549050905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610536575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161052d9190612128565b60405180910390fd5b5f6105498383610544610c31565b610c4a565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146105bf578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016105b69392919061241b565b60405180910390fd5b50505050565b5f6105cf8361070b565b82106106145782826040517fa57d13dc00000000000000000000000000000000000000000000000000000000815260040161060b929190612450565b60405180910390fd5b60075f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b61068383838360405180602001604052805f815250610867565b505050565b5f6106916104ba565b82106106d6575f826040517fa57d13dc0000000000000000000000000000000000000000000000000000000081526004016106cd929190612450565b60405180910390fd5b600982815481106106ea576106e9612477565b5b905f5260205f2001549050919050565b5f61070482610b72565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361077c575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016107739190612128565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6060600180546107d0906123eb565b80601f01602080910402602001604051908101604052809291908181526020018280546107fc906123eb565b80156108475780601f1061081e57610100808354040283529160200191610847565b820191905f5260205f20905b81548152906001019060200180831161082a57829003601f168201915b5050505050905090565b61086361085c610c31565b8383610c5f565b5050565b6108728484846104c6565b61088661087d610c31565b85858585610dc8565b50505050565b606061089782610f74565b9050919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b6001815f015f828254019250508190555050565b5f815f01549050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109bc575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016109b39190612128565b60405180910390fd5b5f6109c883835f610c4a565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a3a575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401610a319190612128565b60405180910390fd5b505050565b8060065f8481526020019081526020015f209081610a5d9190612644565b507ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce782604051610a8d9190611f78565b60405180910390a15050565b5f7f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610b0b5750610b0a82610b12565b5b9050919050565b5f634906490660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610b6b5750610b6a8261107f565b5b9050919050565b5f5f610b7d83611160565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bef57826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610be69190611f78565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610c458383836001611199565b505050565b5f610c56848484611358565b90509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ccf57816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401610cc69190612128565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610dbb919061202b565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115610f6d578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401610e269493929190612765565b6020604051808303815f875af1925050508015610e6157506040513d601f19601f82011682018060405250810190610e5e91906127c3565b60015b610ee2573d805f8114610e8f576040519150601f19603f3d011682016040523d82523d5f602084013e610e94565b606091505b505f815103610eda57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610ed19190612128565b60405180910390fd5b805160208201fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610f6b57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610f629190612128565b60405180910390fd5b505b5050505050565b6060610f7f82610b72565b505f60065f8481526020019081526020015f208054610f9d906123eb565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc9906123eb565b80156110145780601f10610feb57610100808354040283529160200191611014565b820191905f5260205f20905b815481529060010190602001808311610ff757829003601f168201915b505050505090505f611024611472565b90505f81510361103857819250505061107a565b5f8251111561106c578082604051602001611054929190612828565b6040516020818303038152906040529250505061107a565b61107584611488565b925050505b919050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061114957507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806111595750611158826114ee565b5b9050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b80806111d157505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611303575f6111e084610b72565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561124a57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561125d575061125b818461089e565b155b1561129f57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016112969190612128565b60405180910390fd5b811561130157838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f5f611365858585611557565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036113a8576113a384611762565b6113e7565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146113e6576113e581856117a6565b5b5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603611428576114238461187d565b611467565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461146657611465858561193d565b5b5b809150509392505050565b606060405180602001604052805f815250905090565b606061149382610b72565b505f61149d611472565b90505f8151116114bb5760405180602001604052805f8152506114e6565b806114c5846119c1565b6040516020016114d6929190612828565b6040516020818303038152906040525b915050919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f5f61156284611160565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146115a3576115a2818486611a8b565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461162e576115e25f855f5f611199565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16146116ad57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b600980549050600a5f8381526020019081526020015f2081905550600981908060018154018082558091505060019003905f5260205f20015f909190919091505550565b5f6117b08361070b565b90505f60085f8481526020019081526020015f205490505f60075f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905082821461184f575f815f8581526020019081526020015f2054905080825f8581526020019081526020015f20819055508260085f8381526020019081526020015f2081905550505b60085f8581526020019081526020015f205f9055805f8481526020019081526020015f205f90555050505050565b5f60016009805490506118909190612878565b90505f600a5f8481526020019081526020015f205490505f600983815481106118bc576118bb612477565b5b905f5260205f200154905080600983815481106118dc576118db612477565b5b905f5260205f20018190555081600a5f8381526020019081526020015f2081905550600a5f8581526020019081526020015f205f90556009805480611924576119236128ab565b5b600190038181905f5260205f20015f9055905550505050565b5f60016119498461070b565b6119539190612878565b90508160075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055508060085f8481526020019081526020015f2081905550505050565b60605f60016119cf84611b4e565b0190505f8167ffffffffffffffff8111156119ed576119ec611de2565b5b6040519080825280601f01601f191660200182016040528015611a1f5781602001600182028036833780820191505090505b5090505f82602083010190505b600115611a80578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8581611a7557611a746128d8565b5b0494505f8503611a2c575b819350505050919050565b611a96838383611c9f565b611b49575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611b0a57806040517f7e273289000000000000000000000000000000000000000000000000000000008152600401611b019190611f78565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611b40929190612450565b60405180910390fd5b505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611baa577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381611ba057611b9f6128d8565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310611be7576d04ee2d6d415b85acef81000000008381611bdd57611bdc6128d8565b5b0492506020810190505b662386f26fc100008310611c1657662386f26fc100008381611c0c57611c0b6128d8565b5b0492506010810190505b6305f5e1008310611c3f576305f5e1008381611c3557611c346128d8565b5b0492506008810190505b6127108310611c64576127108381611c5a57611c596128d8565b5b0492506004810190505b60648310611c875760648381611c7d57611c7c6128d8565b5b0492506002810190505b600a8310611c96576001810190505b80915050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015611d5657508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611d175750611d16848461089e565b5b80611d5557508273ffffffffffffffffffffffffffffffffffffffff16611d3d83610bf8565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611d9982611d70565b9050919050565b611da981611d8f565b8114611db3575f5ffd5b50565b5f81359050611dc481611da0565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611e1882611dd2565b810181811067ffffffffffffffff82111715611e3757611e36611de2565b5b80604052505050565b5f611e49611d5f565b9050611e558282611e0f565b919050565b5f67ffffffffffffffff821115611e7457611e73611de2565b5b611e7d82611dd2565b9050602081019050919050565b828183375f83830152505050565b5f611eaa611ea584611e5a565b611e40565b905082815260208101848484011115611ec657611ec5611dce565b5b611ed1848285611e8a565b509392505050565b5f82601f830112611eed57611eec611dca565b5b8135611efd848260208601611e98565b91505092915050565b5f5f60408385031215611f1c57611f1b611d68565b5b5f611f2985828601611db6565b925050602083013567ffffffffffffffff811115611f4a57611f49611d6c565b5b611f5685828601611ed9565b9150509250929050565b5f819050919050565b611f7281611f60565b82525050565b5f602082019050611f8b5f830184611f69565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611fc581611f91565b8114611fcf575f5ffd5b50565b5f81359050611fe081611fbc565b92915050565b5f60208284031215611ffb57611ffa611d68565b5b5f61200884828501611fd2565b91505092915050565b5f8115159050919050565b61202581612011565b82525050565b5f60208201905061203e5f83018461201c565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61207682612044565b612080818561204e565b935061209081856020860161205e565b61209981611dd2565b840191505092915050565b5f6020820190508181035f8301526120bc818461206c565b905092915050565b6120cd81611f60565b81146120d7575f5ffd5b50565b5f813590506120e8816120c4565b92915050565b5f6020828403121561210357612102611d68565b5b5f612110848285016120da565b91505092915050565b61212281611d8f565b82525050565b5f60208201905061213b5f830184612119565b92915050565b5f5f6040838503121561215757612156611d68565b5b5f61216485828601611db6565b9250506020612175858286016120da565b9150509250929050565b5f5f5f6060848603121561219657612195611d68565b5b5f6121a386828701611db6565b93505060206121b486828701611db6565b92505060406121c5868287016120da565b9150509250925092565b5f602082840312156121e4576121e3611d68565b5b5f6121f184828501611db6565b91505092915050565b61220381612011565b811461220d575f5ffd5b50565b5f8135905061221e816121fa565b92915050565b5f5f6040838503121561223a57612239611d68565b5b5f61224785828601611db6565b925050602061225885828601612210565b9150509250929050565b5f67ffffffffffffffff82111561227c5761227b611de2565b5b61228582611dd2565b9050602081019050919050565b5f6122a461229f84612262565b611e40565b9050828152602081018484840111156122c0576122bf611dce565b5b6122cb848285611e8a565b509392505050565b5f82601f8301126122e7576122e6611dca565b5b81356122f7848260208601612292565b91505092915050565b5f5f5f5f6080858703121561231857612317611d68565b5b5f61232587828801611db6565b945050602061233687828801611db6565b9350506040612347878288016120da565b925050606085013567ffffffffffffffff81111561236857612367611d6c565b5b612374878288016122d3565b91505092959194509250565b5f5f6040838503121561239657612395611d68565b5b5f6123a385828601611db6565b92505060206123b485828601611db6565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061240257607f821691505b602082108103612415576124146123be565b5b50919050565b5f60608201905061242e5f830186612119565b61243b6020830185611f69565b6124486040830184612119565b949350505050565b5f6040820190506124635f830185612119565b6124706020830184611f69565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026125007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826124c5565b61250a86836124c5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61254561254061253b84611f60565b612522565b611f60565b9050919050565b5f819050919050565b61255e8361252b565b61257261256a8261254c565b8484546124d1565b825550505050565b5f5f905090565b61258961257a565b612594818484612555565b505050565b5b818110156125b7576125ac5f82612581565b60018101905061259a565b5050565b601f8211156125fc576125cd816124a4565b6125d6846124b6565b810160208510156125e5578190505b6125f96125f1856124b6565b830182612599565b50505b505050565b5f82821c905092915050565b5f61261c5f1984600802612601565b1980831691505092915050565b5f612634838361260d565b9150826002028217905092915050565b61264d82612044565b67ffffffffffffffff81111561266657612665611de2565b5b61267082546123eb565b61267b8282856125bb565b5f60209050601f8311600181146126ac575f841561269a578287015190505b6126a48582612629565b86555061270b565b601f1984166126ba866124a4565b5f5b828110156126e1578489015182556001820191506020850194506020810190506126bc565b868310156126fe57848901516126fa601f89168261260d565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f61273782612713565b612741818561271d565b935061275181856020860161205e565b61275a81611dd2565b840191505092915050565b5f6080820190506127785f830187612119565b6127856020830186612119565b6127926040830185611f69565b81810360608301526127a4818461272d565b905095945050505050565b5f815190506127bd81611fbc565b92915050565b5f602082840312156127d8576127d7611d68565b5b5f6127e5848285016127af565b91505092915050565b5f81905092915050565b5f61280282612044565b61280c81856127ee565b935061281c81856020860161205e565b80840191505092915050565b5f61283382856127f8565b915061283f82846127f8565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61288282611f60565b915061288d83611f60565b92508282039050818111156128a5576128a461284b565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220f4b2c29063c6573019842eb1ba45a0a2da913ae262673f1a80b0408adbf83b5a64736f6c634300081e0033",
}

// WnftABI is the input ABI used to generate the binding from.
// Deprecated: Use WnftMetaData.ABI instead.
var WnftABI = WnftMetaData.ABI

// WnftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WnftMetaData.Bin instead.
var WnftBin = WnftMetaData.Bin

// DeployWnft deploys a new Ethereum contract, binding an instance of Wnft to it.
func DeployWnft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wnft, error) {
	parsed, err := WnftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WnftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wnft{WnftCaller: WnftCaller{contract: contract}, WnftTransactor: WnftTransactor{contract: contract}, WnftFilterer: WnftFilterer{contract: contract}}, nil
}

// Wnft is an auto generated Go binding around an Ethereum contract.
type Wnft struct {
	WnftCaller     // Read-only binding to the contract
	WnftTransactor // Write-only binding to the contract
	WnftFilterer   // Log filterer for contract events
}

// WnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type WnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WnftSession struct {
	Contract     *Wnft             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WnftCallerSession struct {
	Contract *WnftCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WnftTransactorSession struct {
	Contract     *WnftTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type WnftRaw struct {
	Contract *Wnft // Generic contract binding to access the raw methods on
}

// WnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WnftCallerRaw struct {
	Contract *WnftCaller // Generic read-only contract binding to access the raw methods on
}

// WnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WnftTransactorRaw struct {
	Contract *WnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWnft creates a new instance of Wnft, bound to a specific deployed contract.
func NewWnft(address common.Address, backend bind.ContractBackend) (*Wnft, error) {
	contract, err := bindWnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wnft{WnftCaller: WnftCaller{contract: contract}, WnftTransactor: WnftTransactor{contract: contract}, WnftFilterer: WnftFilterer{contract: contract}}, nil
}

// NewWnftCaller creates a new read-only instance of Wnft, bound to a specific deployed contract.
func NewWnftCaller(address common.Address, caller bind.ContractCaller) (*WnftCaller, error) {
	contract, err := bindWnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WnftCaller{contract: contract}, nil
}

// NewWnftTransactor creates a new write-only instance of Wnft, bound to a specific deployed contract.
func NewWnftTransactor(address common.Address, transactor bind.ContractTransactor) (*WnftTransactor, error) {
	contract, err := bindWnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WnftTransactor{contract: contract}, nil
}

// NewWnftFilterer creates a new log filterer instance of Wnft, bound to a specific deployed contract.
func NewWnftFilterer(address common.Address, filterer bind.ContractFilterer) (*WnftFilterer, error) {
	contract, err := bindWnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WnftFilterer{contract: contract}, nil
}

// bindWnft binds a generic wrapper to an already deployed contract.
func bindWnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WnftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wnft *WnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wnft.Contract.WnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wnft *WnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wnft.Contract.WnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wnft *WnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wnft.Contract.WnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wnft *WnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wnft *WnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wnft *WnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wnft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Wnft *WnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Wnft *WnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Wnft.Contract.BalanceOf(&_Wnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Wnft *WnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Wnft.Contract.BalanceOf(&_Wnft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Wnft *WnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Wnft *WnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Wnft.Contract.GetApproved(&_Wnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Wnft *WnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Wnft.Contract.GetApproved(&_Wnft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Wnft *WnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Wnft *WnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Wnft.Contract.IsApprovedForAll(&_Wnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Wnft *WnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Wnft.Contract.IsApprovedForAll(&_Wnft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wnft *WnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wnft *WnftSession) Name() (string, error) {
	return _Wnft.Contract.Name(&_Wnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wnft *WnftCallerSession) Name() (string, error) {
	return _Wnft.Contract.Name(&_Wnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Wnft *WnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Wnft *WnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Wnft.Contract.OwnerOf(&_Wnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Wnft *WnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Wnft.Contract.OwnerOf(&_Wnft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Wnft *WnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Wnft *WnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Wnft.Contract.SupportsInterface(&_Wnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Wnft *WnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Wnft.Contract.SupportsInterface(&_Wnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wnft *WnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wnft *WnftSession) Symbol() (string, error) {
	return _Wnft.Contract.Symbol(&_Wnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wnft *WnftCallerSession) Symbol() (string, error) {
	return _Wnft.Contract.Symbol(&_Wnft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Wnft *WnftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Wnft *WnftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Wnft.Contract.TokenByIndex(&_Wnft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Wnft *WnftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Wnft.Contract.TokenByIndex(&_Wnft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Wnft *WnftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Wnft *WnftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Wnft.Contract.TokenOfOwnerByIndex(&_Wnft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Wnft *WnftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Wnft.Contract.TokenOfOwnerByIndex(&_Wnft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Wnft *WnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Wnft *WnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Wnft.Contract.TokenURI(&_Wnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Wnft *WnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Wnft.Contract.TokenURI(&_Wnft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wnft *WnftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wnft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wnft *WnftSession) TotalSupply() (*big.Int, error) {
	return _Wnft.Contract.TotalSupply(&_Wnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wnft *WnftCallerSession) TotalSupply() (*big.Int, error) {
	return _Wnft.Contract.TotalSupply(&_Wnft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Wnft *WnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Wnft *WnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.Contract.Approve(&_Wnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Wnft *WnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.Contract.Approve(&_Wnft.TransactOpts, to, tokenId)
}

// CreateToken is a paid mutator transaction binding the contract method 0x01382807.
//
// Solidity: function createToken(address owner, string _tokenURI) returns(uint256)
func (_Wnft *WnftTransactor) CreateToken(opts *bind.TransactOpts, owner common.Address, _tokenURI string) (*types.Transaction, error) {
	return _Wnft.contract.Transact(opts, "createToken", owner, _tokenURI)
}

// CreateToken is a paid mutator transaction binding the contract method 0x01382807.
//
// Solidity: function createToken(address owner, string _tokenURI) returns(uint256)
func (_Wnft *WnftSession) CreateToken(owner common.Address, _tokenURI string) (*types.Transaction, error) {
	return _Wnft.Contract.CreateToken(&_Wnft.TransactOpts, owner, _tokenURI)
}

// CreateToken is a paid mutator transaction binding the contract method 0x01382807.
//
// Solidity: function createToken(address owner, string _tokenURI) returns(uint256)
func (_Wnft *WnftTransactorSession) CreateToken(owner common.Address, _tokenURI string) (*types.Transaction, error) {
	return _Wnft.Contract.CreateToken(&_Wnft.TransactOpts, owner, _tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Wnft *WnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Wnft *WnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.Contract.SafeTransferFrom(&_Wnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Wnft *WnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.Contract.SafeTransferFrom(&_Wnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Wnft *WnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Wnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Wnft *WnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Wnft.Contract.SafeTransferFrom0(&_Wnft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Wnft *WnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Wnft.Contract.SafeTransferFrom0(&_Wnft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Wnft *WnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Wnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Wnft *WnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Wnft.Contract.SetApprovalForAll(&_Wnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Wnft *WnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Wnft.Contract.SetApprovalForAll(&_Wnft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Wnft *WnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Wnft *WnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.Contract.TransferFrom(&_Wnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Wnft *WnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Wnft.Contract.TransferFrom(&_Wnft.TransactOpts, from, to, tokenId)
}

// WnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wnft contract.
type WnftApprovalIterator struct {
	Event *WnftApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WnftApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WnftApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WnftApproval represents a Approval event raised by the Wnft contract.
type WnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Wnft *WnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*WnftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Wnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WnftApprovalIterator{contract: _Wnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Wnft *WnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Wnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WnftApproval)
				if err := _Wnft.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Wnft *WnftFilterer) ParseApproval(log types.Log) (*WnftApproval, error) {
	event := new(WnftApproval)
	if err := _Wnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Wnft contract.
type WnftApprovalForAllIterator struct {
	Event *WnftApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WnftApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WnftApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WnftApprovalForAll represents a ApprovalForAll event raised by the Wnft contract.
type WnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Wnft *WnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*WnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Wnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &WnftApprovalForAllIterator{contract: _Wnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Wnft *WnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *WnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Wnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WnftApprovalForAll)
				if err := _Wnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Wnft *WnftFilterer) ParseApprovalForAll(log types.Log) (*WnftApprovalForAll, error) {
	event := new(WnftApprovalForAll)
	if err := _Wnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WnftBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Wnft contract.
type WnftBatchMetadataUpdateIterator struct {
	Event *WnftBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WnftBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WnftBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WnftBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WnftBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WnftBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WnftBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Wnft contract.
type WnftBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Wnft *WnftFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*WnftBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Wnft.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &WnftBatchMetadataUpdateIterator{contract: _Wnft.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Wnft *WnftFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *WnftBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Wnft.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WnftBatchMetadataUpdate)
				if err := _Wnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Wnft *WnftFilterer) ParseBatchMetadataUpdate(log types.Log) (*WnftBatchMetadataUpdate, error) {
	event := new(WnftBatchMetadataUpdate)
	if err := _Wnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WnftMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Wnft contract.
type WnftMetadataUpdateIterator struct {
	Event *WnftMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WnftMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WnftMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WnftMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WnftMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WnftMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WnftMetadataUpdate represents a MetadataUpdate event raised by the Wnft contract.
type WnftMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Wnft *WnftFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*WnftMetadataUpdateIterator, error) {

	logs, sub, err := _Wnft.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &WnftMetadataUpdateIterator{contract: _Wnft.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Wnft *WnftFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *WnftMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Wnft.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WnftMetadataUpdate)
				if err := _Wnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Wnft *WnftFilterer) ParseMetadataUpdate(log types.Log) (*WnftMetadataUpdate, error) {
	event := new(WnftMetadataUpdate)
	if err := _Wnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wnft contract.
type WnftTransferIterator struct {
	Event *WnftTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WnftTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WnftTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WnftTransfer represents a Transfer event raised by the Wnft contract.
type WnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Wnft *WnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*WnftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Wnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WnftTransferIterator{contract: _Wnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Wnft *WnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Wnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WnftTransfer)
				if err := _Wnft.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Wnft *WnftFilterer) ParseTransfer(log types.Log) (*WnftTransfer, error) {
	event := new(WnftTransfer)
	if err := _Wnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
