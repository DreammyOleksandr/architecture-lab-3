$port = "http://localhost:17000"

curl -Uri "$port" -Method POST -Body "reset"
curl -Uri "$port" -Method POST -Body "green"
curl -Uri "$port" -Method POST -Body "bgrect 0.25 0.25 0.75 0.75"
curl -Uri "$port" -Method POST -Body "update"