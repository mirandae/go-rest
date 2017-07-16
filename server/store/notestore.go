package store

// NoteStore interfaces the data store for users
type NoteStore interface {
	getNotes()
	getNote(noteID int)
}

type NoteStoreImpl struct {
	Notes map[int]Note
}

type Note struct {
	Title, Body string
}

func (n *NoteStoreImpl) getNotes() []int {
	keys := make([]int, len(n.Notes))

	i := 0
	for k := range n.Notes {
		keys[i] = k
		i++
	}
	return keys
}

func (n *NoteStoreImpl) getNote(noteID int) (Note, error) {
	return Note{}, nil
}
