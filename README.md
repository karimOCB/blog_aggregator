# Blog Aggregator

## Required Installations

### *Go*
  - Install Go 1.22+ from [golang.org](https://go.dev/).

### *Postgres*
  - Ensure the Postgres server is running.
  - Connect to Postgres: `psql -U postgres -h localhost`
    or if you are on linux: `sudo -u postgres psql`
  - Create a database, run: `CREATE DATABASE gator;`
  - Have a connection string ready: `postgres://username:password@localhost:5432/gator`
    replace username and password with your own Postgres credentials.

### Database Migrations
  - Make sure to apply the Goose migrations to create the required tables:
    ```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest https://github.com/pressly/goose/v3/cmd/goose@latest
    cd sql/schema
    goose postgres "postgres://username:password@localhost:5432/gator" up
    ```

### *CLI Installation*
  - `go install github.com/karimOCB/blog_aggregator@latest`

### *Set Config file*
  - In your home directory create a `.gatorconfig.json` 
  - It should contain a struct like: 
    ``` JSON
    {
      "db_url": "postgres://username:password@localhost:5432/gator",
      "current_user_name": "your_username" 
    }
    ```
  - The user can be set by register and login commands

### Essential Commands
- **`gator register <name>`**: Register a new user and log in automatically.
- **`gator login <name>`**: Switch to an existing user.
- **`gator users`**: List all registered users.
- **`gator addfeed <name> <url>`**: Add an RSS feed for the current user to track.
- **`gator feeds`**: List all added RSS feeds.
- **`gator follow <url>`**: Follow an existing feed.
- **`gator following`**: List feeds followed by the logged-in user.
- **`gator agg <time_between_reqs>`**: Start aggregating feeds continuously (e.g., `gator agg 1m` or `1s`).
- **`gator browse <limit>`**: View recent posts collected by the aggregator (defaults to 2).