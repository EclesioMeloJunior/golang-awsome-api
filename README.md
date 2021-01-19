# Golang Open Food Facts Challenge

This project needs to realize synchronization with Open Food Facts open data and allow CRUD operations with data

## Running mongo with docker

To properly execute the application you must have an instance of mongo database running somewhere in your machine. For look we
have (docker)[https://docs.docker.com/engine/install/] and (docker compose)[https://docs.docker.com/compose/install/] to help
us with the execution of third part apps. 

In the root of this project we had a file called **docker-compose.yml**, inside it contains the definition of three services
the first one is the app and; the mongo is the mongo database instance and the mongo-express is a tool to visualize our database data.

To execute the mongo database you must execute: 

```sh
docker-compose up mongo mongo-express
```

The above command will start the database at port 27017 and the visualization at http://localhost:8081

## Running application

To execute the application you must clone the project inside a folder at your machine then you must create a .env file
at the ***project root**, the file .env.example contains the required envs.

The env var **CONN_STR** should have the format _mongodb://root:example@localhost:27017/?authSource=admin_
The env var **DATABASE_NAME** should be _admin_
The env var **TIME_EXEC_IMPORT** should have the format _* * * *_ (Where the first one represents minutes and the second is the hour) eg. _10 13 * *_ = 13h:10m

The other variables represents the email information to be able notify when import fails or succeed

### With docker

To run with docker you must have installed (docker)[https://docs.docker.com/engine/install/] and (docker compose)[https://docs.docker.com/compose/install/] and an instance of mongodb must be running (see _Running mongo with docker_ section) and you have filled the .env file. then at the project root run:

```sh
docker-compose up app
```

or if you have make bin installed at your machine you can run:

```sh
make docker
```

Then the application will be running at http://localhost:8080


### Without docker

To run without docker you must have installed the golang (version 1.11+), an instance of mongodb must be running (see _Running mongo with docker_ section) and you have filled the .env file.

After that in the project root you can execute the app runnig:

```sh
go run ./main.go
```

or if you have make bin installed at your machine you can run:

```sh
make run
```

Then the application will be running at http://localhost:8080