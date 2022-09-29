package datastructures

type Pool[T comparable] struct {
	data []T
	size int
    addFunc func() T
}

func CreatePool[T comparable](intitalElements int, f func() T) Pool[T] {
	p := Pool[T]{
		data: []T{},
		size: 0,
        addFunc: f,
	}

    for i := 0; i < intitalElements; i++ {
        newElement :=  p.addFunc()
	    p.data = append(p.data, newElement)
    }

	return p
}

func (p *Pool[T]) Get() T {
    if len(p.data) == 0{
        println("Added new element")
        newElement :=  p.addFunc()
    	p.data = append(p.data, newElement)
    }

    d := p.data[0]
    p.data = p.data[1:]
    return d 
}

func (p *Pool[T]) Return(val T) {
	p.data = append(p.data, val)
}
