curl -X POST "localhost:5601/api/kibana/dashboards/import" -H 'kbn-xsrf: true' -H 'Content-Type: application/json' -d @config/dashboard.json
