# kv-store ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/svbnbyrk/kv-store)

kv-store is basically REST api that works like a key-value store.

<!-- GETTING STARTED -->

## Installation

1. Clone the repo

   ```sh
   git clone https://github.com/svbnbyrk/kv-store.git
   ```

2. Go inside project

   ```sh
   cd kv-store
   ```

3. Build docker file

   ```sh
   docker build -t kv-store .   
    ```

4. Run container

   ```sh
   docker run -d -p 9090 kv-store  
    ```

Ready to go!

## Usage

### Get Request

`GET /`

    curl -i -H 'Accept: application/json' http://localhost:9090/?key=foo

### Response

    HTTP/1.1 200 OK
    Date: Web, 19 Feb 2021 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    [{"key":"foo","value":"bar"}]

### Post Request

`POST /`

    curl -i -H 'Accept: application/json' -d '{"key":"foo","value":"bar"}' http://localhost:9090

### Response

    HTTP/1.1 200 OK
    Date: Web, 19 Feb 2021 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /thing/1
    Content-Length: 36

### Flush Request

`GET /flush`

    curl -i -H 'Accept: application/json' http://localhost:9090/flush

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 36

## License

Distributed under the MIT License. See `LICENSE` for more information.
