package main

import "io/ioutil"

func readCompany() []byte {
	result, _ := ioutil.ReadFile("./company.json")
	return result
}

func readEmployee() []byte {
	result, _ := ioutil.ReadFile("./employee.json")
	return result
}
