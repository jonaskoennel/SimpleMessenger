'use client';

import Message from "./message_types";
import React, { useState, useEffect } from 'react';

export default function MessageArea({chat, userId, messages, onSendMessage}) {

    const [value, setValue] = useState('');

    function clickEvents() {
        onSendMessage(value)
        setValue('')
    }

    console.log(messages)
      
    return (
        <div className="flex-1">
            <header className="bg-white p-4 text-gray-700">
                <h1 className="text-2xl font-semibold">{chat.Name}</h1>
            </header>
            
            
            <div className="h-screen overflow-y-auto p-4 pb-36 bg-slate-200 content-end">
                {messages.length === 0 ? (
                    <div key="1">
                        <p>Keine Nachrichten Verfügbar</p>
                    </div>
                ) : (
                    messages.map((message) => (
                        <Message message={message} userId={userId}/>
                    ))
                )}
            </div>
            
            
            <footer className="bg-white border-t border-gray-300 p-4 absolute bottom-0 w-3/4">
                <div className="flex items-center">
                    <input value={value} onChange={(e) => {setValue(e.currentTarget.value)}} type="text" placeholder="Type a message..." className="w-full p-2 rounded-md border border-gray-400 focus:outline-none focus:border-blue-500"/>
                    <button className="bg-indigo-500 text-white px-4 py-2 rounded-md ml-2" onClick={() => clickEvents()}>Send</button>
                </div>
            </footer>
        </div>
    )
}