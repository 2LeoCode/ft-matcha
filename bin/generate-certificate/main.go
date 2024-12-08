package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"os"
)

func generatePrivateKey() (privateKey *ecdsa.PrivateKey, err error) {
	return ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
}

func encodePrivateKey(key *ecdsa.PrivateKey) (keyPem []byte, err error) {
	derKey, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}

	block := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derKey,
	}

	return pem.EncodeToMemory(&block), nil
}

func generateCertificateRequest(privateKey *ecdsa.PrivateKey) (csr []byte, err error) {
	return x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            []string{"France"},
			Organization:       []string{"42"},
			OrganizationalUnit: []string{"Development"},
			CommonName:         "Matcha",
			Province:           []string{"Paris"},
		},
	}, privateKey)
}

func signCertificate(csr []byte, privateKey *ecdsa.PrivateKey) (certificate []byte, err error) {
	hash := crypto.SHA512.New()
	hash.Write(csr)
	return ecdsa.SignASN1(rand.Reader, privateKey, hash.Sum([]byte{}))
}

func encodeCertificate(cert []byte) (certPem []byte, err error) {
	block := pem.Block{
		Type:  "EC CERTIFICATE",
		Bytes: cert,
	}

	return pem.EncodeToMemory(&block), nil
}

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if privateKey, err := generatePrivateKey(); err != nil {
	} else if csr, err := generateCertificateRequest(privateKey); err != nil {
	} else if cert, err := signCertificate(csr, privateKey); err != nil {
	} else if privateKeyPem, err := encodePrivateKey(privateKey); err != nil {
	} else if certPem, err := encodeCertificate(cert); err != nil {
	} else {
		os.WriteFile("certs/private-key.pem", privateKeyPem, 0o666)
		os.WriteFile("certs/cert.pem", certPem, 0o666)
	}
}
