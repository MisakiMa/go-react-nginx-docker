import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';

type User = {
  Name: string;
  Id: number;
  Password: string;
}

function App() {
  const [users, setUsers] = useState<User[]>([])
  const [name, setName] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [id, setId] = useState<number>(2)

  useEffect(() => {
    axios.get('http://localhost:5000/api/users').then(res => {
      setUsers(res.data)
      let users: User[] = res.data
      let user = users[users.length - 1]
      setId(user.Id)
    })
  }, [])

  const handleSubmit = async () => {
    await axios.post('http://localhost:5000/api/users/signup', {name: name, id: id + 1, password: password})
    await axios.get('http://localhost:5000/api/users').then(res => {
      setUsers(res.data)
    })
    setName('')
    setPassword('')
    setId((i) => i + 1)
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
          return <p key={user.Id}>id: {user.Id} name: {user.Name}</p>
        })}
        <input type="text" value={name} onChange={(event) => setName(event.target.value)}></input>
        <input type="password" value={password} onChange={(event) => setPassword(event.target.value)}></input>
        <button onClick={handleSubmit}>Submit</button>
      </header>
    </div>
  );
}

export default App;
