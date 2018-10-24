# boggle

An implemantation of the game [Boggle](https://en.wikipedia.org/wiki/Boggle)

## Run in the Terminal

```
go get -d github.com/berto/boggle
cd $GOPATH/github.com/berto/boggle/cmd/boggle
go run boggle.go
```

## Server

`main.go` runs a server to provide a web API to check if a word is valid.

Routes:

- GET /dictionary?word={word}
- GET /delaytionary?word={word} // adds a delay

Deployed [link](https://whispering-falls-21983.herokuapp.com/)