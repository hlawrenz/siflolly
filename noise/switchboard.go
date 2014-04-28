package noise

type command int

const (
	register command = iota
	unregister
	broadcast
	stop
)

type switchCmd struct {
	cmd     command
	message []byte
	ch      chan []byte
}

type Switchboard struct {
	clients map[chan []byte]bool
	cmdCh   chan switchCmd
}

func (sw *Switchboard) Broadcast(msg []byte) {
	sw.cmdCh <- switchCmd{cmd: broadcast, message: msg}
}

func (sw *Switchboard) broadcast(cmd switchCmd) {
	for ch, _ := range sw.clients {
		ch <- cmd.message
	}
}

func (sw *Switchboard) Register() chan []byte {
	c := make(chan []byte)
	sw.cmdCh <- switchCmd{cmd: register, ch: c}
	return c
}

func (sw *Switchboard) register(cmd switchCmd) {
	sw.clients[cmd.ch] = true
}

func (sw *Switchboard) Unregister(c chan []byte) {
	sw.cmdCh <- switchCmd{cmd: unregister, ch: c}
}

func (sw *Switchboard) unRegister(cmd switchCmd) {
	if _, ok := sw.clients[cmd.ch]; !ok {
		return
	}

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
		}
	}

}

func NewSwitchboard() Switchboard {
	sb := Switchboard{clients: make(map[chan []byte]bool), cmdCh: make(chan switchCmd)}
	go sb.run()
	return sb
}
