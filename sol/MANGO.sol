// SPDX-License-Identifier: MIT
//
// ███╗░░░███╗░█████╗░███╗░░██╗░██████╗░░█████╗░
// ████╗░████║██╔══██╗████╗░██║██╔════╝░██╔══██╗
// ██╔████╔██║███████║██╔██╗██║██║░░██╗░██║░░██║
// ██║╚██╔╝██║██╔══██║██║╚████║██║░░╚██╗██║░░██║
// ██║░╚═╝░██║██║░░██║██║░╚███║╚██████╔╝╚█████╔╝
// ╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝░╚═════╝░░╚════╝░
//#magnoMondays 
                                                
                                           
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

 contract MANGO is ERC721Enumerable, Ownable {
    using SafeMath for uint256;
    using Address for address;

    string public PROVENANCE;
    string private baseURI;

    uint256 public maxSupply;
    uint256 public price = 0.1 ether;

    bool public privateSaleActive = false;
    bool public saleActive = false;

    mapping (address => uint256) public whitelist;

    constructor(string memory name, string memory symbol, uint256 supply) ERC721(name, symbol) {
        maxSupply = supply;
    }

    function privateMint(uint256 numberOfMints) public payable {
        uint256 supply = totalSupply();
        uint256 reserved = whitelist[msg.sender];
        require(privateSaleActive, "Sale must be active");
        require(reserved > 0, "No tokens reserved for this address");
        require(numberOfMints <= reserved, "Can't mint more than reserved");
        require(supply.add(numberOfMints) <= maxSupply, "Purchase would exceed max supply of tokens");
        require(price.mul(numberOfMints) == msg.value, "Ether value sent is not correct");
        whitelist[msg.sender] = reserved - numberOfMints;
        for(uint256 i; i < numberOfMints; i++){
            _safeMint( msg.sender, supply + i );
        }
    }

    function mint(uint256 numberOfMints) public payable {
        uint256 supply = totalSupply();
        require(saleActive, "Sale must be active to mint");
        require(numberOfMints > 0 && numberOfMints <= 5, "Invalid purchase amount");
        require(supply.add(numberOfMints) <= maxSupply, "Purchase would exceed max supply of tokens");
        require(price.mul(numberOfMints) == msg.value, "Ether value sent is not correct");
        for(uint256 i; i < numberOfMints; i++) {
            _safeMint(msg.sender, supply + i);
        }
    }

    function editPresale(address[] calldata presaleAddresses, uint256[] calldata amount) public onlyOwner {
        for(uint256 i; i < presaleAddresses.length; i++){
            whitelist[presaleAddresses[i]] = amount[i];
        }
    }

    function walletOfOwner(address owner) external view returns(uint256[] memory) {
        uint256 tokenCount = balanceOf(owner);

        uint256[] memory tokensId = new uint256[](tokenCount);
        for(uint256 i; i < tokenCount; i++){
            tokensId[i] = tokenOfOwnerByIndex(owner, i);
        }
        return tokensId;
    }

    function withdraw() public onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    function togglePrivateSale() public onlyOwner {
        privateSaleActive = !privateSaleActive;
    }

    function toggleSale() public onlyOwner {
        saleActive = !saleActive;
    }

    function setPrice(uint256 newPrice) public onlyOwner {
        price = newPrice;
    }

    function setProvenance(string memory provenance) public onlyOwner {
        PROVENANCE = provenance;
    }
    
    function setBaseURI(string memory uri) public onlyOwner {
        baseURI = uri;
    }
    
    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }
}