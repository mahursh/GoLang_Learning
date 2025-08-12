# Useful VScode commands:

### Terminal Stuff:
#### 1. Allow Scripts for Current User Only 
>``Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned``
>RemoteSigned means:
>  *  Scripts you write locally can run.
>  *  Downloaded scripts must be signed by a trusted publisher.


#### 2. some shortcuts
> ctrl + c : to end the execuation in terminal

> ctrl + l : to clean the terminal.

> ls : to see that files and packages .

#### 3. Running go files
>Runs the script :
>``go run main.go``


#### 4. creating package in go
>`go mod init PackageName`
>`go mod tidy`
---
### Version Control Commands:
#### 1. Clone:
>Press Ctrl+Shift+P (Windows) or Cmd+Shift+P (Mac) to open Command Palette.
>Search for:
>``Git: Clone``


#### 2. Commit:
>Click the Source Control icon on the left sidebar (or press Ctrl+Shift+G).

>You’ll see all changed files listed.

>Click the + icon next to a file to stage it,
or click the + at the top to stage all changes.
