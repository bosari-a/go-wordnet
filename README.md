# Wordnet based Dictionary API

## Description
This package provides a function to get word definitions based on the wordnet database files.

## What is WordNet?
It is a large database for English words, their definitions, and other lexical information. The project is made by Princeton University.

- [WordNet website](https://wordnet.princeton.edu/)

## Usage

The function `GetDefinitions` is exported in the package:
```go
func (word *Word) GetDefinitions(dictPath string) error
```

- `dictPath` is the path to the folder containing the Wordnet database files. Check the [downloads](#downloads) section below for more details on how to get the database files.

## Downloads
### linux 
- Database files: [WNdb-3.0.tar.gz](https://wordnetcode.princeton.edu/3.0/WNdb-3.0.tar.gz)
- [What command do I need to unzip/extract a .tar.gz file?](https://askubuntu.com/a/25348)

### windows 
- Download the [WordNet browser, command-line tool, and database files with InstallShield self-extracting installer](https://wordnetcode.princeton.edu/2.1/WordNet-2.1.exe)

- Follow the installation process and make sure you remember the folder where you installed `WordNet`.

- Go to the `dict` folder and copy the highlighted files (below) into a database folder (e.g. name it `dict`) in your project:
![](https://raw.githubusercontent.com/bosari-a/wordnet-parser/main/assets/windowswordnet.png)

## License
- [WordNet License](https://wordnet.princeton.edu/license-and-commercial-use)
