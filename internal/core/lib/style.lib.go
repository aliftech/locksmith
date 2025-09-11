package lib

import "github.com/fatih/color"

// Font Color
var Cyan = color.New(color.FgCyan).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Yellow = color.New(color.BgHiYellow).SprintFunc()

// Font Weight
var Bold = color.New(color.Bold).SprintFunc()
var Italic = color.New(color.Italic).SprintFunc()
var Underline = color.New(color.Underline).SprintFunc()
