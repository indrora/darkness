#include <std/sys.pat>
#include <std/mem.pat>

/* protobuf impl */



import std.core;
import std.mem;

import type.leb128;

struct ZigZag32 {
	u32 value;
} [[sealed, format("format_zigzag32")]];

fn format_zigzag32(ZigZag32 zigzag) {
	return s32((s32(zigzag.value) << 1) ^ (s32(zigzag.value) >> 31));
};

struct ZigZag64 {
	u64 value;
} [[sealed, format("format_zigzag64")]];

fn format_zigzag64(ZigZag64 zigzag) {
	return s64((s64(zigzag.value) << 1) ^ (s64(zigzag.value) >> 63));
};

enum WireType : u8 {
	Varint 			= 0,
	_64Bit 			= 1,
	LengthDelimited = 2,
	StartGroup 		= 3,
	EndGroup 		= 4,
	_32Bit 			= 5
};

bitfield Key {
	field_number : 5;
	wire_type    : 3;
} [[bitfield_order(std::core::BitfieldOrder::MostToLeastSignificant, 8)]];

union _64Bit {
	u64 fixed64;
	ZigZag64 sfixed64;
	double dbl;
};

union _32Bit {
	u32 fixed32;
	ZigZag32 sfixed32;
	float flt;
};

struct LengthDelimited {
	type::LEB128 length;
	char data[length];
};


struct ProtobufEntry {
	Key key;
	
	if (key.wire_type == WireType::Varint)
		type::LEB128 value;
	else if (key.wire_type == WireType::_64Bit)
		_64Bit value;
	else if (key.wire_type == WireType::LengthDelimited)
		LengthDelimited value;
	else if (key.wire_type == WireType::_32Bit)
		_32Bit value;
};

/* end protobuf impl */

#define RECORD_MAGIC "LELR"

enum MsgType:u8 {
    LightHeader = 0,
    ViewPrefrences = 1,
    GPSData = 2,
    unk1 = 3,
    unk2 = 4
};

union BuffBacked<T, auto sz> {
    u8 buff[sz];
    T  value[while( $ < addressof(buff) + sz)];
};


struct LEBlock {
    char magic[4];
    std::print("magic: {}",magic);
    std::assert(magic == RECORD_MAGIC, "Bad magic");
    u64 blocklen;
    u64 offset;
    u32 msglen;
    MsgType msgtype;
    u8 reserved[7] [[hidden]];
    u8 blockdata [blocklen-32];
    
    //u8 msgdata[msglen] @ (addressof(magic) + offset);
    
    BuffBacked<ProtobufEntry, msglen> message @ addressof(magic) + offset; 
    
    
    std::print("block len = {}, msg len = {}, offset = {}", blocklen, msglen, offset);
};



LEBlock blocks[while(!std::mem::eof())] @ 0x00;