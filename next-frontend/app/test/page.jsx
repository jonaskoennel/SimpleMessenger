import { FormEvent } from 'react'
import LoginForm from '../api/login/login'
 
export default function TestPage() {
    
    return (
      <div className="">
        <LoginForm />
      </div>
    )
}

function ChatRow({chat}) {

    return (
        <tr>
            <div className="item">

            </div>
        </tr>
    )
}

async function fetchChatData() {
    const response = await fetch('http://localhost:8090/chats/user', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        body: "",
    })
      
    var data = await response.json();
    return data
}