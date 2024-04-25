$port = "http://localhost:17000"
$interval = 100
$step = 0.05

$moveX = $step
$moveY = $step


$size = 0.5
$startPosX = $size / 2
$startPosY = $size / 2
$currentPosX = $startPosX
$currentPosY = $startPosX
$endPosX = 1 - $size / 2
$endPosY = 1 - $size / 2

curl -Uri "$port" -Method POST -B "reset"
curl -Uri "$port" -Method POST -B "white"
curl -Uri "$port" -Method POST -B "figure $currentPosX $currentPosY"
curl -Uri "$port" -Method POST -B "update"

while ($true) {

  if ($currentPosX -gt $endPosX) {
    $moveX = -$moveX
    $moveY = -$moveY
  } elseif ($currentPosX -lt $startPosX) {
    $moveX = -$moveX
    $moveY = -$moveY
  }

  $currentPosX += $moveX
  $currentPosY += $moveY
  
  curl -Uri "$port" -Method POST -Body "move $moveX $moveY"
  curl -Uri "$port" -Method POST -Body "update"
  sleep -Milliseconds $interval
}