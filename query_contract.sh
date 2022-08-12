#!/bin/sh

CONFIG_DIR=${CONFIG_DIR:-$HOME/.gotabit}
BIN="${BINARY:-gotabitd} --home $CONFIG_DIR"
CONTROLLER_ADDRESS=gio1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqhhpzlg
RESOLVER_ADDRESS=gio1eyfccmjm6732k7wp4p6gdjwhxjwsvje44j0hfx8nkgrm8fs7vqfsvq7yce

RES=$($BIN query wasm contract-state smart $CONTROLLER_ADDRESS '{"get_nodehash":{"name":"alice101"}}')

NODE=$(ruby -rjson -e "res = JSON.parse('$RES'); puts res['node']")

$BIN tx wasm execute $RESOLVER_ADDRESS "{\"set_text_data\":{\"node\":$NODE,\"key\":\"name\",\"value\":\"alice\"}}"
