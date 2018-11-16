# Reminder

This is a simple project to experiment Golang development.

It is expected to be started daily by cron to remember birthdays and other
recuring events.

## Installation

```shell
go get -u -v github.com/kmmndr/reminder
```

## Usage

```shell
$ reminder --help
Usage of reminder:
  -reminder-conf string
    	Path to reminder.conf (default "reminder.conf")
```

There is a [sample config file](reminder.conf). The syntax is very simple:
```
2018:11:16:Anniversaire de Monsieur Robert # Ignored comments
```
