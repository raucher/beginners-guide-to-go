package main

import "fmt"

type Job struct {
	Position string
	Salary   int64
	Tech     []string
}

type Person struct {
	Name    string
	Age     uint8
	Address string
	Job     Job
}

func (p Person) hasPayedJob() {
	if p.Job.Salary > 0 {
		fmt.Printf("%s has the payed job\n", p.Name)
	} else {
		fmt.Printf("%s doesn't have payed job yet\n", p.Name)
	}
}

// We should use pointers to update structure
func (p *Person) addAge(years uint8) {
	p.Age += years
}

func main() {
	println("Don't panic")

	person := Person{Name: "John", Age: 42, Address: "Elm Street 1428, Springwood, Ohio"}
	fmt.Println(person) // {John 42 Elm Street 1428, Springwood, Ohio}

	person = Person{Name: "John", Address: "Elm Street 1428, Springwood, Ohio"}
	fmt.Println(person) // {John 0 Elm Street 1428, Springwood, Ohio}

	person = Person{Name: "John"}
	fmt.Println(person) // {John 0 }

	// With Job
	personWithJob := Person{
		Name:    "Morris",
		Age:     42,
		Address: "Elm Street 1428, Springwood, Ohio",
		Job:     Job{"Manager", 50000, []string{"MS Office", "E-mail", "Talking"}},
	}

	fmt.Println(personWithJob)     // {John 42 Elm Street 1428, Springwood, Ohio {Manager 50000 [MS Office E-mail Talking]}}
	fmt.Println(personWithJob.Job) // {Manager 50000 [MS Office E-mail Talking]}

	personWithJob.addAge(15)
	fmt.Println(personWithJob)

	person.hasPayedJob()
	personWithJob.hasPayedJob()
}
