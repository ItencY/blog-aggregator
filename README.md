# Gator 🐊

Gator is a powerful, production-ready Command Line Interface (CLI) RSS feed aggregator written in Go and backed by a PostgreSQL database. It allows users to manage user profiles, follow RSS feeds, automatically scrape posts in a continuous background loop, and browse news directly from the terminal.

## Features

- **User Management:** Create accounts and switch between active users.
- **Feed Subscriptions:** Add and follow RSS feeds with automatic cascade deletion.
- **Background Aggregator:** Multi-threaded scraping engine that continuously fetches and updates posts based on a configurable interval.
- **Smart Storage:** Keeps track of post publication times, updates feeds chronologically, and safely avoids duplicate records using `ON CONFLICT` constraints.
- **Offline Capabilities:** Once installed via Go, Gator compiles into a single, standalone static binary that runs independent of the Go toolchain.

---

## Prerequisites

Before installing and running Gator, ensure you have the following software installed on your machine:

1. **Go** (version 1.20 or higher recommended) -> [Install Go](https://go.dev)
2. **PostgreSQL** -> [Install PostgreSQL](https://postgresql.org)

---

## Installation

Since Go programs compile into statically bound binaries, you can install the `gator` CLI globally using the `go install` command. Navigate to your project root and run:

```bash
go install .
```

This will compile the program and place the `gator` executable directly into your `$GOPATH/bin` directory. Ensure that your Go binary path is included in your system's `PATH` variable to execute it from anywhere.

---

## Configuration & Database Setup

### 1. Database Initialization
Create a new PostgreSQL database and apply your SQL migrations to initialize the schema:

```bash
# Connect to Postgres and create the database
createdb gator

# (Optional) If you are using a migration tool like golang-migrate or goose:
# migrate -path sql/schema -database "postgres://localhost:5432/gator?sslmode=disable" up
```

### 2. Configuration File
Gator expects a JSON configuration file located at `~/.gatorconfig.json` in your user's home directory. Create this file and specify your database connection string and default user:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "alex"
}
```
*Replace `username`, `password`, and `localhost:5432` with your actual local PostgreSQL credentials.*

---

## Usage & Available Commands

Once configured, you can run the `gator` production binary directly. Here are some of the key commands available:

### Account & Session Management
*   **Register a new user:**
    ```bash
    gator register <username>
    ```
*   **Switch the current user:**
    ```bash
    gator login <username>
    ```

### Feed Subscriptions
*   **Add a new RSS feed:**
    ```bash
    gator addfeed <feed_name> <feed_url>
    ```
*   **Follow an existing feed by URL:**
    ```bash
    gator follow <feed_url>
    ```
*   **Unfollow a feed:**
    ```bash
    gator unfollow <feed_url>
    ```
*   **List all feeds the current user is following:**
    ```bash
    gator following
    ```

### Aggregating & Browsing Posts
*   **Start the continuous background scraper loop (e.g., checking every 1 minute):**
    ```bash
    gator agg 1m
    ```
*   **Browse recent posts from your followed feeds (default limit is 2):**
    ```bash
    gator browse 5
    ```

---

## Development vs. Production

*   **Development:** Use `go run . <command>` while actively writing code.
*   **Production:** Use the compiled `gator <command>` binary for real-world usage without needing the Go runtime.
