# gator

A command-line RSS aggregator. Written as part of the boot.dev backend development course. It's developed using Go programming language, the sqlc tool to generate SQL-related code and Goose for database migrations.

### Requirements

The program requires a PostgreSQL database and Go language installed.

### Usage

Configuration of the program is handled in the ".gatorconfig.json" file in the user's home directory. A full link to the database should be provided there, as "db_url" property, for example:

`
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
`

The program is used by issuing commands, including parameters for some of them. Type `gator help` to see more information.