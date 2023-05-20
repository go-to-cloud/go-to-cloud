package lang

type Maven struct {
	TplCommon
}

func (m *Maven) Sdk(_ string) string {
	const maven = "maven"

	_ = MavenJDK8
	return maven + ":3.8.6-jdk-8"
}
