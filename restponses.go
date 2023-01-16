package restponses

type method int

const (
	Get method = iota
	Post
	Delete
	Put
)

func Response100Continue(requestMethod method) string {
	return "si"
}

func Response101SwitchingProtocols() {

}

func Response102Processing() {

}
