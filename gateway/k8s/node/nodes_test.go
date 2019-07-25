package node

import (
	"fmt"
	"testing"

	"gitlab.yeepay.com/yce/nodeport/k8s"
)

const (
	// APIServer host in devk8s
	APIServer = `10.151.160.73:6443`

	// ca.cert
	Ca = `-----BEGIN CERTIFICATE-----
	MIIB4zCCAUSgAwIBAgIJAJaQW/tthlLzMAoGCCqGSM49BAMCMBgxFjAUBgNVBAMM
	DWt1YmVybmV0ZXMtY2EwHhcNMTkwMTA3MDM0NTQyWhcNMjkwMTA0MDM0NTQyWjAY
	MRYwFAYDVQQDDA1rdWJlcm5ldGVzLWNhMIGbMBAGByqGSM49AgEGBSuBBAAjA4GG
	AAQBPlIrgNHMSFBraKOloqJtCbAtHADfWyYCL6K3xoW5f4XJmkO8oirXec94eC0+
	J3whalwVq59mVRo9AQw4wemUUicAqrM8IT6uoK7hGoQmaI1AB7+q7sBILeXAXcTO
	uoj0EQiZqYUIk7lIRn70VNEGiS183Yjz6m7dthgBoAvFLc1U4GajNDAyMA8GA1Ud
	EwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgKkMA8GA1UdEQQIMAaHBAqXoEYwCgYI
	KoZIzj0EAwIDgYwAMIGIAkIAonQgpCKZLZUVIT7mNEC6lowe43Bpa0Kuf0QSSU+R
	Rl/XPgMAFiFIADFR2J+fzRkF42x6kY6ej8d93aXEuWnQU5ACQgHmIj2SfoWLzAE2
	eX3j4ZA7z85NOexNex/JyBmX4RtDuAQ3LKtp9fHXXsZuckTm7RzhqMso1hy1dpr3
	OincaCf0QA==
	-----END CERTIFICATE-----`

	// client.cert
	Cert = `-----BEGIN CERTIFICATE-----
	MIICCDCCAWugAwIBAgIJANRcFLAVuZmZMAoGCCqGSM49BAMCMBgxFjAUBgNVBAMM
	DWt1YmVybmV0ZXMtY2EwHhcNMTkwMTA3MDM0NTUwWhcNMjkwMTA0MDM0NTUwWjBB
	MSYwJAYDVQQDDB1rdWJlLWFwaXNlcnZlci1rdWJlbGV0LWNsaWVudDEXMBUGA1UE
	CgwOc3lzdGVtOm1hc3RlcnMwgZswEAYHKoZIzj0CAQYFK4EEACMDgYYABAFrTQ5c
	oR/BcRNXR8B/PVx+aO/yurnkAq2XbQaXK5afQsRLIwhDlFU9DGHVd0Bvefzg9Tf6
	nZS6G7UdRWZQSsFb5AD5zCl52PMmnSSAgSP+ULTYOP0c2MbYf1/QJ+SMdojDczOi
	WKLvSdKu/sAJubElgtXfxQwZEe/Jx0/dthDogvrqd6MyMDAwCQYDVR0TBAIwADAO
	BgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwIwCgYIKoZIzj0EAwID
	gYoAMIGGAkEkPyBejROq7GjBgenfnJTHwPUdty7ptDMgGfsrxdgppNyAem05B/pe
	Bjym2Z7x8sy9YxyDfNWW6PT9TTj/JdAS2gJBeFDX/c6j2iGbl1Ivu7QR4RsuunBV
	ARJ5+VsmcTTPPg4OXundR3B+GhbgaboaxxhkT07GH0gewSMssDwBRX2iLcA=
	-----END CERTIFICATE-----`

	// client.key
	Key = `-----BEGIN EC PRIVATE KEY-----
	MIHcAgEBBEIA9EnUn6nNzDU7ZS8gVTiNbfMckhGx3d070fdIwrlh/9zfmoEnu23a
	6/1L4Vo70eWGDJUjnFs8rCKaYFXBqASqJwmgBwYFK4EEACOhgYkDgYYABAFrTQ5c
	oR/BcRNXR8B/PVx+aO/yurnkAq2XbQaXK5afQsRLIwhDlFU9DGHVd0Bvefzg9Tf6
	nZS6G7UdRWZQSsFb5AD5zCl52PMmnSSAgSP+ULTYOP0c2MbYf1/QJ+SMdojDczOi
	WKLvSdKu/sAJubElgtXfxQwZEe/Jx0/dthDogvrqdw==
	-----END EC PRIVATE KEY-----`
)

func Test_GetRandomReadyNodeList(t *testing.T) {
	client, err := k8s.GetK8sClientFromFiles(APIServer, "ca.crt", "client.crt", "client.key")
	if err != nil {
		t.Errorf("k8s.GetK8sClient error: err=%s", err)
		return
	}

	list, err := GetRandomReadyNodeList(client)
	if err != nil {
		t.Errorf("node.GetRandomReadyNodeList error: err=%s", err)
		return
	}

	for _, v := range list {
		fmt.Printf("Node: name=%s, ip=%s\n", v.Name, v.IP)
	}
}
