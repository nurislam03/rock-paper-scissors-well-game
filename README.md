# rock-paper-scissors-well-game
The rock-paper-scissors-well game is an evolution of the rock-paper-scissors game.

Game rules:
- Scissors beats paper
- Paper beats rock and well 
- Rock beats scissors 
- Well beats rock and scissors

### Test
To run the tests, run the following commands:
```sh
$ go test ./...
  ```

### Run

To run the service, run the following commands:
```sh
$ go build
$ go run main.go
  ```

### Run with Docker Compose
To run the service from docker compose, run the following commands:
```sh
$ docker-compose up
   ```
the application will be running on http://localhost:8080


#### Available Endpoint (e.g. cURL)
```sh
curl --location --request POST 'http://localhost:8080/rpsw/play' \
--header 'Content-Type: application/json' \
--data-raw '{"move": "well"}'
   ```

### Thanks