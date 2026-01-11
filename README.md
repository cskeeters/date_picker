`date_picker` is a TUI utility built on [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) and [cskeeters/bubble-datepicker](https://github.com/cskeeters/bubble-datepicker) (forked from [EthanEFung/bubble-datepicker](https://github.com/EthanEFung/bubble-datepicker))

# Example Use

Use in a script with:

```sh
SELECTED_DATE=$(date_picker -f "%Y-%b-%d" 2026-02-28)
if [[ $? -ne 0 ]]; then
    echo "User canceled"
    exit 1
fi
echo "User selected: $SELECTED_DATE"
```

<img src="https://github.com/cskeeters/i/blob/master/date_picker.webp?raw=true" style="width: intrinsic; zoom: 0.5;">

With the default date selected, SELECTED_DATE will be set to `2026-Feb-28`.


# Features

* The initial date can be provided as the first argument (`2026-01-01`)
* Designed to be used with scripts.
    * The only output to stdout is the selected dated
    * The UI uses `/dev/tty` so it can be used in scripts without issue.
* The format of the date selected can be specified with `-f` in [POSIX format](https://pubs.opengroup.org/onlinepubs/009695399/utilities/date.html) (without `+`).
* Nice to use keyboard navigation


# Keyboard Navigation

| Key      | Effect                                 |
|----------|----------------------------------------|
| `l`      | next day                               |
| `h`      | previous day                           |
| `j`      | +7 days (down)                         |
| `k`      | -7 days (up)                           |
|          |                                        |
| `J`      | Next Month                             |
| `K`      | Previous Month                         |
| `L`      | Next Year                              |
| `H`      | Previous Year                          |
|          |                                        |
| `Tab`    | Focus switch (Month, Year, Calendar)   |
| `Escape` | Exit 1 (Can detect with script)        |
| `Enter`  | Select the date and output to `stdout` |


# Defaults

Defaults to *today* with the format `%Y-%m-%d`.
