from .utils import ADDRS, CONTRACTS, deploy_contract, eth_to_bech32, send_transaction


def test_precompiles(cronos):
    w3 = cronos.w3
    addr = ADDRS["validator"]
    amount = 100
    contract = deploy_contract(w3, CONTRACTS["TestBank"])
    tx = contract.functions.nativeMint(amount).buildTransaction({"from": addr})
    receipt = send_transaction(w3, tx)
    assert receipt.status == 1, "expect success"

    # query balance through contract
    assert contract.caller.nativeBalanceOf(addr) == amount
    # query balance through cosmos rpc
    cli = cronos.cosmos_cli()
    assert cli.balance(eth_to_bech32(addr), "evm/" + contract.address) == amount

    # test exception revert
    tx = contract.functions.nativeMintRevert(amount).buildTransaction(
        {"from": addr, "gas": 210000}
    )
    receipt = send_transaction(w3, tx)
    assert receipt.status == 0, "expect failure"

    # check balance don't change
    assert contract.caller.nativeBalanceOf(addr) == amount
    assert cli.balance(eth_to_bech32(addr), "evm/" + contract.address) == amount