package BinarySearch

import "testing"

func TestRussianDollEnvelope(t *testing.T) {
	var envelopes [][]int

	envelopes = [][]int{{5,4},{6,4},{6,7},{2,3}}
	t.Log(maxEnvelopes(envelopes))

	envelopes = [][]int{{5,6},{2,1},{1,2},{3,4}}
	t.Log(maxEnvelopes(envelopes))

	envelopes = [][]int{{5,6}}
	t.Log(maxEnvelopes(envelopes))

	envelopes = [][]int{{7,9},{5,6}}
	t.Log(maxEnvelopes(envelopes))

	envelopes = [][]int{{7,5},{5,6}}
	t.Log(maxEnvelopes(envelopes))

	envelopes = [][]int{{1,1},{1,1},{1,1}}
	t.Log(maxEnvelopes(envelopes))
}