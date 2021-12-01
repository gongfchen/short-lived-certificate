import crossplane
import os
import re
import socket

def should_extract(key, value, directive):
    return key == "directive" and value == directive

"""Extract nested values from a JSON tree."""
"""Modified from https://hackersandslackers.com/extract-data-from-complex-json-python/"""
def json_extract(obj, directive):

    def extract(obj, directive):
        """Recursively search for value of key in JSON tree."""
        if isinstance(obj, dict):
            for k, v in obj.items():
                if isinstance(v, (dict, list)):
                    value = extract(v, directive)
                    if value:
                        return value
                elif should_extract(k, v, directive):
                    return obj.get("args")
        elif isinstance(obj, list):
            for item in obj:
                value = extract(item, directive)
                if value:
                    return value
        return None

    value = extract(obj, directive)
    return value

def should_replace(key, structure, directive):
    return key == "args" and structure.get("directive", None) == directive

def replace_if_necessary(key, value, structure, directive, args):
    if should_replace(key, structure, directive):
        return [args]
    else:
        return nested_replace(value, directive, args)

"""Modified from https://stackoverflow.com/questions/50631393/python-replace-values-in-unknown-structure-json-file"""
def nested_replace(structure, directive, args):
    if type(structure) == list:
        return [nested_replace(item, directive, args) for item in structure]

    if type(structure) == dict:
        return dict(map(lambda x: (x[0], replace_if_necessary(x[0], x[1], structure, directive, args)), structure.items()))

    return structure


hostname = socket.gethostname()
IPAddr = socket.gethostbyname(hostname)

source_dir = "/home/nobellet/short-lived-cert"

nginx_dir = os.path.join(source_dir, "web-server")
nginx_conf_path = os.path.join(nginx_dir, "nginx.conf")

payload = crossplane.parse(nginx_conf_path)
print(payload)
print()

cert_dir = os.path.join(source_dir, "CA/storage/domain-certificates", IPAddr)
previous_cert_path = json_extract(payload, "ssl_certificate")[0]
previous_cert_filename = os.path.basename(previous_cert_path)
m = re.match(r"cert(?P<cert_num>\d+)\.pem", previous_cert_filename)
previous_cert_num = int(m.group("cert_num"))
new_cert_filename = "cert{}.pem".format(previous_cert_num + 1)
new_cert_path = os.path.join(cert_dir, new_cert_filename)

new_payload = nested_replace(payload, "server_name", IPAddr)
new_payload = nested_replace(new_payload, "ssl_certificate", new_cert_path)

print(new_payload)
config = crossplane.build(new_payload["config"][0]["parsed"])

with open(nginx_conf_path, "w") as f:
    f.write(config)
    f.close()


