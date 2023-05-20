package lang

type Default struct {
	TplCommon
}

func (m *Default) Sdk(env string) string {
	return env
}
