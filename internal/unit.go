package internal

type QueueUnit struct {
	key  string
	next *QueueUnit
	prev *QueueUnit
}

func (u *QueueUnit) Is(key string) bool {
	return u.key == key
}

func (u *QueueUnit) Next() *QueueUnit {
	return u.next
}

func (u *QueueUnit) Prev() *QueueUnit {
	return u.prev
}

func (u *QueueUnit) SwapWith(second *QueueUnit) {
	myPrev := u.prev
	myNext := u.next
	secPrev := second.prev
	secNext := second.next

	if myPrev != second {
		u.next = secNext
		second.prev = myPrev
	} else {
		u.next = second
		second.prev = u
	}

	if myNext != second {
		u.prev = secPrev
		second.next = myNext
	} else {
		u.prev = second
		second.next = u
	}

	if u.next != nil {
		u.next.prev = u
	}
	if u.prev != nil {
		u.prev.next = u
	}

	if second.next != nil {
		second.next.prev = second
	}
	if second.prev != nil {
		second.prev.next = second
	}
}

func (u *QueueUnit) InsertBefore(newUnit *QueueUnit) {
	newUnit.prev = u.prev
	newUnit.next = u
	if u.prev != nil {
		u.prev.next = newUnit
	}
	u.prev = newUnit
}

func (u *QueueUnit) MoveUp() {
	if u.prev != nil {
		u.SwapWith(u.prev)
	}
}

func (u *QueueUnit) MoveDown() {
	if u.next != nil {
		u.SwapWith(u.next)
	}
}

func (u *QueueUnit) Remove() *QueueUnit {
	if u.prev != nil {
		u.prev.next = u.next
	}
	if u.next != nil {
		u.next.prev = u.prev
	}
	return u.next
}
