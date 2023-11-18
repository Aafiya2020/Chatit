# Chatit - Real-Time Chat Application

 ________________
  |               |
  |   Chatit      |
  |_______________|
 / \  / \  / \  / \
( C )( h )( a )( t )
 \_/  \_/  \_/  \_/

Chatit is a real-time chat application that allows users to engage in instant messaging with others. It provides a seamless and interactive chat experience, making communication efficient and enjoyable.

## Features

- **Real-Time Messaging:** Enjoy instant messaging with other users in real-time.
- **User Registration:** Users can create accounts and log in securely.
- **Contact Verification:** Verify the existence of contacts before initiating conversations.
- **Chat History:** View chat history between users within a specified time range.
- **Contact List:** Retrieve and display the list of contacts for a user.
- **User-Friendly Interface:** Intuitive and easy-to-use interface for a smooth user experience.

## Technologies Used

- **Go (Golang):** Backend server development.
- **Redis:** Data storage for user information and chat history.
- **WebSocket:** Real-time communication between clients and the server.
- **HTML, CSS, JavaScript:** Frontend development for the user interface.
- **GitHub Actions:** Continuous Integration (CI) for automated testing.

## Getting Started
### Step 1
Clone the repository

### Step 2
Run `git mod tidy` to install all the Golang dependencies.

### Step 3
Go to `clients` and to install frontend dependencies.

```node
npm install
```

## Run the Application
### Terminal 1
Start HTTP server
```
go run main.go --server=http
```

### Terminal 2
Start WebSocket server

```
go run main.go --server=websocket
```

### Terminal 3
Go to `client` and run

```
npm start
```

Application is live at `localhost:3000`. 

