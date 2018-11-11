#!/bin/bash

PASSWORD=12345678

NAME1=test
NAME2=Kevin

while true
 do
    echo $PASSWORD| iriscli4 send --home=/home/ubuntu/.iriscli --to=faa15mqfse0jrckyvglw3dgk08lwcg9js886jw8jyc --from=$NAME2 --chain-id=fuxi-3001 --amount=1000000000000000000iris  --gas=10000 --fee=2000000000000000iris

    sleep 2s

    echo -n "."

    echo $PASSWORD| iriscli4 send --home=/home/ubuntu/.iriscli  --to=faa1geum38wcdcdkaj88hfuf80r7l7x2etc0jwqtx4 --from=$NAME1 --chain-id=fuxi-3001 --amount=1000000000000000000iris  --gas=10000 --fee=2000000000000000iris

    sleep 2s

    echo -n "."

 done