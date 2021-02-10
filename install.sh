#/bin/bash

if [[ "${GITHUB_USER}" == "" ]]; then
echo "You need to define the GITHUB_USER env"
exit 1
fi

curl -o konector https://github.com/atomgenie/konector/releases/download/1.0.0/konector
chmod u+x konector
sudo mv konector /usr/bin
konector init --username=$GITHUB_USER
sudo konector init-systemctl --user=$user
systemctl enable konector
