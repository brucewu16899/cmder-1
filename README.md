# README -- EDH NOTES

Quite custom build - sourced from 1.2, but massively changed.


add paths for new programs into:

	vendor/profile.ps1 for powershell
	vendor/init.bat for cmd
	vendor/msys2

This line in ConEmu.xml config breaks Far.exe left click:
see: https://github.com/Maximus5/ConEmu/issues/334#event-419601706

	<!-- EVIL LINE; BREAKS FAR.EXE				<value name="WndDragKey" type="dword" data="80808001"/>		-->

Try to maintain portability. I BELIEVE all customizations to be in: (but don't count on it)
	!! also remember that user settings are in HOME DIRECTORY
	
	vendor/profile.ps1
	vendor/init.bat
	config/
	bin/
	github's bin additions
	scripts/helpfile
	vendor/irssi
!!!!	vendor/github-for-windows/cmd/start-ssh-agent.bat
------------> I rewrote this and it is important!!!!!

(put the conemu plugins into farmanager)

# Directory Copy (no files)

This will create a directory of all the directories, empty

	xcopy /t /e "C:\cmder" cmder-dir