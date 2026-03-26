# Threatgator

This project is called "gator" it is a guided project that is a part of the boot.dev curriculum. The focus was to create a CLI blog aggregator that can register users, add feeds, allow feeds to be followed, and print the posts to the CLI. 

## Requirements:
- Postgres
- Go

## Installation
- Git clone the repo to your local system: `git clone https://github.com/HyperFlyGuy/Threatgator.git`
- Change directory to the root of the project run `go build . -o gator`
- Add the binary to your path or reference it directly
- You are good to go!

## Config Setup
Your configuration file is fairly simple. Current user name will change with the register command, which we will explain later. You'll need to replace your user name with the proper one for the database.
```
{
"db_url":"postgres://username:@localhost:5432/gator?sslmode=disable",
"current_user_name":"alice"
}
```
## Usage
login - Login as a user (accepts one arguement)
register - Registers a new user in the database (accepts one arguement)
reset - Resets the database (CAUTION)
users - List the available users
agg - Will go and fetch posts from the followed feeds at a specified interval (time argument)
addfeed - Adds a feed to the database (name and url as args)
feeds - Lists all feeds
follow - Follows a feed as the current user. 
following - Lists feeds that the current user is following
unfollow - Unfollow a feed.
browse - Lists a number of posts
