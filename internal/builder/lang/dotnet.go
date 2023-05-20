package lang

type DotNet struct {
	TplCommon
}

func (m *DotNet) Sdk(env string) string {
	const dotnet = "mcr.microsoft.com/dotnet/sdk"
	switch env {
	case DotNet6:
		return dotnet + ":6.0"
	case DotNet5:
		return dotnet + ":5.0"
	case DotNet7:
		return dotnet + ":7.0"
	case DotNet3:
		return dotnet + ":3.1"
	}

	return dotnet + ":6.0"
}
