$port = "http://localhost:17000"
$interval = 100
$zero = 0
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

  if ($currentPosX -lt $endPosX -and $moveX -gt $zero) {
    $moveX = $step 
    $moveY = $zero 
  } elseif ($currentPosY -lt $endPosY -and $moveY -ne -$step) {
    $moveX = $zero
    $moveY = $step
  } elseif ($startPosX -lt $currentPosX) {
    $moveX = -$step
    $moveY = $zero
  } elseif ($startPosY -lt $currentPosY) {
    $moveX = $zero 
    $moveY = -$step
  } else {
    $moveX = $step
    $moveY = $zero
  }

  $currentPosX += $moveX
  $currentPosY += $moveY
  
  curl -Uri "$port" -Method POST -Body "move $moveX $moveY"
  curl -Uri "$port" -Method POST -Body "update"
  sleep -Milliseconds $interval
}