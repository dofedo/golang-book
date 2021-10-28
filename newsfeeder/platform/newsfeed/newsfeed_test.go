package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{Title: "t1", Body: "b1"})
	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{Title: "t1", Body: "b1"})

	feeds := feed.GetAll()
	if len(feeds) != 1 {
		t.Errorf("Item was not added")
	}
}
