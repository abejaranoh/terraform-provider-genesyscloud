if gpg --list-secret-keys | grep 67C54A5E 
then
    echo "secret exists"
else
    gpg --no-tty --batch --yes --import /tmp/terraform_gpg_secret.asc
    echo "done importing gpg key"
fi

gpg --list-secret-keys