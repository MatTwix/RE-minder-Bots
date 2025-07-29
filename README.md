# RE-minder Bots

RE-minder Bots is a multi-platform reminder bot written in Go. It allows users to receive notifications on various platforms like Telegram, Discord, and Google Chat.

## Features

- **Multi-platform:** Support for multiple bots (Telegram, Discord, Google Chat).
- **REST API:** An API for managing chats and sending messages.
- **Extensible:** Easily add new bots thanks to a modular architecture.
- **Database Migrations:** Built-in migration system for managing the database schema.

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Framework & Libraries:**
  - [Fiber](https://github.com/gofiber/fiber) for the REST API
  - [Telebot](https://github.com/tucnak/telebot) for the Telegram Bot
  - [DiscordGo](https://github.com/bwmarrin/discordgo) for the Discord Bot
  - [pgx](https://github.com/jackc/pgx) for PostgreSQL driver
  - [godotenv](https://github.com/joho/godotenv) for managing environment variables
  - [validator](https://github.com/go-playground/validator) for struct validation

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL

### Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/RE-minder-Bots.git
    cd RE-minder-Bots
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

3. **Configure environment variables:**

    Create a `.env` file by copying the `.env.example` and fill in the required values:

    ```bash
    cp .env.example .env
    ```

    You will need to specify:
    - `PORT`: The port to run the API server on.
    - `API_KEY`: The key to access the API.
    - `DB_URL`: The connection URL for the PostgreSQL database.
    - `TELEGRAM_BOT_TOKEN`: Token for the Telegram bot.
    - `DISCORD_BOT_TOKEN`: Token for the Discord bot.
    - `GOOGLE_CHAT_BOT_CREDENTIALS`: Credentials for the Google Chat bot (file path or JSON).

### Running the Application

Execute the following command to run the application:

```bash
go run main.go
```

The API server will start on the port specified in your `.env` file.

## API

The API is protected by an API key (`API_KEY`), which must be passed in the `X-API-Key` header.

### Endpoints

- `GET /chats`: Get a list of all chats.
- `POST /chats`: Create a new chat.
- `DELETE /chats/:id`: Delete a chat by its ID.
- `POST /chats/message`: Send a message to all chats.

**Example request to send a message:**

```bash
curl -X POST http://localhost:8080/chats/message \
-H "Content-Type: application/json" \
-H "X-API-Key: your_api_key" \
-d '{
      "message": "This is your reminder!"
    }'
```

## Architecture

### Project Structure

- `bot/`: Logic for interacting with different bot platforms.
- `config/`: Application configuration.
- `consumer/`: Handles incoming messages from bots.
- `database/`: Database interaction logic.
- `handlers/`: HTTP request handlers.
- `migrations/`: Database schema migrations.
- `models/`: Data models.
- `routes/`: API route definitions.
- `services/`: Application business logic.

### How to Add a New Bot

1. **Create a New Provider:**
    In the `bot/` directory, create a new folder for your bot (e.g., `bot/slack/`).
    Inside, create a `slack.go` file that implements the `Bot` interface from `bot/interface.go`.

2. **Implement the `Bot` Interface:**

    ```go
    package slack

    import "RE-minder-Bots/models"

    type SlackBot struct {
        // ... fields
    }

    func NewSlackBot(/* ... */) *SlackBot {
        // ... constructor
    }

    func (b *SlackBot) SendMessage(chat models.Chat, message string) error {
        // ... message sending logic
    }
    ```

3. **Add the Bot to the Factory:**
    In `bot/factory.go`, add the new bot to the `NewBotFactory` function:

    ```go
    // ...
    import (
        // ...
        "RE-minder-Bots/bot/slack"
    )

    func NewBotFactory(cfg *config.Config) *BotFactory {
        // ...
        slackBot := slack.NewSlackBot(/* ... */)
        factory.bots[models.BotTypeSlack] = slackBot
        return factory
    }
    ```

4. **Add the Bot Type:**
    In `models/Chat.go`, add the new bot type to the `BotType` enum:

    ```go
    const (
        // ...
        BotTypeSlack BotType = "slack"
    )
    ```
