package cist

import (
	"testing"
)

func TestBuildCISTinKNN(t *testing.T) {
	knn := BuildCISTinKNN(8)
	t.Log(knn)
}
