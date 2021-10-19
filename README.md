# kv-store

kv-store is basically REST api that works like a key-value store.

<!-- GETTING STARTED -->
## Getting Started

kv-store is basically REST api that works like a key-value store.
### Installation & Usage

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
3. Run container
   ```sh
  docker run -d -p 9090 kv-store  
    ```
4. Ready to go!
   ```sh
     curl -X POST -H "Content-Type application/json" \ -d '{"key":"foo", "value":"bar"}' \ http://localhost:9090
    ```

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>