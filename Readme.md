This project acts as a simple text completion/editing/auto-correction tool.
Program is written in Go.
How to use:
The program will receive as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text will be placed in (the output).

Next is a list of possible modifications that program executes:

Every instance of (hex) replaces the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")

Every instance of (bin) replaces the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
    For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

Every instance of the punctuations ., ,, !, ?, : and ; will appear close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
    Except if there are groups of punctuation like: ... or !?. In this case the program will format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
    If there are more than one word between the two ' ' marks, the program will place the marks next to the corresponding words. 


