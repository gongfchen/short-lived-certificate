package main

// Generate domain certificate
// Generate root certificte
// Generate precertificate

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

func (ca Ca) gen_enc_cert(pubkey *rsa.PublicKey, count int, domain_name string, hash_list *[][]byte, hashlist_string *[num_of_cert]string) {
	// Prepare directory
	certpath := "../storage/domain-certificates/" + domain_name
	enc_certpath := "../../CA-middle-daemon-storage/Encrypted Certificates/" + domain_name
	if _, err := os.Stat(certpath); os.IsNotExist(err) {
		os.MkdirAll(certpath, 0744)
	}
	if _, err := os.Stat(enc_certpath); os.IsNotExist(err) {
		os.MkdirAll(enc_certpath, 0744)
	}

	// Get Root certificate
	cert_auth, err := x509.ParseCertificate(ca.certificate.Certificate[0])
	check(err)

	// Generate certificate
	certificate := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject:      pkix.Name{Organization: []string{domain_name}},
		NotBefore:    time.Now().AddDate(0, 0, count),
		NotAfter:     time.Now().AddDate(0, 0, (count + 4)),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		DNSNames:     []string{domain_name},
	}

	if ip := net.ParseIP(domain_name); ip != nil {
		certificate.IPAddresses = append(certificate.IPAddresses, ip)
	} else {
		certificate.DNSNames = append(certificate.DNSNames, domain_name)
	}

	crt, err := x509.CreateCertificate(rand.Reader,
		&certificate,
		cert_auth,
		pubkey,
		ca.certificate.PrivateKey)
	check(err)

	// Export original certificate
	var certOut bytes.Buffer
	certfilepath := certpath + "/cert" + strconv.Itoa(count) + ".pem"
	pem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: crt})
	f, err := os.Create(certfilepath)
	check(err)

	l, err := f.Write(certOut.Bytes()) // Certificate written in byte
	check(err)
	fmt.Println(l, "bytes written successfully --- Original certificate in CA's storage")
	err = f.Close()
	check(err)

	// Encrypt certificate and export
	salt := domain_name + time.Time.String(ca.first_start_time.AddDate(0, 0, count))
	encrypt_key := kdf(ca.master_key, salt)
	enc_cert := encrypt(string(certOut.Bytes()), hex.EncodeToString(encrypt_key[:]))
	enc_certfilepath := enc_certpath + "/enc_cert" + strconv.Itoa(count) + ".txt"
	f, err = os.Create(enc_certfilepath)
	check(err)

	l, err = f.WriteString(enc_cert) // Encrypted certificate written in string
	check(err)

	fmt.Println(l, "bytes written successfully --- Encrypted certificate in shared storage")
	err = f.Close()
	check(err)

	// Add to hash list
	hash_cert := hash(string(certOut.Bytes()))
	hashlist_string[count] = hex.EncodeToString(hash_cert[:]) // Hash list written in string
	*hash_list = append((*hash_list)[:], ByteSlice(hash_cert))

}

func ByteSlice(b []byte) []byte { return b }

func genPreCert(domain_name string, pubkey *rsa.PublicKey) {
	merkle_root_value, err := ioutil.ReadFile("../storage/merkle-roots/" + domain_name + "/merkleroot.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(merkle_root_value))

	certpath := "../storage/precertificate/" + domain_name
	if _, err := os.Stat(certpath); os.IsNotExist(err) {
		os.MkdirAll(certpath, 0744)
	}

	sct_value, err := ioutil.ReadFile("../storage/sct/" + domain_name + "/sct.pem")
	if err != nil {
		log.Fatal(err)
	}

	sharedcertpath := "../../CA-middle-daemon-storage/precertificate/" + domain_name
	if _, err := os.Stat(sharedcertpath); os.IsNotExist(err) {
		os.Mkdir(sharedcertpath, 0744)
	}

	// Load CA
	catls, err := tls.LoadX509KeyPair("../storage/root-certificate/ca_cert.pem", "../storage/root-certificate/ca_key.pem")
	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}

	var extension = make([]pkix.Extension, 3)
	poison_extension := pkix.Extension{
		Id:       asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 3},
		Critical: true,
		Value:    []byte{5, 0},
	}
	extension[0] = poison_extension

	merkle_root := pkix.Extension{
		Id:       asn1.ObjectIdentifier{1, 2, 3, 4, 1},
		Critical: false,
		// Value:    []byte{0xe},
		Value: merkle_root_value,
	}
	extension[1] = merkle_root

	sct_extension := pkix.Extension{
		Id:       asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 2},
		Critical: false,
		Value:    sct_value,
	}
	extension[2] = sct_extension

	template := x509.Certificate{
		SerialNumber:    big.NewInt(time.Now().Unix()),
		Subject:         pkix.Name{Organization: []string{domain_name}},
		NotBefore:       time.Now().Add(time.Hour * 24 * 0),
		NotAfter:        time.Now().Add(time.Hour * 24 * 4),
		KeyUsage:        x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:     []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		DNSNames:        []string{domain_name},
		ExtraExtensions: extension,
	}

	if ip := net.ParseIP(domain_name); ip != nil {
		template.IPAddresses = append(template.IPAddresses, ip)
	} else {
		template.DNSNames = append(template.DNSNames, domain_name)
	}

	// privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	panic(err)
	// }
	crt, err := x509.CreateCertificate(rand.Reader,
		&template,
		ca,
		pubkey,
		catls.PrivateKey)
	if err != nil {
		panic(err)
	}
	var certOut bytes.Buffer
	pem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: crt})
	// pem.Encode(&keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)})
	// ioutil.WriteFile("./precert/key.pem", keyOut.Bytes(), 0644)
	ioutil.WriteFile(certpath+"/precert.pem", certOut.Bytes(), 0744)
	ioutil.WriteFile(sharedcertpath+"/precert.pem", certOut.Bytes(), 0744)

}
