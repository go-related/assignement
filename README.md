# Gin proxy

here i have created enough code to make it run and to show the road that i would follow if i had to solve this challenge.
Normally i would have liked to have much more test coverage(mock the service, the http call etc)  :) 

the code can be found [here](https://github.com/go-related/assignement) <br>
config/nuitee.yaml contains all the configuration that we are gonna need for this application.


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
   go test -v ./...
  ```
- run
  ```
    ./bin/nuitee
  ```

## check the Service

there are 2 endpoints that i have created the status one, and the one for the hotel.

- Status
  <br>this will return true of false if we can reach the other service
    ```shell
      curl --location 'http://localhost:8081/status'
    ```


- Hotel Rates
  <br> the assignment 
  ```shell
  curl --location --globoff 'http://localhost:8081/api/v1/hotels?checkin=2024-07-15&checkout=2024-07-16&currency=USD&guestNationality=US&hotelIds=129410%2C105360%2C106101%2C1762514%2C106045%2C1773908%2C105389%2C1790375%2C1735444%2C1780872%2C1717734%2C105406%2C105328%2C229436%2C105329%2C1753277&occupancies=[{%22rooms%22%3A2%2C%20%22adults%22%3A%202}%2C%20{%22rooms%22%3A1%2C%20%22adults%22%3A%201}]'
  ```
  

