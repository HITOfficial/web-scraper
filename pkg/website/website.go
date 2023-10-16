package website

type Website struct {
	URL   string
	Error error
	Words map[string]int
}
