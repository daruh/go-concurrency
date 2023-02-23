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
}

func (sm safeMap) run() {

}

func (sm safeMap) Insert(s string, i interface{}) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (sm safeMap) Update(s string, updateFunc UpdateFunc) {
	//TODO implement me
	panic("implement me")
}

func (sm safeMap) Close() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}
