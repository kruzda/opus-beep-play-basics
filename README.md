# Plays an opus file with Go  
This project is the result of me working out how to play an Opus file using Beep  
To read the opus file it uses https://github.com/hraban/opus  
At the time Beep did not support PCM input, but there was an additon in the works by @icholy (https://github.com/faiface/beep/pull/109)  
Right now this had to be downloaded separately and referred to in the go.mod file  

Requires the following development headers:  
* libalsa
* libopus
* libopusfile  
