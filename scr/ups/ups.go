package ups

import (
	"fmt"
	"log"
	"reflect"
)

var SceneManager sceneManager

type sceneManager struct {
	ActiveSceneIDs []uint32
	Scenes         []Scene
}

func (m *sceneManager) New() Scene {
	scene := Scene{
		Objects: make(map[string]*Object),
	}
	scene.ID = uint32(len(m.Scenes))

	m.Scenes = append(m.Scenes, scene)
	return scene
}

func (m *sceneManager) Add(scenes ...uint32) {
	for _, sceneID := range scenes {
		if !m.Contains(sceneID) {
			m.ActiveSceneIDs = append(m.ActiveSceneIDs, sceneID)
		}
	}
}

func (m *sceneManager) Remove(scenes ...uint32) {
	var newScenes []uint32
	for _, currentID := range m.ActiveSceneIDs {
		if !m.Contains(currentID) {
			newScenes = append(newScenes, currentID)
		}
	}
	m.ActiveSceneIDs = newScenes
}

func (m *sceneManager) Contains(sceneID uint32) bool {
	for _, id := range m.ActiveSceneIDs {
		if id == sceneID {
			return true
		}
	}
	return false
}

func (m *sceneManager) Set(scenes ...uint32) {
	m.ActiveSceneIDs = scenes
}

func (m *sceneManager) Update(deltaTime float32) {
	for _, scene := range m.ActiveSceneIDs {
		m.Scenes[scene].update(deltaTime)
	}
}

type Scene struct {
	ID      uint32
	Objects map[string]*Object
}

func (s *Scene) update(deltaTime float32) {
	for _, obj := range s.Objects {
		obj.update(deltaTime)
	}
}

func (s *Scene) Add(obj *Object) {
	obj.Scene = s
	s.Objects[obj.Name] = obj
}

func (s *Scene) Delete(name string) {
	delete(s.Objects, name)
}

func (s *Scene) Set(name string, dt Data, syss []System, tags ...string) {
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

func (s *Scene) GetByTag(tag string) []*Object {
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

	Name string
	Data Data

	Systems  []System
	Starters []starter
	Updaters []updater

	Tags Tags

	hasStarted bool
}

func (s *Scene) New(name string, dt Data, syss []System, tags ...string) *Object {
	tagSet := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		tagSet[tag] = struct{}{}
	}
	tagSet[name] = struct{}{}

	obj := &Object{
		Name:    name,
		Data:    dt,
		Systems: syss,
		Tags:    tagSet,
	}
	s.Add(obj)
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

		o.Scene.New(
			o.Name+fmt.Sprint(i+1),
			newData,
			o.Systems,
			o.Name,
		)
	}
}

func (o *Object) update(deltaTime float32) {
	for _, sys := range o.Systems {
		if !o.hasStarted {
			if starter, ok := sys.(starter); ok {
				starter.Start(o)
			}
		}

		if o.hasStarted {
			if updater, ok := sys.(updater); ok {
				updater.Update(o, deltaTime)
			}
		}
	}

	o.hasStarted = true
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

func (d *Data) GetByType(tar interface{}) (string, bool) {
	for key, value := range *d {
		if reflect.TypeOf(value) == tar {
			return key, true
		}
	}
	return "", false
}

func (d *Data) GetByName(tar string) (interface{}, bool) {
	value, exists := (*d)[tar]
	return value, exists
}

func (d *Data) Set(name string, value interface{}) {
	(*d)[name] = value
}

func (d *Data) Get(name string) interface{} {
	if value, exists := (*d)[name]; exists {
		return value
	}
	log.Fatalf("Failed to get data '%s'", name)
	return nil
}

func (d *Data) Delete(key string) {
	delete(*d, key)
}

// Systems
type System interface{}

type starter interface {
	Start(o *Object)
}

type updater interface {
	Update(o *Object, deltaTime float32)
}

// Tags
type Tags map[string]struct{}

func (t *Tags) Add(tags ...string) {
	for _, tag := range tags {
		(*t)[tag] = struct{}{}
	}
}

func (t *Tags) Remove(tags ...string) {
	for _, tag := range tags {
		delete((*t), tag)
	}
}

func (t *Tags) Set(tags ...string) {
	(*t) = make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		(*t)[tag] = struct{}{}
	}
}

func (t *Tags) Has(tag string) bool {
	_, exists := (*t)[tag]
	return exists
}
