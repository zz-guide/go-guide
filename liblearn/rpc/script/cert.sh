#!/usr/bin/env bash

:<<BLOCK

1.签名证书
# 自签证书、服务加入证书验证

BLOCK


function generateX509(){
  #  "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs instead"
  # 需要修改GODEBUG
  openssl
  # 1.生成server.key私钥文件
  # Enter pass phrase for server.key:这一步会要求输入密码，123456
  genrsa -des3 -out server.key 2048
  # 2.创建证书请求文件server.csr,其中Common Name客户端会用到,xulei
  req -new -key server.key -out server.csr

  # 3.删除密码server_no_passwd.key
  rsa -in server.key -out server_no_passwd.key
  # 4.生成服务端证书server.crt
  x509 -req -days 3650 -in server.csr -signkey server_no_passwd.key -out server.crt
}

function generateX5092(){
    # ca
  	openssl genrsa -out ca.key 2048
  	openssl req -new -x509 -days 3650 \
  		-subj "/C=GB/L=China/O=gobook/CN=github.com" \
  		-key ca.key -out ca.crt

  	# server
    openssl genrsa \
      -out server.key 2048
    openssl req -new \
      -subj "/C=GB/L=China/O=server/CN=server.io" \
      -key server.key \
      -out server.csr
    openssl x509 -req -sha256 \
      -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
      -in server.csr \
      -out server.crt

    # client
    openssl genrsa \
      -out client.key 2048
    openssl req -new \
      -subj "/C=GB/L=China/O=client/CN=client.io" \
      -key client.key \
      -out client.csr
    openssl x509 -req -sha256 \
      -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
      -in client.csr \
      -out client.crt
}

function generateSan() {
#    C  => Country
#    ST => State
#    L  => City
#    O  => Organization
#    OU => Organization Unit
#    CN => Common Name (证书所请求的域名)
#    emailAddress => main administrative point of contact for the certificate

  # 生成默认 ca:
  openssl genrsa -out ca.key 2048
  openssl req -x509 -new -nodes -key ca.key -subj "/CN=xuleica.com" -days 5000 -out ca.crt
  # 生成证书
  openssl req -new -sha256 \
      -key ca.key \
      -subj "/C=CN/ST=Beijing/L=Beijing/O=UnitedStack/OU=Devops/CN=xulei1.com" \
      -reqexts SAN \
      -config <(cat /Users/xulei/jungle/golangworkspace/go-guide/liblearn/rpc/keys/openssl.cnf \
          <(printf "[SAN]nsubjectAltName=DNS:xulei1.com,DNS:xulei2.com")) \
      -out server.csr

   # 3.删除密码server_no_passwd.key
  openssl rsa -in ca.key -out server_no_passwd.key

  # 签名证书
  openssl x509 -req -days 365000 \
      -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
      -extfile <(printf "subjectAltName=DNS:xulei1.com,DNS:xulei2.com") \
      -out server.crt
}

generateSan