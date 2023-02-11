# Movies & Actors Rest Api in Golang ⭐️

# golang-movies-restapi

This is a simple REST that connects into a PostgreSQL to perform CRUD operations between actors and movies.

## Getting Started

These basic instructions will give you a copy of the project up and running on
your local machine for development and testing purposes. Watch out for every important step!

### Prerequisites

Requirements for the software and other tools to build, test and push

[![My Skills](https://skillicons.dev/icons?i=postgres,postman,vscode&theme=dark)](https://skillicons.dev)
- [PostgreSQL Software Installed](https://www.postgresql.org/download/)
- [IDE Visual Studio Code](https://code.visualstudio.com/download)
- [Postman](https://www.postman.com/downloads/)

### Installing

A step by step how to open the project.

The first thing to do it's open the project on your IDE. For this on the terminal VSCode screen navigate into your folder where the project remains and type:

    code .

And repeat the same for Movies/Actors (Depending on which one you opened first)

When you are done with this, the next important step it's to configure your PostgreSQL for this project.
First of all we'll need to configure our .env file which is showed in a .envexample file.

Furthermore we run the terminal in our computer and we type

    psql postgres
    
When the database system is up, we need to CREATE the DB manually

    CREATE DATABASE "name you put in your .env file"

If everything is fine until here, the only thing you have to do is going back to your VSCode and type in the terminal on the project you are working on this

    go run .
    
It should appear "Connect succesfull" as the Movies/Actor pattern is already created in the postgres database system.
You can check it out by connecting into the database in your computer terminal and checking its structure.

    psql postgres
    
    \c golang_movies_restapi
    
    \d movies or \d actors
    
    
## Running the tests

Using your own web browser working on http://localhost:8000 or http://0.0.0.0:8000, should be fine to see your results displayed into a JSON format. 
However, I strongly recommend using Postman to do the CRUD operations which errors and displays are more clear.

In main.go you are able to see the HandlerFunctions.

## Authors

  - **Alejandro Milian** - *Telematic Engineering Student* -
    [Alejandro's GitHub💻](https://github.com/amilian89)

