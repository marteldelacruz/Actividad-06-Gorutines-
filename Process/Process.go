package Process

type Processing interface {
	Proceso()
}

type Process struct {
	Processes []Processing
}

type P struct {
	id int
}
