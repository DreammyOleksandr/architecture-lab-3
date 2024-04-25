$port = "http://localhost:17000"

 curl -Uri "$port" -Method POST -Body "reset" 
curl -Uri "$port" -Method POST -Body "white"
curl -Uri "$port" -Method POST -Body "update"