# Renamer CLI Tool

This is a command-line tool written in Go that allows you to rename files and directories by replacing a specified part of their names with another string. The tool can also replace text within the files themselves, preserving the casing of the original text.

## Installation

To install the tool, clone the repository and build the executable:

```sh
git clone https://github.com/christian-gama/renamer.git
cd renamer
make build
```

Usage
To use the tool, run the renamer executable followed by four arguments:

**\<path\>**: The path to the directory containing the files/directories you want to rename.
**\<old_name\>**: The text you want to replace in the names of the files/directories.
**\<new_name\>**: The replacement text for the old text.
**\<replace_inside_files\>**: Set this to "true" if you want to replace text within the files as well.

Example:

```txt
./renamer /path/to/your/files old_text new_text true
```

This will rename all files and directories within /path/to/your/files by replacing old_text with new_text. If replace_inside_files is set to "true", the text within the files will also be replaced.
