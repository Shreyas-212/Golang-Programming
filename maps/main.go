package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
        "Facebook": "https://www.facebook.com",
        "Instagram": "https://www.instagram.com",
	}
	
	fmt.Println(websites)
	fmt.Println(websites["Facebook"])
	websites["LinkedIn"] = "http://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Instagram")
	fmt.Println(websites)
}