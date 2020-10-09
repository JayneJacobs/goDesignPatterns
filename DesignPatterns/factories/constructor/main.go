package constructor

import "fmt"

// Employee struct {
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Factories for specific roles

// functional approach

// NewEmployeeFactory (position string, annualIncome int) func(name string) *Employee {
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// EmployeeFactory struct {
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// NewEmployeeFactory2 (position string, annualIncome int) *EmployeeFactory {
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

// Create (name string) *Employee {
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

// Mainlike () {
func Mainlike() {
	developerFactor := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Mnager", 80000)

	developer := developerFactor("Jayne")
	fmt.Println(developer)

	manager := managerFactory("Dan")
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.AnnualIncome = 11000000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}

// Advantages: An interface tells you the methods and arguements. so that is all you need to use the type.
