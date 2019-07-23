package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"AN ITEM", "TEST BODY"})

	if len(feed.Items) != 1 {
		t.Errorf("Item not added, TestAdd() Failed")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item not added, GetAll() Failed")
	}
}
