//Exercise 1.7 - The function call io.copy(dst,src) reads from src and writes to dst.
//Use it instead of ioutil.ReadAll to copy the response body to os.stdout without requireing a buffer large enough to hold the entire stream.
//Be sure to check the error result of io.Copy

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		dst := os.Stdout
		_, err = io.Copy(dst, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//!-
