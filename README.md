# Negtd

A modified version of the gtd system used with neorg and neovim. It requires:

- Go
- Neovim
- [norg](https://github.com/nvim-neorg/neorg)

# How it Works

Negtd consists of 3 (4 in the future files). They are located in ~/notes/gtd. I haven't implemented a system to change that folder. You can just modify the source code in `main.go` somewhere if you need to change that. These files are:

- `inbox.norg` - just an input for new ideas
- `agenda.norg` - your todo list, by date
- `backlog.norg` - ideas that you're not ready to do yet, but would like to do later

## The Inbox File

The inbox file is just an input for new ideas and tasks you have to do. It could look something like:

```norg
complete my homework
[tomorrow] Do this tomorrow
[?] i'll get around to this
[07/19/2024] Do this on the 19th of July in 2024
[_] I'm garbage
```

In this example, the "complete my homework" task will just stay in the inbox, the "Do this tomorrow" and "Do this on the 19th of July 2024" task will be moved to the agenda, and the "I'm garbage" task will be thrown away. The "I'll get around to this" task is just moved to the backlog.

## The Agenda File

The agenda file is the heart of your system. It contains a to-do list for every day. For example:

```norg
* 2024-12-18
- ( ) Do something
- ( ) Do another thing
- (x) A complete task

* 2024-12-19
- (-) Move me to the backlog
- (_) Trash me
```

What most of these do is fairly self explanatory. These checkboxes follow the norg standard. If you're unfamiliar with norg or neorg, consider checking it out on [github](https://github.com/nvim-neorg/neorg).

## The Backlog File

The backlog file behaves exactly like the inbox. It's a good place to store tasks that you want to do at some point, but not yet. You should try to clean it out for dead tasks whenever your weekly planning occurs.

## Tagging

I use a "postdata" system where I tag my tasks with certain metadata. For example:

```norg
- ( ) this is my task [$project #anything &mom @home]
```

These tags are:

- `$` - a "project" that this task is a part of
- `#` - any general metadata that I feel is important
- `&` - the "who" or "what" this task is for
- `@` - where or when the task can be done

No programmatic interfaces have been built for these yet. I may build some in the future, but I have no explicit plans to.

# Coming Soon

- An events input system
- A calendar view that's displayed on localhost
- A config file to change date format and file paths
- A garbage collection system

# How to Use it

1. Clone the git repo
2. Run `go build`
3. Copy the executable to somewhere you'll be able to find later
4. Run the executable once a night or so
