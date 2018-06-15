# prometheusExporter_selenium_grid
Prometheus exporter for Selenium Grid

Only exports backlog, free slots, and total slots right now. 

```
# HELP selenium_session_backlog Number of Selenium requests waiting for a session
# TYPE selenium_session_backlog gauge
selenium_session_backlog 0

# HELP selenium_slots_free Number of Selenium requests currently occupying a slot
# TYPE selenium_slots_free gauge
selenium_slots_free 20

# HELP selenium_slots_total Number of Selenium slots total
# TYPE selenium_slots_total gauge
selenium_slots_total 20
```
