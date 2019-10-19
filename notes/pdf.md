# Portable Document Format (PDF)

## File Structure

> This section is copied directly from Wikipedia [2]

A PDF file is a 7-bit ASCII file, except for certain elements that may have binary content. 
A PDF file starts with a header containing the magic number and the version of the format such as `%PDF-1.7`. 
The format is a subset of a COS ("Carousel" Object Structure) format. 
A COS tree file consists primarily of objects, of which there are eight types:

- Boolean values, representing true or false
- Numbers
- Strings, enclosed within parentheses (`(...)`), may contain 8-bit characters.
- Names, starting with a forward slash `(/)`
- Arrays, ordered collections of objects enclosed within square brackets (`[...]`)
- Dictionaries, collections of objects indexed by Names enclosed within double pointy brackets (`<<...>>`)
- Streams, usually containing large amounts of data, which can be compressed and binary
- The null object

Furthermore, there may be comments, introduced with the percent sign (`%`). 
Comments may contain 8-bit characters. 

Objects may be either *direct* (embedded in another object) or *indirect*. 
Indirect objects are numbered with an *object number* and a *generation number* and defined between the `obj` and `endobj` keywords. 
An index table, also called the cross-reference table and marked with the `xref` keyword, follows the main body and gives the byte offset of each indirect object from the start of the file.
This design allows for efficient random access to the objects in the file, and also allows for small changes to be made without rewriting the entire file (*incremental update*). 
Beginning with PDF version 1.5, indirect objects may also be located in special streams known as *object streams*. 
This technique reduces the size of files that have large numbers of small indirect objects and is especially useful for *Tagged PDF*. 

At the end of a PDF file is a trailer introduced with the `trailer` keyword. It contains

- A dictionary
- An offset to the start of the cross-reference table (the table starting with the `xref` keyword)
- And the `%%EOF` end-of-file marker.

The dictionary contains

- A reference to the root object of the tree structure, also known as the catalog
- The count of indirect objects in the cross-reference table
- And other optional information.


## Fonts

### Font Encodings

For large fonts or fonts with non-standard glyphs, the special encodings Identity-H (for horizontal writing) and Identity-V (for vertical) are used. 
With such fonts it is necessary to provide a ToUnicode table if semantic information about the characters is to be preserved. [1]

### Emoji

- [Emoji in Adobe PDF (Adobe Forums)](https://forums.adobe.com/message/10767530#10767530)

## Links

- [History of the Portable Document Format (PDF) (Wikipedia)](https://en.wikipedia.org/wiki/History_of_the_Portable_Document_Format_(PDF))
- [PDF (Wikipedia)](https://en.wikipedia.org/wiki/PDF)

[1]: https://en.wikipedia.org/wiki/PDF#Encodings
[2]: https://en.wikipedia.org/wiki/PDF#File_structure
