# pagemerge
A cli tool for recursively merging multiple files into one using tags (with tab indentation). I originally built this for seperating html files into smaller static components then merging together during build time, but this could be useful for anything.

## Usage
pagemerge [headFile] [outputFile]

Inside your files, include {{nameOfFileToInsert}} to have a file injected when running pagemerge.

Keep the filepaths inside your tags relative to the headFile.

## Example
![]("./images/index.png")
