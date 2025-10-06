- ### Installing the dependencies and setup env
  Inside the newly cloned project folder run the following command to install the dependencies:
  ```bash
  $ go mod download
  ```

- ### Running the application
  Inside the project, run the following command to run the application:
  ```bash
  $ API_PORT=:8080 go run main.go
  ```

  $ curl --request GET \
  --url http://127.0.0.1:8080/api/agg/fetch/uuid \
  --header 'Content-Type: application/json'

  ```
