make all

rm -rf ~/.aether

mkdir ~/.aether

aether init --chain-id test test
aether keys add test --recover --keyring-backend test<<<"y
wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please"
aether add-genesis-account test 100000000000000stake --keyring-backend test
aether gentx test 1000000000stake --chain-id test --keyring-backend test
aether collect-gentxs
aether start