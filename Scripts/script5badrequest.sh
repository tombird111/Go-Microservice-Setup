#!/bin/sh
ID="badbad"
AUDIO=`base64 -i "$ID".wav`
RESOURCE=localhost:3001/search
echo "{ \"Audio\":\"$AUDIO\" }" > input
curl -v -X POST -d @input $RESOURCE
read -p "Press any key to continue..."