package websocket

import (
	"fmt"
	"messages/messages/models"
	"messages/messages/utils"
)

type Hub struct {
	// Registered clients.
	//clients map[*Client]bool
	clients map[uint]*Client

	// Inbound messages from the clients.
	send chan models.MessageAPI

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	//connections map[uint]*websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		send:       make(chan models.MessageAPI),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[uint]*Client),
		//connections: make(map[uint]*websocket.Conn),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.ID] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.ID]; ok {
				delete(h.clients, client.ID)
				close(client.send)
			}
		case message := <-h.send:
			fmt.Println(message)
			res, err := utils.IsChatParticipant(message.ChatID, message.SenderID)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res != true {
				fmt.Println("Error")
			}
			participants, err := utils.GetChatParticipants(message.ChatID)
			messageModel := models.Message{ChatId: message.ChatID, Text: message.Text, SenderId: message.SenderID}
			err = utils.CreateNewMessage(&messageModel)
			if err != nil {
				fmt.Println("Could not store message in db")
			}

			for _, client := range h.clients {
				for _, participant := range participants {
					fmt.Printf("ClientId: %d, ParticipantId: %d\n", client.ID, participant.UserID)
					if client.ID == participant.UserID {
						select {
						case client.send <- message:
						default:
							close(client.send)
							//delete(h.clients, client)
							delete(h.clients, client.ID)
						}
					}
				}
			}
			/*
				for client := range h.clients {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			*/
		}
	}
}
