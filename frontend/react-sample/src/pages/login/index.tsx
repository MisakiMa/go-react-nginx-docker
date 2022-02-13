import axios from "axios";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

function LoginPage() {
  let navigate = useNavigate();
	const [name, setName] = useState<string>('')
	const [password, setPassword] = useState<string>('')
  const [id, setId] = useState<number>(2)
  const handleSubmit = async () => {
    await axios.post('http://localhost:5000/api/users/signup', {name: name, id: id + 1, password: password})
    setName('')
    setPassword('')
    setId((i) => i + 1)
  }
	return (
    <>
      <nav>
        <Link to="/">Home</Link>
      </nav>
      <main>
        <input type="text" value={name} onChange={(event) => setName(event.target.value)}></input>
        <input type="password" value={password} onChange={(event) => setPassword(event.target.value)}></input>
        <div>
          <button onClick={handleSubmit}>Submit</button>
        </div>
        <button onClick={() => {
          navigate("/")
        }}>ボタン Homeへ</button>
      </main>
    </>
	);
}

export default LoginPage;