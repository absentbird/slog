Slog
====

###A simple logging application.

Slog takes a string and appends it with a timestamp to a CSV named slog.csv stored in the working directory. If slog.csv doesn't exist it will be created.

Slog can take input in 3 different ways.

1. As an argument: `slog "I'm not talking to you"`
2. From a pipe: `echo "hand-blown glass" | slog`
3. Or you will be asked to enter a record: `Input record: record`
