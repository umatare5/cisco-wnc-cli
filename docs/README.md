# Documentation

This directory contains the documentation for this CLI's commands, configuration, and technical details.

## For Users Docs

- **[Show Commands](command.show.md)**

  The views that read a controller, with the output each one prints.

- **[Action Commands](command.action.md)**

  The commands that write to a controller, with what each one answers.

- **[Other Commands](command.other.md)**

  The commands that name no target: `generate-token` and `save-config`.

- **[Customization](customization.md)**

  The flags, the environment variables, and the configuration file.

- **[Troubleshooting](troubleshooting.md)**

  Every error message this CLI prints, and what to do about each.

- **[Help](help.md)**

  The help text of every command except `completion`.

## For Developers Docs

- **[Architecture](architecture.md)**

  The output contract, the absence rule, the exit codes, and the order every write keeps.

- **[Measurements](measurements.md)**

  Every reading taken on a live controller, and the gaps.

- **[Testing](testing.md)**

  The conventions, the invariants, and the checks that need a live controller.
