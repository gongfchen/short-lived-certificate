#!/bin/bash

if [ $# != 1 ]
then
    echo "Usage: $0 [domain_name]"
    exit 1
fi

source_dir="/home/nobellet/short-lived-cert/"

privkey_dir="${source_dir}/website-daemon/storage/"
privkey_filename="priv_key.pem"
privkey_path="${privkey_dir}/${privkey_filename}"

pubkey_dir="${source_dir}/CA/storage/domain-pubkey/$1/"
pubkey_filename="pub_key.pem"
pubkey_path="${pubkey_dir}/${pubkey_filename}"

for dir in $pubkey_dir $privkey_dir; do
    if [ ! -d ${dir} ]
    then
        echo "Making ${dir}"
        mkdir -p ${dir}
    fi
done

echo "Storing private key at ${privkey_path}"
echo "Storing public key at ${pubkey_path}"
openssl genrsa -out $privkey_path 2048
openssl rsa -in $privkey_path -pubout -out $pubkey_path
