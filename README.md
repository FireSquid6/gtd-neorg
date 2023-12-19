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

Whenever you come up with something that you have to do, you write it down as just a new line in your inbox file.

# Coming Soon

- An events input system
- A calendar view that's displayed on localhost

# How to Use it

1. Clone the git repo
2. Run `go build`
3. Copy the executable to somewhere you'll be able to find later
4. Run the executable once a night or so
