package commponent

import (
	"fmt"
)

type Physics struct {
	// Fields specific to physics component
}

func (p *Physics) Update() {
	// Logic to update physics component
	fmt.Println("Updating physics component")
}

type Render struct {
	// Fields specific to render component
}

func (r *Render) Update() {
	// Logic to update render component
	fmt.Println("Updating render component")
}
