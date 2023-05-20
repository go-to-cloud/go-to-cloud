package lang

type NodeJS struct {
	TplCommon
}

func (m *NodeJS) Sdk(env string) string {
	const nodejs = "node"
	switch env {
	case NodeJs16:
		return nodejs + ":16"
	case NodeJs14:
		return nodejs + ":14"
	case NodeJs18:
		return nodejs + ":18"
	}

	return nodejs + ":16"
}
