// https://thenewstack.io/understanding-golang-type-system/
package main

import (
	"fmt"
	// "time"
)

type person struct {
	name        string
	age         int
	city, phone string
}

//SayHello A person method
func (p person) SayHello() {
	fmt.Printf("Hi, I am %s, from %s\n", p.name, p.city)
}

//GetDetails A person method
func (p person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", p.name, p.age, p.city, p.phone)
}

//Speaker type
type speaker struct {
	person     //type embedding for composition
	speaksOn   []string
	pastEvents []string
}

//GetDetails overrides
func (s speaker) GetDetails() {
	//Call person GetDetails
	s.person.GetDetails()
	fmt.Println("Speaker talks on following technologies:")
	for _, value := range s.speaksOn {
		fmt.Println(value)
	}
	fmt.Println("Presented on the following conferences:")
	for _, value := range s.pastEvents {
		fmt.Println(value)
	}
}

func main() {
	shiju := speaker{person{"Shiju", 35, "Kochi", "+91-94003372xx"},
		[]string{"Go", "Docker", "Azure", "AWS"},
		[]string{"FOSS", "JSFOO", "MS TechDays"}}
	fmt.Printf("%+v\n", shiju)
	// [Name: Shiju, Age: 35, City: Kochi, Phone: +91-94003372xx]
	// before override
	shiju.GetDetails()
}
