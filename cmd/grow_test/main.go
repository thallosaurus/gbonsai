package main

import "github.com/thallosaurus/gbonsai/pkg/gbonsai"

func main() {
	vec := gbonsai.NewGrowingVector(0, 0)

	vec.SetString(0, 0, "Hello World", gbonsai.Black)

}
