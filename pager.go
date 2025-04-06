// DEFINITION PAGE: A PAGE is the smallest unit of data read from or written to disk by a database.The database doesn't write a single word at a time. It always reads/writes in full blocks (PAGE).

// STRUCTURE PAGE: A PAGE is a fixed-size block of data, typically 4KB or 8KB, that the database uses to store rows of data. The structure of a page can vary depending on the database system, but it generally includes the following components:
// +-------------------------+
// | 		Page Header      | ← metadata about this page (type, #rows, etc.)
// +-------------------------+
// |  Row Pointers / Offsets | ← pointers to row positions in data area (slot directory) GROWS DOWNWARD
// +-------------------------+
// | 		Free Space       | ← grows/shrinks as rows are inserted/deleted
// +-------------------------+
// | 	 Actual Row Data     | ← raw serialized rows (structs → binary) GROWS UPWARD
// +-------------------------+

// DEFINITON METADATA: Metadata = “data about data”. It tells you how to interpret or manage actual data — without being the data itself.
// Metadata Field		What it tells you
// Page Type			Is this page a table, index, overflow?
// Number of Records	How many rows are in this page?
// Free Space Offset	Where can we write the next row?
// Slot Table 			Start	Where do row pointers start?
// Page Number (ID)		What is this page’s ID?
// Checksum / CRC		Is this page corrupted? (optional)

// DEFINITION ROW POINTERS / OFFSETS (ADDRESS): Row pointers are offsets that point to the location of each row in the page. They are used to quickly access and manage rows within a page. The row pointers are typically stored in a slot directory, which is a separate area within the page that contains the offsets for each row. The slot directory allows for efficient access to rows without having to scan the entire page. The row pointers are usually stored as integers or fixed-length binary values, and they indicate the starting position of each row within the page. When a row is inserted or deleted, the row pointers are updated accordingly to reflect the new positions of the rows. Row pointers are stored just after the page header and before the actual row data. 
// QUESTION: If a new row pointer is inserted, do we need to shift the data? ANSWER: No,  No, we don’t shift the row data. When a new row is inserted: The actual row data is placed at the bottom of the page, growing upward. A new slot (pointer) is inserted into the slot directory, which grows downward from the top. The slot just points to where the new row's data is stored — no need to touch or move existing row data.	

// STRUCTURE ROW POINTERS / OFFSETS :
// +----------------------+
// |	 Page Header      |
// +----------------------+
// | Slot 0 (offset: 8050)|
// | Slot 1 (offset: 8020)|
// +----------------------+
// |      Free Space      |
// +----------------------+
// |  Row 1 Data @ 8020   |
// |  Row 0 Data @ 8050   |
// +----------------------+
package main