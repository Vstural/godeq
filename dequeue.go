package godeque

type item struct {
	nextItem *item
	lastItem *item
	val      interface{}
}

type Dequeue struct {
	f item // front
	r item // rear
}

func NewDequeue() *Dequeue {
	deq := new(Dequeue)
	deq.f.lastItem = &deq.r
	deq.r.nextItem = &deq.f
	return deq
}

func (d *Dequeue) Foreach(runner func(interface{}) bool) {
	cur := d.r.nextItem
	for {
		// fmt.Println("cur")
		// fmt.Println(cur)
		// fmt.Println("&d.f")
		// fmt.Println(&d.f)
		// fmt.Println(&d.f == cur)
		if cur != &d.f {

			if !(runner(cur.val)) {
				return
			}
			cur = cur.nextItem
		} else {
			return
		}
	}
}

func (d *Dequeue) PushFront(i interface{}) {
	newItem := new(item)
	front := d.f.lastItem

	front.nextItem = newItem
	newItem.lastItem = front

	d.f.lastItem = newItem

	newItem.nextItem = &d.f
	d.f.lastItem.val = i
}

func (d *Dequeue) PushBack(i interface{}) {
	newItem := new(item)
	raer := d.r.nextItem

	raer.lastItem = newItem
	newItem.nextItem = raer

	d.r.nextItem = newItem

	newItem.lastItem = &d.r
	d.r.nextItem.val = i
}

func (d *Dequeue) PopFront() interface{} {

	if d.f.lastItem == &d.r {
		return nil
	}

	res := d.f.lastItem.val
	d.f.lastItem = d.f.lastItem.lastItem
	d.f.lastItem.nextItem = &d.f
	return res
}

func (d *Dequeue) PopBack() interface{} {
	if d.r.nextItem == &d.f {
		return nil
	}

	res := d.r.nextItem.val
	d.r.nextItem = d.r.nextItem.nextItem
	d.r.nextItem.lastItem = &d.r
	return res
}

func (d *Dequeue) Back() interface{} {
	return d.r.nextItem.val
}

func (d *Dequeue) Front() interface{} {
	return d.f.lastItem.val
}
