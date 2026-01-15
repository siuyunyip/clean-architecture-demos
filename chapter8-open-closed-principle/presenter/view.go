package presenter

type ScreenRenderer interface {
	Render(model *ScreenViewModel) error
}
