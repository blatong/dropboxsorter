# dropboxsorter

Sorts files from sources to destinations
=======

This is a cron-run binary that moves files from a source directory to a destination directory by crafting and running an `rsync` commandline

## Example

dropboxsorter.json
```
[
	{"source": "~/dropbox/Archive", "destination": "/volume1/Archive"},
	{"source": "/volume1/Upload/Dropbox", "destination": "/volume1/Upload"}
]
```

`dropboxsorter` on this file would rsync all the files from `${HOME}/dropbox/Archive` into `/volume1/Archive` and all the files from `/volume1/Upload/Dropbox/` into `/volume1/Upload`

## Building

	go build ./...

