param(
  [string]$Image = "ghcr.io/hyperledger/bevel-build"
)

$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectDir = Resolve-Path (Join-Path $ScriptDir "..")
$RepoRoot = Resolve-Path (Join-Path $ProjectDir "..\..")
$KubeConfig = Join-Path $ProjectDir "config"

if (-not (Test-Path $KubeConfig)) {
  throw "Missing kubeconfig. Copy it to $KubeConfig before deploying."
}

docker run --rm -it `
  -v "${RepoRoot}:/home/bevel" `
  -v "${ProjectDir}:/home/bevel/build" `
  $Image `
  /home/bevel/run.sh
