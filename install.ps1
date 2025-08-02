
Write-Output "Installing HoloDebug..."
New-Item -ItemType Directory -Force -Path "C:\HoloDebug"
Copy-Item -Recurse * "C:\HoloDebug"
$env:Path += ";C:\HoloDebug\bin"
Write-Output "Installed to C:\HoloDebug"
