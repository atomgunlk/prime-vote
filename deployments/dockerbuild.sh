#!/bin/sh
appversion=$(cat appversion.txt)

cd ../
echo "Building $appversion"
docker build --no-cache -t prime-vote:$appversion .
if [ $? -eq 0 ]; then
    echo "Build OK"
else
    echo "Build FAIL"
	return
fi

echo "Build version $appversion Success"
