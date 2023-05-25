package ksherPayment

import (
	"github.com/ksher-api/ksher-sdk/go/src/KsherGO"
)

func GenerateSign(privateKeyPath string) (string, error) {
	privateKeyData := KsherGO.ReadPrivateKeyFromPath(privateKeyPath)
	return KsherGO.KsherSign(nil, privateKeyData)
}

func VerifySig(response KsherGO.KsherResp) error {
	publicKey := []byte(`
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAL7955OCuN4I8eYNL/mixZWIXIgCvIVE
ivlxqdpiHPcOLdQ2RPSx/pORpsUu/E9wz0mYS2PY7hNc2mBgBOQT+wUCAwEAAQ==
-----END PUBLIC KEY-----`)
	return KsherGO.KsherVerify(response, publicKey)
}
