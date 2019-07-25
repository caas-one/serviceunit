package setup

const (
	// DebugImage assiocation
	// DebugImageRepository constant
	DebugImageRepository = "artifact.paas.yp/yeepay-docker-dev-local/troubleshooting"

	// DebugImageTag constant
	DebugImageTag = "201809271441"

	// Resources assiocation
	// ResourceLimitsCPU constant
	ResourceLimitsCPU = "200m"
	// ResourceLimitsMemory constant
	ResourceLimitsMemory = "200M"

	// ResourceRequestsCPU constant
	ResourceRequestsCPU = "50m"
	// ResourceRequestsMemory constant
	ResourceRequestsMemory = "100M"

	// ImagePullPolicy constant
	ImagePullPolicy = "Always"

	// ImagePullSecrets constant
	ImagePullSecrets = "myregistrykey"
)

const (
	// CICD ApiServer
	APIServer = "10.151.33.87:6443"
)

const (
	Ca = `-----BEGIN CERTIFICATE-----
MIIB0jCCATOgAwIBAgIJAN9Tn8zjvt6XMAoGCCqGSM49BAMCMBgxFjAUBgNVBAMM
DWt1YmVybmV0ZXMtY2EwHhcNMTkwNTA2MDcwMzMzWhcNMjkwNTAzMDcwMzMzWjAY
MRYwFAYDVQQDDA1rdWJlcm5ldGVzLWNhMIGbMBAGByqGSM49AgEGBSuBBAAjA4GG
AAQA8Ag52q0dIvFRfCwSkdGyFYr6JnBzPj8d/iZoUyNiggcB3TaoWU0WpVWRc3Jj
zhVuMHu2zqPyMxPGWmQKtXTmCoMA8yc55kJoHdWH2XMPhP6N8VT5t4PL30NAJtPT
eVDXN/MY2smAn6dVx/1DLdYnpQw2WDhwKTvZ7DSrcl1ihj3oLUCjIzAhMA8GA1Ud
EwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgKkMAoGCCqGSM49BAMCA4GMADCBiAJC
AXnbqj6DuHoR3yTyiwfGfLXAQppokNZm8geklU8dYSQnY65JmXtZvxCyH8JOE62J
FoSkwxwx1TIbjhNSldih/zYrAkIA3K6635JH+bhX6+owb15oS8WLBGEg9AVdvU6G
i+SqVStiQKHIz1oydktXSI36lQB1Fwdj8UGzOHJNHMBJ2ferQ+s=
-----END CERTIFICATE-----`

	Cert = `-----BEGIN CERTIFICATE-----
MIIB/TCCAV6gAwIBAgIJAI3dihwTm+SgMAoGCCqGSM49BAMCMBgxFjAUBgNVBAMM
DWt1YmVybmV0ZXMtY2EwHhcNMTkwNTA2MDcwNTE0WhcNMjkwNTAzMDcwNTE0WjA0
MRkwFwYDVQQDDBBrdWJlcm5ldGVzLWFkbWluMRcwFQYDVQQKDA5zeXN0ZW06bWFz
dGVyczCBmzAQBgcqhkjOPQIBBgUrgQQAIwOBhgAEASe0wLw7psFSzCGbRs6BFz9q
xATXb57n8X/fIPONWckoGiLQhl2rzSsLT8crqkhlCWn46qG3gLuzHtYw8Rh2bQgd
AdH7iIlyHZsfU2LNxpcqWUGEb/Ar/i1TWPGarB5RzcyKzs/kdnUF1bqMdhGmd6Bh
hgfXQwxHWfAEPdKNGHOhIaqQozIwMDAJBgNVHRMEAjAAMA4GA1UdDwEB/wQEAwIF
oDATBgNVHSUEDDAKBggrBgEFBQcDAjAKBggqhkjOPQQDAgOBjAAwgYgCQgG4+Ho8
r4AHtYf+G9aeGNSopHg7lUhilsWVAl1tqoBclJKgo9W1dn3HIWv9wGgil3WG8weF
W9/o6mWE2KEmuLMbegJCAJAio0Suyb1f8+8M1kibLKxiw3DQp+6GPlUB/o8oG6Kj
YM9LxoWPe7emDtzTpdOz1eTB44bVet7pN0b0Fg6Nkpr2
-----END CERTIFICATE-----`

	Key = `-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBHKA61QxbUILlbftxL/xd718tnvbqMA3czvq1Y+pmJ39E/NNjewno
9MWsLq1NcIKyL65EuaPIf6c5rrtpEK17frigBwYFK4EEACOhgYkDgYYABAEntMC8
O6bBUswhm0bOgRc/asQE12+e5/F/3yDzjVnJKBoi0IZdq80rC0/HK6pIZQlp+Oqh
t4C7sx7WMPEYdm0IHQHR+4iJch2bH1NizcaXKllBhG/wK/4tU1jxmqweUc3Mis7P
5HZ1BdW6jHYRpnegYYYH10MMR1nwBD3SjRhzoSGqkA==
-----END EC PRIVATE KEY-----`
)
