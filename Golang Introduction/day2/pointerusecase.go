package main

import "fmt"

// Define a structure Employee
type Employee struct {
    empName string // Employee's name
    salary  int    // Employee's salary
}

// Function to update the salary of an Employee
func updateSal(emp *Employee, newSal int) {
    emp.salary = newSal // Update the salary of the Employee using pointer to Employee
}

func main() {
    // Create an Employee instance
    emp := Employee{"Raghu", 30000}

    // Print the details before updating
    fmt.Println("Before update:", emp)

    // Update the salary by passing the address of the Employee instance
    updateSal(&emp, 60000)

    // Print the details after updating
    fmt.Println("After update:", emp)
}
