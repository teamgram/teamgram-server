// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
photoSizeList size_id:long sizes:Vector<PhotoSize> dc_id:int = PhotoSizeList;
videoSizeList size_id:long sizes:Vector<VideoSize> dc_id:int = VideoSizeList;
mediaResolvedFileObject object_id:string read_lease:bytes size:long mime_type:string dc_id:int storage_file_type:int = MediaResolvedFileObject;

---functions---
media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;
media.uploadProfilePhotoFile flags:# owner_id:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = Photo;
media.getPhoto photo_id:long = Photo;
media.getPhotoSizeList size_id:long = PhotoSizeList;
media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
media.getVideoSizeList size_id:long = VideoSizeList;
media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
media.getDocument id:long = Document;
media.getDocumentList id_list:Vector<long> = Vector<Document>;
media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
media.getEncryptedFile id:long access_hash:long = EncryptedFile;
media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;
media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;
media.uploadRingtoneFile flags:# owner_id:long file:InputFile mime_type:string file_name:string = Document;

media.uploadedProfilePhoto owner_id:long photo_id:long = Photo;
media.resolveFileLocation location:InputFileLocation viewer_id:long = MediaResolvedFileObject;

// LAYER 0
