package main

type StellarObjectData struct {
	Dec []struct {
		Value string `json:"value"`
	} `json:"dec"`
	Ra []struct {
		Value string `json:"value"`
	} `json:"ra"`
	Discoverdate []struct {
		Value string `json:"value"`
	} `json:"discoverdate"`
}
