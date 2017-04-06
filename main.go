package main

func main() {
	gk := getGK()
	rk := getRK()

	n := news{}
	n = append(gk, rk...)

	cli(n)
}
