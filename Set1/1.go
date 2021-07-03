package main

import (
	//"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main(){
	hexstring := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decodehex, _ := hex.DecodeString(hexstring)
	fmt.Println("Decoded hex raw" , decodehex)
	fmt.Println("Decoded hex: ", string(decodehex))
	enc64 := base64.StdEncoding.EncodeToString(decodehex)
	fmt.Println("Encoded base64:", enc64)


	/*
	010010 010010 011101 101101
	18      18     29     45

	S        S      d     t	

	*/
}