package main

import (
	"encoding/hex"
)

func main(){
	fixedxor("1c0111001f010100061a024b53535009181c","686974207468652062756c6c277320657965")
}

func fixedxor(input, second string) string {
	decodeinput, _ := hex.DecodeString(input) // get byte representation
	decodesecond, _ := hex.DecodeString(second) // get byte representation
	output := make([] byte, 0)

	for i := range decodeinput {		
		output = append(output, decodeinput[i] ^ decodesecond[i]) 
	}
	return hex.EncodeToString(output)
}