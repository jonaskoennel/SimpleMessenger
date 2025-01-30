import { FormEvent } from 'react'
import LoginForm from '../api/login/login'
 
export default function LoginPage() {
    return (
      <div className="h-screen w-screen place-content-center bg-slate-300">
        <LoginForm />
      </div>
    )
}