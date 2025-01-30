'use client'

import { useRouter } from "next/navigation";
/*
export default function Chat({ chat }) {
    var router = useRouter()
    return (
        <button onClick={() => router.push('/login')} key={chat.id} className="flex bg-slate-700 h-20 border border-stone-900 justify-items-start align-items-center content-center p-2 w-full">
            <div className="imageContainer flex place-content-center">
                <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/avatar-8.webp" alt="avatar" className="rounded-full"/>
            </div>
            <div className="flex relative basis-4/5 place-content-center w-auto h-auto">
                <strong>{chat.name}</strong>
            </div>
        </button>
    );
}
*/
export default function Chat({chat, setchat}) {
    //var router = useRouter()
    //console.log(chat)
    return (
        <div key={chat.ID} id={chat.ID} onClick={() => setchat(chat)} className="flex items-center mb-4 cursor-pointer hover:bg-gray-100 p-2 rounded-md">
              <div className="w-12 h-12 bg-gray-300 rounded-full mr-3">
                <img src="https://placehold.co/200x/ffa8e4/ffffff.svg?text=ʕ•́ᴥ•̀ʔ&font=Lato" alt="User Avatar" className="w-12 h-12 rounded-full"/>
              </div>
              <div className="flex-1">
                <h2 className="text-lg font-semibold">{chat.Name}</h2>
                {chat.Messages.length === 0 ? (
                    <p className="text-gray-600">...</p>
                ) : (
                    <p className="text-gray-600">{chat.Messages[0].text}</p>
                )}
              </div>
        </div>
    );
}

async function handleClick(id) {
    //const router = useRouter()
    //router.push('/login')
}