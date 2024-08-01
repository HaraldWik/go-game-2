package ups

import (
	"fmt"
	"log"
	"reflect"
)

var SceneManager sceneManager

type sceneManager struct {
	CurrentSceneIDs []uint32
	Scenes          []Scene
}

func (m *sceneManager) NewScene() Scene {
	scene := Scene{
		Objects: make(map[string]*Object),
	}
	scene.ID = uint32(len(m.Scenes))

	m.Scenes = append(m.Scenes, scene)
	return scene
}

func (m *sceneManager) SetCurrentScenes(scenes ...uint32) {
	m.CurrentSceneIDs = scenes
}

func (m *sceneManager) Update(deltaTime float32) {
	for _, scene := range m.CurrentSceneIDs {
		m.Scenes[scene].update(deltaTime)
	}
}

type Scene struct {
	ID      uint32
	Objects map[string]*Object
}

func (s *Scene) update(deltaTime float32) {
	for _, obj := range s.Objects {
		obj.Scene = s

		if !obj.hasStarted {
			obj.start()
		} else {
			obj.update(deltaTime)
		}
	}
}

func (s *Scene) AddObject(obj *Object) {
	s.Objects[obj.Name] = obj
}

func (s *Scene) DeleteObject(name string) {
	delete(s.Objects, name)
}

func (s *Scene) SetObject(name string, dt Data, syss []System, tags ...string) {
	obj, exists := s.Objects[name]
	if !exists {
		log.Fatalf("Object '%s' does not exist", name)
		return
	}
	obj.Data = dt
	obj.Systems = syss
	obj.Tags.Set(tags...)
}

func (s *Scene) GetObject(name string) *Object {
	return s.Objects[name]
}

func (s *Scene) FindByTag(tag string) []*Object {
	var filteredObjects []*Object
	for _, obj := range s.Objects {
		if obj.Tags.Has(tag) {
			filteredObjects = append(filteredObjects, obj)
		}
	}
	return filteredObjects
}

// Object
type Object struct {
	Scene *Scene

	Name    string
	Data    Data
	Systems []System
	Tags    Tags

	hasStarted bool
}

func (s *Scene) NewObject(name string, dt Data, syss []System, tags ...string) *Object {
	tagSet := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		tagSet[tag] = struct{}{}
	}
	tagSet[name] = struct{}{}

	obj := &Object{
		Scene: s,

		Name:    name,
		Data:    dt,
		Systems: syss,
		Tags:    tagSet,
	}
	s.AddObject(obj)
	return obj
}

func (o *Object) Clone(dt ...Data) {
	for i, data := range dt {
		newData := make(Data)

		for key, value := range o.Data {
			newData[key] = value
		}

		for key, value := range data {
			newData.Set(key, value)
		}

		o.Scene.NewObject(
			o.Name+fmt.Sprint(i+1),
			newData,
			o.Systems,
			o.Name,
		)
	}
}

func (o *Object) start() {
	if !o.hasStarted {
		for _, sys := range o.Systems {
			sys.Start(o)
		}
	}
	o.hasStarted = true
}

func (o *Object) update(deltaTime float32) {
	if o.hasStarted {
		for _, sys := range o.Systems {
			sys.Update(o, deltaTime)
		}
	} else {
		log.Printf("Object '%s' has not been able to run its start function\n", o.Name)
		o.start()
	}
}

// Data
type Data map[string]interface{}

func (d *Data) Clone() Data {
	cloned := Data{}

	for key, value := range *d {
		cloned[key] = value
	}

	return cloned
}

func (d *Data) FindByType(tar interface{}) (string, bool) {
	for key, value := range *d {
		if reflect.TypeOf(value) == tar {
			return key, true
		}
	}
	return "", false
}

func (d *Data) FindByName(tar string) (interface{}, bool) {
	value, exists := (*d)[tar]
	return value, exists
}

func (d *Data) Set(key string, value interface{}) {
	(*d)[key] = value
}

func (d *Data) Get(key string) interface{} {
	if value, exists := (*d)[key]; exists {
		return value
	}
	log.Fatalf("Failed to get data '%s'", key)
	return nil
}

func (d *Data) Delete(key string) {
	delete(*d, key)
}

// Systems
type System interface {
	Start(obj *Object)
	Update(obj *Object, deltaTime float32)
}

// Tags
type Tags map[string]struct{}

func (t *Tags) Add(tags ...string) {
	for _, tag := range tags {
		(*t)[tag] = struct{}{}
	}
}

func (t *Tags) Delete(tags ...string) {
	for _, tag := range tags {
		delete((*t), tag)
	}
}

func (t *Tags) Has(tag string) bool {
	_, exists := (*t)[tag]
	return exists
}

func (t *Tags) Set(tags ...string) {
	(*t) = make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		(*t)[tag] = struct{}{}
	}
}
