package render

import (
	//nolint:gosec
	"crypto/md5"
	"fmt"

	"github.com/ghodss/yaml"
	mcfgv1 "github.com/uccps-samples/machine-config-operator/pkg/apis/machineconfiguration.uccp.io/v1"
)

var (
	// salt is 80 random bytes.
	// The salt was generated by `od -vAn -N80 -tu1 < /dev/urandom`. Do not change it.
	salt = []byte{
		16, 124, 206, 228, 139, 56, 175, 175, 79, 229, 134, 118, 157, 154, 211, 110,
		25, 93, 47, 253, 172, 106, 37, 7, 174, 13, 160, 185, 110, 17, 87, 52,
		219, 131, 12, 206, 218, 141, 116, 135, 188, 181, 192, 151, 233, 62, 126, 165,
		64, 83, 179, 119, 15, 168, 208, 197, 146, 107, 58, 227, 133, 188, 238, 26,
		33, 26, 235, 202, 32, 173, 31, 234, 41, 144, 148, 79, 6, 206, 23, 22,
	}
)

// Given a config from a pool, generate a name for the config
// of the form rendered-<poolname>-<hash>
func getMachineConfigHashedName(pool *mcfgv1.MachineConfigPool, config *mcfgv1.MachineConfig) (string, error) {
	if config == nil {
		return "", fmt.Errorf("empty machineconfig object")
	}

	data, err := yaml.Marshal(config.Spec)
	if err != nil {
		return "", err
	}

	h, err := hashData(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("rendered-%s-%x", pool.GetName(), h), nil
}

func hashData(data []byte) ([]byte, error) {
	//nolint:gosec
	hasher := md5.New()
	if _, err := hasher.Write(salt); err != nil {
		return nil, fmt.Errorf("error computing hash: %v", err)
	}
	if _, err := hasher.Write(data); err != nil {
		return nil, fmt.Errorf("error computing hash: %v", err)
	}
	return hasher.Sum(nil), nil
}
