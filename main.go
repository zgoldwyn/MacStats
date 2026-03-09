package main

func main() {
	ch := make(chan Stats, 1)
	go pollStats(ch)
	startUI(ch)
}
