# Getting Started 

The purpose of this github repo is to give you a sample for getting up and running with Heroku and Slack. 

Had a bit of trouble with guides so said I'd put this up for anyone who needs it. 

In your slack chat simply to the Apps section and create an app.

From there you can create a bot and add it to a channel, take the API token for later.


# Godep use
Once you take this repo you MUST RUN godep 

if you do not have godep please use this link: 

> [GODEP](https://github.com/tools/godep)

Make sure to follow their guide on using godep on a new project.

```bash
$ godep save
```

If the error is something to do with adding dependencies (happened me) try this:

```bash
$ go get package/name
```

Entering package name that the error throws (if any)

# Folder Structure

After running "godep" successfully you should be left with a structure like this:
```
ProjectName/
    Godeps/
        Godeps.json
        readme
    vendor/ 
        github repos of code used her
    main.go 
    Procfile 
    readme.md
```

# Procfile!

Within the Procfile you will need to change the worker to the **Name** of the project folder for when to push it to Heroku! 

So from the example above it would be:

```
worker: ProjectName
```

# Creating Heroku App

> [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)

Link above on how to install CLI. 

After installing CLI simply use: 
```
$ heroku create ProjectName
```

Will ask for for information to login to your Heroku account if you haven't used CLI before. 

Once you've the app created and all that jazz above, you are on the final stretch! 

Using git you initialise an empty repo and commit, for those who don't know it's this:
```
$ git init 
$ git add . 
$ git commit -m "inital commit" 
``` 
Now here is where you've got one more thing to do, if you look at the main.go file you'll see that there's this:
```Go
api := os.Getenv("api")
```
Now, to set this in your heroku project, from the terminal enter the following:
```
$ heroku set:config api=api-key-here
```
Enter the api key given when you setup the slack bot 

Once that's done its simply: 

```
$ git push heroku master 
```


Boom! You should be up and running! Hopefully this works for you!