#!/bin/sh
RESOURCE=localhost:3000/tracks
curl -v -X GET $RESOURCE
read -p "Press any key to continue..."