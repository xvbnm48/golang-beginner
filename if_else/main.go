package main

func main() {
	println("if statements in go")
	var customerHeight int = 131
	if customerHeight >= 150 {
		print("can access")
	} else if customerHeight >= 130 {
		println("can access but with parents")
	} else {
		println("customer is to small")
	}
}
