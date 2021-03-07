package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var configDefault = `{
  "autosheller": {
    "asp": false,
    "php": true
  },
  "core": {
    "batchmode": false,
    "threads": 100,
    "timeouts": 10
  },
  "dumper": {
    "keepblanks": false,
    "minrowcount": 500,
    "minrows": false,
    "targeted": false,
    "usedios": true
  },
  "exploiter": {
    "dbms": {
      "mssql": true,
      "mysql": true,
      "oracle": true,
      "postgresql": true
    },
    "heuristics": true,
    "intensity": 2,
    "techniques": {
      "blind": false,
      "error": true,
      "stacked": false,
      "union": true
    }
  },
  "generator": {
    "limit": true,
    "max": 50000,
    "parameters": [
      {
        "Data": null,
        "FilePath": "kw.txt",
        "Name": "Keywords",
        "Prefix": "(KW)"
      },
      {
        "Data": null,
        "FilePath": "pf.txt",
        "Name": "Page Formats",
        "Prefix": "(PF)"
      },
      {
        "Data": null,
        "FilePath": "pt.txt",
        "Name": "Page Types",
        "Prefix": "(PT)"
      },
      {
        "Data": null,
        "FilePath": "pp.txt",
        "Name": "Page Parameter",
        "Prefix": "(PP)"
      },
      {
        "Data": null,
        "FilePath": "sf.txt",
        "Name": "Search Functions",
        "Prefix": "(SF)"
      },
      {
        "Data": null,
        "FilePath": "de.txt",
        "Name": "Domains",
        "Prefix": "(DE)"
      }
    ],
    "patterns": [
      "(KW).(PT)?(PP)=",
      "(KW) / php?(PP)=",
      "\"(KW)\" / (DE).(PT)?(PP)=",
      ".(PT)? \"(KW)\" / (DE) (PP)=",
      "(KW)? \"(DE)\""
    ]
  },
  "scraper": {
    "engines": {
      "aol": false,
      "bing": false,
      "duckduckgo": false,
      "ecosia": false,
      "google": false,
      "googleapi": false,
      "mywebsearch": false,
      "qwant": false,
      "startpage": false,
      "yahoo": false,
      "yandex": false
    },
    "filter": true,
    "pages": 2
  }
}`

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		ioutil.WriteFile("config.json", []byte(configDefault), os.ModeAppend)
		viper.ReadInConfig()
	}
}