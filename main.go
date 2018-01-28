package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/darmanuals/blockchain/controllers"
	"github.com/darmanuals/blockchain/db"
	"crypto/rsa"
	"crypto/rand"
	"log"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func main() {
	db.Load()


	privateKey, err := rsa.GenerateKey(rand.Reader, 512)

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := []byte("text")
	newhash := crypto.SHA256

	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	pssh2 := newhash.New()
	pssh2.Write([]byte("no"))
	h2 := pssh2.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, newhash, hashed, &opts)
	log.Printf("pub: %x | %v", privateKey.D, err)

	err = rsa.VerifyPSS(&privateKey.PublicKey, newhash, h2, signature, &opts)
	log.Println(err)

	//===== Pub key
	PubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		// do something about it
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})

	ioutil.WriteFile("key.pub", pubBytes, 0644)
	//====== Private key
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type: "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	ioutil.WriteFile("key.pem", pemdata, 0644)
	// end

	router := mux.NewRouter()
	router.HandleFunc("/blockchain/get_blocks/{count}", controllers.GetBlocks).Methods("GET")
	router.HandleFunc("/blockchain/receive_update", controllers.Update).Methods("POST")
	router.HandleFunc("/management/add_transaction", controllers.AddTransaction).Methods("POST")
	router.HandleFunc("/management/add_link", controllers.AddLink).Methods("POST")
	router.HandleFunc("/management/status", controllers.GetStatus).Methods("GET")
	router.HandleFunc("/management/sync", controllers.Sync).Methods("GET")

	http.ListenAndServe(":3000", router)
}
