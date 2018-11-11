#!/bin/bash

PASSWORD=12345678

NAME=iris

ACCOUNT=cosmos1shuqhpl273t96yg6nnqvyfeewj3ew3mdlgstsp

OPERATOR=cosmosvaloper1shuqhpl273t96yg6nnqvyfeewj3ew3md6uy7uj

DENOM=steak

MIN=10

while true
 do
    
    withdraw=$(echo PASSWORD| gaiacli tx withdraw-rewards --chain-id=gaia-9001 --from=$NAME --is-validator=true) 
	
    current_balance=$(gaiacli query account $ACCOUNT --trust-node=true| jq .value.coins[0].amount | tr -d '"')
	echo "balance is "  $current_balance

	if [ !"$current_balance" ] && [ "$current_balance"  -gt "$MIN" ];
	then

      echo " got more steaks!"
      
      output=$(echo $PASSWORD| gaiacli tx delegate --amount=$current_balance$DENOM --chain-id=gaia-9001 --from=$NAME  --validator=$OPERATOR --trust-node=true)
	fi

	if [ !"$output"  ]
	then
        echo -n "Error in delegation"
	else
        echo "delegation success"
	fi
    sleep 30s
 done