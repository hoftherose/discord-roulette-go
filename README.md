# discord-roulette
Discord bot to kill myself (or someone else). Russian roulette bot with powerups and various modes.

## Setup Bot

To setup the bot you will need to have `golang 1.19` installed. Run the following to get all of the dependencies.

``` golang
go mod download
go mod verify
go run main.go
```

If instead you want to run the bot on docker you can use the following commands instead. For this you already need to have docker setup.

``` bash
docker build -t discord_bot_dev -f Dockerfile .
docker run --env-file .env discord_bot_dev
```

For local deployment using the docker-compose would be more comfortable since this is a simple one line setup and provides auto-reload. That said, it is not recommended for production due to the larger image size and dynamic nature.

``` bash
docker-compose up -d
```
