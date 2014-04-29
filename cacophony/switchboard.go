package cacophony

import (
	"fmt"
)

type command int

const (
	register command = iota
	unregister
	broadcast
	stop
	setnick
)

type switchCmd struct {
	cmd     command
	message Message
	ch      chan Message
	nick    string
}

type Switchboard struct {
	clients map[chan Message]string
	cmdCh   chan switchCmd
	nicks   map[string]bool
}

func nickMaker() chan string {
	c := make(chan string)
	i := 0

	go func() {
		for {
			i += 1
			nick := fmt.Sprintf("MOOK_%d", i)
			c <- nick
		}
	}()

	return c
}

var nicker chan string = nickMaker()

func (sw *Switchboard) Broadcast(msg Message) {
	sw.cmdCh <- switchCmd{cmd: broadcast, message: msg}
}

func (sw *Switchboard) broadcast(cmd switchCmd) {
	for ch, _ := range sw.clients {
		ch <- cmd.message
	}
}

func (sw *Switchboard) Register() chan Message {
	c := make(chan Message)
	sw.cmdCh <- switchCmd{cmd: register, ch: c}
	return c
}

func (sw *Switchboard) register(cmd switchCmd) {
	var nick string
	for {
		nick = <-nicker
		if _, ok := sw.nicks[nick]; !ok {
			break
		}
	}

	sw.clients[cmd.ch] = nick
	sw.nicks[nick] = true
	cmd.ch <- Message{
		Op:      "nick",
		From:    "@@kingmook",
		To:      nick,
		Payload: []string{nick}}
}

func (sw *Switchboard) SetNick(c chan Message, msg Message) {
	sw.cmdCh <- switchCmd{cmd: setnick, ch: c, nick: msg.Payload[0]}
}

func (sw *Switchboard) setNick(cmd switchCmd) {
	if cmd.nick == "@@kingmook" {
		/* @@kingmook is the servers name */
		cmd.ch <- Message{
			Op:      "error",
			From:    "@@kingmook",
			To:      sw.clients[cmd.ch],
			Payload: []string{"That's MY NAME!"}}
	} else if _, ok := sw.nicks[cmd.nick]; ok {
		/* nick exists */
		cmd.ch <- Message{
			Op:      "error",
			From:    "@@kingmook",
			To:      sw.clients[cmd.ch],
			Payload: []string{"Nick taken"}}
	} else {
		/* nick doesn't exist */
		sw.clients[cmd.ch] = cmd.nick
		sw.nicks[cmd.nick] = true
		cmd.ch <- Message{
			Op:      "nick",
			From:    "@@kingmook",
			To:      cmd.nick,
			Payload: []string{cmd.nick}}
	}
}

func (sw *Switchboard) Unregister(c chan Message) {
	sw.cmdCh <- switchCmd{cmd: unregister, ch: c}
}

func (sw *Switchboard) unRegister(cmd switchCmd) {
	if _, ok := sw.clients[cmd.ch]; !ok {
		return
	}

	delete(sw.nicks, sw.clients[cmd.ch])
	delete(sw.clients, cmd.ch)
}

func (sw *Switchboard) run() {

	for c := range sw.cmdCh {
		switch c.cmd {
		case register:
			sw.register(c)
		case unregister:
			sw.unRegister(c)
		case broadcast:
			sw.broadcast(c)
		case setnick:
			sw.setNick(c)
		}
	}

}

func NewSwitchboard() Switchboard {
	sb := Switchboard{
		clients: make(map[chan Message]string),
		cmdCh:   make(chan switchCmd),
		nicks:   make(map[string]bool)}
	go sb.run()
	return sb
}
