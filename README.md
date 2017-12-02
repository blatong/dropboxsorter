# dropboxsorter

Sorts files from sources to destinations
=======

This is a cron-run binary that moves files from a source directory to a destination directory by crafting and running an `rsync` commandline.

Really, all it does is eliminate some boilerplate from scripting this up.  Instead of maintaining a multiline script, the user maintains a multi-structure JSON config that should be smaller and more concise.  It can be checked using any JSON validation tools, and can be viewed in a "dry run" mode to validate before exeuction.  Additionally, Environment variables are replaced based on the environment of the user running the utility.

## Example

dropboxsorter.json
```
[
	{"source": "~/dropbox/Archive", "destination": "/volume1/Archive"},
	{"source": "/volume1/Upload/Dropbox", "destination": "/volume1/Upload"}
]
```

`dropboxsorter` on this file would rsync all the files from `${HOME}/dropbox/Archive` into `/volume1/Archive` and all the files from `/volume1/Upload/Dropbox/` into `/volume1/Upload`

Keep in mind: currently, environment variables such as `${HOME}` (or `$HOME` if you live dangerously) are expanded to the values currently present in the environment when executed; "~" is not.  `~/dropbox/Archive` above is expanded by the `rsync` command running, not `dropboxsorter` preparing the command.

The same command leveraging pre-expansion by `dropboxsorter` would be:

dropboxsorter.json
```
[
	{"source": "${HOME}/dropbox/Archive", "destination": "/volume1/Archive"},
	{"source": "/volume1/Upload/Dropbox", "destination": "/volume1/Upload"}
]
```

## Building

	go build -ldflags '-w -extldflags "-static"' -o dropboxsorter .
