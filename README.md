# Prometheus cadvisor proxy (for nodes)

This is a debugging tool used to expose cadvisor metrics on a node via an ssh+http proxy.

## Usage

1. Make sure you have downloaded Prometheus and extracted it to `prometheus-x.y.z`.
2. Go to the `prometheus-tools` directory.
3. Generate a `targets.json` by running:
```sh
python3 gen-targets.py > targets.json
```
4. Copy `prometheus.yml` and `targets.json` into your Prometheus config. (If you are familiar with Prometheus, you can merge `prometheus.yml` with your existing config, otherwise it will work on its own.)
5. Run Prometheus!
