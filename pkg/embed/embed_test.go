package embed

import (
	"testing"

	"github.com/calvinit/algorithms/internal/algs4test"
)

func TestReadFile(t *testing.T) {
	print(algs4test.GetJobs())
}

func TestReadFileBytes(t *testing.T) {
	print(string(algs4test.GetJobsBytes()))
}

func TestEmbedFS(t *testing.T) {
	folder := algs4test.GetListm123FS()
	listContent, _ := folder.ReadFile("testdata/list.txt")
	print(string(listContent))
	m1Content, _ := folder.ReadFile("testdata/m1.txt")
	print(string(m1Content))
	m2Content, _ := folder.ReadFile("testdata/m2.txt")
	print(string(m2Content))
	m3Content, _ := folder.ReadFile("testdata/m3.txt")
	print(string(m3Content))
}
