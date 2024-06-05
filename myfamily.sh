#!/bin/bash
curl -s https://platform.zone01.gr/assets/superhero/all.json | jq --arg heroID "$HERO_ID" ' .[] | select(.id ==($heroID|tonumber)) | .connections.relatives' | tr -d '\n""'
