package main

import
"fmt"

type bill struct {
	name string
	items map[string] float64
	tip float64
}

//new bills

func newBill (name string) bill{
     b:=bill {
		name : name,
		items : map[string] float64 {
		"pie":       5.45,
		"soup":      2.26,
		"schnitzel": 8.95,},
		
		tip : 0.2,
	 }

	 return b 
}

//format the bill
func (b bill) format() string {  //Receiver funtion to specifically format the bill custom struct 
	fs:= "Bill breakdown: \n"
	var total float64 = 0
	//name
	fmt.Println(b.name)
	
	//list items
	for k,v:= range b.items{
		fs+= fmt.Sprintf("%v.....$%v \n",k+":",v) //utilizng fs as a non-pointer value 
		total += v
		
	}
	
	
	//total
	fs +=fmt.Sprintf("%v.....$%0.2f \n", "total:", total)

	// tip
	tipAmount := total * b.tip
	fs += fmt.Sprintf("%v.....$%0.2f \n", "tip:", tipAmount)
	
	return fs 
}