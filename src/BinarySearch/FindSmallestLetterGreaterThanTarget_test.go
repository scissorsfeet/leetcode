package BinarySearch

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	t.Log(string(nextGreatestLetter([]byte("cfj"), byte('a'))))
	t.Log(string(nextGreatestLetter([]byte("cfj"), byte('c'))))
	t.Log(string(nextGreatestLetter([]byte("cfj"), byte('d'))))
	t.Log(string(nextGreatestLetter([]byte("cfj"), byte('g'))))
	t.Log(string(nextGreatestLetter([]byte("cfj"), byte('j'))))
	t.Log(string(nextGreatestLetter([]byte("cfj"), byte('k'))))
	t.Log(string(nextGreatestLetter([]byte("eeeeeennnnn"), byte('e'))))
	t.Log(string(nextGreatestLetter([]byte("eeeeeennnnn"), byte('n'))))
}
