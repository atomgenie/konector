#!/bin/sh

if [ "${GITHUB_USER}" = "" ]; then
echo "You need to define the GITHUB_USER env"
exit 1
fi

echo "DOWNLOAD"
curl -L -o konector https://github.com/atomgenie/konector/releases/download/1.0.0/konector
echo "CHMOD"
chmod u+x konector
echo "/usr/bin"
sudo mv konector /usr/bin
echo "Konector init"
konector init --username=$GITHUB_USER
echo "Konector systemctl"
sudo konector init-systemctl --user=$USER
echo "systemctl enable & start"
sudo systemctl enable konector
sudo systemctl start konector
