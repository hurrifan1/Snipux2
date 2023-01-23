package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// App struct
type App struct {
	ctx  context.Context
	conf AppConfig
	data []Snippet
	// db  *bolt.DB
}

// ===============================================================================================
//  THEMEMODE
// ===============================================================================================

// New ThemeMode type
type ThemeMode int64

// Define new Enums for ThemeMode
const (
	System ThemeMode = iota
	Light
	Dark
)

// Add a String method to the ThemeMode enum type
func (s ThemeMode) String() string {
	switch s {
	case System:
		return "System"
	case Light:
		return "Light Mode"
	case Dark:
		return "Dark Mode"
	}
	return "Unknown Theme Mode"
}

// ===============================================================================================
//  LANGUAGES
// ===============================================================================================

// New CodeLanguage type
type CodeLanguage int64

// Define new Enums for CodeLanguage
const (
	Apple ThemeMode = iota
	Argdown
	Asm
	Audio
	Babel
	Bazel
	Bicep
	Bower
	Bsl
	CSharp
	C
	Cake_php
	Cake
	CheckboxUnchecked
	Checkbox
	Cjsx
	Clock
	Clojure
	CodeClimate
	CodeSearch
	Coffee_erb
	Coffee
	Coldfusion
	Config
	Cpp
	Crystal_embedded
	Crystal
	Css
	Csv
	Cu
	D
	Dart
	Db
	Default
	DeprecationCop
	Docker
	Editorconfig
	Ejs
	Elixir_script
	Elixir
	Elm
	Error
	Eslint
	Ethereum
	FSharp
	Favicon
	Firebase
	Firefox
	Folder
	Font
	Git_folder
	Git_ignore
	Git
	Github
	Gitlab
	Go
	Go2
	Godot
	Gradle
	Grails
	Graphql
	Grunt
	Gulp
	Hacklang
	Haml
	Happenings
	Haskell
	Haxe
	Heroku
	Hex
	Html_erb
	Html
	Ignored
	Illustrator
	Image
	Info
	Ionic
	Jade
	Java
	Javascript
	Jenkins
	Jinja
	Js_erb
	Json
	Julia
	Karma
	Kotlin
	Less
	License
	Liquid
	Livescript
	Lock
	Lua
	Makefile
	Markdown
	Maven
	Mdo
	Mustache
	NewFile
	Nim
	Notebook
	Npm_ignored
	Npm
	Nunjucks
	Ocaml
	Odata
	Pddl
	Pdf
	Perl
	Photoshop
	Php
	Pipeline
	Plan
	Platformio
	Powershell
	Prisma
	Project
	Prolog
	Pug
	Puppet
	Purescript
	Python
	R
	Rails
	React
	Reasonml
	Rescript
	Rollup
	Ruby
	Rust
	Salesforce
	Sass
	Sbt
	Scala
	Search
	Settings
	Shell
	Slim
	Smarty
	Spring
	Stylelint
	Stylus
	Sublime
	Svelte
	Svg
	Swift
	Terraform
	Tex
	TimeCop
	Todo
	Tsconfig
	Twig
	Typescript
	Vala
	Video
	Vue
	Wasm
	Wat
	Webpack
	Wgt
	Windows
	Word
	Xls
	Xml
	Yarn
	Yml
	Zig
	Zip
)

// Add a String method to the CodeLanguage enum type
// func (s CodeLanguage) String() string {
// 	switch s {
// 	case System:
// 		return "System"
// 	case Light:
// 		return "Light Mode"
// 	case Dark:
// 		return "Dark Mode"
// 	}
// 	return "Unknown Theme Mode"
// }

// ===============================================================================================
//  APP SETTINGS
// ===============================================================================================

// Add a new AppConfig type to handle app settings
type AppConfig struct {
	NumOfLinesToPreview int       `json:"NumOfLinesToPreview"`
	PrimaryColor        string    `json:"PrimaryColor"`
	ColorMode           ThemeMode `json:"ColorMode"`
}

// ===============================================================================================
//  SNIPPETS DATA
// ===============================================================================================

// Add a new SnippetList and Snippet type to handle app settings
// type SnippetList struct {
// 	SnippetList  []Snippet `json:"NumOfLinesToPreview"`
// 	PrimaryColor string    `json:"PrimaryColor"`
// }

type Snippet struct {
	Description     string `json:"description"`
	Name            string `json:"name"`
	LastModifiedGMT int    `json:"lastModifiedGMT"`
	Language        string `json:"language"`
	Code            string `json:"code"`
}

// ===============================================================================================
//
//	MISC
//
// ===============================================================================================
const settingsFilePath string = "frontend/src/data/snipux-settings.json"
const snippetsFilePath string = "frontend/src/data/snipux-snippets.json"

var settingsObj AppConfig
var snippetsObj []Snippet

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	json2Struct(settingsFilePath, &settingsObj)
	json2Struct(snippetsFilePath, &snippetsObj)

	log.Println("###### Well fuck me asunder...")

	a.ctx = ctx
	a.conf = settingsObj
	a.data = snippetsObj
}

func (a *App) GetCurrentSettings() AppConfig {
	// json2Struct(settingsFilePath, &settingsObj)

	return a.conf
}

func (a *App) GetCurrentSnippets() []Snippet {
	// json2Struct(snippetsFilePath, &snippetsObj)

	return a.data
}

// ===============================================================================================
//
//	HELPER FUNCTIONS
//
// ===============================================================================================
func json2Struct(fp string, t any) {
	rawFile, err1 := os.Open(fp)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Successfully Opened snipux-settings.json")

	defer rawFile.Close()

	rawFileByte, err2 := io.ReadAll(rawFile)
	if err2 != nil {
		fmt.Println(err2)
	}

	json.Unmarshal(rawFileByte, t)
}
