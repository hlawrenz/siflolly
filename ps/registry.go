package noise

type Registry struct {
	clients map[chan interface{}]bool
}


func (reg *Registry) add(ch chan interface{}) {
	reg.clients[ch] = true
}

func (reg *Registry) remove(ch chan interface{}) {
	if _, ok := reg.clients[ch]; !ok {
		return
	}

	delete(reg.clients, ch)

}
