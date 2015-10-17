﻿# if CMDER_ROOT not set (for some reason in admin mode this doesn't happen)
if( -not $env:CMDER_ROOT ){
	$env:CMDER_ROOT = Split-Path -Path $PSScriptRoot -Parent 
}

# Compatibility with PS major versions <= 2
if(!$PSScriptRoot) {
    $PSScriptRoot = Split-Path $Script:MyInvocation.MyCommand.Path
}

# Add Cmder modules directory to the autoload path.
$CmderModulePath = Join-path $PSScriptRoot "psmodules/"

# Add GOPATH
if (Test-Path -Path "C:\Go" ) {
	$env:GOROOT = "C:\Go"
	$env:Path += ";$env:GOROOT\bin\"
	$env:GOPATH = (Join-Path $env:USERPROFILE "\dev")
}
if (Test-Path -Path "~\dev\lib\go" ) {
	$env:GOROOT = (Join-Path $env:USERPROFILE "\dev\lib\go")
	$env:Path += ";$env:GOROOT\bin\"
	$env:GOPATH = (Join-Path $env:USERPROFILE "\dev")
}

if( -not $env:PSModulePath.Contains($CmderModulePath) ){
    $env:PSModulePath = $env:PSModulePath.Insert(0, "$CmderModulePath;")
}

try {
    # Check if git is on PATH, i.e. Git already installed on system
    Get-command -Name "git" -ErrorAction Stop >$null
} catch {
	$env:Path += ";$env:CMDER_ROOT\vendor\git-for-windows\bin"
	# add binaries
	$env:Path += ";$env:CMDER_ROOT\vendor\git-for-windows\usr\bin"
	# and setup aliases
	#set-alias vim bashcall
	set-alias vi bashcall
	set-alias gunzip bashcall
	set-alias irssi bashcall
}
# set alias (for no arguments ) // functions (for arguments)
Remove-Item alias:wget
function wget { wget.exe --no-check-certificate $args }

# set-alias -passthru vim bashcall # debug
function bashcall {
	#echo "entered bashcall"
	#echo "args="
	#echo $args
	#echo "last command="
	#echo $$
	#echo "myinvocation="
	#echo $myinvocation
	#echo "myinvocation.invocationname="
	#echo $myinvocation.invocationname
	& ( $env:CMDER_ROOT + "\vendor\git-for-windows\bin\bash.exe" ) $myinvocation.invocationname $args
}

try {
    Import-Module -Name "posh-git" -ErrorAction Stop >$null
    $gitStatus = $true
} catch {
    Write-Warning "Missing git support, install posh-git with 'Install-Module posh-git' and restart cmder."
    $gitStatus = $false
}

function checkGit($Path) {
    if (Test-Path -Path (Join-Path $Path '.git/') ) {
        Write-VcsStatus
        return
    }
    $SplitPath = split-path $path
    if ($SplitPath) {
        checkGit($SplitPath)
    }
}

# Set up a Cmder prompt, adding the git prompt parts inside git repos
function global:prompt {
    $realLASTEXITCODE = $LASTEXITCODE
    $Host.UI.RawUI.ForegroundColor = "White"
    Write-Host $pwd.ProviderPath -NoNewLine -ForegroundColor Green
    if($gitStatus){
        checkGit($pwd.ProviderPath)
    }
    $global:LASTEXITCODE = $realLASTEXITCODE
    Write-Host "`nλ" -NoNewLine -ForegroundColor "DarkGray"
    return " "
}

# Run the GIT- Start Agent Script rather than Post-Git
if ($gitStatus) {
#	Start-SshAgent
    & ( $env:CMDER_ROOT + "\vendor\git-for-windows\cmd\start-ssh-agent.cmd" )
#	$Env:SSH_AUTH_SOCK=$env:TMP + "\ssh-agent.sock"
}

# Move to the wanted location
if (Test-Path Env:\CMDER_START) {
    Set-Location -Path $Env:CMDER_START
} elseif ($Env:CMDER_ROOT -and $Env:CMDER_ROOT.StartsWith($pwd)) {
    Set-Location -Path $Env:USERPROFILE
}

# Enhance Path
$env:Path = "$Env:CMDER_ROOT\bin;$env:Path;$Env:CMDER_ROOT"