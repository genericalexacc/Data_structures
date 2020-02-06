package main

//Queue datastructure
type Queue struct {
	Elements []int
	H        int
}

//Add an element to the end of the queue
func (q *Queue) enqueue(e int) {
	//append element at the end
	q.Elements = append(q.Elements, e)
	//increase the height
	q.H++
}

//Remove the first element added to the queue and return value
func (q *Queue) dequeue() int {
	//panic if queue is empty
	if q.H < 1 {
		panic("Queue IS EMPTY")
	}
	//Get first element
	e := q.Elements[0]
	//Remove the element
	q.Elements = q.Elements[1:]
	//Reduce height
	q.H--
	//Return the last element
	return e
}

//Check if the queue is empty
func (q *Queue) empty() bool {
	return q.H == 0
}
