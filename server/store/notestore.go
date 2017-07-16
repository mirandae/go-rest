package store

// NoteStore interfaces the data store for users
type NoteStore interface {
	getNotes()
	getNote(noteID int)
}

// NoteStoreImpl implements NoteStore
// NOT safe for concurrent use (TODO change impl)
type NoteStoreImpl struct {
	Notes map[int]Note
}

// Note is a basic note-type
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

// TODO
func (n *NoteStoreImpl) getNote(noteID int) (Note, error) {
	return Note{}, nil
}
