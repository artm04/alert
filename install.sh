#!/bin/bash

copy () {
	if [[ $1 == "yes" ]] ; then
		sudo chmod +x alert
		sudo mv alert /bin/
	elif [[ $1 == "no" ]] ; then
		exit
		else
		echo 'invalid installation type'
	fi
}

PS3='Installation type: '
types=("yes" "no")

select type in "${types[@]}" ; do
	if [[ -n $types ]] ; then
		break
	else
		echo 'Please select between user and system installation'
	fi
done

sleep 2

copy $type
