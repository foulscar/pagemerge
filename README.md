# pagemerge
A cli tool for recursively merging multiple files into one using tags (with tab indentation). I originally built this for seperating html files into smaller static components then merging together during build time, but this could be useful for anything.

## Installation
You can simply run the following to install pagemerge
```bash
go install github.com/foulscar/pagemerge
```

## Usage
```
Usage: pagemerge [OPTION] [HEADFILE] [OUTPUTFILE]

Options:
  help        Show this help message and exit

Examples:
  pagemerge help
  pagemerge index.html ../dist/index.html
```

Inside your files, include {{nameOfFileToInsert}} to have a file injected when running pagemerge and keep the filepaths inside your tags relative to the headFile.

## Example
![index.html](images/index.png)
![mainWrapper.html](images/mainWrapper.png)
![navbar.html](images/navbar.png)
![content.html](images/content.png)
![pagemerge](images/command.png)
![output.html](images/output.png)
