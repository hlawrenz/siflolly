package cacophony

/*
	Types of communication:
	send a message
	set nick

*/
type Message struct {
	/*
		op can be:
			say
			nick
			catchup
	*/
	Op      string
	From    string
	To      string
	Payload []string
}
