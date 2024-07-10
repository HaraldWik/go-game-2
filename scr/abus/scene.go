package abus //Array based update system ABUS

type sceneManager struct {
	Scenes   []Scene
	CurScene Scene
}

var SceneManager sceneManager

func (manager sceneManager) GetCurrentScene() Scene {
	return SceneManager.CurScene
}

type Scene struct {
	Objs   []Obj
	CurObj Obj
}

func (scene *Scene) GetCurrentObject() Obj {
	SceneManager.CurScene = *scene
	return scene.CurObj
}

func (manager sceneManager) NewScene() Scene {
	manager.Scenes = append(manager.Scenes, Scene{})
	return manager.Scenes[len(manager.Scenes)-1]
}

func (scene *Scene) Update() {
	SceneManager.CurScene = *scene
	for i := range scene.Objs {
		scene.CurObj = scene.Objs[i]
		scene.CurObj.Update()
	}
}

func (scene *Scene) Instance(obj Obj) Obj {
	return scene.Create(obj.Mods...)
}

func (scene *Scene) Create(mods ...Mod) Obj {
	scene.Objs = append(scene.Objs, Obj{ID: uint32(len(scene.Objs)), Mods: mods})
	return scene.Objs[len(scene.Objs)-1]
}

func (scene *Scene) Delete(id uint32) {
	for i, obj := range scene.Objs {
		if obj.ID == id {
			scene.Objs = append(scene.Objs[:i], scene.Objs[i+1:]...)
			return
		}
	}
}
