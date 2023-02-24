package safemap

// SafeMap interface
type SafeMap interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

type UpdateFunc func(interface{}, bool) interface{}

type commandAction int

const (
	remove commandAction = iota
	end
	find
	insert
	length
	update
)

type safeMap chan commandData

func New() SafeMap {
	sm := make(safeMap)
	go sm.run()
	return sm
}

type commandData struct {
	action  commandAction
	key     string
	value   interface{}
	result  chan<- interface{}
	data    chan<- map[string]interface{}
	updater UpdateFunc
}

func (sm safeMap) run() {

	store := make(map[string]interface{})
	for command := range sm {
		switch command.action {
		case insert:
			store[command.key] = command.value
		case length:
			command.result <- len(store)
		case end:
			close(sm)
			command.data <- store
		}
	}
}

func (sm safeMap) Insert(key string, value interface{}) {
	sm <- commandData{action: insert, key: key, value: value}
}

func (sm safeMap) Delete(s string) {
	//TODO implement me
	panic("implement me")
}

func (sm safeMap) Find(s string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}

func (sm safeMap) Len() int {
	reply := make(chan interface{})
	sm <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (sm safeMap) Update(s string, updateFunc UpdateFunc) {
	//TODO implement me
	panic("implement me")
}

func (sm safeMap) Close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	sm <- commandData{action: end, data: reply}
	return <-reply
}
