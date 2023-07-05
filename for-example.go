package main 


type CreditCart struct {
	cartNumber string
	shaba string
}

func main(){
	
	carts := []CreditCart{
		{cartNumber : "32893290849230" , shaba : "384829349234"},
		{cartNumber : "32893290343244" , shaba :"344449349234"},
	}
	

	for _,cart := range carts{
		println("this is cart " ,cart.cartNumber , cart.shaba)
	}

}
