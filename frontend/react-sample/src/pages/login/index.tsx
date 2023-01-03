import axios from "axios";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import toast, { Toaster } from 'react-hot-toast';
import { User } from "../../type/user";

function LoginPage() {
  const navigate = useNavigate();
	const [password, setPassword] = useState<string>('')
  const [id, setId] = useState<number>(0)
  const [isLogin, setIsLogin] = useState<boolean>(false)
  const signinClick = async () => {
    try {
      const result = await axios.post<User>('http://localhost:8000/api/users/signin', {id, password})
      const user = result.data
      setIsLogin(true)
    } catch {
      toast('errorが起きました!!')
    }
  }
	return (
    <>
      <nav>
        <Link to="/">Home</Link>
      </nav>
      <main>
        <Toaster />
        {isLogin && <h2>ログインしました!!!!!!</h2>}
        id<input type="number" value={id} onChange={(event) => setId(event.target.valueAsNumber)} />
        password<input type="password" value={password} onChange={(event) => setPassword(event.target.value)} />
        <div>
          <button type="button" onClick={signinClick}>signin</button>
        </div>
        <button type="button" onClick={() => {
          navigate("/")
        }}>ボタン Homeへ</button>
      </main>
    </>
	);
}

export default LoginPage;