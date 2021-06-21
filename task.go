package main

import "fmt"

func main() {
	fmt.Print("Hello")

	d := Dad{
		Person: Person{
			Name: "BUDI",
			Age:  60,
		},
	}

	m := Mom{
		Person: Person{
			Name: "IBU BUDI",
			Age:  50,
		},
	}

	fmt.Print(d.Name)
	FindJob(d)

	fmt.Print(m.Name)
	FindJob(m)
}

type Person struct {
	Name string
	Age  int
}

type Dad struct {
	Person
}

type Mom struct {
	Person
	Husband Dad
}

type Child struct {
	Person
	Dad Dad
	Mom Mom
}

type JobsInterface interface {
	Job() string
}

type StatusInterface interface {
	Status() string
}

func (d Dad) Job() string {
	return "DAGANG"
}

func (m Mom) Job() string {
	return "MEMASAK"
}

func FindJob(j JobsInterface) {
	fmt.Print(j.Job())
}
