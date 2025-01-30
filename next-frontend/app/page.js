'use client'

import Image from "next/image";
import React, { useState, useEffect, useCallback } from 'react';
import Navbar from "./header/header";
import dynamic from 'next/dynamic';
import MessageArea from "./message/message_section";
import { ChatSidebar } from "./chats/[chatId]/chats";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter()
  const [selectedChat, setSelectedChat] = useState(null);
  const [messages, setMessages] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [socket, setSocket] = useState(null); // WebSocket Verbindung
  const [unreadChats, setUnreadChats] = useState(new Set()); // Set für nicht gelesene Chats
  const [userId, setUserId] = useState(null); // Zustand für die UserId
  const [chats, setChats] = useState([]);


  const cookies = require('js-cookie')
  const authCookie = cookies.get('Authorization')

  if(authCookie == undefined) {
    //router.push("/")
  }

  const getUserId = async () => {
    try {
      const response = await fetch(`http://localhost:8080/validate`, {
        method: 'GET',
        credentials: 'include',
        headers: { 
            'Content-Type': 'application/json'
            //'Cookie': `${authorization.name}=${authorization.value}`
        }
      });

      if (!response.ok) {
        throw new Error('Fehler beim Abrufen der UserId');
      }

      const data = await response.json();
      //console.log(data.sub)
      setUserId(data.sub); // UserId speichern
    } catch (err) {
      setError('Fehler beim Abrufen der UserId');
      console.error(err);
    }
  }
  

  useEffect(() => {
    getUserId(); // UserId nach dem Laden der Komponente abrufen
  }, []);

  const loadMessages = useCallback(async (chatId) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch(`http://localhost:8090/messages/${chatId}`, {
        method: 'GET',
        credentials: 'include',
        headers: { 
            'Content-Type': 'application/json'
            //'Cookie': `${authorization.name}=${authorization.value}`
        }
      });
      if (!response.ok) {
        throw new Error('Fehler beim Laden der Nachrichten');
      }
      const data = await response.json();
      setMessages(data.messages);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  }, []);

  const handleSelectChat = (chat) => {
    console.log(chat)
    setSelectedChat(chat);
    setUnreadChats((prev) => {
      const newUnreadChats = new Set(prev);
      newUnreadChats.delete(chat.id); // Markiert den Chat als "gelesen"
      return newUnreadChats;
    });
    loadMessages(chat.ID)
  };

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8090/ws');
    setSocket(ws);

    // WebSocket Nachrichten empfangen
    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      const { chatId, senderId, Test } = message;
      console.log(JSON.stringify(message))
      chats.forEach((chat) => {
        if(chatId == chat.ID) {
          if(chat.Messages.length === 0) {
            chat.Messages.push(message)
          } else {
            chat.Messages[0] = message
          }
        }
      })
      // Wenn der Chat ausgewählt ist, füge die Nachricht direkt hinzu
      if (selectedChat && selectedChat.ID === chatId) {
        setMessages((prevMessages) => [...prevMessages, message]);
      } else {
        // Wenn der Chat nicht ausgewählt ist, markiere den Chat als ungelesen
        setUnreadChats((prev) => new Set(prev).add(chatId));
      }
    };

    // WebSocket-Verbindung schließen
    return () => {
      ws.close();
    };
  }, [selectedChat]); // WebSocket wird nur einmal beim Laden initialisiert

  // Nachricht über den WebSocket senden
  const handleSendMessage = (text) => {
    if (!socket || !selectedChat) return;
    const message = {
      chatId: selectedChat.ID,
      senderId: userId,
      text,
    };

    socket.send(JSON.stringify(message));
    // Direkt in der Nachrichtenliste anzeigen
    //setMessages((prevMessages) => [...prevMessages, message]);
  };

  return (
    <div>
      
      <div className="flex h-screen overflow-hidden">
        <ChatSidebar chats={chats} setChats={setChats} setchat={handleSelectChat}/>
        {selectedChat ? (
            <MessageArea chat={selectedChat} userId={userId} messages={messages} onSendMessage={handleSendMessage}/>
          ) : (
            <div key="1">
              <p>Wähle einen Chat aus</p>
            </div>
          )}
      </div>
    </div>
  );
}