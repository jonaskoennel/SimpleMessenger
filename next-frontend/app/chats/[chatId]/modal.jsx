
import React, { useState, useEffect } from 'react';

export default function Modal({setChats, setError, closeModal, setChatName, chatName, userToInvite, setUserToInvite}) {
    const [loading, setLoading] = useState(false);

    const handleFetchRequest = async() => {
      setLoading(true);  // Setze Loading-Status auf true, um den Ladevorgang anzuzeigen
      setError(null);    // Lösche mögliche vorherige Fehler
  
      try {
        // Beispiel für den POST-Request
        const response = await fetch('http://localhost:8090/chats/create', {
          method: 'POST',
          credentials: 'include',
          headers: { 
              'Content-Type': 'application/json'
              //'Cookie': `${authorization.name}=${authorization.value}`
           },
           body: JSON.stringify({name: chatName, Username: userToInvite})
        });
  
        if (!response.ok) {
          throw new Error('Fehler beim Abrufen der Daten');
        }
  
        const data = await response.json();  // Parst die Antwort als JSON
        //setResponseData(data);  // Speichere die Antwort im State
        setChats((prevChats) => [...prevChats, data.chat]);
        closeModal();
        console.log("Data: " + JSON.stringify(data))
      } catch (err) {
        setError(err.message);  // Setze den Fehler im State
      } finally {
        setLoading(false);  // Setze den Loading-Status zurück
      }
    };

    return (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center z-50">
        <div className="bg-white p-6 rounded-lg shadow-lg w-80">
          <h3 className="text-xl font-semibold mb-4">Neuen Chat erstellen</h3>

          <form>
            <div className="mb-4">
              <label htmlFor="chatName" className="block text-sm font-medium text-black">
                Chat Name
              </label>
              <input
                id="chatName"
                type="text"
                value={chatName}
                onChange={(e) => setChatName(e.target.value)}
                className="w-full p-2 mt-1 border border-gray-300 rounded-md text-black focus:outline-none focus:ring-2 focus:ring-indigo-500"
                placeholder="Geben Sie den Chat-Namen ein"
              />
            </div>

            <div className="mb-4">
              <label htmlFor="userToInvite" className="block text-sm font-medium text-black">
                Benutzer einladen
              </label>
              <input
                id="userToInvite"
                type="text"
                value={userToInvite}
                onChange={(e) => setUserToInvite(e.target.value)}
                className="w-full p-2 mt-1 border border-gray-300 text-black rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                placeholder="Benutzername"
              />
            </div>

            <div className="flex justify-end space-x-2">
              {/* Abbrechen-Button */}
              <button
                type="button"
                onClick={() => closeModal()}
                className="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400"
              >
                Abbrechen
              </button>
              {/* Chat erstellen-Button */}
              <button
                type="button"
                onClick={handleFetchRequest}
                className="bg-indigo-500 text-white px-4 py-2 rounded-md hover:bg-indigo-600"
              >
                Chat erstellen
              </button>
            </div>
          </form>
        </div>
      </div> 
    )
}