package entity

import (
	"testing"

	"github.com/goinfinite/os/src/domain/valueObject"
)

func TestSslCertificate(t *testing.T) {
	t.Run("ValidSslCertificate (Root)", func(t *testing.T) {
		rootSslCrtContentStr := `-----BEGIN CERTIFICATE-----
MIIFFjCCAv6gAwIBAgIRAJErCErPDBinU/bWLiWnX1owDQYJKoZIhvcNAQELBQAw
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMjAwOTA0MDAwMDAw
WhcNMjUwOTE1MTYwMDAwWjAyMQswCQYDVQQGEwJVUzEWMBQGA1UEChMNTGV0J3Mg
RW5jcnlwdDELMAkGA1UEAxMCUjMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQC7AhUozPaglNMPEuyNVZLD+ILxmaZ6QoinXSaqtSu5xUyxr45r+XXIo9cP
R5QUVTVXjJ6oojkZ9YI8QqlObvU7wy7bjcCwXPNZOOftz2nwWgsbvsCUJCWH+jdx
sxPnHKzhm+/b5DtFUkWWqcFTzjTIUu61ru2P3mBw4qVUq7ZtDpelQDRrK9O8Zutm
NHz6a4uPVymZ+DAXXbpyb/uBxa3Shlg9F8fnCbvxK/eG3MHacV3URuPMrSXBiLxg
Z3Vms/EY96Jc5lP/Ooi2R6X/ExjqmAl3P51T+c8B5fWmcBcUr2Ok/5mzk53cU6cG
/kiFHaFpriV1uxPMUgP17VGhi9sVAgMBAAGjggEIMIIBBDAOBgNVHQ8BAf8EBAMC
AYYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBIGA1UdEwEB/wQIMAYB
Af8CAQAwHQYDVR0OBBYEFBQusxe3WFbLrlAJQOYfr52LFMLGMB8GA1UdIwQYMBaA
FHm0WeZ7tuXkAXOACIjIGlj26ZtuMDIGCCsGAQUFBwEBBCYwJDAiBggrBgEFBQcw
AoYWaHR0cDovL3gxLmkubGVuY3Iub3JnLzAnBgNVHR8EIDAeMBygGqAYhhZodHRw
Oi8veDEuYy5sZW5jci5vcmcvMCIGA1UdIAQbMBkwCAYGZ4EMAQIBMA0GCysGAQQB
gt8TAQEBMA0GCSqGSIb3DQEBCwUAA4ICAQCFyk5HPqP3hUSFvNVneLKYY611TR6W
PTNlclQtgaDqw+34IL9fzLdwALduO/ZelN7kIJ+m74uyA+eitRY8kc607TkC53wl
ikfmZW4/RvTZ8M6UK+5UzhK8jCdLuMGYL6KvzXGRSgi3yLgjewQtCPkIVz6D2QQz
CkcheAmCJ8MqyJu5zlzyZMjAvnnAT45tRAxekrsu94sQ4egdRCnbWSDtY7kh+BIm
lJNXoB1lBMEKIq4QDUOXoRgffuDghje1WrG9ML+Hbisq/yFOGwXD9RiX8F6sw6W4
avAuvDszue5L3sz85K+EC4Y/wFVDNvZo4TYXao6Z0f+lQKc0t8DQYzk1OXVu8rp2
yJMC6alLbBfODALZvYH7n7do1AZls4I9d1P4jnkDrQoxB3UqQ9hVl3LEKQ73xF1O
yK5GhDDX8oVfGKF5u+decIsH4YaTw7mP3GFxJSqv3+0lUFJoi5Lc5da149p90Ids
hCExroL1+7mryIkXPeFM5TgO9r0rvZaBFOvV2z0gp35Z0+L4WPlbuEjN/lxPFin+
HlUjr8gRsI3qfJOQFy/9rKIJR0Y/8Omwt/8oTWgy1mdeHmmjk7j1nYsvC9JSQ6Zv
MldlTTKB3zhThV1+XWYp6rjd5JW1zbVWEkLNxE7GJThEUG3szgBVGP7pSWTUTsqX
nLRbwHOoq7hHwg==
-----END CERTIFICATE-----`
		validSslCrtContent, _ := valueObject.NewSslCertificateContent(rootSslCrtContentStr)

		_, err := NewSslCertificate(validSslCrtContent)
		if err != nil {
			t.Errorf("Expected no error for Root SSL certificate, got '%s'", err.Error())
		}
	})

	t.Run("ValidSslCertificate (Let's Encrypt)", func(t *testing.T) {
		letsEncryptSslCrtContent := `-----BEGIN CERTIFICATE-----
MIIE9DCCA9ygAwIBAgISAxXgq9lFBTcrWpYfYpeAUa1XMA0GCSqGSIb3DQEBCwUA
MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
EwJSMzAeFw0yNDAyMDQxMzU1NTNaFw0yNDA1MDQxMzU1NTJaMBYxFDASBgNVBAMT
C3NwZWVkaWEubmV0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqS8h
K15VfvROwLk1zranIayaWytBuJukBv0mT72S4gi1JGh564pUOP06e0zNlyaCLmHh
pNgyHN7Imp6uhmlqsif9l0KXJspQ2WwCMndqPkuZoY/5myGPUsB/yNsQtKs4+BUO
DnmoNKIfY8ckd+w+tRjw0gwVx6CnRXlvU8mD4d/0TXV+nY7+w8gn1GOc9tmG41Sy
6hDrLZ4urKnNQujyafvIZA5Foo8flDg7LyUshczt2/wBTwvkR9CkBcstKN7qC//K
uTCKWB8ovLxjxB0btpHrci9QAql26gNcGzWGSJLMP3+jlMMgkpMDJfxm23XQLq/s
TRB8NyJq/or9pp4jawIDAQABo4ICHjCCAhowDgYDVR0PAQH/BAQDAgWgMB0GA1Ud
JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQW
BBQEwEoC4meZAVmjb2gGrQ3sXei2/TAfBgNVHSMEGDAWgBQULrMXt1hWy65QCUDm
H6+dixTCxjBVBggrBgEFBQcBAQRJMEcwIQYIKwYBBQUHMAGGFWh0dHA6Ly9yMy5v
LmxlbmNyLm9yZzAiBggrBgEFBQcwAoYWaHR0cDovL3IzLmkubGVuY3Iub3JnLzAn
BgNVHREEIDAeggtzcGVlZGlhLm5ldIIPd3d3LnNwZWVkaWEubmV0MBMGA1UdIAQM
MAowCAYGZ4EMAQIBMIIBBAYKKwYBBAHWeQIEAgSB9QSB8gDwAHYASLDja9qmRzQP
5WoC+p0w6xxSActW3SyB2bu/qznYhHMAAAGNdJ5hRgAABAMARzBFAiEA1LATXDGG
KaHW6dKgdWjsPVJgDIliYfgE97BLWZQ4w4oCIEU/ptSab0O1sx9c1WYaiIu8YOYM
DNXKmuSbd2fst9SKAHYAO1N3dT4tuYBOizBbBv5AO2fYT8P0x70ADS1yb+H61BcA
AAGNdJ5hTQAABAMARzBFAiEA9xze/WAMNSQ3u+mgb3MHRAGXvJhDpSdft+qlvbFg
XnUCIFZXfubeel5UQRwH+B6Dg8vQ16ha6pb2r/ut40DIzVEpMA0GCSqGSIb3DQEB
CwUAA4IBAQAO2TeYxfRngEQk2aa+6XJT6GiXhYD9T35bxhD/ZKWjMpBK5BenV5X3
cJsKSttMOIjgnf0PY3JKvGADATYOV7/vxKkSbm1WJhwSRD/29AFrZJVfkahG3IgZ
0pPa8Y+3swTqY/ZmwWB6KswsW0zS1W7469gLJ04CNIIAXW7FCTXWrJz30Vt8rZRx
QI8kt4rcJUZ7r9cSQilkDCgGxalMt6GPAGnbmqdkbrDq39swMVapwhjEAa00QSW+
nq8q70Xi28NCV/Jepz6Ud1GM2lGfnsSZptXviFSD/43Tt9jmRYc4+fyQUZmoUeKg
qVJbJOmn5TZEn+aQetPL0lrd4Ke+zZPA
-----END CERTIFICATE-----`
		validSslCrtContent, _ := valueObject.NewSslCertificateContent(letsEncryptSslCrtContent)

		_, err := NewSslCertificate(validSslCrtContent)
		if err != nil {
			t.Errorf("Expected no error for Let's Encrypt SSL certificate, got '%s'", err.Error())
		}
	})

	t.Run("ValidSslCertificate (Self-signed)", func(t *testing.T) {
		selfSignedSslCrtContent := `-----BEGIN CERTIFICATE-----
MIIDmTCCAoGgAwIBAgIUIS59eqnSjoGxRRbFkfwTMVvFyZYwDQYJKoZIhvcNAQEL
BQAwXDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAcM
Ckxvc0FuZ2VsZXMxDTALBgNVBAoMBEFjbWUxFDASBgNVBAMMC3NwZWVkaWEubmV0
MB4XDTI0MDMyMDAzMjkwMFoXDTI1MDMyMDAzMjkwMFowXDELMAkGA1UEBhMCVVMx
EzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAcMCkxvc0FuZ2VsZXMxDTALBgNV
BAoMBEFjbWUxFDASBgNVBAMMC3NwZWVkaWEubmV0MIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEA1MkNo2kVjN+UerukCe5luoH0n10Odi9zzRXxFZtdGHBM
83t6mcgWhH0EjOBfWNpNAvX3jt6mBh4Mf2zKeAmARfuAFNn0lIx8acyX3eST442o
c51b2Lk0ORVeHjoxKubI3acJrnNNOI7Mc9WHPlkSoiXgu+eK3JoeEVmUN3hVkQEz
x27H5UgjQ543qm+Xvf0DFdr1COHS/i19cvuqjSChN5jOjOkWStgvbQb9yI+fKMBI
3iT+Ls9X/K+wdUws17+pKc7U+5eeDJi1GpC85chUbEmg+pTK1eeIhbRnE8phPcqb
VK4BIdm/w8Xq2/XeRSW9cmN6Nen+GRoJtRytQLSVIQIDAQABo1MwUTAdBgNVHQ4E
FgQUXLjzpubGYM3Qk8mWJS84sqdSZrMwHwYDVR0jBBgwFoAUXLjzpubGYM3Qk8mW
JS84sqdSZrMwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAp5Ae
ZbjA0Kjg8zyPVS62rOFtJoWv+MEKDoZHkklvkKprjCKcWKe7BOo2LZtEhH4HyFWv
xfx/jOZsvH+QZyCdLHXY4OiJDAbRKy1bcRBn/+j/WfwlUOoI2VavFa/HSerHnb+Y
4DsPlNm3QeCEW9VCtLKQK7olVL461DesJ4gac+RLnYfMiTosaHkC4terIuPQzeVy
L/kj32P0VIrkoAM54vbMoqKG0y0pgqO2W4bFRypypCgPbFOlx4s28j/ZUdvgpNNr
Ga23e8x0F9dhysmCyc1S13JptCjQ5Z7RNeU354dHUR8alE0CEvXSwFZPi220OKfj
hxh2OsC7/1mvbLs/WA==
-----END CERTIFICATE-----`
		validSslCrtContent, _ := valueObject.NewSslCertificateContent(selfSignedSslCrtContent)

		_, err := NewSslCertificate(validSslCrtContent)
		if err != nil {
			t.Errorf("Expected no error for Self-signed SSL certificate, got '%s'", err.Error())
		}
	})

	t.Run("ValidSslCertificate (Google)", func(t *testing.T) {
		googleSslCrtContent := `-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
-----END CERTIFICATE-----`
		validSslCrtContent, _ := valueObject.NewSslCertificateContent(googleSslCrtContent)

		_, err := NewSslCertificate(validSslCrtContent)
		if err != nil {
			t.Errorf("Expected no error for Google SSL certificate, got '%s'", err.Error())
		}
	})

	t.Run("InvalidSslCertificate", func(t *testing.T) {
		invalidSslCertContentsStr := []string{
			`-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
-----END CERTIFICATE-----`,
			`MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==`,
		}
		for invalidSslCrtContentStrIndex, invalidSslCertContentStr := range invalidSslCertContentsStr {
			invalidSslCrtContent, _ := valueObject.NewSslCertificateContent(invalidSslCertContentStr)
			_, err := NewSslCertificate(invalidSslCrtContent)
			if err == nil {
				t.Errorf(
					"Expected error for '%d' invalid SSL certificate index, got nil",
					invalidSslCrtContentStrIndex,
				)
			}
		}
	})
}
