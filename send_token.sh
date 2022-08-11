#!/bin/sh

CONFIG_DIR=${CONFIG_DIR:-$HOME/.gotabit}
BIN="${BINARY:-gotabitd} --home $CONFIG_DIR"
FLAGS="--keyring-backend=test"
MNEMONIC=${MNEMONIC:-"dinner proud piano mention silk plunge forest fold trial duck electric define"}
USERNAME="sender"

if ! $BIN keys show $USERNAME $FLAGS; then
    echo MNEMONIC $MNEMONIC
    (echo $MNEMONIC; echo $PASSWORD; echo $PASSWORD ) | $BIN keys add $USERNAME --recover $FLAGS
fi

$BIN keys add user $FLAGS

$BIN tx bank send validator $($BIN keys show -a $USERNAME $FLAGS) 1000000000stake $FLAGS --chain-id=test --broadcast-mode=block --yes

$BIN tx bank send $($BIN keys show -a $USERNAME $FLAGS) $($BIN keys show -a user $FLAGS) 600000000stake $FLAGS --chain-id=test --broadcast-mode=block --yes

# $BIN tx bank send $($BIN keys show -a $USERNAME $FLAGS) $($BIN keys show -a user $FLAGS) 600000000ugtb $FLAGS --chain-id=test --broadcast-mode=block --yes
