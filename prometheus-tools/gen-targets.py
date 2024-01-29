import json
import requests
import re

r = requests.get("https://auth.sagecontinuum.org/manifests/")
r.raise_for_status()
items = r.json()
items.sort(key=lambda item: item["vsn"])

targets = [
    {
        "labels": {
            "vsn": item["vsn"].upper(),
            "node": item["name"].lower(),
            "project": (item.get("project") or "").lower(),
        },
        "targets": [
            "localhost:9911",
        ],
    }
    for item in items
    if item["name"]
    and item["phase"] == "Deployed"
    and re.match(r"^W[0-9A-Z]+$", item["vsn"])
]

print(json.dumps(targets, sort_keys=True))
