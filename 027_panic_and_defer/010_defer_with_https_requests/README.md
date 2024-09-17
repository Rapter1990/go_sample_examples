Here’s the `README.md` for the `defer with HTTP requests` example following the same format:

---

# Go Sample Example - Defer with HTTP Requests

This repository demonstrates how to use `defer` in Go to manage resources when making HTTP requests. It shows how `defer` can be used to ensure that resources like HTTP response bodies are closed properly after use.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of `defer` to handle resource cleanup in HTTP requests.</li>
  <li>It highlights how to use `defer` to close the HTTP response body, ensuring that resources are released even if an error occurs during the request or response processing.</li>
  <li>This example involves making an HTTP GET request and reading the response body.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// using defer to close the response body of an HTTP request

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/010_defer_with_https_requests` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/010_defer_with_https_requests
   ```

4. Run the Go program:

   ```bash
   go run 010_defer_with_https_requests.go
   ```

### 📦 Output
```bash
   html response from https://www.google.com
```
