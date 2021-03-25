package main

import (
	"encoding/base64"
	"fmt"
)

// Base64
// For eg, if you want to store text in double qoutes into a cookie,
// e.g. JSON, cookies don't like double qoutes, best way is encode the JSON
// or string in Base64, put that in the cookie, then decode to use

func main() {
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
