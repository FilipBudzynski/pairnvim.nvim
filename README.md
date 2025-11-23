# Goal

The goal of PairVim is to allow for collaborative editing of code in Neovim.
The end result and experience should be similar to what you would get in Google Docs when collaborating with someone on a document.

# Roadmap

- [x] Basic pairing functionality
  - [x] Spawn a golang server
  - [x] Connect to it as host
  - [x] Enable connecting for other users by `:Pair` command
  - [] Send a simple message on `:Ping`
- [ ] Send current buffer to connecting users
- [ ] (AI) Add support for multiple selections
- [ ] (AI) Add support for multiple lines
- [ ] (AI) Add support for multiple files
- [ ] (AI) Add support for multiple windows
- [ ] (AI) Add support for multiple buffers
- [ ] (AI) Add support for multiple tabs
- [ ] (AI) Add support for multiple splits
- [ ] (AI) Add support for multiple tabs
- [ ] (AI) Add support for multiple splits

## Installation

### Lazy Installation

```lua
return {
  "filipbudzynski/pairvim.nvim",
  config = function()
    require("pairvim").setup()
  end,
}

```
