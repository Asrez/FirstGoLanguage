package main



type List[T any] struct{
	Items []T
}

func main() {}


func (l *List[T]) Add(item T){
	l.Items = append(l.Items, item)
}