#!/bin/bash

set -e

ca() {

  # Generate CA private key 
  # Attention: this is the key used to sign the certificate requests
  # anyone holding this can sign certificates on your behalf. 
  # So keep it in a safe place!
  # 根证书私钥
  echo ""
  echo "prepare to create ca.key ..."
  
  #If you want a non password protected key just remove the -des3 option
  openssl genrsa -des3 -out ca.key 4096 


  # Create and self sign the Root Certificate
  # Here we used our root key to create the root certificate that needs to be distributed in all the computers 
  # that have to trust us.
  echo ""
  echo "prepare to create ca.crt ..."
  openssl req -x509 -new -nodes -key ca.key -sha256 -days 10240 -out ca.crt

}

server() {

  # This procedure needs to be followed for each server/appliance 
  # that needs a trusted certificate from our CA

  # Create the certificate key
  # 服务器端私钥
  echo ""
  echo "prepare to create server.key ..."
  openssl genrsa -out server.key 2048 

  # Create the signing (csr)
  # The certificate signing request is where you specify the details for the certificate you want to generate. 
  # This request will be processed by the owner of the Root key (you in this case since you create it earlier) 
  # to generate the certificate.

  # Important: Please mind that while creating the signign request is important 
  # to specify the Common Name providing the IP address or domain name for the service, 
  # otherwise the certificate cannot be verified.
  echo ""
  echo "prepare to create server.csr ..."
  openssl req -new -key server.key -out server.csr

  # Verify the csr's content
  echo ""
  echo "prepare to verify the csr's content ..."
  openssl req -in server.csr -noout -text

  # Generate the certificate using the csr and key along with the CA Root key
  echo ""
  echo "prepare to create server.crt ..."
  openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 500 -sha256

  # Verify the certificate's content
  echo ""
  echo "prepare to veriry certificate's content ..."
  openssl x509 -in server.crt -text -noout 
}

client() {

  # This procedure needs to be followed for each server/appliance 
  # that needs a trusted certificate from our CA

  # Create the certificate key
  # 服务器端私钥
  echo ""
  echo "prepare to create client.key ..."
  openssl genrsa -out client.key 2048 

  # Create the signing (csr)
  # The certificate signing request is where you specify the details for the certificate you want to generate. 
  # This request will be processed by the owner of the Root key (you in this case since you create it earlier) 
  # to generate the certificate.

  # Important: Please mind that while creating the signign request is important 
  # to specify the Common Name providing the IP address or domain name for the service, 
  # otherwise the certificate cannot be verified.
  echo ""
  echo "prepare to create client.csr ..."
  openssl req -new -key client.key -out client.csr

  # Verify the csr's content
  echo ""
  echo "prepare to verify the csr's content ..."
  openssl req -in client.csr -noout -text

  # Generate the certificate using the csr and key along with the CA Root key
  echo ""
  echo "prepare to create client.crt ..."
  openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 500 -sha256

  # Verify the certificate's content
  echo ""
  echo "prepare to veriry certificate's content ..."
  openssl x509 -in client.crt -text -noout 
}


while getopts 'rcs' OPT; do
  case $OPT in
    r)
      ca
      ;;
    c)
      client
      ;;
    s)
      server
      ;;
    ?)
      echo "wrong options"
      ;;
  esac
done
