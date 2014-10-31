Slog
====

###A simple logging application.

Slog takes a string and appends it with a timestamp to a CSV named slog.csv stored in the working directory. If slog.csv doesn't exist it will be created.

Slog can take input in 3 different ways.

1. As an argument: `slog "I'm not talking to you"`
2. From a pipe: `echo "hand-blown glass" | slog`
3. Or you will be asked to enter a record: `Input record: record`

###Why?

Slog can be used to keep track of time spend on a project or to keep track of when your co-workers use the bathroom. The uses are endless! Any one-line event or memo that you want to attach to a timestamp is a contender.

If you can't think of a good use for it just call it minimalist micro-journaling.
