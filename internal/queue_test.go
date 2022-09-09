package internal

import "testing"

func PrepareQueue() *Queue {
	queue := NewQueue(3)
	queue.Add("a", "1")
	queue.Add("b", "2")
	return queue
}

func TestPrepare(t *testing.T) {
	queue := PrepareQueue()
	if queue.maxSize != 3 {
		t.Errorf("expected max size %d, got %d", 3, queue.maxSize)
	}
	if queue.size != 2 {
		t.Errorf("expected size %d, got %d", 2, queue.size)
	}
	if !queue.first.Is("b") {
		t.Errorf("expected highest priority key %q, got %q", "b", queue.first.key)
	}
	if !queue.last.Is("a") {
		t.Errorf("expected lowest priority key %q, got %q", "a", queue.first.key)
	}
}

func TestQueue_Get(t *testing.T) {
	t.Run("existed key", func(t *testing.T) {
		queue := PrepareQueue()
		val, ok := queue.Get("a")
		if !ok {
			t.Error("key 'a' not found")
		}
		if val != "1" {
			t.Errorf("expected value %q, got %q", "1", val)
		}
	})
	t.Run("non-existed key", func(t *testing.T) {
		queue := PrepareQueue()
		val, ok := queue.Get("x")
		if ok {
			t.Error("key 'x' found")
		}
		if val != "" {
			t.Errorf("expected value empty value, got %q", val)
		}
	})
}

func TestQueue_Contains(t *testing.T) {
	t.Run("existed key", func(t *testing.T) {
		queue := PrepareQueue()
		ok := queue.Contains("a")
		if !ok {
			t.Error("key 'a' not found")
		}
		if !queue.first.Is("a") {
			t.Errorf("expected highest priority key %q, got %q", "a", queue.first.key)
		}
	})
	t.Run("non-existed key", func(t *testing.T) {
		queue := PrepareQueue()
		ok := queue.Contains("x")
		if ok {
			t.Error("key 'x' found")
		}
		if !queue.first.Is("b") {
			t.Errorf("expected highest priority key %q, got %q", "b", queue.first.key)
		}
	})
}

func TestQueue_Add(t *testing.T) {
	t.Run("non-existed key", func(t *testing.T) {
		queue := PrepareQueue()
		ok := queue.Add("c", "3")
		if !ok {
			t.Error("key 'c' not added")
		}
		if !queue.first.Is("c") {
			t.Errorf("expected highest priority key %q, got %q", "c", queue.first.key)
		}
		if !queue.Contains("c") {
			t.Error("key 'c' not added")
		}
	})
	t.Run("existed key", func(t *testing.T) {
		queue := PrepareQueue()
		ok := queue.Add("a", "18")
		if ok {
			t.Error("key 'a' added successfully")
		}
		if !queue.first.Is("b") {
			t.Errorf("expected highest priority key %q, got %q", "b", queue.first.key)
		}
	})
	t.Run("oversize queue", func(t *testing.T) {
		queue := PrepareQueue()
		queue.Add("c", "3")
		queue.Add("d", "4")
		if queue.size != queue.maxSize {
			t.Errorf("expected queue size %d, got %d", queue.maxSize, queue.size)
		}
		if queue.Contains("a") {
			t.Errorf("expected key 'a' truncated")
		}
	})
}

func TestQueue_Remove(t *testing.T) {
	t.Run("remove any", func(t *testing.T) {
		queue := PrepareQueue()
		queue.Add("c", "3")
		ok := queue.Remove("b")
		if !ok {
			t.Error("expected 'b' remove return true, got false")
		}
		if queue.Contains("b") {
			t.Error("expected 'b' not existed")
		}
		if queue.size != 2 {
			t.Errorf("expected queue size %d, got %d", 2, queue.size)
		}
		if queue.first.key != "c" {
			t.Errorf("expected first key to be %q, got %q", "c", queue.first.key)
		}
		if queue.last.key != "a" {
			t.Errorf("expected first key to be %q, got %q", "a", queue.last.key)
		}
	})
	t.Run("remove first", func(t *testing.T) {
		queue := PrepareQueue()
		ok := queue.Remove("b")
		if !ok {
			t.Error("expected 'b' remove return true, got false")
		}
		if queue.first != queue.last {
			t.Errorf("expected first == last")
		}
		if queue.first.key != "a" {
			t.Errorf("expected first key to be %q, got %q", "a", queue.first.key)
		}
	})
	t.Run("remove last", func(t *testing.T) {
		queue := PrepareQueue()
		ok := queue.Remove("a")
		if !ok {
			t.Error("expected 'a' remove return true, got false")
		}
		if queue.first != queue.last {
			t.Errorf("expected first == last")
		}
		if queue.first.key != "b" {
			t.Errorf("expected first key to be %q, got %q", "b", queue.first.key)
		}
	})
	t.Run("remove all", func(t *testing.T) {
		queue := PrepareQueue()
		if !queue.Remove("a") {
			t.Error("expected 'a' remove return true, got false")
		}
		if !queue.Remove("b") {
			t.Error("expected 'b' remove return true, got false")
		}
		if queue.first != queue.last {
			t.Errorf("expected first == last")
		}
		if queue.first != nil {
			t.Errorf("expected first nil, got %v", queue.first)
		}
	})
	t.Run("remove non-existed", func(t *testing.T) {
		queue := PrepareQueue()
		if queue.Remove("x") {
			t.Error("expected 'x' remove return false, got true")
		}
		if queue.first.key != "b" {
			t.Errorf("expected first not changed, got %q", queue.first.key)
		}
		if queue.last.key != "a" {
			t.Errorf("expected last not changed, got %q", queue.last.key)
		}
	})
}
