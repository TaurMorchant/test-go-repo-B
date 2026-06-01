package repob

import repoa "github.com/TaurMorchant/test-go-repo-A"

func GreetFromB(name string) string {
	return repoa.Greet(name) + " (from B)"
}
