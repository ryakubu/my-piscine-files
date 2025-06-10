#!/bin/bash

if [ -z "$HERO_ID" ]; then
  echo "Error: HERO_ID environment variable is not set."
  exit 1
fi

relatives=$(curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json | \
jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives")

# Convert actual newlines to literal "\n"
relatives_escaped=$(echo "$relatives" | sed ':a;N;$!ba;s/\n/\\n/g')

if [ -z "$relatives" ] || [ "$relatives" = "null" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID"
else
  echo "$relatives_escaped"
fi

