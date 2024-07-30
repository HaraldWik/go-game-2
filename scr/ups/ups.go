package ups

import (
	"fmt"
	"log"
)

// Engine
var Engine engine = engine{
	Objects: make(map[string]*Object),
	curObj:  &Object{},
}

type engine struct {
	Objects map[string]*Object
	curObj  *Object
}

func (m *engine) GetParent() *Object {
	return m.curObj
}

func (m *engine) Update(deltaTime float32) {
	for _, obj := range m.Objects {
		*m.curObj = *obj
		if !obj.hasStarted {
			obj.start()
			obj.hasStarted = true
		}
		m.curObj.update(deltaTime)
	}
}

func (m *engine) AddObject(obj *Object) {
	m.Objects[obj.Name] = obj
}

func (m *engine) DeleteObject(name string) {
	delete(m.Objects, name)
}

func (m *engine) SetObject(name string, dt Data, syss []System, tags ...string) {
	obj := m.Objects[name]
	obj.Data = dt
	obj.Systems = syss
	obj.Tags = tags
}

func (m *engine) GetObject(name string) *Object {
	return m.Objects[name]
}

// Object
type Object struct {
	Name       string
	Data       Data
	Systems    []System
	Tags       []string
	hasStarted bool
}

func NewObject(name string, dt Data, syss []System, tags ...string) *Object {
	obj := &Object{
		Name:    name,
		Data:    dt,
		Systems: syss,
		Tags:    tags,
	}
	Engine.AddObject(obj)
	return obj
}

func (o *Object) Clone(times uint32, dt ...Data) {
	o.AddTags(o.Name)
	for i := 0; i < int(times); i++ {
		if i > len(dt)-1 {
			new := make(map[string]interface{})
			for key, value := range o.Data {
				new[key] = value
			}

			NewObject(
				o.Name+fmt.Sprintf("-%02d", i+1),
				new,
				o.Systems,
				o.Name,
			)
		} else {
			NewObject(
				o.Name+fmt.Sprintf("-%02d", i+1),
				dt[int(i)],
				o.Systems,
				o.Name,
			)
		}
	}
}

func (o *Object) start() {
	for _, sys := range o.Systems {
		sys.Start()
	}
}

func (o *Object) update(deltaTime float32) {
	for _, sys := range o.Systems {
		sys.Update(deltaTime)
	}
}

// Data

type Data map[string]interface{}

func (d *Data) Set(tar string, new interface{}) {
	(*d)[tar] = new
}

func (d Data) Get(tar string) interface{} {
	if value, exists := d[tar]; exists {
		return value
	} else {
		log.Fatalf("Failed to get data '%s', please attach the data to the object '%s'\n", tar, Engine.GetParent().Name)
		return nil
	}
}

func (d *Data) Delete(tar string) {
	d.Set(tar, nil)
}

// Systems
func (o *Object) AddSystemss(syss ...System) {
	o.Systems = append(o.Systems, syss...)
}

func (o *Object) RemoveSystems(syss ...System) {
	set := make(map[System]struct{}, len(syss))
	for _, sys := range syss {
		set[sys] = struct{}{}
	}

	var updatedSystems []System
	for _, s := range o.Systems {
		if _, exists := set[s]; !exists {
			updatedSystems = append(updatedSystems, s)
		}
	}
	o.Systems = updatedSystems
}

type System interface {
	Start()
	Update(deltaTime float32)
}

// Tags
func (o *Object) AddTags(tags ...string) {
	o.Tags = append(o.Tags, tags...)
}

func (o *Object) RemoveTags(tags ...string) {
	tagMap := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		tagMap[tag] = struct{}{}
	}

	var updatedTags []string
	for _, tag := range o.Tags {
		if _, found := tagMap[tag]; !found {
			updatedTags = append(updatedTags, tag)
		}
	}

	o.Tags = updatedTags
}

func (m *engine) FindTag(tar string) []*Object {
	var filteredObjects []*Object

	for _, obj := range m.Objects {
		for _, tag := range obj.Tags {
			if tag == tar {
				filteredObjects = append(filteredObjects, obj)
			}
		}
	}

	return filteredObjects
}
