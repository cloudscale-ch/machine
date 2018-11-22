package cloudscale
import (
	"testing"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/stretchr/testify/assert"
)

func TestSetConfigFromFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"cloudscale-token": "TOKEN",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)

	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)

	assert.Equal(t, driver.ResolveStorePath("id_rsa"), driver.GetSSHKeyPath())
}

func TestDefaultSSHUserAndPort(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"cloudscale-token": "TOKEN",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)

	sshPort, err := driver.GetSSHPort()
	assert.Equal(t, "root", driver.GetSSHUsername())
	assert.Equal(t, 22, sshPort)
	assert.NoError(t, err)
}

func TestCustomSSHUserAndPort(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"cloudscale-token": "TOKEN",
			"cloudscale-ssh-user":     "user",
			"cloudscale-ssh-port":     2222,
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)

	sshPort, err := driver.GetSSHPort()
	assert.Equal(t, "user", driver.GetSSHUsername())
	assert.Equal(t, 2222, sshPort)
	assert.NoError(t, err)
}

func TestVolumeSizeGB(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"cloudscale-token": "TOKEN",
			"cloudscale-volume-size-gb": 100,
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)
	assert.Equal(t, 100, driver.VolumeSizeGB)
}

func TestAntiAffinityWith(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"cloudscale-token": "TOKEN",
			"cloudscale-anti-affinity-with": "some uuid",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)
	assert.Equal(t, "some uuid", driver.AntiAffinityWith)
}
