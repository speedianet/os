package entity

import "github.com/speedianet/sam/src/domain/valueObject"

type Ssl struct {
	Id               valueObject.SslId            `json:"id"`
	Hostname         valueObject.Fqdn             `json:"hostname"`
	IssuedAt         *valueObject.UnixTime        `json:"issuedAt,omitempty"`
	ExpireAt         *valueObject.UnixTime        `json:"expireAt,omitempty"`
	Certificate      valueObject.SslCertificate   `json:"certificate"`
	Key              valueObject.SslPrivateKey    `json:"key"`
	RootCertificates []valueObject.SslCertificate `json:"rootCertificates"`
}

func NewSsl(
	id valueObject.SslId,
	hostname valueObject.Fqdn,
	issuedAt *valueObject.UnixTime,
	expireAt *valueObject.UnixTime,
	certificate valueObject.SslCertificate,
	key valueObject.SslPrivateKey,
	rootCertificates []valueObject.SslCertificate,
) Ssl {
	return Ssl{
		Id:               id,
		Hostname:         hostname,
		IssuedAt:         issuedAt,
		ExpireAt:         expireAt,
		Certificate:      certificate,
		Key:              key,
		RootCertificates: rootCertificates,
	}
}
