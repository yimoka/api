// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: fault/fault.proto

package fault

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	ErrorReason_BAD_REQUEST ErrorReason = 0
	// 微信 api 出错
	ErrorReason_WE_CHAT_API_FAILED ErrorReason = 300
	// 当前请求需要用户验证。
	ErrorReason_UNAUTHORIZED ErrorReason = 1
	// 此响应码保留以便将来使用，创造此响应码的最初目的是用于数字支付系统，然而现在并未使用
	ErrorReason_PAYMENT_REQUIRED ErrorReason = 2
	// 服务器已经理解请求，但是拒绝执行它。
	ErrorReason_FORBIDDEN ErrorReason = 3
	// 请求失败，请求所希望得到的资源未被在服务器上发现。
	ErrorReason_NOT_FOUND ErrorReason = 4
	// 请求行中指定的请求方法不能被用于请求相应的资源。
	ErrorReason_METHOD_NOT_ALLOWED ErrorReason = 5
	// 请求的资源的内容特性无法满足请求头中的条件，因而无法生成响应实体。
	ErrorReason_NOT_ACCEPTABLE ErrorReason = 6
	// 与401响应类似，只不过客户端必须在代理服务器上进行身份验证。
	ErrorReason_PROXY_AUTHENTICATION_EQUIRED ErrorReason = 7
	// 请求超时。客户端没有在服务器预备等待的时间内完成一个请求的发送。客户端可以随时再次提交这一请求而无需进行任何更改。
	ErrorReason_REQUEST_TIMEOUT ErrorReason = 8
	// 由于和被请求的资源的当前状态之间存在冲突，请求无法完成。这个代码只允许用在这样的情况下才能被使用：用户被认为能够解决冲突，并且会重新提交新的请求。该响应应当包含足够的信息以便用户发现冲突的源头。
	ErrorReason_CONFLICT ErrorReason = 9
	// 被请求的资源在服务器上已经不再可用，而且没有任何已知的转发地址。这样的状况应当被认为是永久性的。如果可能，拥有链接编辑功能的客户端应当在获得用户许可后删除所有指向这个地址的引用。
	ErrorReason_GONE ErrorReason = 10
	// 服务器拒绝在没有定义 Content-Length 头的情况下接受请求。
	ErrorReason_LENGTH_REQUIRED ErrorReason = 11
	// 服务器在验证在请求的头字段中给出先决条件时，没能满足其中的一个或多个。这个状态码允许客户端在获取资源时在请求的元信息（请求头字段数据）中设置先决条件，以此避免该请求方法被应用到其希望的内容以外的资源上。
	ErrorReason_PRECONDITION_FAILED ErrorReason = 12
	// 服务器拒绝处理当前请求，因为该请求提交的实体数据大小超过了服务器愿意或者能够处理的范围。
	ErrorReason_PAYLOAD_TOO_LARGE ErrorReason = 13
	// 请求的URI长度超过了服务器能够解释的长度，因此服务器拒绝对该请求提供服务。
	ErrorReason_URI_TOO_LONG ErrorReason = 14
	// 对于当前请求的方法和所请求的资源，请求中提交的实体并不是服务器中所支持的格式，因此请求被拒绝。
	ErrorReason_UNSUPPORTED_MEDIA_TYPE ErrorReason = 15
	// 如果请求中包含了 Range
	// 请求头，并且Range中指定的任何数据范围都与当前资源的可用范围不重合，同时请求中又没有定义
	// If-Range 请求头，那么服务器就应当返回416状态码。
	ErrorReason_RANGE_NOT_SATISFIABLE ErrorReason = 16
	// 此响应代码意味着服务器无法满足 EXPECT 请求标头字段指示的期望值。
	ErrorReason_EXPECTATION_FAILED ErrorReason = 17
	// 服务器拒绝尝试用 “茶壶冲泡咖啡”。
	ErrorReason_IM_A_TEAPOT ErrorReason = 18
	// 该请求针对的是无法产生响应的服务器这可以由服务器发送，该服务器未配置为针对包含在请求URI中的方案和权限的组合产生响应
	ErrorReason_MISDIRECTED_REQUEST ErrorReason = 21
	// 请求格式良好，但由于语义错误而无法遵循。
	ErrorReason_UNPROCESSABLE_ENTITY ErrorReason = 42
	// 正在访问的资源被锁定。
	ErrorReason_LOCKED ErrorReason = 23
	// 由于先前的请求失败，所以此次请求失败。
	ErrorReason_FAILED_DEPENDENCY ErrorReason = 24
	// 服务器不愿意冒着风险去处理可能重播的请求。
	ErrorReason_Too_Early ErrorReason = 25
	// 服务器拒绝使用当前协议执行请求
	ErrorReason_UPGRADE_REQUIRED ErrorReason = 26
	// 原始服务器要求该请求是有条件的。
	// 旨在防止“丢失更新”问题，即客户端获取资源状态，修改该状态并将其返回服务器，同时第三方修改服务器上的状态，从而导致冲突。
	ErrorReason_PRECONDITION_REQUIRED ErrorReason = 28
	// 用户在给定的时间内发送了太多请求（“限制请求速率”）。
	ErrorReason_TOO_MANY_REQUESTS ErrorReason = 29
	// 服务器不愿意处理请求，因为它的 请求头字段太大
	ErrorReason_REQUEST_HEADER_FIELDS_TOO_LARGE ErrorReason = 31
	// 用户请求非法资源，例如：由政府审查的网页。
	ErrorReason_UNAVAILABLE_FOR_LEGAL_REASONS ErrorReason = 51
	// 服务器遇到了不知道如何处理的情况
	ErrorReason_INTERNAL_SERVER_ERROR ErrorReason = 100
	// 此请求方法不被服务器支持且无法被处理。只有GET和HEAD是要求服务器支持的，它们必定不会返回此错误代码
	ErrorReason_NOT_IMPLEMENTED ErrorReason = 101
	// 此错误响应表明服务器作为网关需要得到一个处理这个请求的响应，但是得到一个错误的响应
	ErrorReason_BAD_GATEWAY ErrorReason = 102
	// 服务器没有准备好处理请求。 常见原因是服务器因维护或重载而停机。
	ErrorReason_SERVICE_UNAVAILABLE ErrorReason = 103
	// 当服务器作为网关，不能及时得到响应时返回此错误代码。
	ErrorReason_GATEWAY_TIMEOUT ErrorReason = 104
	// 服务器不支持请求中所使用的HTTP协议版本。
	ErrorReason_HTTP_VERSION_NOT_SUPPORTED ErrorReason = 105
	// 服务器有一个内部配置错误：对请求的透明内容协商导致循环引用
	ErrorReason_VARIANT_ALSO_NEGOTIATES ErrorReason = 106
	// 服务器有内部配置错误：所选的变体资源被配置为参与透明内容协商本身，因此不是协商过程中的适当端点。
	ErrorReason_INSUFFICIENT_STORAGE ErrorReason = 107
	// 服务器在处理请求时检测到无限循环
	ErrorReason_LOOP_DETECTED ErrorReason = 108
	// 客户端需要对请求进一步扩展，服务器才能实现它。服务器会回复客户端发出扩展请求所需的所有信息。
	ErrorReason_NOT_EXTENDED ErrorReason = 110
	// 状态码指示客户端需要进行身份验证才能获得网络访问权限。
	ErrorReason_NETWORK_AUTHENTICATION_REQUIRED ErrorReason = 111
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:   "BAD_REQUEST",
		300: "WE_CHAT_API_FAILED",
		1:   "UNAUTHORIZED",
		2:   "PAYMENT_REQUIRED",
		3:   "FORBIDDEN",
		4:   "NOT_FOUND",
		5:   "METHOD_NOT_ALLOWED",
		6:   "NOT_ACCEPTABLE",
		7:   "PROXY_AUTHENTICATION_EQUIRED",
		8:   "REQUEST_TIMEOUT",
		9:   "CONFLICT",
		10:  "GONE",
		11:  "LENGTH_REQUIRED",
		12:  "PRECONDITION_FAILED",
		13:  "PAYLOAD_TOO_LARGE",
		14:  "URI_TOO_LONG",
		15:  "UNSUPPORTED_MEDIA_TYPE",
		16:  "RANGE_NOT_SATISFIABLE",
		17:  "EXPECTATION_FAILED",
		18:  "IM_A_TEAPOT",
		21:  "MISDIRECTED_REQUEST",
		42:  "UNPROCESSABLE_ENTITY",
		23:  "LOCKED",
		24:  "FAILED_DEPENDENCY",
		25:  "Too_Early",
		26:  "UPGRADE_REQUIRED",
		28:  "PRECONDITION_REQUIRED",
		29:  "TOO_MANY_REQUESTS",
		31:  "REQUEST_HEADER_FIELDS_TOO_LARGE",
		51:  "UNAVAILABLE_FOR_LEGAL_REASONS",
		100: "INTERNAL_SERVER_ERROR",
		101: "NOT_IMPLEMENTED",
		102: "BAD_GATEWAY",
		103: "SERVICE_UNAVAILABLE",
		104: "GATEWAY_TIMEOUT",
		105: "HTTP_VERSION_NOT_SUPPORTED",
		106: "VARIANT_ALSO_NEGOTIATES",
		107: "INSUFFICIENT_STORAGE",
		108: "LOOP_DETECTED",
		110: "NOT_EXTENDED",
		111: "NETWORK_AUTHENTICATION_REQUIRED",
	}
	ErrorReason_value = map[string]int32{
		"BAD_REQUEST":                     0,
		"WE_CHAT_API_FAILED":              300,
		"UNAUTHORIZED":                    1,
		"PAYMENT_REQUIRED":                2,
		"FORBIDDEN":                       3,
		"NOT_FOUND":                       4,
		"METHOD_NOT_ALLOWED":              5,
		"NOT_ACCEPTABLE":                  6,
		"PROXY_AUTHENTICATION_EQUIRED":    7,
		"REQUEST_TIMEOUT":                 8,
		"CONFLICT":                        9,
		"GONE":                            10,
		"LENGTH_REQUIRED":                 11,
		"PRECONDITION_FAILED":             12,
		"PAYLOAD_TOO_LARGE":               13,
		"URI_TOO_LONG":                    14,
		"UNSUPPORTED_MEDIA_TYPE":          15,
		"RANGE_NOT_SATISFIABLE":           16,
		"EXPECTATION_FAILED":              17,
		"IM_A_TEAPOT":                     18,
		"MISDIRECTED_REQUEST":             21,
		"UNPROCESSABLE_ENTITY":            42,
		"LOCKED":                          23,
		"FAILED_DEPENDENCY":               24,
		"Too_Early":                       25,
		"UPGRADE_REQUIRED":                26,
		"PRECONDITION_REQUIRED":           28,
		"TOO_MANY_REQUESTS":               29,
		"REQUEST_HEADER_FIELDS_TOO_LARGE": 31,
		"UNAVAILABLE_FOR_LEGAL_REASONS":   51,
		"INTERNAL_SERVER_ERROR":           100,
		"NOT_IMPLEMENTED":                 101,
		"BAD_GATEWAY":                     102,
		"SERVICE_UNAVAILABLE":             103,
		"GATEWAY_TIMEOUT":                 104,
		"HTTP_VERSION_NOT_SUPPORTED":      105,
		"VARIANT_ALSO_NEGOTIATES":         106,
		"INSUFFICIENT_STORAGE":            107,
		"LOOP_DETECTED":                   108,
		"NOT_EXTENDED":                    110,
		"NETWORK_AUTHENTICATION_REQUIRED": 111,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_fault_fault_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_fault_fault_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_fault_fault_proto_rawDescGZIP(), []int{0}
}

var File_fault_fault_proto protoreflect.FileDescriptor

var file_fault_fault_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb0, 0x09, 0x0a, 0x0b, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x0b, 0x42, 0x41,
	0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x1d, 0x0a, 0x12, 0x57, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x41, 0x50, 0x49,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xac, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03,
	0x12, 0x16, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44,
	0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x04,
	0xa8, 0x45, 0x92, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45,
	0x4e, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c,
	0x0a, 0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x18, 0x0a, 0x0e,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x06,
	0x1a, 0x04, 0xa8, 0x45, 0x96, 0x03, 0x12, 0x26, 0x0a, 0x1c, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x97, 0x03, 0x12, 0x19,
	0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x43, 0x4f, 0x4e,
	0x46, 0x4c, 0x49, 0x43, 0x54, 0x10, 0x09, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x0e, 0x0a,
	0x04, 0x47, 0x4f, 0x4e, 0x45, 0x10, 0x0a, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x19, 0x0a,
	0x0f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x0b, 0x1a, 0x04, 0xa8, 0x45, 0x9b, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x50, 0x52, 0x45, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x9c, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x50, 0x41, 0x59, 0x4c, 0x4f,
	0x41, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x0d, 0x1a, 0x04,
	0xa8, 0x45, 0x9d, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x55, 0x52, 0x49, 0x5f, 0x54, 0x4f, 0x4f, 0x5f,
	0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x0e, 0x1a, 0x04, 0xa8, 0x45, 0x9e, 0x03, 0x12, 0x20, 0x0a, 0x16,
	0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0f, 0x1a, 0x04, 0xa8, 0x45, 0x9f, 0x03, 0x12, 0x1f,
	0x0a, 0x15, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x41, 0x54, 0x49,
	0x53, 0x46, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x10, 0x1a, 0x04, 0xa8, 0x45, 0xa0, 0x03, 0x12,
	0x1c, 0x0a, 0x12, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x11, 0x1a, 0x04, 0xa8, 0x45, 0xa1, 0x03, 0x12, 0x15, 0x0a,
	0x0b, 0x49, 0x4d, 0x5f, 0x41, 0x5f, 0x54, 0x45, 0x41, 0x50, 0x4f, 0x54, 0x10, 0x12, 0x1a, 0x04,
	0xa8, 0x45, 0xa2, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x4d, 0x49, 0x53, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x15, 0x1a, 0x04, 0xa8,
	0x45, 0xa5, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x55, 0x4e, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x2a, 0x1a, 0x04, 0xa8,
	0x45, 0xa6, 0x03, 0x12, 0x10, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x17, 0x1a,
	0x04, 0xa8, 0x45, 0xa7, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f,
	0x44, 0x45, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x18, 0x1a, 0x04, 0xa8, 0x45,
	0xa8, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x54, 0x6f, 0x6f, 0x5f, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x10,
	0x19, 0x1a, 0x04, 0xa8, 0x45, 0xa9, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x55, 0x50, 0x47, 0x52, 0x41,
	0x44, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x1a, 0x1a, 0x04, 0xa8,
	0x45, 0xaa, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x1c, 0x1a, 0x04,
	0xa8, 0x45, 0xac, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x10, 0x1d, 0x1a, 0x04, 0xa8, 0x45, 0xad,
	0x03, 0x12, 0x29, 0x0a, 0x1f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x48, 0x45, 0x41,
	0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c,
	0x41, 0x52, 0x47, 0x45, 0x10, 0x1f, 0x1a, 0x04, 0xa8, 0x45, 0xaf, 0x03, 0x12, 0x27, 0x0a, 0x1d,
	0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x53, 0x10, 0x33, 0x1a,
	0x04, 0xa8, 0x45, 0xc3, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x64,
	0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d,
	0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x65, 0x1a, 0x04, 0xa8, 0x45, 0xf5,
	0x03, 0x12, 0x15, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59,
	0x10, 0x66, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x67, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x47, 0x41, 0x54, 0x45, 0x57,
	0x41, 0x59, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x68, 0x1a, 0x04, 0xa8, 0x45,
	0xf8, 0x03, 0x12, 0x24, 0x0a, 0x1a, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x69, 0x1a, 0x04, 0xa8, 0x45, 0xf9, 0x03, 0x12, 0x21, 0x0a, 0x17, 0x56, 0x41, 0x52, 0x49,
	0x41, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x53, 0x4f, 0x5f, 0x4e, 0x45, 0x47, 0x4f, 0x54, 0x49, 0x41,
	0x54, 0x45, 0x53, 0x10, 0x6a, 0x1a, 0x04, 0xa8, 0x45, 0xfa, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x49,
	0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52,
	0x41, 0x47, 0x45, 0x10, 0x6b, 0x1a, 0x04, 0xa8, 0x45, 0xfb, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x4c,
	0x4f, 0x4f, 0x50, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x6c, 0x1a, 0x04,
	0xa8, 0x45, 0xfc, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x6e, 0x1a, 0x04, 0xa8, 0x45, 0xfe, 0x03, 0x12, 0x29, 0x0a, 0x1f,
	0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x6f, 0x1a, 0x04, 0xa8, 0x45, 0xff, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x3b, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x3b, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_fault_fault_proto_rawDescOnce sync.Once
	file_fault_fault_proto_rawDescData = file_fault_fault_proto_rawDesc
)

func file_fault_fault_proto_rawDescGZIP() []byte {
	file_fault_fault_proto_rawDescOnce.Do(func() {
		file_fault_fault_proto_rawDescData = protoimpl.X.CompressGZIP(file_fault_fault_proto_rawDescData)
	})
	return file_fault_fault_proto_rawDescData
}

var file_fault_fault_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fault_fault_proto_goTypes = []any{
	(ErrorReason)(0), // 0: yimoka.api.fault.ErrorReason
}
var file_fault_fault_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fault_fault_proto_init() }
func file_fault_fault_proto_init() {
	if File_fault_fault_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fault_fault_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fault_fault_proto_goTypes,
		DependencyIndexes: file_fault_fault_proto_depIdxs,
		EnumInfos:         file_fault_fault_proto_enumTypes,
	}.Build()
	File_fault_fault_proto = out.File
	file_fault_fault_proto_rawDesc = nil
	file_fault_fault_proto_goTypes = nil
	file_fault_fault_proto_depIdxs = nil
}
