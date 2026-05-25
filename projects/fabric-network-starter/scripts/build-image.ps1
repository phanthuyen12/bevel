param(
  [string]$Image = "ghcr.io/hyperledger/bevel-build"
)

$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectDir = Resolve-Path (Join-Path $ScriptDir "..")
$RepoRoot = Resolve-Path (Join-Path $ProjectDir "..\..")

docker build $RepoRoot -t $Image
