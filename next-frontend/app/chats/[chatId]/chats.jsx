import React, { useState, useEffect } from 'react';

import Chat from './chatbox';
import Modal from './modal';

export function ChatSidebar({setchat, chats, setChats}) {
    const [loading, setLoading] = useState(true); // Zustand für Ladeanzeige
    const [error, setError] = useState(null); // Zustand für Fehlerbehandlung
    const [dropdownOpen, setDropdownOpen] = useState(false);
    const [chatName, setChatName] = useState('');
    const [isModalOpen, setIsModalOpen] = useState(false);
    
    const openModal = () => {
        setIsModalOpen(true);
        setDropdownOpen(false); // Schließt das Dropdown-Menü
    };

    const closeModal = () => {
        setIsModalOpen(false);
        setChatName('');
        setUserToInvite('');
    };

    useEffect(() => {
        // Asynchrone Funktion innerhalb von useEffect
        const fetchChats = async () => {
          try {
            const response = await fetch('http://localhost:8090/chats/user', {
                method: 'GET',
                credentials: 'include',
                headers: { 
                    'Content-Type': 'application/json'
                    //'Cookie': `${authorization.name}=${authorization.value}`
                 }
              }) // Beispiel-API-Endpoint
            if (!response.ok) {
              throw new Error('Fehler beim Laden der Chats');
            }
            const data = await response.json();
            console.log(data)
            setChats(data.chats); // Chats in den Zustand setzen
          } catch (err) {
            setError(err.message); // Fehler im Zustand speichern
          } finally {
            setLoading(false); // Ladeanzeige beenden
          }
        };
    
        fetchChats(); // Die Funktion aufrufen
      }, []); // leeres Abhängigkeitsarray -> nur einmal beim Initialisieren der Komponente ausführen
      
      useEffect(() => {

      }, [chats])
      // Wenn die Daten noch geladen werden
      if (loading) return <div>Loading...</div>;
    
      // Wenn ein Fehler aufgetreten ist
      if (error) return <div>Fehler: {error}</div>;

    return (
        <div className="w-1/4 bg-white border-r border-gray-300">
            <header className="p-4 border-b border-gray-300 flex justify-between items-center bg-indigo-600 text-white">
                <h1 className="text-2xl font-semibold">Chats</h1>
                    <div className="relative">
                    <button onClick={() => setDropdownOpen((prev) => !prev)} id="menuButton" className="focus:outline-none">
                        <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-gray-100" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                        <path d="M2 10a2 2 0 012-2h12a2 2 0 012 2 2 2 0 01-2 2H4a2 2 0 01-2-2z" />
                        </svg>
                    </button>
                    {dropdownOpen && (
                        <div id="menuDropdown" className="absolute right-0 mt-2 w-48 bg-white border border-gray-300 rounded-md shadow-lg">
                            <ul className="py-2 px-3">
                            <li><button onClick={() => openModal()} className="block px-4 py-2 text-gray-800 hover:text-gray-400">Neuer Chat</button></li>
                            <li><a href="#" className="block px-4 py-2 text-gray-800 hover:text-gray-400">Logout</a></li>
                            </ul>
                        </div>
                    )}
                    </div>
                    {isModalOpen && (
                        <Modal closeModal={closeModal}/>
                    )}
            </header>
            <div className="overflow-y-auto h-screen p-3 mb-9 pb-20">
                {chats.length === 0 ? (
                    <div key="1">
                        <p>Keine Chats verfügbar</p>
                    </div>
                ) : (
                    chats.map((chat) => (
                        <Chat chat={chat} setChats={setChats} setchat={setchat}/>
                    ))
                )}
            </div>
        </div>
    )
}