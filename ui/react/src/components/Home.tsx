import * as React from 'react';
import { connect } from 'react-redux';
import { getAll } from '../apis/petsApi';
import "../custom.css";

const Home = () => {
  const [pets, setPets] = React.useState<any[]>([]);
  const [searchString, setSearchString] = React.useState("");

  const handleSearch = async () => {
    fetch("/api/pets/"+searchString)
      .then(res => res.json())
      .then(
        (result) => {
          setPets(result);
          console.log(result)
      })
  };

  React.useEffect(() => {
  fetch("/api/pets")
      .then(res => res.json())
      .then(
        (result) => {
          setPets(result);
          console.log(result)
        })
}, [])

return (
<div>
  <div>
  <input
          className="loginInput"
          type="text"
          title="Search"
          onChange={(e) => setSearchString(e.target.value)}
          value={searchString}
          placeholder="pet type"
  />
  </div>
  <div>
    <button className="loginButton" onClick={handleSearch}>
      Search
    </button>
  </div>
  <div>
    <br></br>
  </div>
  <div>
  <table>
    <thead>
      <th>Pet Name </th>
      <th>Pet Type </th>
      <th>Missing Since </th>
    </thead>
    <tbody>
    {pets.map(pet => (
    <tr>
      <td>{pet.name}</td>
      <td>{pet.petType}</td>
      <td>{pet.missingSince}</td>
      </tr>
    ))}
    </tbody>
  </table>
  </div>
</div>)
}

export default connect()(Home);
