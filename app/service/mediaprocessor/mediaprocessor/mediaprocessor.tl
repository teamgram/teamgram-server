// Server side api

///////////////////////////////
///////// Nebula server API
///////////////////////////////

---types---
processorDerivative kind:string object_id:string file_name:string mime_type:string size:long width:int height:int bytes:bytes = ProcessorDerivative;
processedPhoto original_object_id:string sizes:Vector<ProcessorDerivative> = ProcessedPhoto;
processedDocument object_id:string mime_type:string size:long attributes:bytes thumbs:Vector<ProcessorDerivative> = ProcessedDocument;

---functions---
mediaProcessor.processPhoto owner_id:long object_id:string read_lease:bytes file_name:string profile:Bool = ProcessedPhoto;
mediaProcessor.processGif owner_id:long object_id:string read_lease:bytes file_name:string thumb_object_id:string thumb_read_lease:bytes = ProcessedDocument;
mediaProcessor.processMp4 owner_id:long object_id:string read_lease:bytes file_name:string attributes:bytes = ProcessedDocument;

// LAYER 0
