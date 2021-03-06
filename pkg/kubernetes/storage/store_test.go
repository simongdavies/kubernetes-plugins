package storage

import (
	"os"
	"testing"

	"get.porter.sh/plugin/kubernetes/pkg/kubernetes/config"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
)

var logger hclog.Logger = hclog.New(&hclog.LoggerOptions{
	Name:   PluginInterface,
	Output: os.Stderr,
	Level:  hclog.Error})

func Test_NoNamespace(t *testing.T) {
	k8sConfig := config.Config{}
	store := NewStore(k8sConfig, logger)
	t.Run("Test No Namepsace", func(t *testing.T) {
		_, err := store.Read("type", "test")
		require.Error(t, err)
		require.EqualError(t, err, "open /var/run/secrets/kubernetes.io/serviceaccount/namespace: no such file or directory")
	})
}
