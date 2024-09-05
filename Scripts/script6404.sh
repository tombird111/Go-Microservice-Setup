#!/bin/sh
ID="~good+4+u"
AUDIO=`base64 -i "$ID".wav`
URL=localhost:3002/cooltown
echo "{ \"Audio\":\"$AUDIO\" }" > input
curl -v -X POST -d @input $URL
read -p "Press any key to continue..."