package main

type Person struct {
	ability int32
}

type Member struct {
	Person
	orgnization Orgnization
	role        OrgnizationRole
}
