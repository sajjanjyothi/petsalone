import React from "react";
import { useState } from "react";
import { authenticate } from "../apis/usersApi";
import "../custom.css";

export const Login: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [showInvalidUserMessage, setShowInvalidUserMessage] = useState(false);

  const handleLogin = async () => {
    setShowInvalidUserMessage(false);
    const result = await authenticate(username, password);
    if(!result.token) setShowInvalidUserMessage(true);
    else{
      window.location.href = "/";
      console.log(result.token)
    }
  };

  return (
    <div className="loginContainer">
      <h2 className="loginHeader">Welcome to Pets Alone! -- Create missing pets</h2>
      <div>
        <input
          className="loginInput"
          type="text"
          title="username"
          onChange={(e) => setUsername(e.target.value)}
          value={username}
          placeholder="username"
        />
      </div>
      <div>
        <input
          className="loginInput"
          type="password"
          title="username"
          onChange={(e) => setPassword(e.target.value)}
          value={password}
          placeholder="password"
        />
      </div>
      {showInvalidUserMessage && (
        <div>
          <label style={{color: "red"}}>Invalid Username or Password</label>
        </div>
      )}
      <div>
        <button className="loginButton" onClick={handleLogin}>
          Login
        </button>
      </div>
    </div>
  );
};

export default Login;
