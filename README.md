# Jadwalin User Server Management

User Server Management backend program

## Project Setup

To start the project, you need to make sure that you have Go (Golang) and Docker installed. Run these command:

```
git clone https://github.com/watermelonreborn/jadwalin-user-server.git
cd jadwalin-user-server
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/spf13/viper
go get gorm.io/gorm
docker-compose up
```

After "docker-compose up", a docker container will created.

### Modify Configuration

This app will take configurations from environment variables or `.env` file in `/config`. If you're running this app locally, it is recommended to duplicate the `.env.sample` file in `/config` and rename it to `.env`.

## Running the Program

You can run the program by simply running these commands:

```
go mod tidy
go run main.go
```

By default, you can access the app from `http://localhost:8080/`
