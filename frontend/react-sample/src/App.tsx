import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';

type User = {
  Name: string;
  Id: Number;
}

function App() {
  const [users, setUsers] = useState<User[]>([])
  const [name, setName] = useState<string>('')
  const [id, setId] = useState<number>(2)

  useEffect(() => {
    axios.get('http://localhost:5000/api/users').then(res => {
      setUsers(res.data)
    })
  }, [])

  const handleSubmit = () => {
    setId((i) => i + 1)
    axios.post('http://localhost:5000/api/users', {name: name, id: id})
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        {users.map((user) => {
          return <p>id: {user.Id} name: {user.Name}</p>
        })}
        <input type="text" value={name} onChange={(event) => setName(event.target.value)}></input>
        <button onClick={handleSubmit}>Submit</button>
      </header>
    </div>
  );
}

export default App;
