package main

import (
    "github.com/neovim/go-client/nvim/plugin"
)

func Hello(args []string) (string, error) {
    return "Hello, World!", nil
}

func main() {
    plugin.Main(func(p *plugin.Plugin) error {
        p.HandleFunction(&plugin.FunctionOptions{Name: "Hello"}, Hello)
        return nil
    })
}

