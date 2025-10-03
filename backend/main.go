package main

import "ecommerce/cmd"

//import "encoding/base64"

//import "ecommerce/cmd"

func main() {

	cmd.Serve()

	// var s string
	// s = "a"

	// byteArr := []byte(s)
	// fmt.Println(s)
	// fmt.Println(byteArr)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(byteArr)
	// fmt.Println(b64Str)

	// decodedStr, err := enc.DecodeString(b64Str)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(decodedStr)
	// data := []byte("hello guys")
	// hash := sha256.Sum256(data)
	// fmt.Println("hash after sha-256", hash)

	// secret := []byte("my-secret")
	// msg := []byte("hello my world")

	// h := hmac.New(sha256.New, secret)
	// h.Write(msg)
	// text := h.Sum(nil)
	// fmt.Println(text)

	// jwt, err := utility.CreateJwt("my-secret", utility.Payload{
	// 	Sub:         54,
	// 	FirstName:   "Alif",
	// 	LastName:    "boss",
	// 	Email:       "alif@gmail.com",
	// 	IsShopOwner: false,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(jwt)
}
