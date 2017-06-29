package identityservice

message GetSupportedVersionsRequest {
}

message GetSupportedVersionsResponse {
message Result {
// All the versions that the Plugin supports. This field is
// REQUIRED.
repeated Version supported_versions = 1;
}

// One of the following fields MUST be specified.
oneof reply {
Result result = 1
Error error = 2
}
}

// Specifies the version in Semantic Version 2.0 format.
type Version struct {
major uint32 = 1  // This field is REQUIRED.
uint32 minor = 2;  // This field is REQUIRED.
uint32 patch = 3;  // This field is REQUIRED.
}