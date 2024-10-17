import React, { useState, useEffect } from "react";
import { ethers } from "ethers";
import contractABI from './artifacts/LAIRegistrationABI.json';
const contractAddress = process.env.REACT_APP_CONTRACT_ADDRESS;

if (!contractAddress) {
    throw new Error("Contract address not found in environment variables");
}

const App: React.FC = () => {
  const [key, setKey] = useState<string>("");
  const [tgId, setTgId] = useState<number>(0);
  const [status, setStatus] = useState<string>("");
  const [currentAccount, setCurrentAccount] = useState<string | undefined>();

  useEffect(() => {
    const searchParams = new URLSearchParams(window.location.search);
    const tgIdFromURL = searchParams.get("tgid");

    // Set tgId from URL if present, otherwise set to 0
    if (tgIdFromURL) {
      setTgId(Number(tgIdFromURL));
    }
  }, []);

  const generateApiKey = (): string => {
    const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
    let randomLetters = "";
    for (let i = 0; i < 5; i++) {
      randomLetters += letters.charAt(Math.floor(Math.random() * letters.length));
    }
    const timestamp = Date.now(); 
    return randomLetters + timestamp;
  };

  // Use this as reference to decrypt
  const decryptApiKey = async (encryptedApiKey: string): Promise<string> => {
    const privateKeyBase64 = process.env.REACT_APP_PRIVATE_KEY;
    if (!privateKeyBase64) {
      throw new Error("Private key not found in .env file");
    }
  
    // Decode the base64 private key
    const privateKeyBuffer = Uint8Array.from(atob(privateKeyBase64), c => c.charCodeAt(0));
  
    // Import the private key for decryption
    const privateKey = await window.crypto.subtle.importKey(
      "pkcs8", // PKCS8 format
      privateKeyBuffer.buffer, // DER-encoded private key
      {
        name: "RSA-OAEP",
        hash: "SHA-256",
      },
      true,
      ["decrypt"]
    );
  
    // Decode the base64 encrypted API key
    const encryptedKeyBuffer = Uint8Array.from(atob(encryptedApiKey), c => c.charCodeAt(0));
  
    // Decrypt the API key
    const decryptedKey = await window.crypto.subtle.decrypt(
      {
        name: "RSA-OAEP",
      },
      privateKey,
      encryptedKeyBuffer.buffer // Encrypted byte array
    );
  
    // Convert decrypted byte array to string
    return new TextDecoder().decode(decryptedKey);
  };

  const encryptApiKey = async (apiKey: string): Promise<string> => {
    const publicKeyBase64 = process.env.REACT_APP_PUBLIC_KEY;
    if (!publicKeyBase64) {
      throw new Error("Public key not found in .env file");
    }

    // Decode the base64 public key
    const publicKeyBuffer = Uint8Array.from(atob(publicKeyBase64), c => c.charCodeAt(0));

    // Import the public key for encryption
    const publicKey = await window.crypto.subtle.importKey(
      "spki", 
      publicKeyBuffer.buffer, 
      {
        name: "RSA-OAEP",
        hash: "SHA-256",
      }, 
      true, 
      ["encrypt"]
    );

    // Encrypt the API key
    const encryptedKey = await window.crypto.subtle.encrypt(
      {
        name: "RSA-OAEP",
      },
      publicKey,
      new TextEncoder().encode(apiKey)
    );

    // Convert encrypted key to base64 for easy transmission
    return btoa(String.fromCharCode(...new Uint8Array(encryptedKey)));
  };

  const handleRegister = async () => {
    if (!window.ethereum) {
      setStatus("MetaMask is not installed");
      return;
    }

    try {
      // Generate and encrypt the API key
      const apiKey = generateApiKey();
      const encryptedKey = await encryptApiKey(apiKey);
      console.log(tgId);
      setKey(encryptedKey); // Optionally set it to state for reference

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

      // Send the registration transaction with the encrypted key
      const tx = await contract.registration(encryptedKey, tgId, { value: fee });
      await tx.wait();

      setStatus("Registration successful! API Key: " + apiKey);
    } catch (error) {
      console.error("Transaction error:", error);

      if (typeof error === 'object' && error !== null && 'code' in error) {
        const typedError = error as { code: string };

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
      <br />
      <button onClick={handleRegister} style={{ padding: "10px 20px" }}>
        Register
      </button>
      <p>{status}</p>
    </div>
  );
};

export default App;
