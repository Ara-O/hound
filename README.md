## Hound - CLI Tool 🤖

### Totally not ChatGPT Created Description

Imagine a nifty command-line tool that's your go-to sidekick for diving into documentation. You simply feed it a link to any documentation page, and voila! It becomes your trusty guide, ready to answer all your burning questions about that documentation

<img src="https://github.com/Ara-O/hound/assets/67078991/479bdc7a-5813-4522-ba9f-484f3503f7db" style="width: 60%">

### Features
- Easily query through documentation to find what you want
- Saves previously queried sites for future reference
- Easily works with obscure documentation sites

### How to run
- Install Golang - the latest version preferably
- Clone the repository into an IDE of your choice
- If you would like an executable, in the terminal run go build . -o hound
-- This will create an executable file that you can run with the command ./hound in the terminal
- If you would like to be able to run hound from wherever in your computer terminal, add the compiled binary ( the result of running go build -o hound ) to your environment path ( the process is a bit different for different operating systems ). Or you can run go install ., which will install the binary into the $GOPATH/bin folder. If this folder is part of your environment path, then you should be able to run it anywhere in your terminal now.

