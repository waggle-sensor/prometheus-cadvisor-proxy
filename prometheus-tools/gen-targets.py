import json
import requests

vsns = [
    "W021",
    "W026",
    "W02C",
    "W02E",
    "W015",
    "W01B",
    "W022",
    "W017",
    "W023",
]

r = requests.get("https://api.sagecontinuum.org/api/state")
items = r.json()["data"]
node_for_vsn = {item["vsn"].upper(): item["id"].lower() for item in items if item.get("vsn") and item.get("id")}

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
