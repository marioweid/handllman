![GopherImageLogo](./img/logo_with_text.png)

The best Tool when you need shitty advice from a buddy that has no idea of how to do handy work. 
Written in go!

## Abstract
1. Have go based api that talks to gemini and gives you the ability to communicate with and bot. (Backend)
2. Whatever (Frontend)
3. Postgresql with pgVector that stores manuals of machiens that you own in an vector database (mounted via `/local/manuals:/pg/manuals` or whatever)
4. Endpoint to upload manuals
5. (If im feeling super fancy) WebSearch for the manual given the name of the machine or asking the user question about which year how money buckz they spended on it blaa to determine which machine it could be
6. Giving Advice what could be broken
7. Giving Advice how it could be fixed also maybe a shopping list of what we need to fix it

## What's the point of this project?
- I want to do go 
- I suck at fixing stuff
- Reading manuals is lame
- Handymans are expensive

## Prerequisites

- Docker
- Docker Compose
- Go (for local development)

## Project Structure

```
gin-docker-project
├── src
│   ├── main.go          # Entry point of the Gin application
│   ├── go.mod           # Module dependencies
│   ├── go.sum           # Checksums for module dependencies
│   └── config
│       └── database.go  # Database connection configuration
├── docker-compose.yaml   # Defines services for the application
├── Dockerfile            # Docker image build instructions
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd gin-docker-project
   ```

2. **Build and run the application:**
   Use Docker Compose to build and run the application with the following command:
   ```bash
   docker-compose up --build
   ```

3. **Access the application:**
   Once the containers are up and running, you can access the Gin application at `http://localhost:8080`.

## Hot Reload

The Gin application is set up with hot reload using [air](https://github.com/cosmtrek/air). Any changes made to the source code will automatically restart the application.

## Database Configuration

The PostgreSQL database is configured in the `docker-compose.yaml` file. Ensure that the connection details in `src/config/database.go` match the environment variables set in the Docker Compose file.

## Usage

- To stop the application, press `CTRL+C` in the terminal where Docker Compose is running.
- To remove the containers, run:
  ```bash
  docker-compose down
  ```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.