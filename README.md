# Gin proxy

here i have created enough code to make it run and to show the road that i would follow if i had to solve this challenge.
Normally i would have liked to have much more test coverage :) 

## Create the executeble
- create a bin for linux
    ```
    GOOS=linux GOARCH=amd64 go build -o bin/nuitee cmd/main/*go
    ```
- create a bin for mac
  ```
   go build -o bin/nuitee cmd/main/*go
    ```
## Run the tests

- run the test
  ```
   go test -v -race ./...
  ```
- run
  ```
    ./bin/nuitee
  ```
