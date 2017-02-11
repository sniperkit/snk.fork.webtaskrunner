package commandhandler

type responseLine struct {
	Status int
	Line   string
}
type responseError struct {
	Status int
	Error  string
}
