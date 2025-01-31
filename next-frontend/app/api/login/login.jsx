'use client'
import { redirect } from 'next/dist/server/api-utils'
import { FormEvent } from 'react'
import { useRouter } from 'next/navigation'; 
import Link from 'next/link'
 
export default function LoginForm() {
  return (
    <div className="flex flex-col gap-2 w-80 place-self-center">
      <div className="place-self-center place-content-center rounded-xl bg-slate-700 w-80 h-80 p-10 shadow-lg">
        <form className="flex flex-col gap-8 w-7/8 justify-center" onSubmit={handleSubmit}>
          <input className="rounded-xl bg-slate-400 p-2 placeholder:text-white shadow-lg" type="email" name="email" placeholder="Email" required />
          <input className="rounded-xl bg-slate-400 p-2 placeholder:text-white shadow-lg" type="password" name="password" placeholder="Password" required />
          <div className="justify-self-center justify-center bg-blue-700 rounded-xl w-1/2 shadow-lg">
            <button className="size-full" type="submit">Login</button>
          </div>
        </form>
      </div>
      <div className="text-center text-slate-700">
        Noch nicht angemeldet? <Link href="/register"><u>Registrieren</u></Link>
      </div>
    </div>
  )
}

async function handleSubmit(event) {
    //const router = useRouter();
    event.preventDefault()
 
    const formData = new FormData(event.target)
    const email = formData.get('email')
    const password = formData.get('password')
 
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    })
    
    var data = await response.json();
    console.log(data)

    if (response.ok) {+
      console.log("Erfolgreich!")
      //router.push('/');
    } else {
      // Handle errors
    }
}