import json
import requests

#  missing???

vsns = sorted(set("""
W02E
W02C
W026
W030
W039
W014
W01E
W01B
W018
W024
W021
W023
W028
W019
W022
W02B
W02F
W029
W02D
W015
W01C
W016
W073
W06F
W084
W06D
W057
W059
W045
W06B
""".split()))

r = requests.get("https://api.sagecontinuum.org/api/state")
items = r.json()["data"]
node_for_vsn = {item["vsn"].upper(): item["id"].lower() for item in items if item.get("vsn") and item.get("id")}
# missing???
node_for_vsn["W06B"] = node_for_vsn.get("W06B", "000048B02D3AE349")

targets = [
    {
        "labels": {
            "vsn": vsn,
            "node": node_for_vsn[vsn],
        },
        "targets": [
            "localhost:9911",
        ]
    }
    for vsn in vsns
]

print(json.dumps(targets, sort_keys=True))
