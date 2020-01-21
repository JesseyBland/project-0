## Tic Tac Toe - Project-0

Tictactoe is a CLI application that takes in user input and plays tictactoe. You can play by yourself against the built in Ai, You can play against the another player with two player option. You can not play at all and watch the outcome of Ai vs Ai. 

## Requirements

Go is a a requirement and is availbe for download at https://golang.org

Windows users will need to use bash to build this project. Bash is incorperated into git and is avalible at https://gitforwindows.org/ 

## Installation

First, open terminal / bash.

Second, install package with command: `go get -u github.com/JesseyBland/project-0`

Thrid, Navagate to the `project-0` directory run both of these commands.

CLI
```bash
go build cmd/tictactoe/tictactoe.go
```
HTTP server
```bash
go build cmd/tictactoed/tictactoed.go
```

## Command Line Interface.

Tic tac toe accepts numerous Command-line flags:
**-tp**
Two Player option
**-av**
Ai vs Ai option

Example:
```bash
./tictactoe -tp
```
## Http server 

The server runs off Localhost:8080 includes directories /tt1 /ttt2 /ttt3

To start the server navigate in terminal/bash to the tictactoed and run the program.
Now look for the Server status.
Example:
```bash
./tictactoed 
Server Status:Listening Host:localhost Port:8080 
```
Once the server has started navigate in your browser to localhost:8080

3 available play modes 

Player vs AI
Player vs Player
Ai vs Ai

## Features Road map

- [x] User can Play against AI
- [x] 2 player flag option
- [x] Respond to username
- [x] Add http browser play 
- [ ] Add config file for customization of map appearance and board size.
- [ ] Add Client to connect to http server via the command line.
