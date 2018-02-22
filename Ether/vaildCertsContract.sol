pragma solidity ^0.4.18;
import "github.com/Arachnid/solidity-stringutils/strings.sol";

contract VaildCerts {
    
    using strings for *;
  
    struct certData {
        address creator;
        string hash;
    }

    mapping (uint32 => mapping(uint16 => certData)) map;
    mapping (string => bool) editors;
    address creator;
    
    function VaildCerts() public {
        creator = msg.sender;
    }    

    // https://ethereum.stackexchange.com/questions/8346/convert-address-to-string
    function toAsciiString(address x) returns (string) {
        bytes memory s = new bytes(40);
        for (uint i = 0; i < 20; i++) {
            byte b = byte(uint8(uint(x) / (2**(8*(19 - i)))));
            byte hi = byte(uint8(b) / 16);
            byte lo = byte(uint8(b) - 16 * uint8(hi));
            s[2*i] = char(hi);
            s[2*i+1] = char(lo);            
        }
        return string(s);
    }

    function char(byte b) returns (byte c) {
        if (b < 10) return byte(uint8(b) + 0x30);
        else return byte(uint8(b) + 0x57);
    }

    
    modifier checkCreator() {
        require(msg.sender == creator);
        _;                      
    }
    
    
    function isEditor() public returns(bool result){
        return editors[toAsciiString(msg.sender)];
    }

    function isEditor(string toCheck) public constant returns(bool result){
        return editors[toCheck];
    }
    
    modifier checkIsEditor() {
        require(isEditor() || msg.sender == creator);
        _;                      
    }
    
    function addEditor(string toAdd) checkCreator public {
        var s = toAdd.toSlice();
        var foo = s.split("0x".toSlice());

        editors[s.toString()] = true;
    }
    
    function removeEditor(string toRemove) checkCreator public {
        delete editors[toRemove];
    }

    function add(uint32 studentID, uint16 documentId, string hash) checkIsEditor public {
        map[studentID][documentId] = certData(msg.sender, hash);
    }
    
    function revoke(uint32 studentID, uint16 documentId) checkIsEditor public {
       map[studentID][documentId] = certData(msg.sender, "REVOKED");
    }

    function getHash(uint32 studentID, uint16 documentId) public constant returns (string result) {
        return map[studentID][documentId].hash;
    }
    
}
 