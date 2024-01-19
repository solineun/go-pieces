package main

type MyData struct {
	Name string `csv:"name"`
	Age int 	`csv:"age"`
	HasPet bool	`csv:"has_pet"`
}

var data = `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"
`