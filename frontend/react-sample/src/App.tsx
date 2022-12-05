import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';
import { Link, Route, Routes } from 'react-router-dom';
import logo from './logo.svg';
import LoginPage from './pages/login';

type User = {
  userName: string;
  userId: string;
  id: string;
  password: string;
}

function UsersList() {

  const [users, setUsers] = useState<User[]>([])
  const [userName, setName] = useState<string>('')
  const [userId, setUserId] = useState<string>('')
  const [password, setPassword] = useState<string>('')

  useEffect(() =>  {
    const getUsers = async () => {
      const res = await axios.get<User[]>('http://localhost:8000/api/users')
      setUsers(res.data);
    }
    void getUsers()
  }, [])

  const resetForm = () => {
    setName('')
    setUserId('')
    setPassword('')
  }

  const handleSubmit = async () => {
    const requestdata = {
      "userName" : userName,
      "userId" : userId,
      password
    };
    await axios.post('http://localhost:8000/api/users/signup', requestdata)
    await axios.get<User[]>('http://localhost:8000/api/users').then(res => {
      setUsers(res.data)
    })
    resetForm();
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
        {users.map((user) => <p key={user.id}>id: {user.id} name: {user.userName} userId: {user.userId}</p>)}
        UserID <input type="text" value={userId} onChange={(event) => setUserId(event.target.value)} />
        UserName <input type="text" value={userName} onChange={(event) => setName(event.target.value)} />
        Password <input type="password" value={password} onChange={(event) => setPassword(event.target.value)} />
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
