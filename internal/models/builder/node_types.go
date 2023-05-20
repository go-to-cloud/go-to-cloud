package builder

type NodeType int

// 1：k8s；2：docker；3: windows；4：linux；5：macos
const (
	K8s     NodeType = 1
	Docker  NodeType = 2
	Windows NodeType = 3
	Linux   NodeType = 4
	MacOs   NodeType = 5
)
