package composite

import "testing"

func TestComposite(t *testing.T) {
	root := &Department{
		List: make(map[string]*Department),
		Name: "root",
		Id:   "0010",
	}
	c1 := &Department{
		List: make(map[string]*Department),
		Name: "user1",
		Id:   "00010",
	}
	c2 := &Department{
		List: make(map[string]*Department),
		Name: "user2",
		Id:   "00011",
	}
	c3 := &Department{
		List: make(map[string]*Department),
		Name: "user3",
		Id:   "000010",
	}
	root.Add(c1, c2)
	c1.Add(c3)
	root.Find("000010").ReadList()
	root.ReadList()
	root.Remove("00010")
	root.ReadList()
}
