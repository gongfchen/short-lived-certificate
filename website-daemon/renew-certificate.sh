#!/bin/bash

domain_name=$(dig +short myip.opendns.com @resolver1.opendns.com)

source_dir="/home/nobellet/short-lived-cert/"

website_daemon_dir="${source_dir}/website-daemon/"
nginx_dir="${source_dir}/web-server/"
nginx_conf_filename="nginx.conf"
nginx_conf_path="${nginx_dir}/${nginx_conf_filename}"

# Update certificate path in web server config
python3 ${website_daemon_dir}/update-nginx-config.py

# Reload the web server to pick up the new config
$HOME/.local/src/nginx-1.8.1/objs/nginx -s reload -c ${nginx_conf_path} -p ${nginx_dir}
