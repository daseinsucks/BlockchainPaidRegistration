import React, { useState } from "react";
import { ethers } from "ethers";

const contractAddress = "0x59721ADab008a1354630DBDd99006F9bc6a71312";
const contractABI = [
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "key",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "tgId",
        "type": "uint256"
      }
    ],
    "name": "registration",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  }
];

const App: React.FC = () => {
  const [key, setKey] = useState<string>("");
  const [tgId, setTgId] = useState<number>(0);
  const [status, setStatus] = useState<string>("");
  const [currentAccount, setCurrentAccount] = useState<string | undefined>()


  
  const handleRegister = async () => {
    if (!window.ethereum) {
      setStatus("MetaMask is not installed");
      return;
    }
  
    try {
      // Request user's wallet connection
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      
      // Request account access from MetaMask
      const accounts = await provider.send("eth_requestAccounts", []);
      if (accounts.length === 0) {
        setStatus("No accounts found.");
        return;
      }
  
      // Set the current account
      const account = accounts[0];
      setCurrentAccount(account);
  
      // Check the balance of the user's account
      const balance = await provider.getBalance(account);
      const balanceInEther = ethers.utils.formatEther(balance);
  
      console.log("Wallet balance:", balanceInEther);
  
      // Get fee from the contract (hardcoded as 0.001 ETH)
      const fee = ethers.utils.parseEther("0.001");
  
      // Check if the user has enough balance to cover the fee + gas
      if (balance.lt(fee)) {
        setStatus(`Insufficient funds. Your balance: ${balanceInEther} ETH, required: 0.001 ETH + gas fees.`);
        return;
      }
  
      const signer = provider.getSigner();
      const contract = new ethers.Contract(contractAddress, contractABI, signer);
  
      // Send the registration transaction
      const tx = await contract.registration(key, tgId, { value: fee });
      await tx.wait();
  
      setStatus("Registration successful!");
    } catch (error) {
      console.error("Transaction error:", error);
  
      // Type guard to ensure error is an object with a code property
      if (typeof error === 'object' && error !== null && 'code' in error) {
        const typedError = error as { code: string };  // Safely typecast error to have a code
  
        if (typedError.code === '-32000') {
          setStatus("Error: Insufficient funds to cover gas fees.");
        } else if (typedError.code === "UNPREDICTABLE_GAS_LIMIT") {
          setStatus("Transaction failed: Unpredictable gas limit. Try again later.");
        } else {
          setStatus("An error occurred during the registration process.");
        }
      } else {
        setStatus("An unexpected error occurred.");
      }
    }
  };
  

  return (
    <div style={{ textAlign: "center", padding: "50px" }}>
      <h2>Register for LAI</h2>
      
      <input
        type="text"
        placeholder="API Key"
        value={key}
        onChange={(e) => setKey(e.target.value)}
        style={{ padding: "10px", marginBottom: "10px", width: "300px" }}
      />
      <br />
      <h3>Telegram ID (optional)</h3>
      <input
        type="text"
        placeholder="Telegram ID (Optional)"
        value={tgId}
        onChange={(e) => setTgId(Number(e.target.value))}
        style={{ padding: "10px", marginBottom: "10px", width: "300px" }}
      />
      <br />
      <button onClick={handleRegister} style={{ padding: "10px 20px" }}>
        Register
      </button>
      <p>{status}</p>
    </div>
  );
};

export default App;