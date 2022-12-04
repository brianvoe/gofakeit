package data

var Error = map[string][]string{
	"object": {
		"server",
		"service",
		"database",
		"buffer",
		"connection",
		"pointer",
		"parameter",
		"argument",
		"protocol",
		"hostname",
		"port",
		"url",
		"uri",
		"method",
		"request",
		"response",
		"header",
		"tag",
		"signature",
		"undefined",
	},
	"inputField": { // TODO: move to html generator
		"title",
		"firstName",
		"middleName",
		"lastName",
		"suffix",
		"addressLine1",
		"addressLine2",
		"postalCode",
		"city",
		"state",
		"province",
		"country",
		"dateOfBirth",
		"cardNumber",
		"description",
		"message",
		"status",
	},
	"httpMethod": {
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
	},
	"generic": {
		"error",
		"syntax error",
		"requested {object} is unavailable",
		"failed to {hackerverb} {object}",
		"expected {object} is undefined",
		"[object Object]",
		"no such variable",
		"variable not initialized",
		"variable assigned before declaration",
	},
	"database": {
		"sql error",
		"database connection lost",
		"table does not exist",
		"unique key constraint",
		"table migration failed",
		"bad connection",
		"destination pointer is nil",
	},
	"grpc": {
		"rpc error",
		"connection refused",
		"connection closed",
		"connection is shut down",
		"client protocol error",
	},
	"http": {
		"http error",
		"cross-origin-resource-policy error",
		"feature not supported",
		"trailer header without chunked transfer encoding",
		"no multipart boundary param in Content-Type",
		"request Content-Type isn't multipart/form-data",
		"header too long",
		"entity body too short",
		"missing ContentLength in HEAD response",
		"named cookie not present",
		"invalid method",
		"connection has been hijacked",
		"request method or response status code does not allow body",
		"wrote more than the declared Content-Length",
		"{httpMethod} not allowed",
	},
	"http-client": { // 400s
		"bad request",                   // 400
		"unauthorized",                  // 401
		"payment required",              // 402
		"forbidden",                     // 403
		"not found",                     // 404
		"method not allowed",            // 405
		"not acceptable",                // 406
		"proxy authentication required", // 407
		"request timeout",               // 408
		"conflict",                      // 409
		"gone",                          // 410
		"length required",               // 411
		"precondition failed",           // 412
		"payload too large",             // 413
		"URI too long",                  // 414
		"unsupported media type",        // 415
		"im a teapot",                   // 418
	},
	"http-server": { // 500s
		"internal server error",           // 500
		"not implemented",                 // 501
		"bad gateway",                     // 502
		"service unavailable",             // 503
		"gateway timeout",                 // 504
		"http version not supported",      // 505
		"variant also negotiates",         // 506
		"insufficient storage",            // 507
		"loop detected",                   // 508
		"not extended",                    // 510
		"network authentication required", // 511
	},
	"input": {
		"input error",
		"validation error",
		"verification error",
		"invalid format",
		"missing required field",
		"{inputField} is required",
		"{inputField} max length exceeded",
		"{inputField} must be at exactly 16 characters",
		"{inputField} must be at exactly 32 bytes",
		"failed to parse {inputField}",
		"date is in the past",
		"payment details cannot be verified",
	},
	"runtime": {
		"panic: runtime error: invalid memory address or nil pointer dereference",
		"address out of bounds",
		"undefined has no such property 'length'",
		"not enough arguments",
		"expected 2 arguments, got 3",
	},
}
