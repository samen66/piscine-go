CURL=$(curl -L -s https://01.alem.school/assets/superhero/all.json | jq ' .[] | select( .id == '$HERO_ID' ).connections.relatives')
CURL=${CURL//$'\"'/}
echo $CURL
