package gset

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// Utility function to convert string record to
// a sha516 string value to be used as key
func string_to_sha512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	sha512_hash := hex.EncodeToString(h.Sum(nil))
	return sha512_hash
}

// Create gset
func Create() map[string]string {
	return make(map[string]string)
}

// Prints entire gset
func Get(gset map[string]string) {
	for _, value := range gset {
		fmt.Println(value)
	}
}

// Checks if a given record exists in the gset
func Exists(gset map[string]string, record string) bool {
	if strings.Contains(record, ".") {
		record = strings.Split(record, ".")[2]
	}
	hash := string_to_sha512(record)
	if _, exists := gset[hash]; exists {
		return true
	}
	return false
}

// Appends record to gset
func Add(gset map[string]string, record string) {
	if strings.Contains(record, ".") {
		record = strings.Split(record, ".")[2]
	}
	// create a sha512 value based on the record
	sha512_hash := string_to_sha512(record)
	gset[sha512_hash] = record
}

func GsetToString(gset map[string]string, verbose bool) string {
	if len(gset) == 0 {
		return "{}"
	}
	var s = ""
	if verbose {
		for k, v := range gset {
			s = s + "{key:" + k + ", value:" + v + "},"
		}
	} else {
		for _, v := range gset {
			s = s + "{" + v + "},"
		}
	}
	s = s[:len(s)-1]
	return s
}
