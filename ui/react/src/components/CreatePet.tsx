import React from "react";
import { useState } from "react";
import { authenticate } from "../apis/usersApi";
import "../custom.css";

export const CreatePet: React.FC = () => {
  const [petName, setPetName] = useState("");
  const [petType, setPetType] = useState("");
  const [missingTime, setMissingTime] = useState("");
  const [showFailedMessage, setFailedMessage] = useState(false);

  const handleCreate = async () => {
    setFailedMessage(false);
    fetch("/api/pets", {
        method: 'POST', 
        cache: 'no-cache', 
        credentials: 'same-origin',
        headers: {
          'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify({ name: petName,petType:petType,missingSince:missingTime })
      }).then(response => {
        if (!response.ok){
          setFailedMessage(true)
        }else{
          console.log(response.json())   
        }
      })
      .then(result => {
        window.location.href = "/"
        console.log("success:",result)
      });
  }

  return (
    <div className="loginContainer">
      <h2 className="loginHeader">Welcome to Pets Alone! Create missing pets</h2>
      <div>
        <input
          className="loginInput"
          type="text"
          title="petname"
          onChange={(e) => setPetName(e.target.value)}
          value={petName}
          placeholder="petname"
        />
      </div>
      <div>
        <input
          className="loginInput"
          type="text"
          title="pettype"
          onChange={(e) => setPetType(e.target.value)}
          value={petType}
          placeholder="petType"
        />
      </div>
      <div>
        <input
          className="loginInput"
          type="text"
          title="missing time"
          onChange={(e) => setMissingTime(e.target.value)}
          value={missingTime}
          placeholder="missing time"
        />
      </div>
      {showFailedMessage && (
        <div>
          <label style={{color: "red"}}>Cannot create pet</label>
        </div>
      )}
      <div>
        <button className="loginButton" onClick={handleCreate}>
          Create
        </button>
      </div>
    </div>
  );
};

export default CreatePet;
