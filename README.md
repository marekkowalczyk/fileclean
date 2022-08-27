# fileclean
Clean up text in a file by removing ill-formed and non-printing characters and by normalizing composite Unicode characters to their NFC forms for easy processing with LaTeX etc.

Basic usage: ````fileclean /path/to/file.txt```` emits output to ````STDOUT````.

Throws error if no files are give or if more than one file is given. Does not accept ````STDIN```` input or text as argument (yet).

In order to modify a file in place, use: ````fileclean file.txt | sponge file.txt````

To see what ````sponge```` does, go to <https://manned.org/sponge>.
