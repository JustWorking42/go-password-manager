rm *.pem

mkdir ./server
mkdir ./client
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=RU/ST=SPB/L=SPB/O=Yandex/OU=Education/CN=main"

echo "CA's self-signed certificate"

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout ./server/server-key.pem -out ./server/server-req.pem -subj "/C=RU/ST=SPB/L=SPB/O=Yandex/OU=Education/CN=server"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in ./server/server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out ./server/server-cert.pem -extfile server-ext.cnf -extensions v3_req

echo "Server's signed certificate"


# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout ./client/client-key.pem -out ./client/client-req.pem -subj "/C=RU/ST=SPB/L=SPB/O=Yandex/OU=Education/CN=client"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in ./client/client-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out ./client/client-cert.pem -extfile client-ext.cnf -extensions v3_req

echo "Client's signed certificate"
