import React, { useState, useEffect } from "react";
import { ethers } from "ethers";

const contractAddress = process.env.REACT_APP_CONTRACT_ADDRESS;
const APIAddress = process.env.REACT_APP_API_ADDRESS;
const RegisterEndpoint = process.env.REACT_APP_REGISTER_ENDPOINT;
if (!contractAddress) {
    throw new Error("Contract address not found in environment variables");
}

const App: React.FC = () => {
  const [key, setKey] = useState<string>("");
  const [tgId, setTgId] = useState<string>("");
  const [status, setStatus] = useState<string>("");
  const [currentAccount, setCurrentAccount] = useState<string | undefined>();

  const encryptApiKey = async (apiKey: string): Promise<string> => {
    const publicKeyBase64 = process.env.REACT_APP_PUBLIC_KEY;
    if (!publicKeyBase64) {
      throw new Error("Public key not found in .env file");
    }

    const publicKeyBuffer = Uint8Array.from(atob(publicKeyBase64), (c) => c.charCodeAt(0));

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

    const encryptedKey = await window.crypto.subtle.encrypt(
      {
        name: "RSA-OAEP",
      },
      publicKey,
      new TextEncoder().encode(apiKey)
    );

    return btoa(String.fromCharCode(...new Uint8Array(encryptedKey)));
  };

  const handleRegister = async () => {
    if (!window.ethereum) {
      setStatus("MetaMask is not installed");
      return;
    }

    if (!key) {
      setStatus("Please enter an API key.");
      return;
    }

  const tgIdNumber = Number(tgId)
    if (tgId && isNaN(Number(tgIdNumber))) {
      setStatus("Please enter a valid number for Telegram ID.");
      return;
    }

    try {
      const encryptedKey = await encryptApiKey(key);

      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const accounts = await provider.send("eth_requestAccounts", []);
      if (accounts.length === 0) {
        setStatus("No accounts found.");
        return;
      }
      const account = accounts[0];
      setCurrentAccount(account);

      try {
        const res = await fetch(`${APIAddress}${RegisterEndpoint}`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            wallet: account,
            apikey: key,
            telegram: tgId ? tgId : "", // Optional Telegram ID
          }),
        });
        const data = await res.json();
      } catch (error) {
        console.error("Error:", error);
      }

      setStatus(`Registration successful! API Key: ${key}`);
    } catch (error) {
      console.error("Transaction error:", error);
      setStatus("Error occurred.");
    }
  };


  return (
    <div style={{ textAlign: "center", padding: "50px" }}>
      <h2>Register for LAI</h2>
      <input
        type="text"
        placeholder="Enter your key"
        value={key}
        onChange={(e) => setKey(e.target.value)}
        style={{ padding: "10px", width: "300px", marginBottom: "20px" }}
      />
      <br />
      <input
        type="text"
        placeholder="Optional Telegram ID"
        value={tgId}
        onChange={(e) => setTgId(e.target.value)}
        style={{ padding: "10px", width: "300px", marginBottom: "20px" }}
      />
      <br />
      <button onClick={handleRegister} disabled={!key} style={{ padding: "10px 20px" }}>
        Register
      </button>
      <p>{status}</p>
    </div>
  );
};

export default App;