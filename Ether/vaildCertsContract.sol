pragma solidity ^0.4.18;
contract VaildCerts {

    mapping (uint32 => mapping(uint16 => string)) map;

    address _creator;

    function VaildCerts() public {
        _creator = msg.sender;
    }
    
    modifier onlyCreator() {
        require(msg.sender == _creator); // If it is incorrect here, it reverts.
        _;                              // Otherwise, it continues.
    } 

    function add(uint32 studentID, uint16 documentId, string value) onlyCreator public {
        map[studentID][documentId] = value;
    }
    
    function revoke(uint32 studentID, uint16 documentId) onlyCreator public {
       map[studentID][documentId] = "REVOKED";
    }

    function getValue(uint32 studentID, uint16 documentId) public constant returns (string result) {
        return map[studentID][documentId];
    }
}
