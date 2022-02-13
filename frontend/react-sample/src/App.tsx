import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import { Link, Route, Routes } from 'react-router-dom';
import LoginPage from './pages/login';

type User = {
  Name: string;
  Id: number;
  Password: string;
}


function UsersList() {

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
      <nav>
        <Link to="login">Login</Link>
      </nav>
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

function App() {
  return (
    <div className="App">
      <h1>React Router!</h1>
      <Routes>
        <Route path="/" element={<UsersList/>}/>
        <Route  path="login" element={<LoginPage/>}/>
        <Route
          path="*"
          element={
            <main style={{ padding: "1rem"}}>
              <p>There's noting here!</p>
            </main>
          }
        />
      </Routes>
    </div>
  )
}

export default App;
