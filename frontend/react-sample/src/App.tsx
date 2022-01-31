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

  useEffect(() => {
    axios.get('http://localhost:5000/api/users').then(res => {
      setUsers(res.data)
    })
  }, [])

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
        <p>
          id: {users[0]?.Id} name: {users[0]?.Name}
        </p>
      </header>
    </div>
  );
}

export default App;
