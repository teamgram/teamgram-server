// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
fileFinalizedObject object_id:string upload_session_id:string bucket:string key:string size:long mime_type:string sha256:bytes read_lease:bytes dc_id:int = FileFinalizedObject;
fileHashChunk offset:long limit:int hash:bytes = FileHashChunk;

---functions---
dfs.commitUpload upload_session_id:string owner_id:long file:InputFile purpose:string = FileFinalizedObject;
dfs.putFile owner_id:long purpose:string file_name:string mime_type:string bytes:bytes = FileFinalizedObject;
dfs.getFileByReadLease read_lease:bytes offset:long limit:int = upload.File;
dfs.getFileHashesByReadLease read_lease:bytes offset:long limit:int = Vector<FileHash>;

dfs.writeFilePartData flags:# creator:long file_id:long file_part:int bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
dfs.uploadPhotoFileV2 creator:long file:InputFile = Photo;
dfs.uploadProfilePhotoFileV2 flags:# creator:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = Photo;
dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
dfs.downloadFile location:InputFileLocation offset:long limit:int = upload.File;
dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
dfs.uploadGifDocumentMedia creator:long media:InputMedia = Document;
dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
dfs.uploadWallPaperFile creator:long file:InputFile mime_type:string admin:Bool = Document;
dfs.uploadThemeFile flags:# creator:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
dfs.uploadRingtoneFile creator:long file:InputFile mime_type:string file_name:string = Document;

dfs.uploadedProfilePhoto creator:long photo_id:long = Photo;

// LAYER 0
