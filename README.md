# DHT11TemperaturePoller

A simple GO application that reads data from a given serial port and publishes
metrics that can be scraped by for example Prometheus.

This is intended to run in pair with [this code](https://github.com/fowlie/dht11)
running on the measuring device.