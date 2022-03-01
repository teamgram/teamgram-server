# Uploading and Downloading Files
> Forwarded from [Uploading and Downloading Files](https://core.telegram.org/api/files)

When working with the API, it is sometimes necessary to send a relatively large file to the server. For example, when sending a message with a photo/video attachment or when setting the current user’s profile picture.

## Uploading files
There are a number of API methods to save files. The schema of the types and methods used is presented below:

```
inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
inputFileBig id:long parts:int name:string md5_checksum:string = InputFile;

inputEncryptedFileUploaded id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
inputEncryptedFileBigUploaded id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;

inputMediaUploadedPhoto#2dc53a7d file:InputFile = InputMedia;
inputMediaUploadedVideo#4847d92a file:InputFile duration:int w:int h:int = InputMedia;
inputMediaUploadedThumbVideo#e628a145 file:InputFile thumb:InputFile duration:int w:int h:int = InputMedia;

inputChatUploadedPhoto#94254732 file:InputFile crop:InputPhotoCrop = InputChatPhoto;

---functions---

messages.sendMedia#a3c85d76 peer:InputPeer media:InputMedia random_id:long = messages.StatedMessage;
messages.sendEncryptedFile peer:InputEncryptedChat random_id:long data:bytes file:InputEncryptedFile = messages.SentEncryptedMessage;

photos.uploadProfilePhoto#d50f9c88 file:InputFile caption:string geo_point:InputGeoPoint crop:InputPhotoCrop = photos.Photo;

upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
upload.saveBigFilePart file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
```

Before transmitting the contents of the file itself, the file has to be assigned a unique 64-bit client identifier: file_id.

The file’s binary content is then split into parts. All parts must have the same size ( part_size ) and the following conditions must be met:

- part_size % 1024 = 0 (divisible by 1KB)
- 524288 % part_size = 0 (512KB must be evenly divisible by part_size)

The last part does not have to satisfy these conditions, provided its size is less than part_size.

Each part should have a sequence number, file_part, with a value ranging from 0 to 2,999.

After the file has been partitioned you need to choose a method for saving it on the server. Use upload.saveBigFilePart in case the full size of the file is more than 10 MB and upload.saveFilePart for smaller files.

Each call saves a portion of the data in a temporary location on the server to be used later. The storage life of each portion of data is between several minutes and several hours (depending on how busy the storage area is). After this time, the file part will become unavailable. To increase the time efficiency of a file save operation, we recommend using a call queue, so X pieces of the file are being saved at any given moment in time. Each successful operation to save a part invokes the method call to save the next part. The value of X can be tuned to achieve maximum performance.

When using one of the methods mentioned above to save file parts, one of the following data input errors may be returned:

- `FILE_PARTS_INVALID` - Invalid number of parts. The value is not between 1..3000
- `FILE_PART_INVALID`: The file part number is invalid. The value is not between 0 and 2,999.
- `FILE_PART_TOO_BIG`: The size limit (512 KB) for the content of the file part has been exceeded
- `FILE_PART_EMPTY`: The file part sent is empty
- `FILE_PART_SIZE_INVALID` - 512KB cannot be evenly divided by part_size
- `FILE_PART_SIZE_CHANGED` - The part size is different from the size of one of the previous parts in the same file

While the parts are being uploaded, an MD5 hash of the file contents can also be computed to be used later as the md5_checksum parameter in the inputFile constructor.
After the entire file is successfully saved, the final method may be called and passed the generated inputFile object. In case the upload.saveBigFilePart method is used, the inputFileBig constructor must be passed, in other cases use inputFile.

The file save operation may return one of the following data input errors:

- `FILE_PARTS_INVALID`: The number of file parts is invalid The value is not between 1 and 3,000.
- `FILE_PART_Х_MISSING`: Part X (where X is a number) of the file is missing from storage. Try repeating the method call to resave the part.
- `MD5_CHECKSUM_INVALID`: The file’s checksum did not match the md5_checksum parameter

## Downloading files
There are methods available to download files which have been successfully uploaded. The schema of the types and methods used is presented below:

```
inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
inputVideoFileLocation#3d0364ec id:long access_hash:long = InputFileLocation;

upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;

storage.fileUnknown#aa963b05 = storage.FileType;
storage.fileJpeg#7efe0e = storage.FileType;
storage.fileGif#cae1aadf = storage.FileType;
storage.filePng#a4f63c0 = storage.FileType;
storage.fileMp3#528a0677 = storage.FileType;
storage.fileMov#4b09ebbc = storage.FileType;
storage.filePartial#40bc6f52 = storage.FileType;
storage.fileMp4#b3cea0e4 = storage.FileType;
storage.fileWebp#1081464c = storage.FileType;

---functions---

upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
```

Any file can be downloaded by calling upload.getFile. The data for the input parameter of the InputFileLocation type is taken from the video constructor for video recordings, otherwise it is taken from the fileLocation constructor. The size of each file in bytes is available, which makes it possible to download the file in parts using the parameters offset and limit, similar to the way files are uploaded. The parameter offset must be divisible by 1 KB.

The file download operation may return one of the following data input errors:

- `LOCATION_INVALID`: The file address is invalid
- `OFFSET_INVALID`: The offset value is invalid. The value is not divisible by 1 KB.
- `LIMIT_INVALID`: The limit value is invalid
- `FILE_MIGRATE_X`: The file is in Data Center No. X

## General Considerations
It is recommended that large queries (upload.getFile, upload.saveFilePart) be handled through a separate session and a separate connection, in which no methods other than these should be executed. If this is done, then data transfer will cause less interference with getting updates and other method calls.