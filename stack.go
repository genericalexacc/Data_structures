package main

//Stack datastructure
type Stack struct {
	Elements []int
	H        int
}

//Add an element to the end of the stack
func (s *Stack) push(e int) {
	s.Elements = append(s.Elements, e)
	//increase the height
	s.H++
}

//Remove the last element added to the stack
func (s *Stack) pop() int {

	//panic if stack is empty
	if s.H < 1 {
		panic("STACK IS EMPTY")
	}

	//Get last element
	e := s.Elements[len(s.Elements)-1]
	//Remove the element
	s.Elements = s.Elements[:len(s.Elements)-1]
	//Reduce height
	s.H--
	//Return the last element
	return e
}

//Check if the stack is empty
func (s *Stack) empty() bool {
	return s.H == 0
}
