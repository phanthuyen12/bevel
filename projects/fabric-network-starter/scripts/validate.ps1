param(
  [string]$Image = "ghcr.io/hyperledger/bevel-build"
)

$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectDir = Resolve-Path (Join-Path $ScriptDir "..")
$RepoRoot = Resolve-Path (Join-Path $ProjectDir "..\..")

docker run --rm `
  -v "${RepoRoot}:/home/bevel" `
  -v "${ProjectDir}:/home/bevel/build" `
  $Image `
  ajv validate `
  -s /home/bevel/platforms/network-schema.json `
  -d /home/bevel/build/network.yaml
