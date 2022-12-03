import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';
import { Link, Route, Routes } from 'react-router-dom';
import logo from './logo.svg';
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

  useEffect(() =>  {
    const getUsers = async () => {
      const res = await axios.get<User[]>('http://localhost:8000/api/users')
      setUsers(res.data);
      const usersData = res.data;
      const user = usersData[usersData.length - 1]
      setId(user.Id)
    }
    void getUsers()
  }, [])

  const handleSubmit = async () => {
    await axios.post('http://localhost:8000/api/users/signup', {name, id: id + 1, password})
    await axios.get<User[]>('http://localhost:8000/api/users').then(res => {
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
        {users.map((user) => <p key={user.Id}>id: {user.Id} name: {user.Name}</p>)}
        <input type="text" value={name} onChange={(event) => setName(event.target.value)} />
        <input type="password" value={password} onChange={(event) => setPassword(event.target.value)} />
        <button type="button" onClick={handleSubmit}>Submit</button>
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
              <p>There&apos;s noting here!</p>
            </main>
          }
        />
      </Routes>
    </div>
  )
}

export default App;
