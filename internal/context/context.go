package context

type LauncherContext interface {
	AppVersion() string

	AppBuildNum() string

	AppName() string

	AppDirname() string

	AppHomeEnvVar() string

	UsernameEnvVar() string

	PasswordEnvVar() string

	AuthTokenEnvVar() string

	LogLevelEnvVar() string

	DebugFlagsEnvVar() string

	ConfigurationFileEnvVar() string

	RemoteConfigurationUrlEnvVar() string

	CmdPackageDirEnvVar() string
}
