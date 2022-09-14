This task will teach you how to set up a new golang project, or module, and then a small "hello world" program. Modules are a way of disconnecting the go code you write from the libraries you would download to support it. Every language you write in has to manage shared code in some way. Sometimes that means pulling a library from a repo and putting it into a folder in the project, like with node, or putting all your libraries in one big folder structure and organizing them so there isn't duplicate traffic for pulling something you already have. Go is the later. This use to mean that if you wanted to make a new project in go you would have to build it in that structure. Modules remove that requirement so you can build them anywhere.

Steps for this session:
1. Make a new folder 'mkdir test-module'
2. Jump into that folder 'cd test-module'
3. Create the go.mod file 'go mod init test-module'

You should now see a file called go.mod in the folder. It should show your current version of go and the name of your "package".

That is not super fun though, because we need to prove it works. Let's set up a small go program to make sure it works.
1. create a new file in the test-module folder called main.go. 'touch main.go' or create it how you want. 
2. Use the contents of main.go in 1-init to fill it. (I would recommend copy/pasting the first time and slowly start doing it by memory).
3. Run it. 'go run main.go'

This should output "Hello World". Congrats you did it!

Now for this to be effective you should delete the test-module folder and do this at the beginning of your day and at the end until you no longer need to read the readme to get this right. Learning takes time, if you need to get frustrated then do that, if you need to get mad or sad or whatever then do what you need to. Learning is hard and computers aren't a second nature thing. Take your time and you will get it. Practice, Practice, Practice. You got this.