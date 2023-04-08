// Package internal contains environment variables
package internal

import (
	"fmt"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"
	"github.com/spf13/viper"
)

const (
	// PortVar is name of environment variable containing port
	PortVar     = "port"
	defaultPort = 8090

	// MetricsPortVar is name of environment variable containing port used for metrics
	MetricsPortVar     = "metricsport"
	defaultMetricsPort = 8095

	// LogLevelVar is name of environment variable containing loglevel
	LogLevelVar     = "loglevel"
	defaultLogLevel = "Warn"

	// TargetRepoVar is name of environment variable containing target repo URL
	TargetRepoVar = "targetrepo"

	// LocalDirVar is name of environment variable container local repo path
	LocalDirVar = "localdir"

	// KubeConfigPathVar is name of environment variable for kube config path
	KubeConfigPathVar = "kubeconfigpath"

	// NamespaceVar is name of environment variable for kube namespace
	NamespaceVar = "namespace"

	// InitialCloneVar is name of environment variable for the initial clone flag
	InitialCloneVar = "initialclone"

	// InitialBuildVar is name of environment variable for the initial clone flag
	InitialBuildVar = "initialbuild"

	// WebhookListenVar is name of environment variable for the webhook listen flag
	WebhookListenVar = "webhooklisten"

	// InitialBranchVar is the name environment variable for the webhook listen flag
	InitialBranchVar = "initialbranch"

	// StartupMode is the name of the environment variable for the startup mode
	StartupMode = "startupmode"

	// ServiceURLVar is the name of the environment variable for the service url
	ServiceURLVar = "serviceurl"

	// DBWriteEnabledVar name of var to enable db writing
	DbWriteEnabledVar = "dbwriteenabled"

	// DbCosmosUrlVar is the name of the url variable key
	DbCosmosUrlVar = "dbcosmosurl"

	// DbCosmosUrlVar is the variable name for the key to the cosmosdb
	DbCosmosKeyVar = "dbcosmoskey"
)

var (
	// Port is the port set in environment for serving http traffic
	Port int

	// MetricsPort is the port set in environment for metrics
	MetricsPort int

	// LogLevel is read from env
	LogLevel string

	// TargetRepo Url to target repo
	TargetRepo string

	// LocalDir absolute path to local dir
	LocalDir string

	// KubeConfigPath is the path to a valid KubeConfig file
	KubeConfigPath string

	// Namespace is the kubernetes namespace to create resources in
	Namespace string

	// InitialClone indicates if an initial clone should be performed at startup time
	InitialClone bool

	// InitialBuild indicates if the source should be built at startup time
	InitialBuild bool

	// WebhookListen indicates if the webhook should listener should be run at startup time
	WebhookListen bool

	// InitialBranch holds the branch that should be cloned on startup
	InitialBranch string

	// ServiceURL is the url of the service
	ServiceURL string

	// DbWriteEnabled new logs received should be written to db
	DbWriteEnabled bool

	// DbCosmosUrl is the url to the cosmosdb cloud instance
	DbCosmosUrl string

	// DbCosmosKey is the key to the cosmosdb cloud instance
	DbCosmosKey string
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(PortVar, defaultPort)
	viper.SetDefault(MetricsPortVar, defaultMetricsPort)
	viper.SetDefault(LogLevelVar, defaultLogLevel)
	viper.SetDefault(InitialBuildVar, true)
	viper.SetDefault(InitialCloneVar, true)
	viper.SetDefault(WebhookListenVar, true)
	viper.SetDefault(DbWriteEnabledVar, false)

	Port = viper.GetInt(PortVar)
	MetricsPort = viper.GetInt(MetricsPortVar)
	LogLevel = viper.GetString(LogLevelVar)
	TargetRepo = viper.GetString(TargetRepoVar)
	LocalDir = viper.GetString(LocalDirVar)
	KubeConfigPath = viper.GetString(KubeConfigPathVar)
	Namespace = viper.GetString(NamespaceVar)
	InitialClone = viper.GetBool(InitialCloneVar)
	InitialBuild = viper.GetBool(InitialBuildVar)
	WebhookListen = viper.GetBool(WebhookListenVar)
	InitialBranch = viper.GetString(InitialBranchVar)
	ServiceURL = viper.GetString(ServiceURLVar)
	DbWriteEnabled = viper.GetBool(DbWriteEnabledVar)
	DbCosmosUrl = viper.GetString(DbCosmosUrlVar)
	DbCosmosKey = viper.GetString(DbCosmosKeyVar)
}

// ValidateEnv validates environment variables
func ValidateEnv() error {
	clarkezoneLog.Debugf("ValidateEnv called")
	if Port == 0 {
		clarkezoneLog.Debugf("ValudateEnv() error port == 0")
		return fmt.Errorf("bad port")
	}
	if MetricsPort == 0 {
		clarkezoneLog.Debugf("ValudateEnv() error etricsport == 0")
		return fmt.Errorf("bad port")
	}
	if TargetRepo == "" {
		clarkezoneLog.Errorf("TargetRepo empty")
		return fmt.Errorf("TargetRepo empty")
	}
	if LocalDir == "" {
		clarkezoneLog.Errorf("LocalDir empty")
		return fmt.Errorf("LocalDir empty")
	}
	return nil
}
