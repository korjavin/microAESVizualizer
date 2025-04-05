# Micro AES Visualizer

A web-based interactive visualization of the AES (Advanced Encryption Standard) algorithm built with Go and WebAssembly.

## Features

- Interactive step-by-step visualization of AES encryption/decryption
- Simplified implementation focusing on core AES operations:
  - SubBytes (substitution)
  - ShiftRows
  - MixColumns
  - AddRoundKey
- Go backend compiled to WebAssembly
- Single-page web application

## Screenshots

(Add screenshots here once available)

## Running Locally

### Prerequisites

- Go 1.16 or higher

### Building and Running

1. Clone the repository:
   ```bash
   git clone https://github.com/korjavin/microAESVizualizer.git
   cd microAESVizualizer
   ```

2. Build the WebAssembly module:
   ```bash
   GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/main.go
   ```

3. Build and run the server:
   ```bash
   go build -o server cmd/server/main.go
   ./server
   ```

4. Open your browser and navigate to [http://localhost:8080](http://localhost:8080)

## Running with Docker

### Using Pre-built Image

```bash
docker run -p 8080:8080 ghcr.io/korjavin/microaesvizualizer:latest
```

Then open your browser and navigate to [http://localhost:8080](http://localhost:8080)

### Building the Docker Image Locally

```bash
docker build -t microaesvizualizer .
docker run -p 8080:8080 microaesvizualizer
```

## How to Use

1. Enter your text and encryption key in the input fields
2. Click "Initialize" to set up the initial state
3. Use the operation buttons (SubBytes, ShiftRows, MixColumns, AddRoundKey) to see how each step transforms the data
4. Navigate through the operations history with Previous/Next buttons

## Implementation Details

- This is a simplified implementation of AES for educational purposes
- The visualization shows how data is transformed at each step
- The state is represented as a 4×4 grid of bytes
- For simplicity, the implementation focuses on showing just the core operations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.