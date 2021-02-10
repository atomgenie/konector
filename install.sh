#/bin/bash

if [[ -z "${GITHUB_USER}" ]]; then
echo "You need to define the GITHUB_USER env"
exit 1
fi

sudo curl -o /usr/bin/konector https://github.com/atomgenie/konector/releases/download/1.0.0/konector
konector init --username=$GITHUB_USER
sudo konector init-systemctl --user=$user
systemctl enable konector
