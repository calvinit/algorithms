package algs4test

import "embed"

//go:embed testdata/jobs.txt
var jobs string

//go:embed testdata/jobs.txt
var jobsBytes []byte

//go:embed testdata/list.txt
//go:embed testdata/m?.txt
var listm123FS embed.FS

func GetJobs() string         { return jobs }
func GetJobsBytes() []byte    { return jobsBytes }
func GetListm123FS() embed.FS { return listm123FS }
