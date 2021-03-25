package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

// HMAC, Hash-based Message Authentication
// The same input should produce the same output
// How we can use HMAC
// If we store a value in a users machine, and want to verify that
// number has not been tempered with,
// 1. We can take the value we want to store, run it through HMAC + Private key,
//    this creates a hash
// 2. We can then then take the value + hash, pin the 2 together, store on them
//    on the users machine
// 3. When we get the value and hash from the user, we can run it throut the
//    hash algorithm with our private key, this produces a hash
// 4. We can then compare that hash to the one stored on the users machine
// These should match if values are not tempered with

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	// hmac.New(), takes a shar256 hash + a private key
	// and returns a hash
	h := hmac.New(sha256.New, []byte("ourkey"))
	// Here we write the passed string into the hash
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
