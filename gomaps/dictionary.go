package dictionary

// Dictionary type map key string values string
type Dictionary map[string]string

// ErrDictionary type
type ErrDictionary string

const (
	// ErrNotFound const
	ErrNotFound = ErrDictionary("could not find the word you were looking for")
	// ErrWordExist const
	ErrWordExist = ErrDictionary("cannot add word because already exist")
	// ErrWordDoesNotExist const
	ErrWordDoesNotExist = ErrDictionary("cannot update word because it does not exist")
)

// Erorr method on ErrDictionary
func (e ErrDictionary) Error() string {
	return string(e)
}

// Search method on Dictionary
func (d Dictionary) Search(word string) (string, error) {
	// a map can return a boolean as a second value that idicates if the key was found successfully
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add method on Dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

// Update method on Dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete method on Dictionary
func (d Dictionary) Delete(word string) {
	// delete is a build-in function in go
	delete(d, word)
}
