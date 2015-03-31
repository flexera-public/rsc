package main

var TestCases = []string{
	"cm15 index clouds",
	"cm15 index /api/clouds",
	"cm15 show /api/clouds/6",
	"--x1 .cloud_type cm15 show /api/clouds/6",
	"--xm .cloud_type cm15 index clouds",
	"--xj .cloud_type cm15 index clouds",
}
