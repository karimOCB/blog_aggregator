# Required Installations

1. *Go*
2. *Postgres*
  - Ensure the Postgres server is running
  - Connect to Postgres: `psql -U postgres -h localhost`
    or if you are on linux: `sudo -u postgres psql`
  - Create a database, run: `CREATE DATABASE gator;`
  - Have a connection string ready: `postgres://username:password@localhost:5432/gator`
    replace username and password with your own Postgres credentials.
3. *CLI Installation*
go install github.com/karimOCB/blog_aggregator@latest