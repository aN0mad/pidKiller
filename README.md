# pidKiller
A go project to handle killing a process by PID.

## Makefile
There is a makefile in the repository that makes building and cleaning the folder structure easier

```bash
root@f6a960422947:/workspace# make help
 help: print this help message
 tidy: format code and tidy modfile
 build: build the unix version
 buildwin: build the windows version
 all: build all applications for unix and windows
 clean: clean the repository
```

## Build
The easiest way is to compile with the included Docker container or just on Linux in general. Golang cross-compilation makes this easy.
### Linux
```bash
make build
```

### Windows
```bash
make buildwin
```

## Usage
```bash
➜  bin ./pidKiller.elf -c data.yaml
2024/10/29 18:44:52 INFO Starting pidKiller
2024/10/29 18:44:52 INFO Start time: 2024-10-29 18:44:52.932254625 +0000 UTC m=+0.045336682
Time remaining: 00:00:00
2024/10/29 18:46:52 INFO Killed process: Responder, PID: 1823
```

## Help
```bash
➜  bin ./pidKiller.elf -h
Usage of ./pidKiller.elf:
  -c string
        Config file to read (default "/workspace/example/example.yaml")
  -debug
        Enable debug logging
```

## Config yaml file
- `terminate` - Bucket to hold terminate data
- `signal` - Not used at this time
- `hours` - Hours to run before executing kill commands
- `minutes` - Minutes to run before executing kill commands
- `seconds` - Seconds to run before executing kill commands
- `processes` - Bucket to hold process type data
- `name` - Name to be used for the PID, only used for logging
- `pid` - PID to be killed

```yaml
terminate:
  signal: 9 
  hours: 0
  minutes: 0
  seconds: 10
processes:
  - name: Responder
    pid: 182
  - name: ntlmrelayx
    pid: 8219
```

## High-level Code flow
1. Read config files
2. Save start time
3. Calculate total duration from config file
4. Loop until total duration from start time has elapsed
5. Loop all processes in config file looking for OS process PIDs that match
6. If PID matches, the process is killed