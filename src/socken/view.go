package socken

type View interface {
	Flash(msg string)
}

type DummyView struct{}

func (d DummyView) Flash(_ string) {}
