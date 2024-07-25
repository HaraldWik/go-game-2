package ups

// MANAGER

var Manager manager = manager{
	Objs:   make(map[string]*Object),
	curObj: &Object{}, // Initialize curObj
}

type manager struct {
	Objs   map[string]*Object
	curObj *Object
}

func (m *manager) GetParent() Object {
	return *m.curObj
}

func (m *manager) Update() {
	for _, obj := range m.Objs {
		*m.curObj = *obj
		m.curObj.Update()
	}
}

func (m *manager) AddObject(name string, obj Object) {
	m.Objs[name] = &obj
}

func (s *manager) DeleteObject(name string) {
	delete(s.Objs, name)
}

// OBJ

type Object struct {
	Data *Data
	Syss []Sys
}

func NewObject() Object {
	dt := make(map[string]interface{})
	return Object{Data: (*Data)(&dt)}
}

func (o *Object) Update() {
	for _, sys := range o.Syss {
		sys.Update()
	}
}

func (o *Object) AddData(key string, value interface{}) {
	(*o.Data)[key] = value
}

func (o *Object) DeleteData(key string) {
	delete(*o.Data, key)
}

func (o *Object) AddSystems(sys ...Sys) {
	o.Syss = append(o.Syss, sys...)
}

func (o *Object) RemoveSystem(sys Sys) {
	for i, s := range o.Syss {
		if s == sys {
			o.Syss = append(o.Syss[:i], o.Syss[i+1:]...)
		}
	}
}

// DATA

type Data map[string]interface{}

func (d *Data) Set(tar string, new interface{}) {
	(*d)[tar] = new
}

func (d *Data) Get(tar string) interface{} {
	return (*d)[tar]
}

// SYS

type Sys interface {
	Update()
}
