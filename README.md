# Hangman
[![Build Status](https://cloud.drone.io/api/badges/elahe-dastan/Hangman/status.svg)](https://cloud.drone.io/elahe-dastan/Hangman)

### Implemented Hangman game using Go

There is a game loop in the main of the program which breaks whenever the player wins (there is no looser :joy:)<br/>
This code contains a directory called bank containing words which a random one is selected from, and a handler package <br/>
that is responsible for checking if the letter the player guesses is included in the word or not and also checks if the <br/>
game is finished.
