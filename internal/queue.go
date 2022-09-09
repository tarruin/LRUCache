package internal

type Queue struct {
	maxSize int
	size    int
	first   *QueueUnit
	last    *QueueUnit
	keys    map[string]*QueueUnit
	db      map[string]string
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		maxSize: maxSize,
		db:      make(map[string]string),
		keys:    make(map[string]*QueueUnit),
	}
}

func (q *Queue) Contains(key string) bool {
	val, ok := q.keys[key]
	if ok {
		val.MoveUp()
		for q.first.Prev() != nil {
			q.first = q.first.Prev()
		}
		for q.last.Next() != nil {
			q.last = q.last.Next()
		}
	}
	return ok
}

func (q *Queue) isKeyExists(key string) bool {
	_, ok := q.keys[key]
	return ok
}

func (q *Queue) Get(key string) (string, bool) {
	if q.Contains(key) {
		return q.db[key], true
	}
	return "", false
}

func (q *Queue) Add(key string, value string) bool {
	if !q.isKeyExists(key) {
		newUnit := &QueueUnit{key: key}
		if q.first != nil {
			q.first.InsertBefore(newUnit)
		}
		if q.last == nil {
			q.last = q.first
		}
		q.first = newUnit
		q.addToDb(key, value)
		q.keys[key] = newUnit
		return true
	}
	return false
}

func (q *Queue) addToDb(key, value string) {
	q.db[key] = value
	q.size++
	for q.size > q.maxSize {
		q.TruncLast()
	}
}

func (q *Queue) removeFromDb(key string) {
	delete(q.db, key)
	q.size--
}

func (q *Queue) TruncLast() {
	if q.last != nil {
		q.removeFromDb(q.last.key)
		delete(q.keys, q.last.key)
		q.last = q.last.Prev()
		if q.last != nil {
			q.last.Next().Remove()
		} else {
			q.first = nil
		}
	}
}

func (q *Queue) TruncFirst() {
	if q.first != nil {
		q.removeFromDb(q.first.key)
		delete(q.keys, q.first.key)
		q.first = q.first.Next()
		if q.first != nil {
			q.first.Prev().Remove()
		} else {
			q.last = nil
		}
	}
}

func (q *Queue) Remove(key string) bool {
	if q.last.Is(key) {
		q.TruncLast()
		return true
	}
	if q.first.Is(key) {
		q.TruncFirst()
		return true
	}
	for i := q.first.Next(); i != q.last; i = i.Next() {
		if i.Is(key) {
			q.removeFromDb(key)
			delete(q.keys, key)
			i.Remove()
			return true
		}
	}
	return false
}
