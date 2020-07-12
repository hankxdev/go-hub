# Gohub - add, commit and push with one line command

This is just a super simple tool written in Golang to commit and push code at the same time.

## Without it, if you want to commit and push your code

- git add . 
- git commit -m 'some message'
- git push

or you can do this:
```bash
git add .; git commit -m 'some message' ; git push
```


All that matters is the commit message, but you need to input others every time. Even if you enabled autocomplete in bash, you still need to locate to the message and modify it.

# with Gohub

You just need to do:

```bash
go-hub -m 'some message'
```

And it will do all the things above for you.


This is just a suoer simple tool, and I'm barely new to writing GO and CLI. Make sure you really want to add all your changed files before you run this command.
