package jsoner

import "testing"

func TestJSON(t *testing.T) {
	s := `
	{
		"name": "Ther",
		"age": 22,
		"hobbies": ["PC", "Game", "Go"]
	}
`
	j, e := New([]byte(s))
	if e != nil {
		panic(e)
	}
	t.Log(j.KeysInOrder())
}
